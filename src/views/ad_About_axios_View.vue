<!-- axios 연결 시 관리자 메인화면 -->

<template>
    <div class="main">  
       <div className="Rectangle3 ">
        <div className="Rectangle4"></div>
        <div className="Rectangle5"></div>
        <div class="ad_name">{{ name }}</div>
      <div style="width: 115px; height: 20px; left: 160px; top: 236px; position: absolute; color: #b3b3b3; font-size: 11px; font-family: Inter; font-style: italic; font-weight: 600; word-wrap: break-word">{{ email }}</div>
      <div style="width: 115px; height: 20px; left: 156px; top: 264px; position: absolute; color: #686868; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">{{ studentNumber }}</div>

    </div>
    <img class="PknuLogo1" :src="require('@/assets/pknu_logo.png')"/>
    <div class="Group1" style="width: 317px; height: 442px; left: 0px; top: 0px; position: relative;">      
    <div class="Rectangle10" style="width: 100%; height: 100%; relative; background: white; box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.25); border-radius: 19px"> 
    <div v-for="book in filteredreport_List" :key="book.id" @click="navigateToreport(book.id)" style="margin: 10px;">
        <router-link :to="'/admin/report/' + book.id" class="book" style="width: 78px; height: 90px;">
          <p style="color: black; font-size: 14px; ">{{ book.name }}</p>          
        </router-link>
      </div>

<!-- 여기서부터는 백엔드 연결되면 지워도 됨둥 -->
<router-link to="/admin/report" style="width: 76.84px; left: 0.18px; position: absolute; text-align: center; color: #507BBC; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">
          <div class="Rectangle17" style="width: 260px; height: 50px; left: 20px; top: 58px; position: absolute; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px"></div>
          <div class="Oo20xxxxxxx" style="width: 193px; height: 25px; left: 0px; top: 33px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 500; word-wrap: break-word">김OO | 20XXXXXXX</div>
          <div style="width: 170px; height: 25px; left: 20px; top: 75px; position: absolute; text-align: center; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">어린왕자를 읽고</div>
        </router-link>
 
        <router-link to="/admin/report" style="width: 76.84px; left: 0.18px; position: absolute; text-align: center; color: #507BBC; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">
          <div class="Rectangle18" style="width: 260px; height: 50px; left: 20px; top: 167px; position: absolute; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px"></div>
          <div class="Oo20xxxxxxx" style="width: 193px; height: 25px; left: 0px; top: 142px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 500; word-wrap: break-word">이OO | 20XXXXXXX</div>
          <div style="width: 170px; height: 25px; left: 20px; top: 184px; position: absolute; text-align: center; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">백설공주를 읽고</div>
        </router-link>

        <router-link to="/admin/report" style="width: 76.84px; left: 0.18px; position: absolute; text-align: center; color: #507BBC; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">
          <div class="Rectangle19" style="width: 260px; height: 50px; left: 20px; top: 276px; position: absolute; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px"></div>
          <div class="Oo20xxxxxxx" style="width: 193px; height: 25px; left: 0px; top: 251px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 500; word-wrap: break-word">박OO | 20XXXXXXX</div>
          <div style="width: 170px; height: 25px; left: 10px; top: 293px; position: absolute; text-align: center; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">1984를 읽고</div>
        </router-link> 

      
        <router-link to="/admin/report" style="width: 76.84px; left: 0.18px; position: absolute; text-align: center; color: #507BBC; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">
          <div class="Rectangle20" style="width: 260px; height: 50px; left: 20px; top: 385px; position: absolute; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px"></div>
          <div class="Oo20xxxxxxx" style="width: 193px; height: 25px; left: 0px; top: 360px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 500; word-wrap: break-word">박OO | 20XXXXXXX</div>
          <div style="width: 170px; height: 25px; left: 10px; top: 402px; position: absolute; text-align: center; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">1984를 읽고</div>
        </router-link> 

            
        <router-link to="/admin/report" style="width: 76.84px; left: 0.18px; position: absolute; text-align: center; color: #507BBC; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">
          <div class="Rectangle21" style="width: 260px; height: 50px; left: 20px; top: 494px; position: absolute; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px"></div>
          <div class="Oo20xxxxxxx" style="width: 193px; height: 25px; left: 0px; top: 469px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 500; word-wrap: break-word">박OO | 20XXXXXXX</div>
          <div style="width: 170px; height: 25px; left: 10px; top: 511px; position: absolute; text-align: center; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">1984를 읽고</div>
        </router-link> 
      </div>
    </div>
  </div>
<!-- 여기까지 지워도 됨 -->
  <!-- 백엔드 연결되면 /div 태그 2개 추가해야함 -->
  
  </template>
  
  
  <script>
import report_data from '@/assets/js/report.js';
import axios from 'axios';

  export default {
    data() {
      return {
      // 모달을 숨겨 주는 변수선언
      modalOpen: false,
      name: '관리자',
      email: 'library@pukyong.ac.kr',
      studentNumber: '201212345',
      // NEW!!!!!      
      reportList: [], // 새로운 데이터 배열 추가
      };
  },
  methods: {
    async fetchUserData() {
      try {
        // axios를 사용하여 서버에서 데이터를 가져옵니다.
        const response = await axios.get("/api/user-data"); // 실제 API 엔드포인트로 변경
        const userData = response.data; // 서버에서 받아온 데이터 객체

        // 서버에서 받아온 데이터를 컴포넌트의 데이터에 할당합니다.
        this.name = userData.name || "";
        this.email = userData.email || "";
        this.studentNumber = userData.studentNumber || "";        
      } catch (error) {
        console.error("데이터를 가져오는 중 에러 발생:", error);
      }
    },
    async fetchreportList() {
      try {
        const response = await axios.get("/api/report-list"); // 실제 API 엔드포인트로 변경
        this.reportList = response.data.filter(book => book.approval === false); // approval이 false인 책만 필터링
      } catch (error) {
        console.error("데이터를 가져오는 중 에러 발생:", error);
      }
    },
    navigateToreport(reportId) {
      this.$router.push({ name: '/admin/report', params: { id: reportId } });
    },
  },
  computed: {
    filtered_reportList() {
      // 필터링된 책 리스트를 반환하는 computed 속성
      return this.reportList.filter(book => book.approval === false);
    },
  },
  mounted() {
    this.fetchUserData();
    this.fetch_reportList(); // 책 리스트를 가져오는 메소드 호출
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
  height: 850px;
  background: #f2f2f2;
  }
  .Rectangle3{
    width: 360px;
    height: 360px; 
    top: -20px; 
    position: relative; 
    background: #507BBC;
  }
  .Rectangle4{
    width: 308px; 
    height: 200px; 
    left: 28px; 
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
    top: 75px;
    position: absolute;
  }

  .Rectangle10{
  width: 123px; 
  height: 140px; 
  left: 25px; 
  top: 10px; 
  position: relative; 
  background: #D9D9D9;
  overflow-y: scroll;
}

  
  .modal-exit-btn {
    margin-top: 30px;
  }
  
  .modal-exit-btn:hover {
    cursor: pointer;
  }
  .group1{
    width: 317px;
    height: 442px;
    left: 20px;
    top: 321px;
    position: relative;
  }
  .ad_name
  {
    width: 135px;
    height: 28px;
    left: 133px;
    top: 200px;
    position: absolute;
    color: black;
    font-size: 18px;
    font-family: Inter;
    font-weight: 700;
    word-wrap: break-word;
  }

  
  </style>