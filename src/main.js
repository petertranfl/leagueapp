import Vue from 'vue'
import router from './router'
import App from './App.vue'
import Vuex from 'vuex'
import store from './store/index'
import 'es6-promise/auto'

import AWS from 'aws-sdk'

Vue.config.productionTip = false
Vue.use(AWS)
Vue.use(Vuex)

new Vue({
  render: h => h(App),
  store,
  router,
}).$mount('#app')
