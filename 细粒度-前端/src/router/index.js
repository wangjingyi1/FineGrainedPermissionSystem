import Vue from "vue";
import Router from "vue-router";

// utils

// views
import Home from "../views/home/index";
import RegisterLogin from "../views/registerlogin/index";
import Register from "../views/register/index";
import Main from "../views/main/index";

Vue.use(Router);

const routes = [
  { path: "/", name: "Home", component: Home, redirect: "/home", meta: { requireAuth: false } },
  {
    path: "/home",
    name: "HomeIndex",
    component: Home,
    meta: { requireAuth: false }
  },
  {
    path: "/registerlogin",
    name: "RegisterLogin",
    component: RegisterLogin,
    meta: { requireAuth: false }
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
    meta: { requireAuth: false }
  },
  {
    path: "/main",
    name: "Main",
    component: Main,
    meta: { requireAuth: true }
  },
];
const router = new Router({
  routes
});

// router.beforeEach((to, from, next) => {
//   const needPageScroll = to.matched.some(record => record.meta.needPageScroll);
//   store.dispatch("setNeedPageScroll", needPageScroll);
//   if (to.matched.some(record => record.meta.requiresAuth)) {
//     const isLogin = getToken();
//     // this route requires auth, check if logged in
//     // if not, redirect to login page.
//     if (!isLogin) {
//       next({
//         path: "/login",
//         query: { redirect: to.fullPath }
//       });
//     } else {
//       next();
//     }
//   } else {
//     next(); // make sure to always call next()!
//   }
// });
export default router;
