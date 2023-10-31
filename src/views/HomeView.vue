<template>
  <div className="AndroidLarge1">
      <img class="pknu_logo" :src="require('@/assets/white_logo.png')">
      <div className="Title">부경대학교<br/>모바일 도서관</div>
      <a href="googleOAuthUrl">
      <img class="google_login" :src="require('@/assets/sign_google.png')"/>
      </a>
        <nav>
      <!-- <router-link to="/">Home</router-link> | -->
    <router-link to="/about">메인 화면</router-link>
  </nav>
  <router-view/>
  </div>
</template>

<script>
export default{
computed:{
  googleOAuthUrl() {
      return `https://accounts.google.com/o/oauth2/v2/auth?scope=https://www.googleapis.com/auth/userinfo.email&response_type=code&client_id=${this.client_id}&redirect_uri=${this.redirect_uri}`;
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