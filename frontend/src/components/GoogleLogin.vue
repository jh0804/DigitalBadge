<template>
    <div class="row">
      <div class="col-md-6">
        <a href="/posts/save" role="button" class="btn btn-primary">글 등록</a>
        <div v-if="userName">
          Logged in as: <span id="user">{{ userName }}</span>
          <a href="/logout" class="btn btn-info active" role="button" @click="logout">Logout</a>
        </div>
        <div v-else>
          <a href="/oauth2/authorization/google" class="btn btn-success active" role="button">Google Login</a>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        userName: null,
      };
    },
    created() {
      // Make an API call to fetch user information
      axios.get('/api/user').then((response) => {
        if (response.data) {
          this.userName = response.data.name;
        }
      });
    },
    methods: {
      logout() {
        // Make an API call or perform necessary actions for logout
        // For example, you can make an API call to invalidate the session on the server-side
        // After logout, you can redirect the user to the desired page
        // window.location.href = '/logout';
      },
    },
  };
  </script>
  