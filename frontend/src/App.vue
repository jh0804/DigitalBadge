<template>
  <div>
    <button @click="getData">데이터 가져오기</button>
    <p>{{ responseData }}</p>
    <router-link to="/">Home</router-link>
    <router-link to="/google-login">구글 로그인</router-link> <!-- 구글 로그인 버튼 -->
    <router-view></router-view>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "App",
  data() {
    return {
      responseData: "", 
    };
  },
  methods: {
    getData() {
      axios
        .get("http://localhost:8082/api/test") 
        .then((response) => {
          this.responseData = response.data;
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>

<style></style>