<template>
  <div class="w-full sm:max-w-120 mx-auto p-14 h-screen flex flex-col items-center bg-white rounded-lg shadow-sm">
    <Icon name="mdi:email-check-outline" class="size-12 bg-stone-600 mb-5" />
    <Text as="h1" class="text-2xl font-extrabold mb-2">Confirme seu e-mail</Text>

    <Text as="p" variant="secondary" class="text-sm mb-5 text-center">
      Insira o código que enviamos para <strong>{{ form.email }}</strong>
    </Text>

    <p v-if="state.error" class="text-red-600 text-sm mb-4">
      {{ state.error }}
    </p>

    <form class="w-full" @submit.prevent="handleConfirm">
      <VerificationCodeInput class="mb-6" v-model="form.code" :hasError="state.error.length > 0" />

      <Button :disabled="state.isLoading" type="submit" class="w-full rounded-md py-1 mt-2">
        <template v-if="state.isLoading">
          <Icon name="mdi:loading" class="animate-spin size-5" />
          Verificando...
        </template>
        <template v-else>
          Confirmar
        </template>
      </Button>
    </form>

    <button class="text-indigo-500 hover:underline mt-4 text-sm" :disabled="state.isResending" @click="handleResend">
      <template v-if="state.isResending">Reenviando...</template>
      <template v-else>Reenviar código</template>
    </button>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { useRoute } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
import Button from '~/components/Button.vue'
import Text from '~/components/Text.vue'
import VerificationCodeInput from '~/components/VerificationCodeInput.vue'

definePageMeta({ layout: 'default' })

const route = useRoute()
const { confirmEmail, resendCode } = useAuth()

const form = reactive({
  email: decodeURIComponent((route.query.email as string) || ''),
  code: '',
})

const state = reactive({
  isLoading: false,
  isResending: false,
  error: '',
})

async function handleConfirm() {
  state.error = ''
  state.isLoading = true

  const result = await confirmEmail(form.code, form.email)
  state.isLoading = false

  if (result.success) {
    navigateTo('/login')
  } else {
    state.error = result.message || 'Erro ao confirmar e-mail.'
    redirectToLogin()
  }
}

async function handleResend() {
  state.error = ''
  state.isResending = true

  const result = await resendCode(form.email)
  state.isResending = false

  if (!result.success) {
    state.error = result.message || 'Erro ao reenviar o código.'
    redirectToLogin()
  }
}

function redirectToLogin() {
  setTimeout(() => {
    navigateTo('/login')
  }, 600)
}
</script>
