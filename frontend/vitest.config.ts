import react from '@vitejs/plugin-react';
import TsconfigPaths from 'vite-tsconfig-paths';
import { defineConfig } from 'vitest/config';

export default defineConfig({
  plugins: [react(), TsconfigPaths()],
  test: {
    include: ['./tests/**/*.{test,spec}.{ts,tsx}'],
    environment: 'jsdom',
    globals: true,
    setupFiles: './tests/setup.ts',
  },
});
