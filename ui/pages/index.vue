<script setup lang="ts">
import { Icon } from "#components";
import Button from "~/components/Button.vue";
import Text from "~/components/Text.vue";

definePageMeta({
  layout: "default",
});
</script>

<template>
  <main class="w-full h-dvh flex flex-col">
    <header
      class="flex items-center justify-between container mx-auto py-8 px-8 md:px-8"
    >
      <Text as="h1">Curso</Text>
      <Button size="sm" @click="navigateTo('/auth/login')">
        <Icon name="mdi:location-enter" class="size-5"/>
        Entrar
      </Button>
    </header>

    <section class="flex flex-col items-center pt-40">
      <Text
        as="span"
        class="text-xs font-bold border border-dashed border-stone-300 flex items-center gap-2 p-1 px-3 rounded-2xl mb-5 text-stone-600"
      >
        <Icon name="mdi:calendar-blank-outline" class="size-4 bg-stone-600" />
        Previsão de lançamento: 2025
      </Text>
      <Text as="h1" class="text-6xl font-calistoga"
        >Seu próximo passo começa aqui</Text
      >
      <Text as="p" variant="secondary" class="mt-3"
        >Aprenda com especialistas, pratique com projetos reais e evolua no seu
        tempo.</Text
      >
      <Button class="mt-8" @click="navigateTo('/auth/register')">
        <Icon
          name="mdi:arrow-top-right-thin-circle-outline"
          class="size-6 bg-white"
        />
        Cadastre-se
      </Button>
    </section>
  </main>
</template>
