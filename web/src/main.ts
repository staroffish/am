import { createApp } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'
import axios from 'axios'
import App from './App.vue'

axios.interceptors.response.use(
  r => r,
  err => {
    const msg = err.response?.data?.error || err.message || '请求失败'
    alert('后端错误: ' + msg)
    return Promise.reject(err)
  }
)

const routes = [
  { path: '/', component: () => import('./views/Home.vue') },
  { path: '/collection', component: () => import('./views/Home.vue') },
  { path: '/anime/:id', component: () => import('./views/AnimeDetail.vue') },
  { path: '/anime/:id/edit', component: () => import('./views/AnimeEdit.vue') },
  { path: '/anime/new', component: () => import('./views/AnimeEdit.vue') },
  { path: '/tasks', component: () => import('./views/Tasks.vue') },
  { path: '/tasks/new', component: () => import('./views/TaskEdit.vue') },
  { path: '/tasks/:id/edit', component: () => import('./views/TaskEdit.vue') },
  { path: '/downloads', component: () => import('./views/Downloads.vue') },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

const app = createApp(App)
app.use(router)
app.mount('#app')
