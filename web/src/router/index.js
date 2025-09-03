import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '../store'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: {
      title: '登录',
      requiresAuth: false
    }
  },
  {
    path: '/',
    redirect: '/dashboard'
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('../views/Layout.vue'),
    meta: {
      title: '后台管理',
      requiresAuth: true
    },
    children: [
      {
        path: '',
        redirect: '/articles'
      },
      {
        path: '/articles',
        name: 'Articles',
        component: () => import('../views/Articles.vue'),
        meta: {
          title: '文章列表',
          requiresAuth: true
        }
      },
      {
        path: '/write',
        name: 'WriteArticle',
        component: () => import('../views/WriteArticle.vue'),
        meta: {
          title: '写文章',
          requiresAuth: true
        }
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

/**
 * 路由守卫 - 检查用户登录状态
 */
router.beforeEach((to, from, next) => {
  // 设置页面标题
  if (to.meta.title) {
    document.title = `${to.meta.title} - GoBlog 后台管理系统`
  }
  
  // 检查路由是否需要登录权限
  if (to.meta.requiresAuth) {
    // 检查用户是否已登录
    const isLoggedIn = store.getters.getLoginStatus
    const token = localStorage.getItem('token')
    
    if (!isLoggedIn || !token) {
      // 未登录，跳转到登录页
      next({
        path: '/login',
        query: { redirect: to.fullPath } // 保存原始路径，登录后可以跳转回来
      })
    } else {
      // 已登录，允许访问
      next()
    }
  } else {
    // 不需要登录权限的路由，直接允许访问
    if (to.path === '/login' && store.getters.getLoginStatus) {
      // 如果已登录用户访问登录页，重定向到首页
      next('/dashboard')
    } else {
      next()
    }
  }
})

/**
 * 路由后置守卫 - 设置当前激活菜单
 */
router.afterEach((to) => {
  // 更新store中的当前激活菜单
  if (to.path !== '/login') {
    store.dispatch('setActiveMenu', to.path)
  }
})

export default router