<template>
  <div
    class="w-full max-w-96 mx-auto mt-20 p-6 flex flex-col items-center bg-white rounded-lg shadow-sm"
  >
    <Icon name="mdi:longitude" class="size-12 bg-stone-600 mb-5" />
    <Text as="h1" class="text-2xl font-extrabold mb-2"> Acesse sua conta </Text>
    <Text as="p" variant="secondary" class="text-sm mb-5"
      >Não tem uma conta?
      <NuxtLink to="/register" class="text-indigo-500 hover:underline"
        >Cadastre-se</NuxtLink
      ></Text
    >

    <p v-if="error" class="text-red-600 text-sm mb-4">
      {{ error }}
    </p>

    <form class="w-full" @submit.prevent="handleLogin">
      <FormInput
        id="email"
        label="E-mail"
        type="email"
        v-model="email"
        :error="validationErrs.email"
        placeholder="Digite seu e-mail"
      />
      <FormInput
        id="password"
        label="Senha"
        type="password"
        v-model="password"
        :error="validationErrs.password"
        placeholder="Digite sua senha"
      />
      <Button type="submit" class="w-full rounded-md py-1"> Entrar </Button>

      <div class="flex my-3 items-center gap-3">
        <hr class="flex-grow border-t border-gray-300" />
        <Text as="span" class="text-gray-600">ou</Text>
        <hr class="flex-grow border-t border-gray-300" />
      </div>
    </form>
  </div>
</template>

<script setup>
import Button from "~/components/Button.vue";
import FormInput from "~/components/FormInput.vue";
import Text from "~/components/Text.vue";
import { useAuth } from "~/composables/useAuth";

definePageMeta({
  layout: "default",
});

const { login } = useAuth();
const email = ref("");
const password = ref("");
const validationErrs = ref({});
const error = ref("");

const handleLogin = async () => {

  const result = await login(email.value, password.value);
  if (result.success) {
    validationErrs.value = {};
    alert("Login realizado com sucesso!");
    navigateTo("/dashboard");
  } else if (result.errors) {
    for (const err of result.errors) {
      validationErrs.value[err.field] = err.message;
    }
  } else {
    error.value = result.message || "E-mail ou senha inválidos";
  }
};
</script>
