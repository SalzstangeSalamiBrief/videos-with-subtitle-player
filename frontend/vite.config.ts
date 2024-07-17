import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { TanStackRouterVite } from '@tanstack/router-vite-plugin';
import TsconfigPaths from 'vite-tsconfig-paths';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), TanStackRouterVite(), TsconfigPaths({ root: '.' })],
  server: {
    port: 4200,
  },
  build: {
    outDir: '../backend/public',
    minify: true,
    cssMinify: true,
  },
});
