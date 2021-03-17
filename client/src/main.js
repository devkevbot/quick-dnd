import Axios from 'axios';
import customAxios from '@/api/config';
import Vue from 'vue';
import Vuelidate from 'vuelidate';
import App from './App.vue';
import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css';
import '@mdi/font/css/materialdesignicons.css';

Vue.config.productionTip = false;
Axios.defaults.baseURL = customAxios.defaults.baseURL;
Vue.prototype.$http = Axios.create(customAxios);
Vue.use(Vuelidate);

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
