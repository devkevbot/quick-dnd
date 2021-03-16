import Vue from 'vue';
import VueRouter from 'vue-router';
import store from '@/store/index';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      layout: 'Main',
      protected: true,
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: {
      layout: 'Main',
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
