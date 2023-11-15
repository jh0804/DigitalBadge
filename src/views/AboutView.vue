<!-- 학생 메인 화면 -->

<template>
  <div class="main">
    <div className="badge_box">
    </div>
    <!-- badge jason에서 받아와서 입력 -->
    <router-link to="/Badge_info" className="badge" style="width: 78px; height: 90px; left: 52px; top: 424px; position: absolute; background: #D9D9D9; border-radius: 6px">
      <li v-for="badge in badgeList" :key="badge">
        <img :src="require(badge.img)"  alt="Badge Image">
        <span>{{ badge.name }}</span>
      </li>
    </router-link>
      <div style="width: 83px; height: 25px; left: 48px; top: 523px; position: absolute; text-align: center; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">배지 이름</div>
     <div className="Rectangle3">
      <div className="Rectangle4"></div>
      <div className="Rectangle5"></div>
      <!-- 여기서부터 세개는 입력받는 데이터 -->
      <div style="width: 135px; height: 28px; left: 133px; top: 200px; position: absolute; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">이름</div>
      <div style="width: 115px; height: 20px; left: 141px; top: 234px; position: absolute; color: #686868; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">학번</div>
      <div style="width: 115px; height: 20px; left: 141px; top: 264px; position: absolute; color: #686868; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">학과</div>

        <!-- 밑에 버튼 누르면 데이터 베이스로 데이터 보내지게 해야 함 -->
        <button @click="modalOpen = false" style=" width: 70px; height:40px; top: 0px; position: relative; background: #507BBC; border-radius: 20px; border:none; text-align: center; " class="modal-exit-btn">
          <h4 style="color : white; text-align: center; ">확인</h4>
        </button>
      </div>
    </div>
    <img class="PknuLogo1" :src="require('@/assets/pknu_logo.png')"/>
    <router-link to="/activity/report" style=" text-align: center;" class="reporterbtn"> 
    <h4 style="color:white; font-family: Inter; font-weight: 700; word-wrap: break-word;" >활동 보고</h4>
  </router-link>
</template>

<script>
import badge_data from '@/assets/badge.json';

export default {
  data() {
    return {
      badgeList : badge_data,
    };
},

methods: {
  submitForm2() {
    console.log('dd');
  },
},
};

</script>

<style>
.main
{width: 360px;
height: 800px;
background: #f2f2f2;
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
  left: 21px;
  top: 386px;
  position: absolute;
  background: white;
  box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.25);
  border-radius: 19px;
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
.reporterbtn{
  width: 317px;
  height: 49px;
  left: 21px;
  top: 726px;
  position: absolute;
  background: #507BBC;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 11px
}
</style>