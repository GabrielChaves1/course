<template>
  <div class="w-full sm:max-w-120 mx-auto p-14 h-screen flex flex-col items-center bg-white rounded-lg shadow-sm">
    <Icon name="mdi:longitude" class="size-12 bg-stone-600 mb-5" />
    <Text as="h1" class="text-2xl font-extrabold mb-2"> Crie sua conta </Text>
    <Text as="p" variant="secondary" class="text-sm mb-5">
      Já tem uma conta?
      <NuxtLink to="/login" class="text-indigo-500 hover:underline">Entrar</NuxtLink>
    </Text>

    <p v-if="state.error" class="text-red-600 text-sm mb-4">
      {{ state.error }}
    </p>

    <form class="w-full flex flex-col gap-6" @submit.prevent="handleRegister">
      <FormInput id="email" label="E-mail" type="email" v-model="form.email" :error="state.validationErrs.email"
        placeholder="Digite seu e-mail" />
      <FormInput id="password" label="Senha" type="password" v-model="form.password"
        :error="state.validationErrs.password" placeholder="Crie uma senha" />
      <FormInput id="confirmPassword" label="Confirme a Senha" type="password" v-model="form.confirmPassword"
        :error="state.validationErrs.confirmPassword" placeholder="Repita sua senha" />
      <PhoneNumberInput id="phoneNumber" label="Número de telefone" type="tel" v-model="form.phoneNumber"
        :error="state.validationErrs.phoneNumber" placeholder="Digite seu número" />

      <Button :disabled="state.isLoading" type="submit" class="w-full rounded-md py-1">
        <template v-if="state.isLoading">
          <Icon name="mdi:loading" class="animate-spin size-5" />
          Carregando...
        </template>
        <template v-else>
          Cadastrar
        </template>
      </Button>
      <p class="text-sm text-center text-gray-500">
        Ao se cadastrar, você concorda com nossos
        <NuxtLink to="#" class="text-indigo-500 hover:underline">Termos de Serviço</NuxtLink> e
        <NuxtLink to="#" class="text-indigo-500 hover:underline">Política de Privacidade</NuxtLink>.
      </p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import Button from '~/components/Button.vue'
import FormInput from '~/components/FormInput.vue'
import PhoneNumberInput from '~/components/PhoneNumberInput.vue'
import Text from '~/components/Text.vue'
import { useAuth } from '~/composables/useAuth'

definePageMeta({ layout: 'default' })

const { register } = useAuth()

const form = reactive({
  email: '',
  password: '',
  confirmPassword: '',
  phoneNumber: '',
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

const handleRegister = async () => {
  state.error = ''
  state.validationErrs = {}

  if (form.password !== form.confirmPassword) {
    state.validationErrs.confirmPassword = 'As senhas não coincidem'
    return
  }

  state.isLoading = true
  const result = await register(form.email, form.password, form.confirmPassword, form.phoneNumber)
  state.isLoading = false

  if (result.success) {
    navigateTo('/confirm?email=' + encodeURIComponent(form.email))
  } else if (result.errors) {
    setValidationErrors(result.errors)
  } else {
    state.error = result.message || 'Erro ao registrar'
  }
}
</script>
