import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#083D77',
        secondary: '#DDDDDD',
        accent: '#FFFFFF',
        error: '#CA3C25',
        info: '#5998C5',
        success: '#4C9F70',
        warning: '#ED7D3A',
      },
    },
  },
});
