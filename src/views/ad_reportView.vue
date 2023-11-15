<!-- axios테스트 -->

<template>
    <v-container>
      <v-card style="width: 30%">
        <v-text-field v-model="b1" label="1 or 2"></v-text-field>
        <v-btn outlined @click="button">Click</v-btn>
        <p>{{a}}</p>
      </v-card>
    </v-container>
  </template>
  
  <script>
  export default {
    name: 'App',
    data: () => ({
      b1: '',
      a: ''
    }),
    methods: {
      button () {
        this.axios.get('/example', { params: { tEst: this.b1 } }).then(result => {
          console.log(result.data)
          this.a = result.data
        })
      }
    }
  }
  </script>