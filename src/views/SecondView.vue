<!-- Intro 화면 -->

<template>
    <div class="main">        
      <TextAnime2 v-if="anime2" class="intro"/>
      </div>

</template>
  
<script>
  import TextAnime2 from './FirstView.vue'
  export default {
    components: {
      TextAnime2,
    },
    data() {
      return {
        anime2: true,
        autoplay: true
        
      }
    }
  }
</script>

<style>
.main
{width: 360px;
height: 800px;
background: #f2f2f2;
font-family: 'nanumgothic';
font-weight: 400;
border: 1px black;
position : absolute;
}
.intro {
  position: absolute; /* 절대 위치로 설정 */
  top: 100%; /* 상위 가운데 정렬 */
  left: 50%; /* 좌측 가운데 정렬 */
  transform: translate(-50%, 700%); /* 세로, 가로 모두 중앙으로 이동 */
}
</style>