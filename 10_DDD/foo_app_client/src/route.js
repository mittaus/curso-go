import Vue from 'vue'
import VueRouter from 'vue-router'
import Register from './components/Auth/Register'
import Login from './components/Auth/Login'
import AllFood from './components/Food/AllFood'
import SingleFood from './components/Food/SingleFood'
import EditFood from './components/Food/EditFood'
import CreateFood from './components/Food/CreateFood'

Vue.use(VueRouter)

const router = new VueRouter({
  
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'all_food',
      component: AllFood
    },
    {
      path: '/single_food/:id',
      name: 'single_food',
      component: SingleFood
    },
    {
      path: '/edit/single_food/:id',
      name: 'edit_food',
      component: EditFood,
      meta: { requiresAuth: true }
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/register',
      name: 'register',
      component: Register
    },
    {
      path: '/food',
      name: 'createFood',
      component: CreateFood,
      meta: { requiresAuth: true }
    },
  ]
});

router.beforeEach((to, from, next) => {
  if(to.path.startsWith("/login") || to.path.startsWith("/register")) {
    if (window.localStorage.getItem("access_token")) {
      next({
        path: "/"
      })
    }
  }
  if (to.meta.requiresAuth) {
    if (!window.localStorage.getItem("access_token")) {
        next({
            path: "/login",
            query: {
                redirect: to.path
            }
        });
    } else {
        next();
    }
  } else {
    next();
  } 
})

export default router;