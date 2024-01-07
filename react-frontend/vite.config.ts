import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 4200,
  },
  build: {
    outDir: "../backend/public",
    minify: true,
    cssMinify: true,
  },
});
