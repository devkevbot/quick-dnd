import Vue from 'vue';
import Vuex from 'vuex';
import authentication from './modules/authentication';
import notifications from './modules/notifications';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    authentication,
    notifications,
  },
});
