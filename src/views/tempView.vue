<template>
  <div class="AndroidLarge1">
    <img class="pknu_logo" :src="require('@/assets/white_logo.png')">
    <div class="Title">부경대학교<br/>모바일 도서관</div>

    <!-- 로그인 버튼 수정 -->
    <a :href="getLoginUrl('onSuccess', 'onFailure')">
      <img class="google_login" :src="require('@/assets/sign_google.png')"/>
    </a>
  </div>
</template>

<script>
export default {
  methods: {
    getLoginUrl(onSuccessCallback, onFailureCallback) {
      const baseUrl = 'https://accounts.google.com/o/oauth2/v2/auth/oauthchooseaccount';
      const clientId = '682858095425-cithjug9d5i265s3npr1m78jmhg95lmq.apps.googleusercontent.com';
      const redirectUri = 'http://localhost:8080/login/oauth2/code/google';
      const onSuccessRedirect = 'http://localhost:8080/main'; // 로그인 성공 시 리디렉션할 주소
      // const onFailureRedirect = 'http://localhost:8082/login/failure';
      // console.log('onFailureRedirect:', onFailureRedirect);

      // 사용자 정의된 콜백 함수와 리디렉션 주소를 전달
      const url = `${baseUrl}?response_type=code&client_id=${clientId}&scope=profile email&state=EFKjA32UrTx63Bs5wRIDqkE72JZyS_CiFjtVtAqLCY8%3D&redirect_uri=${redirectUri}&service=lso&o2v=2&theme=glif&flowName=GeneralOAuthFlow&${onSuccessCallback}=${onSuccessRedirect}&${onFailureCallback}=${this.handleLoginFailure}`;
      return url;
    },
  }
};
</script>




<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.AndroidLarge1
{width: 360px;
height: 800px;
background: #173970;
}
.pknu_logo
{
  width: 185.31px;
  height: 250px;
  left: 87px;
  margin-top: 107px;
}
.Title
{
  width: 247px;
  height: 69px;
  margin-left: 57px;
  /* margin-top: 200px; */
  text-align: 'center';
  color: white;
  font-Size: 25px;
  /* fontFamily: 'Inter';
  fontWeight: '700';
  wordWrap: 'break-word'; */
}
</style>