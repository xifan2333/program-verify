import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import '@unocss/reset/tailwind.css'
import 'virtual:uno.css'
import toast from './plugins/toast'


const app = createApp(App)
app.use(router)
app.use(toast)
app.mount('#app')
