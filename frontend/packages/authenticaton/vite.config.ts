import { createRequire } from 'node:module';
import { resolve } from 'path';
import { defineConfig } from 'vite';
import dts from 'vite-plugin-dts';
const require = createRequire(import.meta.url);
const pkg = require('./package.json') as {
  dependencies?: Record<string, string>;
  peerDependencies?: Record<string, string>;
};

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
      external: [
        ...Object.keys(pkg.dependencies ?? {}),
        ...Object.keys(pkg.peerDependencies ?? {}),
      ],
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
