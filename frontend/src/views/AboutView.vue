<!-- 학생 메인 화면 -->

<template>
  <div class="main" >
    <div className="badge_box" style="display: flex; justify-content: space-around; flex-wrap: wrap; ">
  <!-- badge jason에서 받아와서 입력 -->
  <div v-for="badge in badgeList" :key="badge.id" @click="navigateToBadge(badge.id)" style="margin: 10px;">
    <router-link :to="'/Badge_info/' + badge.id" className="badge" style="width: 78px; height: 90px;">
      <img :src="badge.img" alt="Badge Image" style="width: 78px; height: 100px;"/>
      <p style="color: black; font-size: 14px; ">{{ badge.name }}</p>
    </router-link>
  </div>
</div>
     <div className="Rectangle3">
      <div className="Rectangle4"></div>
      <div className="Rectangle5"></div>
      <!-- 여기서부터 세개는 입력받는 데이터 -->
      <div style="width: 135px; height: 28px; left: 133px; top: 200px; position: absolute; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">이름</div>
      <div style="width: 115px; height: 20px; left: 141px; top: 234px; position: absolute; color: #686868; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">학번</div>
      <div style="width: 115px; height: 20px; left: 141px; top: 264px; position: absolute; color: #686868; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">학과</div>
      </div>
    </div>
    <img class="PknuLogo1" :src="require('@/assets/pknu_logo.png')"/>
    <router-link to="/activity/report" style=" text-align: center;" class="reporterbtn1"> 
    <h4 style="color:white; font-family:'nago'; font-weight: 700; text-align: center; word-wrap: break-word;" >활동 보고</h4>
  </router-link>
</template>

<script>
import badge_data from '@/assets/badge.js';

export default {
  data() {
    return {
      badgeList : badge_data,
    };        
},
methods: {
    navigateToBadge(badgeId) {
      // 예시: 해당 라우터로 이동하는 코드
      // 실제로는 적절한 라우터로 이동하는 로직을 추가해야 합니다.
      this.$router.push({ name: '/Badge_info', params: { id: badgeId } });
    },
  },

};

</script>

<style>
@font-face{
  font-family: 'nanumgothic';
  src: url('@/assets/fonts/NanumGothic.ttf');
}
.main
{width: 360px;
height: 800px;
background: #f2f2f2;
font-family: 'nanumgothic';
font-weight: 600;
}

.Rectangle3{
  width: 360px;
  height: 369px; 
  top: 0px; 
  position: absolute; 
  background: #507BBC;
}
.Rectangle4{
  width: 308px; 
  height: 221px; 
  left: 26px; 
  top: 116px; 
  position: absolute; 
  background: white; 
  border-radius: 12px;
}
.Rectangle5{
  width: 123px; 
  height: 140px; 
  left: 48px; 
  top: 156px; 
  position: absolute; 
  background: #D9D9D9;
}
.PknuLogo1
{
  width: 140px;
  height: 35px;
  left: 25px;
  top: 69px;
  position: absolute;
}
.badge_box{
  width: 317px;
  height: 321px;
  left: 24px;
  top: 386px;
  position: absolute;
  background: white;
  box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.25);
  border-radius: 20px;
}
.black-bg {
  display: flex;
  align-items: center;
  width: 360px;
  height: 800px;
  background-color: rgba(0, 0, 0, 0.432);
  position: fixed;
  padding: 20px;
}

.white-bg1 {
  width: 300px;
  height: 400px;
  background-color: white;
  padding: 20px;
  border-radius: 18px;
}

.modal-exit-btn {
  margin-top: 30px;
}

.modal-exit-btn:hover {
  cursor: pointer;
}
.reporterbtn1{
  width: 317px;
  height: 50px;
  left: 21px;
  top: 726px;
  position: absolute;
  background: #507BBC;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 50px;
  text-align: center;
}


</style>