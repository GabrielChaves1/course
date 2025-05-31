import { ref } from 'vue'
import { useApiFetch } from '~/composables/useFetch'

type ValidationError = {
  field: string
  message: string
}

type ParsedError = {
  success: false
  message?: string
  errors?: ValidationError[]
}

function parseApiError(error: unknown): ParsedError {
  const defaultMsg = 'Erro inesperado. Tente novamente mais tarde.'
  const responseData = (error as any)?.response?._data
  const errors: ValidationError[] | undefined = responseData?.context?.errors
  const message: string | undefined = responseData?.detail ?? defaultMsg

  if (errors) return { success: false, errors }
  return { success: false, message }
}

export function useAuth() {
  const user = ref(null)

  const login = async (
    email: string,
    password: string
  ): Promise<{ success: boolean; errors?: ValidationError[]; message?: string }> => {
    try {
      await useApiFetch('/auth/sign-in', {
        method: 'POST',
        body: { email, password },
      })
      return { success: true }
    } catch (error) {
      return parseApiError(error)
    }
  }

  const register = async (
    email: string,
    password: string,
    confirmPassword: string,
    phoneNumber: string
  ): Promise<{ success: boolean; errors?: ValidationError[]; message?: string }> => {
    try {
      await useApiFetch('/auth/sign-up', {
        method: 'POST',
        body: { email, password, confirmPassword, phoneNumber },
      })
      return { success: true }
    } catch (error) {
      return parseApiError(error)
    }
  }

  const fetchUser = async () => {
    try {
      user.value = await useApiFetch('/api/dashboard')
    } catch (error) {
      console.error('Erro ao buscar usuário:', error)
      user.value = null
    }
  }

  const logout = async () => {
    try {
      await useApiFetch('/api/logout', { method: 'POST' })
      user.value = null
      navigateTo('/login')
    } catch (error) {
      console.error('Erro no logout:', error)
    }
  }

  const confirmEmail = async (code: string, email: string): Promise<{ success: boolean; message?: string }> => {
    try {
      await useApiFetch('/auth/confirm', {
        method: 'POST',
        body: { code, email },
      })
      return { success: true }
    } catch (error) {
      const parsedError = parseApiError(error)
      return { success: false, message: parsedError.message }
    }
  }

  const resendCode = async (email: string): Promise<{ success: boolean; message?: string }> => {
    try {
      await useApiFetch('/auth/resend-confirmation-code', {
        method: 'POST',
        body: { email },
      })
      return { success: true }
    } catch (error) {
      const parsedError = parseApiError(error)
      return { success: false, message: parsedError.message }
    }
  }

  return {
    user,
    login,
    register,
    fetchUser,
    logout,
    confirmEmail,
    resendCode,
  }
}

