<!-- 로그인 화면 -->

<template>
  <div class="AndroidLarge1">
    <button @click="modalOpen = true" style="height: 35px; width: 45px; position: absolute; background-color:#173970; top:110px; left: 295px; border: none;">
    <img class="app_info_icon" :src="require('@/assets/Infocircle.png')" style="height: 100%; width: 100%;">
    </button>
    <div class="black-bg" v-if="modalOpen === true">
          <div >
            <h3>▼ APP info ▼</h3>
            <img class="app_info" :src="require('@/assets/App_info.jpg')" style=";">
            <h3>▼ github ▼</h3>
            <img class="app_info_icon" :src="require('@/assets/qrcode.png')" >
            <button @click="modalOpen = false" style="background-color: none; height: 40px; width: 40px; position: absolute; top:30px; right: 60px;" >
            <img class="cross" :src="require('@/assets/Cross.png')" style="height: 100%; width: 100%;">
        </button>
          </div>          
      </div>
    <img class="pknu_logo" :src="require('@/assets/white_logo.png')">
    <div class="Title">부경대학교<br/>모바일 도서관</div>

    <!-- 로그인 버튼 수정 -->
    <a :href="getLoginUrl('onSuccess', 'onFailure')">
      <div style="height: 50px; width: 220px; left:70px; background-color: #fffafa; position: relative; border-radius: 30px;"></div>
      <img class="google_login" :src="require('@/assets/sign_google.png') "/>
    </a>
    <router-link to="/main" style=" text-align: center; top: 582px; left: 100px; position: absolute; color:white;" > 기존 아이디로 시작하기▶ </router-link>
  </div>
</template>

<script>
export default {
  data(){
    return{
    modalOpen: false,
    };
  },
  methods: {
    getLoginUrl(onSuccessCallback, onFailureCallback) {
      const baseUrl = 'https://accounts.google.com/o/oauth2/v2/auth/oauthchooseaccount';
      const clientId = '682858095425-cithjug9d5i265s3npr1m78jmhg95lmq.apps.googleusercontent.com';
      const redirectUri = 'http://localhost:8080/login/oauth2/code/google';
      const onSuccessRedirect = 'http://localhost:8080/main'; // 로그인 성공 시 리디렉션할 주소
      const url = `${baseUrl}?response_type=code&client_id=${clientId}&scope=profile email&state=EFKjA32UrTx63Bs5wRIDqkE72JZyS_CiFjtVtAqLCY8%3D&redirect_uri=${redirectUri}&service=lso&o2v=2&theme=glif&flowName=GeneralOAuthFlow&${onSuccessCallback}=${onSuccessRedirect}&${onFailureCallback}=${this.handleLoginFailure}`;
      return url;
    },
    onSuccess() {
      // 로그인 성공 시 수행할 작업
      const onSuccessRedirect = 'http://localhost:8080/main'; // 로그인 성공 시 리디렉션할 주소
      window.location.href = onSuccessRedirect;
    },
  },
};
</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #507BBC;
  margin-top: 60px;
}
@font-face{
  font-family: 'nanumgothic';
  src: url('@/assets/fonts/NanumGothic.ttf');
}
.AndroidLarge1
{width: 360px;
height: 800px;
background: #173970;
}
.black-bg {
  display: flex;
  align-items: center;
  width: 815px;
  height: 1200px;
  top: 0px;
  background-color: rgba(0, 0, 0, 0.432);
  position: relative;
  padding: 20px;
}
.pknu_logo
{
  width: 185.31px;
  height: 250px;
  left: 87px;
  margin-top: 115px;
}
.Title
{
  width: 247px;
  height: 69px;
  margin-left: 57px;
  margin-bottom: 20px;
  text-align: 'center';
  color: white;
  font-Size: 25px;
  font-Family: 'nanumgothic';
  font-weight: 600;
}
.google_login
{
  position: relative;
  top:-50px;
}
</style>