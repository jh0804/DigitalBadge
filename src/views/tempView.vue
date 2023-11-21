<!-- 로그인 테스트화면 -->

<template>
  <div class="AndroidLarge1">
    <img class="pknu_logo" :src="require('@/assets/white_logo.png')">
    <div class="Title">부경대학교<br/>모바일 도서관</div>

    <!-- 로그인 버튼 수정 -->
    <a :href="getLoginUrl()" @click.prevent="handleLogin">
      <img class="google_login" :src="require('@/assets/sign_google.png')"/>
    </a>
    <router-link to="/main" style=" text-align: center; top: 565px; left: 100px;position: absolute;" > 기존 아이디로 시작하기 </router-link>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  methods: {
    getLoginUrl() {
      const clientId = '682858095425-cithjug9d5i265s3npr1m78jmhg95lmq.apps.googleusercontent.com';
      const redirectUri = 'http://localhost:8082/login/oauth2/code/google';
      const baseUrl = 'http://localhost:8082/oauth2/authorization/google';
      const onSuccessRedirect = 'http://localhost:8080/main';
      const url = `${baseUrl}?response_type=code&client_id=${clientId}&scope=profile email&state=EFKjA32UrTx63Bs5wRIDqkE72JZyS_CiFjtVtAqLCY8%3D&redirect_uri=${redirectUri}&service=lso&o2v=2&theme=glif&flowName=GeneralOAuthFlow&onSuccess=${onSuccessRedirect}`;
      // return url;
      this.showSocialLoginPopup(url)
    },
    handleLogin() {
      // 로그인 버튼 클릭 시 수행할 작업
      window.location.href = this.getLoginUrl();
    },
    onSuccess() {
      // 로그인 성공 시 수행할 작업

      // URL에서 코드 파싱
      const urlParams = new URLSearchParams(window.location.search);
      const authorizationCode = urlParams.get('code');

      axios.post('http://localhost:8082/oauth2/authorization/google', {
        code: authorizationCode,
      })
        .then(response => {
          // 백엔드로부터 받은 추가 정보를 처리
          const userData = response.data;
          console.log(userData);
        })
        .catch(error => {
          console.error('백엔드와 통신 중 오류 발생:', error);
        });
    },
  },
};
</script>
