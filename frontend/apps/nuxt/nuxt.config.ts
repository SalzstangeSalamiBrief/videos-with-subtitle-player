// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from '@tailwindcss/vite';

export default defineNuxtConfig({
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
  modules: ['@nuxt/eslint', '@nuxt/icon', '@nuxt/test-utils'],
  devServer: {
    port: 4200,
  },
  vite: {
    plugins: [tailwindcss()],
  },
  css: ['~/assets/app.css'],
});
