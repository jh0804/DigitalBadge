import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios';
import router from './router';
import GoogleLogin from './components/GoogleLogin.vue';

const app = createApp(App);

app.use(router);

app.provide('$axios', axios);

app.mount('#app');

const indexApp = createApp(GoogleLogin); // Create a new app for the Index component
indexApp.mount('#GoogleLogin');

