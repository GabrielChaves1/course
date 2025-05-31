<template>
  <div class="w-full sm:max-w-120 mx-auto p-14 h-screen flex flex-col items-center bg-white rounded-lg shadow-sm">
    <Icon name="mdi:longitude" class="size-12 bg-stone-600 mb-5" />
    <Text as="h1" class="text-2xl font-extrabold mb-2"> Acesse sua conta </Text>
    <Text as="p" variant="secondary" class="text-sm mb-5">
      Não tem uma conta?
      <NuxtLink to="/register" class="text-indigo-500 hover:underline">Cadastre-se</NuxtLink>
    </Text>

    <p v-if="state.error" class="text-red-600 text-sm mb-4">
      {{ state.error }}
    </p>

    <form class="w-full flex flex-col gap-6" @submit.prevent="handleLogin">
      <FormInput id="email" label="E-mail" type="email" v-model="form.email" :error="state.validationErrs.email"
        placeholder="Digite seu e-mail" />
      <FormInput id="password" label="Senha" type="password" v-model="form.password"
        :error="state.validationErrs.password" placeholder="Digite sua senha" />

      <Button :disabled="state.isLoading" type="submit" class="w-full rounded-md py-1">
        <template v-if="state.isLoading">
          <Icon name="mdi:loading" class="animate-spin size-5" />
          Carregando...
        </template>
        <template v-else>
          Entrar
        </template>
      </Button>
    </form>
    <p class="mt-2 text-sm">Esqueceu sua senha?
      <NuxtLink to="/reset-password" class="text-indigo-500 hover:underline">Clique aqui</NuxtLink>
    </p>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import Button from '~/components/Button.vue'
import FormInput from '~/components/FormInput.vue'
import Text from '~/components/Text.vue'
import { useAuth } from '~/composables/useAuth'

definePageMeta({ layout: 'default' })

const { login } = useAuth()

const form = reactive({
  email: '',
  password: '',
})

const state = reactive({
  isLoading: false,
  error: '',
  validationErrs: {} as Record<string, string>,
})

function setValidationErrors(errors: { field: string; message: string }[]) {
  for (const err of errors) {
    state.validationErrs[err.field] = err.message
  }
}

const handleLogin = async () => {
  state.error = ''
  state.validationErrs = {}

  state.isLoading = true
  const result = await login(form.email, form.password)
  state.isLoading = false

  if (result.success) {
    navigateTo('/dashboard')
  } else if (result.errors) {
    setValidationErrors(result.errors)
  } else {
    state.error = result.message || 'E-mail ou senha inválidos'
  }
}
</script>
