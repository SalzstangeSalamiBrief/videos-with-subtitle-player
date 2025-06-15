import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';
import TsconfigPaths from 'vite-tsconfig-paths';

export default defineConfig({
  plugins: [tailwindcss(), sveltekit(), TsconfigPaths({ root: '.' })],
  server: {
    port: 4200,
  },
  build: {
    minify: true,
    cssMinify: true,
    target: 'esnext',
  },
});
