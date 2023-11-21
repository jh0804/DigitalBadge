// index.js

import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import UpdateView from '../views/updateView.vue';
import FirstView from '../views/FirstView.vue';
import SecondView from '../views/SecondView.vue';


const routes = [
  {
    path: '/',
    name: 'Intro',
    component: SecondView,
  },
  {
    path: '/first',
    name: 'first',
    component: FirstView,
  },
  {
    //있어빌리티 페이지
    path: '/login',
    name: 'login',
    component: HomeView,
  },
  {
    path: '/student/main',
    name: '/student/main',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')

  },
  {
    path: '/activity/report',
    name: '/activity/report',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "bogoseo" */ '../views/bogoseoView.vue')
  },
  {
    path: '/Badge_info/:id',
    name: '/Badge_info',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "bogoseo" */ '../views/Badge_infoView.vue')
  },
  {
    //정보수정 페이지임
    path: '/main',
    name: 'Main',
    component: UpdateView,
  },
  {
    path: '/admin/main',
    name: '/admin/main',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/ad_AboutView.vue')
  },
  {
    path: '/admin/report/:id',
    name: '/admin/report',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/ad_reportView.vue')
  },
  {
    path: '/verify',
    name: '/verfication',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/verifyView.vue')
  },
  {
    path: '/temp',
    name: '/temp',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/tempView.vue')
  },

]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
