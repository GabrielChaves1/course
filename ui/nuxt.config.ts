// https://nuxt.com/docs/api/configuration/nuxt-config

import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  compatibilityDate: "2025-05-15",
  devtools: { enabled: true },
  modules: ["@nuxt/fonts", "@nuxt/icon"],
  css: ["~/assets/css/tailwind.css"],
  ssr: false,

  fonts: {
    families: [
      { name: "Poppins", provider: "google" },
      { name: "Calistoga", provider: "google" },
    ],
  },
  icon: {
    mode: "css",
    cssLayer: "base",
  },

  vite: {
    plugins: [
      tailwindcss(),
    ],
  },

  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL || "http://localhost:8080",
    },
  },
});
