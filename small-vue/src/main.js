import Vue from 'vue';
import App from './App.vue';
import router from './router';
import './mint-ui';
import './cube-ui';
import './global';
import './mycomponents';

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
