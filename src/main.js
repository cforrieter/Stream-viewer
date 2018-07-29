import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import VueGAPI from 'vue-gapi'
import VueYoutube from 'vue-youtube'

Vue.config.productionTip = false;


Vue.use(VueYoutube)

const apiConfig = {
  apiKey: 'API_KEY',
  clientId: 'CLIENT_ID.apps.googleusercontent.com',
  discoveryDocs: ['https://www.googleapis.com/discovery/v1/apis/youtube/v3/rest'],
  scope: 'https://www.googleapis.com/auth/youtube'
  // see all available scopes here: https://developers.google.com/identity/protocols/googlescopes'
}

Vue.use(VueGAPI, apiConfig)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
