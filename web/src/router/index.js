import { createRouter, createWebHashHistory } from 'vue-router'


const router = createRouter({
  history: createWebHashHistory(import.meta.env.VITE_CREATEWEBHISTORY_URL),
  routes: [
    {
      path: '/',
      name: 'files',
      component: () => import('../pages/files/index.vue')
    },
  ]
})

export default router
