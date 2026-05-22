import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/',
    component: () => import('../views/Home.vue'),
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: () => import('../views/Dashboard.vue')
      },
      {
        path: 'inventory',
        name: 'Inventory',
        component: () => import('../views/Inventory.vue')
      },
      {
        path: 'receipt',
        name: 'Receipt',
        component: () => import('../views/Receipt.vue')
      },
      {
        path: 'iqc',
        name: 'IQC',
        component: () => import('../views/IQC.vue')
      },
      {
        path: 'inbound',
        name: 'Inbound',
        component: () => import('../views/Inbound.vue')
      },
      {
        path: 'outbound',
        name: 'Outbound',
        component: () => import('../views/Outbound.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.name !== 'Login' && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router
