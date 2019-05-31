import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'

import AWS from 'aws-sdk'

Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(AWS)

new Vue({
  render: h => h(App),
}).$mount('#app')
