import Vue from 'vue'
import App from './App'
import router from './router'

import VueSession from "./etc/session.js"
import VueResource from "vue-resource"

Vue.use(VueSession)
Vue.use(VueResource)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
