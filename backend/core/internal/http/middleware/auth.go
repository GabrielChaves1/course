package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	appcontext "github.com/GabrielChaves1/course/internal/application/context"
	"github.com/GabrielChaves1/course/internal/domain/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/patrickmn/go-cache"
)

const (
	DevUserIDStr = "c00505f3-3b80-4792-93d2-fcb0cd467d6f"
)

type JWK struct {
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

func getAccessTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("access_token")
	if err != nil {
		return "", err
	}

	return token, nil
}

var publicKeyCache = cache.New(1*time.Hour, 10*time.Minute)

func Auth(environment types.Environment, cognitoIssuer, cognitoAppClientId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if environment == types.Development {
			userID, err := types.NewUserIDFromString(DevUserIDStr)
			if err != nil {
				c.Error(err)
				return
			}

			ctx := appcontext.NewContextWithUserID(c.Request.Context(), userID)
			c.Request = c.Request.WithContext(ctx)

			c.Next()
			return
		}

		accessToken, err := getAccessTokenFromCookie(c)
		if err != nil {
			c.Error(err)
		}

		decodedJwt, err := decodeJWT(cognitoIssuer, cognitoAppClientId, accessToken)
		if err != nil {
			var statusCode int
			var errMsg string

			if errors.Is(err, jwt.ErrTokenExpired) {
				statusCode = http.StatusUnauthorized
				errMsg = "Token expired"
			} else {
				var validationErr *jwt.ValidationError
				if errors.As(err, &validationErr) {
					statusCode = http.StatusUnauthorized
					errMsg = "Invalid token"
				} else {
					statusCode = http.StatusUnauthorized
					errMsg = "Authentication error"
				}
			}

			c.JSON(statusCode, gin.H{
				"error":   "Unauthorized",
				"message": errMsg,
			})

			c.Abort()
			return
		}

		if claims, ok := decodedJwt.Claims.(jwt.MapClaims); ok {
			userIDStr, _ := claims["sub"].(string)

			userID, err := types.NewUserIDFromString(userIDStr)
			if err != nil {
				c.Error(err)
				return
			}

			ctx := appcontext.NewContextWithUserID(c.Request.Context(), userID)
			c.Request = c.Request.WithContext(ctx)
		}

		c.Next()
	}
}

func decodeJWT(cognitoIssuer, cognitoAppClientId, tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, nil)
	if err != nil && !strings.Contains(err.Error(), "No keyfuncs was provided") {
		return nil, err
	}

	if token == nil {
		return nil, errors.New("failed to parse token")
	}

	kid, ok := token.Header["kid"].(string)
	if !ok || kid == "" {
		return nil, errors.New("token is missing 'kid' header")
	}

	publicKey, err := getCognitoPublicKey(cognitoIssuer, kid)
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if !claims.VerifyIssuer(cognitoIssuer, true) {
			return nil, jwt.NewValidationError("invalid issuer claim", jwt.ValidationErrorIssuer)
		}

		clientID, ok := claims["client_id"].(string)
		if !ok || clientID != cognitoAppClientId {
			return nil, jwt.NewValidationError("invalid client_id claim", jwt.ValidationErrorClaimsInvalid)
		}
	} else {
		return nil, errors.New("invalid token claims")
	}

	return parsedToken, nil
}

func getCognitoPublicKey(cognitoIssuer, kid string) (interface{}, error) {
	if cached, found := publicKeyCache.Get("jwks"); found {
		jwks := cached.(JWKS)
		return findKeyByKid(jwks, kid)
	}

	resp, err := http.Get(fmt.Sprintf("%s/.well-known/jwks.json", cognitoIssuer))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var jwks JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, err
	}

	publicKeyCache.Set("jwks", jwks, cache.DefaultExpiration)

	return findKeyByKid(jwks, kid)
}

func findKeyByKid(jwks JWKS, kid string) (interface{}, error) {
	for _, key := range jwks.Keys {
		if key.Kid == kid {
			return jwt.ParseRSAPrivateKeyFromPEM([]byte(fmt.Sprintf(`
				-----BEGIN PUBLIC KEY-----
				%s
				-----END PUBLIC KEY-----
			`, key.N)))
		}
	}
	return nil, nil
}
