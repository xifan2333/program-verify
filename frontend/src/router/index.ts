import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/Login.vue')
    },
    {
      path: '/',
      name: 'Layout',
      component: () => import('../views/Layout.vue'),
      children: [
        {
          path: '',
          name: 'Home',
          component: () => import('../views/Home.vue')
        },
        {
          path: 'products',
          name: 'Products',
          component: () => import('../views/Products.vue')
        },
        {
          path: 'licenses',
          name: 'Licenses',
          component: () => import('../views/Licenses.vue')
        },
      ]
    }
  ]
})

// 路由守卫
router.beforeEach((to, _, next) => {
  const token = localStorage.getItem('token')
  if (to.path !== '/login' && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router