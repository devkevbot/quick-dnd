import axios from 'axios';

const states = {
  token: '',
  player: {
    username: '',
  },
};

const getters = {
  isPlayerLoggedIn: (state) => !!state.token,
  getToken: (state) => state.token,
  getPlayer: (state) => state.player,
};

const mutations = {
  onLoginSuccess(state, { token, username }) {
    state.token = token;
    state.player = { username };
  },
  onLogout(state) {
    state.token = null;
    state.player = null;
  },
};

const actions = {
  /**
   * @param {Object} commit - Used to commit changes to the state.
   * @param {String} username - The username used to log in.
   * @param {String} password - The password used to log in.
   */
  async loginPlayer({ commit }, { username, password }) {
    const requestURI = 'login';
    const data = { username, password };
    const method = 'POST';
    const resp = await axios({ url: requestURI, data, method });
    commit('onLoginSuccess', { token: resp.data.data.token, username: resp.data.data.username });
  },
  /**
   * Deletes stored authentication data for the current player.
   * @param {Object} commit - Used to commit changes to the state.
   */
  logoutPlayer({ commit }) {
    commit('onLogout');
  },
};

export default {
  namespaced: true,
  state: states,
  getters,
  mutations,
  actions,
};
