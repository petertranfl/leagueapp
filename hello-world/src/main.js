import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import Vuex from 'vuex'
import 'es6-promise/auto'

import AWS from 'aws-sdk'

Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(AWS)
Vue.use(Vuex)

new Vue({
  render: h => h(App),
}).$mount('#app')
