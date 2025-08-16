import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegisterAuth from '@/views/auth/RegisterAuth.vue'
import LoginAuth from '@/views/auth/LoginAuth.vue'
import DashboardAdmin from '@/views/admin/DashboardAdmin.vue'
import UnauthorizedError from '@/views/error/UnauthorizedError.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: { 
        showNav: true
      }
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue'),
      meta: {
        showNav:true
      }
    },
    {
      path: '/project',
      name: 'project',
      component: () => import('../views/ProjectView.vue'),
      meta: {
        showNav:true
      }
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterAuth,
      meta: {
        showNav:false
      },
    },
    {
      path: '/login',
      name: 'login',
      component: LoginAuth,
      meta: {
        showNav:false
      },
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardAdmin,
      meta: {
        showNav:false
      },
    },
    {
      path: '/unauthorized',
      name: 'unauthorized',
      component: UnauthorizedError,
      meta: {
        showNav:false
      },
    },
  ],
})

export default router
