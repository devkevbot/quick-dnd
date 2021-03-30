import axios from 'axios';
import defaultAPIPath from '@/api/config';
import Vue from 'vue';
import Vuelidate from 'vuelidate';
import App from './App.vue';
import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css';
import '@mdi/font/css/materialdesignicons.css';

Vue.config.productionTip = false;
axios.defaults.baseURL = defaultAPIPath;
axios.interceptors.request.use(
  (config) => {
    const authConfig = config;
    authConfig.headers.Authorization = `Bearer ${store.getters['authentication/getToken']}`;
    return authConfig;
  },
  (error) => {
    Promise.reject(error);
  },
);

Vue.prototype.$http = axios;
Vue.use(Vuelidate);

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
