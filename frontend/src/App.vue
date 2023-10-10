<template>
  <div>
    <button @click="getData">데이터 가져오기</button>
    <p>{{ responseData }}</p>
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

