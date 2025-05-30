import { ref } from "vue";

export function useAuth() {
  const user = ref(null);

  const login = async (email: string, password: string) => {
    try {
      await $fetch("http://localhost:8080/auth/sign-in", {
        method: "POST",
        body: { email, password },
        credentials: "include",
      });
      return { success: true };
    } catch (error: any) {
      if(error?.response?.status === 400) {
        return {
          success: false,
          errors: error.response._data.context.errors,
        }
      }
      return {
        success: false,
        message: error.response._data.detail,
      };
    }
  };

  const register = async (email: string, password: string) => {
    try {
      await $fetch("http://localhost:8080/api/register", {
        method: "POST",
        body: { email, password, confirmPassword: password },
        credentials: "include",
      });
      return true;
    } catch (error) {
      console.error("Erro no cadastro:", error);
      return false;
    }
  };

  const fetchUser = async () => {
    try {
      user.value = await $fetch("http://localhost:8080/api/dashboard", {
        credentials: "include",
      });
    } catch (error) {
      console.error("Erro ao buscar usuário:", error);
      user.value = null;
    }
  };

  const logout = async () => {
    try {
      await $fetch("http://localhost:8080/api/logout", {
        method: "POST",
        credentials: "include",
      });
      user.value = null;
      navigateTo("/login");
    } catch (error) {
      console.error("Erro no logout:", error);
    }
  };

  return { user, login, register, fetchUser, logout };
}
