import { createRouter, createWebHistory } from 'vue-router';
import GoogleLogin from '../components/GoogleLogin.vue';

const routes = [
  { path: '/google-login', component: GoogleLogin },
  // 다른 경로 및 컴포넌트를 추가할 수 있습니다.
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
