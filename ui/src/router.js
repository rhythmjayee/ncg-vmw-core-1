import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import store from './store';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      meta: { requiresAuth: true },
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('./views/Login.vue'),
      meta: { guest: true },
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('./views/Register.vue'),
      meta: { guest: true },
    },
    {
      path: '/build',
      name: 'build',
      component: () => import('./views/Build.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/functions',
      name: 'functions',
      component: () => import('./views/Functions.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/playground',
      name: 'functions',
      component: () => import('./views/Run.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/route',
      name: 'route',
      component: () => import('./views/Route.vue'),
      meta: { requiresAuth: true },
    },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (store.getters.isAuthenticated) {
      next();
      return;
    }
    next('/login');
  } else {
    next();
  }
});

router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.guest)) {
    if (store.getters.isAuthenticated) {
      next('/');
      return;
    }
    next();
  } else {
    next();
  }
});

export default router;
