import { createApp } from 'vue'
import App from './App.vue'
import router from './router'  //설치한 라우터 가져오기
// import GAuth from 'vue-google-oauth2'
// import axios from 'axios'
// import gauthOption from 'vue-google-oauth2'

// .use(axios).use(GAuth, gauthOption) 쓸 때 적어야 함
createApp(App).use(router).mount('#app') //사용 선언  axios 쓸때 여기에 뭐 선언해줘야 하는 듯
