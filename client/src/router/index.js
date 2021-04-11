import Vue from 'vue';
import VueRouter from 'vue-router';
import store from '@/store/index';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    redirect: '/dashboard',
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Registration.vue'),
    meta: {
      layout: 'Guest',
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: {
      layout: 'Guest',
    },
  },
  {
    path: '/campaign',
    name: 'Campaign',
    component: () => import('../views/Campaign.vue'),
    meta: {
      layout: 'Main',
      protected: true,
    },
  },
  {
    path: '/characters',
    name: 'Characters',
    component: () => import('../views/Characters.vue'),
    meta: {
      layout: 'Main',
      protected: true,
    },
  },
  {

    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('../views/Dashboard.vue'),
    meta: {
      layout: 'Main',
      protected: true,
    },
  },
  {

    path: '/account',
    name: 'Account',
    component: () => import('../views/Account.vue'),
    meta: {
      layout: 'Main',
      protected: true,
    },
  },
  {
    path: '/spells',
    name: 'Spells',
    component: () => import('../views/Spells.vue'),
    meta: {
      layout: 'Main',
      protected: true,
    },
  },
];

const router = new VueRouter({
  routes,
});

router.beforeEach(async (to, _, next) => {
  const isProtectedRoute = to.matched.some((record) => record.meta.protected);
  const isPlayerLoggedIn = store.getters['authentication/isPlayerLoggedIn'];
  const playerCanAccess = !isProtectedRoute || (isProtectedRoute && isPlayerLoggedIn);

  if (playerCanAccess) next();
  else next('/login');
});

export default router;
