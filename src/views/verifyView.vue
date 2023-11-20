<!-- 검증화면 -->

<template>
  <div class="container">
    <div class="Verified">verified</div>
    <img class="check" :src="require('@/assets/check.png')"/>
    <div class="Valid">Your badge is valid !</div>
    <button @click="goBack" class="Back_btn">
      <img class="cross" :src="require('@/assets/Cross.png')" alt="Back Button"/>
    </button>
  </div>
</template>

<script>
export default {
  methods: {
    goBack() {
      this.$router.go(-1); // 이전 페이지로 이동
    }
  }
}
</script>

<style>
.container {
  width: 100%;
  height: 100vh; /* 100% 뷰포트 높이 */
  position: relative;
  background: white;
}

.Verified {
  width: 100%;
  text-align: center;
  color: black;
  font-size: 25px;
  font-family: Inter;
  font-weight: 700;
  word-wrap: break-word;
  position: absolute;
  top: 45%;
}

.Valid {
  width: 100%;
  text-align: center;
  color: #27214D;
  font-size: 20px;
  font-family: Alata;
  font-weight: 400;
  line-height: 30px;
  word-wrap: break-word;
  position: absolute;
  top: 50%;
}

.check {
  position: absolute;
  left: 50%;
  top: 30%;
  transform: translate(-50%, -50%);
}

/* 미디어 쿼리 */
@media (max-width: 768px) {
  .Verified {
    font-size: 20px;
  }

  .Valid {
    font-size: 16px;
  }
}
.Back_btn
{
  top: 65px;
  right: 200px;
  width : 40px;
  height: 40px;
  background-color: white;
  border: none;
  position: absolute;

}
.cross{
  position: relative;
  width : 50px;
  height: 50px;
  top:0px;

}
</style>