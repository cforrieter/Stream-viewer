import Vue from "vue";
import Router from "vue-router";
import Login from "./views/Login.vue";


Vue.use(Router);

export default new Router({
  mode: 'history',
  base: "/ui/",
  routes: [
    {
      path: "/",
      name: "login",
      component: Login
    },
    {
      path: "/home",
      name: "home",
      component: () =>
        import(/* webpackChunkName: "home" */ "./views/Home.vue")
    },
    {
      path: "/stream/:videoId",
      name: "stream",
      props: true,
      component: () =>
        import(/* webpackChunkName: "home" */ "./views/Stream.vue")
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/About.vue")
    }
  ]
});
