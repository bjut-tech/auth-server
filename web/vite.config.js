import { join } from 'path'

import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  appType: 'mpa',
  build: {
    rollupOptions: {
      input: {
        main: join(__dirname, 'index.html'),
        login: join(__dirname, 'login.html')
      }
    }
  },
  plugins: [],
  server: {
    proxy: {
      '^/authorize$': 'http://localhost:8080',
      '^/login$': 'http://localhost:8080',
      '^/logout$': 'http://localhost:8080',
      '^/userinfo': 'http://localhost:8080'
    }
  }
})
