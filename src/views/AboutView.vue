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
      <div >
        <img class="profile_pic" :src="require('@/assets/profile.jpg')"/>
      </div>
      <img class="icon" :src="require('@/assets/User.png')" style="width: 22px; height: 22px; left: 0px; top: 175px; position: relative;"/>
      <h3 style="position: relative; top: 140px; left:46px; font-size: 14px; color: cornflowerblue;">학부 재학생 </h3>
      <!-- 여기서부터 세개는 서버에서 입력받는 데이터 -->
      <div style="width: 135px; height: 28px; left: 134px; top: 212px; position: absolute; color: black; font-size: 24px; font-family: nanumgothic; font-weight: 900; word-wrap: break-word">{{ name }}</div>
      <div style="width: 140px; height: 20px; left: 160px; top: 240px; position: absolute; color: #b3b3b3; font-size: 11px; font-family: Inter; font-style: italic; font-weight: 600; word-wrap: break-word">{{ email }}</div>
      <div style="width: 115px; height: 20px; left: 156px; top: 264px; position: absolute; color: #686868; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">{{ roleId }}</div>
      </div>
    </div>
    <img class="PknuLogo1" :src="require('@/assets/pknu_logo.png')"/>
    <router-link to="/activity/report" class="reporterbtn1"> 
    <h4 style="color:white; font-family:nanumgothic; font-weight: 700; text-align: center; word-wrap: break-word;" >활동 보고</h4>
  </router-link>
</template>

<script>
import badge_data from '@/assets/js/badge.js';
import axios from 'axios';

export default {
  data() {
    return {
      badgeList : badge_data,
      name: '김규린',
      email: 'zpdldptmdpf0498@gmail.com',
      roleId: '201912345',
      badgeId: null, // 새로운 badgeId 추가
    };  
  },
  mounted() {
    this.fetchUserInfo();          
},
methods: {
  async fetchUserInfo() {
      try {
        const response = await axios.get('https://reqres.in/api/users?/page=2'); // 실제 API 엔드포인트로 대체
        const userData = response.data; // 서버가 사용자 정보를 담은 객체를 반환한다고 가정
        this.name = userData.name ||'김규린';
        this.email = userData.email || 'axios@pukyong.ac.kr2';
        this.roleId = userData.roleId || '201912345';
        this.badgeId = userData.badgeId || null;
      } catch (error) {
        console.error('사용자 정보를 불러오는 중 오류 발생:', error);
        // 오류 처리, 사용자에게 메시지 표시 등을 수행합니다.
      }
    },
    updateBadgeList() {
      if (this.badgeId) {
        // badgeId가 존재할 경우, 해당 ID에 맞는 배지 정보 가져오기
        const selectedBadge = badge_data.find(badge => badge.id === this.badgeId);
        if (selectedBadge) {
          // 선택한 배지 정보가 존재하면 badgeList 업데이트
          this.badgeList = [selectedBadge];
        }
      }
    },
    navigateToBadge(badgeId) {
      // 예시: 해당 라우터로 이동하는 코드
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
.profile_pic{
  width: 123px; 
  height: 140px; 
  left: 48px; 
  top: 156px; 
  position: absolute; 
  border: 1px;
  border-color: #507BBC;
}
.PknuLogo1
{
  width: 140px;
  height: 35px;
  left: 25px;
  top: 130px;
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
  left: 29px;
  top: 786px;
  position: absolute;
  background: #507BBC;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  border-radius: 50px;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: center;
}


</style>