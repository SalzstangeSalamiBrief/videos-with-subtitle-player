import { TanStackRouterVite } from '@tanstack/router-vite-plugin';
import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';
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
