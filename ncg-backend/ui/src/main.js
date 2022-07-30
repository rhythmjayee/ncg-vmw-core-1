import Vue from 'vue';
import BootstrapVue from 'bootstrap-vue';
import axios from 'axios';
import VueAxios from 'vue-axios';
import VueCodemirror from 'vue-codemirror';
import App from './App.vue';

import router from './router';
import store from './store';
import 'bootstrap/dist/css/bootstrap.min.css';

Vue.config.productionTip = false;
Vue.prototype.$endpoint = process.env.VUE_APP_DIAAS;
Vue.filter('formatId', (value) => value.slice(0, 8));
Vue.use(BootstrapVue);
Vue.use(VueCodemirror);
Vue.use(VueAxios, axios);
Vue.use(require('vue-moment'));

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount('#app');
