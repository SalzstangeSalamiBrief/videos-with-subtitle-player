import { resolve } from 'path';
import { defineConfig } from 'vite';
import dts from 'vite-plugin-dts';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [dts({ tsconfigPath: './tsconfig.json' })],
  server: {
    port: 4200,
  },
  build: {
    lib: {
      entry: resolve(__dirname, 'src/index.ts'),
      formats: ['es'],
    },
    rollupOptions: {
      external: [],
      output: {
        entryFileNames: '[name].js',
      },
    },

    minify: true,
    cssMinify: true,
    sourcemap: true,
    emptyOutDir: true,
    target: 'esnext',
  },
});
