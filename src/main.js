import { createApp } from 'vue' 
import App from './App.vue'
import router from './router' // make sure to import your router this will allow you to run MyBody component through here 

createApp(App).use(router).mount('#app') // put .use(router) 
