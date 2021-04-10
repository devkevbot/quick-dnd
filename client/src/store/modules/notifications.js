/* Stores data to display notifications. */

const states = {
  message: '',
  color: '',
  timeout: -1,
};

const mutations = {
  updateState(state, { message, color, timeout }) {
    state.message = message;
    state.color = color;
    state.timeout = timeout;
  },
};

const actions = {
  display({ commit }, data) {
    commit('updateState', data);
  },
};

export default {
  namespaced: true,
  state: states,
  mutations,
  actions,
};
