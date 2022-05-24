import { createRouter, createWebHistory } from "vue-router";
import MyBody from './../components/MyBody.vue'
import MyLogin from '../components/MyLogin.vue'


const routes = [
  {
    path: '/',
    name: 'Home',
    component: MyBody,
  },
  {
    path: '/login',
    name: 'Login',
    component: MyLogin,
  }
]

const router = createRouter({history: createWebHistory(), routes})
export default router