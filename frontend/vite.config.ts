import { reactRouter } from '@react-router/dev/vite'
import tailwindcss from '@tailwindcss/vite'
import { defineConfig } from 'vite'

export default defineConfig({
  plugins: [tailwindcss(), reactRouter()],

  define: {
    'import.meta.env.VITE_TITLE': '"OPEN POS"',
  },

  server: {
    port: 5555,
  },

  resolve: {
    tsconfigPaths: true,
    alias: {},
  },

  build: {
    outDir: 'dist/client',
  },
})
