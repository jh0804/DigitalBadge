import { createApp } from 'vue';
// import { createRouter, createWebHistory } from 'vue-router';
import router from './router'
import App from './App.vue';
import axios from 'axios';
import store from './store';

// const routes = [
//   { path: '/oauth2/authorization/google', redirect: 'https://accounts.google.com/o/oauth2/v2/auth/oauthchooseaccount?response_type=code&client_id=682858095425-cithjug9d5i265s3npr1m78jmhg95lmq.apps.googleusercontent.com&scope=profile%20email&state=EFKjA32UrTx63Bs5wRIDqkE72JZyS_CiFjtVtAqLCY8%3D&redirect_uri=http%3A%2F%2Flocalhost%3A8082%2Flogin%2Foauth2%2Fcode%2Fgoogle&service=lso&o2v=2&theme=glif&flowName=GeneralOAuthFlow' },
//   // 추가적인 라우트 설정을 필요에 따라 추가할 수 있습니다.
// ];

// const router = createRouter({
//   history: createWebHistory(),
//   routes
// });

const app = createApp(App);
// axios를 전역적으로 사용하도록 설정합니다.
app.config.globalProperties.axios = axios;

// 라우터와 스토어를 사용합니다.
app.use(router).use(store);

// 앱을 마운트합니다.
app.mount('#app');