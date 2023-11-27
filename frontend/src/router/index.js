import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import(/* webpackChunkName: "about" */ '../views/home.vue')
  },
  {
    path: '/question',
    name: 'question',
    component: () => import(/* webpackChunkName: "about" */ '../views/question.vue')
  },
  {
    path: '/ranking',
    name: 'ranking',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/ranking.vue')
  },
  {
    path: '/YWRtaW4/users',
    name: 'users',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/users.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})
// router.beforeEach((to,from,next) => {
//   var cookie = document.cookie;
// })


export default router
