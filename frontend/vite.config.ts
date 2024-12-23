// frontend/vite.config.ts
import { defineConfig } from 'vite'
import { sveltekit } from '@sveltejs/kit/vite'

export default defineConfig({
  plugins: [sveltekit()],
  resolve: {
    alias: {
      $src: '/src', // Example alias configuration
    },
  },
  server: {
    port: 5170,
    host: true, // This enables network access
    watch: {
      usePolling: true // This helps with hot reload in Docker
    }
  }
})
