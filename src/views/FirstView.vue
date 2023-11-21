<!-- Intro 화면의 애니메이션을 담당하는 vue -->
<template>
    <div class="TextAnime1" @animationend="animeEnd">
      <span
        v-for="(t, index) in text"
        :key="index"
        class="item"
        :style="{animationDelay: index*100+'ms'}"
        v-text="t"
        />
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        text: 'Lib Badge'
      };
    },
  methods: {
    animeEnd() {
        setTimeout(() => {
        this.$router.push('/login');// 애니메이션이 끝났을 때 호출되는 메소드
      // 다음 화면으로 전환하거나 원하는 로직 수행 가능
      this.$emit('animationEnd'); // 부모 컴포넌트로 이벤트 발송
    }, 1500);
    }
  }
};
  </script>
  
  <style scoped>
  @keyframes text-in {
    0% {
      transform: translate(0, -20px);
      opacity: 0;
    }
  }
  .item {
    /* top: 150px; */
    display: inline-block;
    min-width: 0.3em;
    font-size: 2rem;
    color:173970;
    animation: text-in .8s cubic-bezier(0.22, 0.15, 0.25, 1.43) 0s backwards;
  }
  .TextAnime1
  {
top:80px;
  }
  </style>