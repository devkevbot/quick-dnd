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
    state.player.username = username;
  },
  onLogout(state) {
    state.token = '';
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
    axios.defaults.headers.common.Authorization = resp.data.data.token;
    commit('onLoginSuccess', { token: resp.data.data.token, username: resp.data.data.username });
  },
};

export default {
  namespaced: true,
  state: states,
  getters,
  mutations,
  actions,
};
