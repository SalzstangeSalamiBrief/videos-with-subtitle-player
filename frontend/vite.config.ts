import tailwindcss from '@tailwindcss/vite';
import { devtools } from '@tanstack/devtools-vite';
import { tanstackRouter } from '@tanstack/router-plugin/vite';
import viteReact from '@vitejs/plugin-react';
import { defineConfig } from 'vite';
import TsconfigPaths from 'vite-tsconfig-paths';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    devtools(),
    tanstackRouter({
      target: 'react',
      autoCodeSplitting: true,
    }),
    TsconfigPaths({ root: '.' }),
    viteReact(),
    tailwindcss(),
  ],
  server: {
    port: 4200,
  },
  build: {
    minify: true,
    cssMinify: true,
    target: 'esnext',
  },
});
