<!-- 관리자 메인 화면 -->

<template>
    <div class="main">
      <div class="Rectangle3">
        <div class="Rectangle4">
        <!-- axios로 부터 받아온 사용자 정보 -->
        <img class="icon" :src="require('@/assets/User.png')" style="width: 24px; height: 24px; left: -112px; top: 35px; position: relative;"/>
        <h3 style="position: relative; top: -5px; left:-48px; font-size: 16px; color: cornflowerblue;">도서관 관리자</h3>
        <div class="ad_name">{{ name }}</div>
        <div style="width: 120px; height: 20px; left: 32px; top: 0px; position:relative; color: #b3b3b3; font-size: 11px; font-family: Inter; font-style: italic; font-weight: 600; word-wrap: break-word">{{ email }}</div>
        <div style="width: 115px; height: 20px; left: 17px; top: 0px; position: relative; color: #686868; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">{{ studentNumber }}</div>
    </div> 
    </div>
      <img class="PknuLogo1" :src="require('@/assets/pknu_logo.png')" />
      <div class="Group1" style="width: 317px; height: 442px; left: 0px; top: -5px; position: relative;">
        <!-- approval이 false인 책 리스트 -->
        <div class="Rectangle10" style="width: 100%; height: 100%; relative; background: white; box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.25); border-radius: 19px">
          <div v-for="book in filtered_reportList" :key="book.id" style="margin: 10px;">
            <router-link :to="'/admin/report/'+ book.id" class="book">
              <div class="student_info" >{{ book.name }} | {{ book.roleId }}</div>
              <div class="Rectangle21" style="width: 260px; height: 50px; left: 10px; top:0px; position:relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px">
              <div class="book_title">{{ book.title }}</div>
            </div>
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import report_data from '@/assets/js/report.js';
  import axios from 'axios';
  
  export default {
    data() {
      return {
        modalOpen: false,
        name: '김부경',
        email: 'library@pukyong.ac.kr',
        studentNumber: '201212345',
        reportList: report_data,
      };
    },
    methods: {
      async fetchUserData() {
        try {
          const response = await axios.get("/api/user-data");
          const userData = response.data;
          this.name = userData.name || "";
          this.email = userData.email || "";
          this.studentNumber = userData.studentNumber || "";        
        } catch (error) {
          console.error("데이터를 가져오는 중 에러 발생:", error);
        }
      },
      async fetchreportList() {  // 메서드 이름 수정
        try {
          const response = await axios.get("/api/report-list");
          this.reportList = response.data.filter(book => book.approval === false);
        } catch (error) {
          console.error("데이터를 가져오는 중 에러 발생:", error);
        }
      },
      navigateToReport(reportId) {
        this.$router.push({ name: '/admin/report', params: { id: reportId } });
      },
    },
    computed: {
      filtered_reportList() {
        return this.reportList.filter(book => book.approval === false);
      },
    },
    mounted() {
      this.fetchUserData();
    },
  };
  </script>
  
  <style>
  @font-face{
  font-family: 'nanumgothic';
  src: url('@/assets/fonts/NanumGothic.ttf');
  }
  .main
  {
    width: 360px;
    height: 800px;
    background: #f2f2f2;
    font-family: 'nanumgothic';
    font-weight: 600;
    }

  .Rectangle3{
    width: 360px;
    height: 310px; 
    top: -20px; 
    position: relative; 
    background: #507BBC;
  }
  .Rectangle4{
    width: 308px; 
    height: 170px; 
    left: 28px; 
    top: 120px; 
    position: absolute; 
    background: white; 
    border-radius: 12px;
  }
  .PknuLogo1
  {
    width: 140px;
    height: 35px;
    left: 25px;
    top: 55px;
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
    left: 2px;
    top: -4px;
    position: relative;
    color: black;
    font-size: 25px;
    font-family: nanumgothic;
    font-weight: 700;
    word-wrap: break-word;
  }
  .book {
    top:-4px;
  width: 260px;
  height: 50px;
  margin: 15px; /* 책 사이의 간격을 위해 여백 추가 */
  position: relative; /* 상대 위치 지정 사용 */
  background: #EFEFEF;
  box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25);
  border-radius: 11px;
  text-decoration: none;
}
.book_title{
    width: 100px;
    height: 30px;
    left: 0px;
    top: 14px;
    position: absolute;
    text-align: center;
    color: #4D4D4D;
    font-size: 18px;
    font-family: nanumgothic;
    font-weight: 700;
    word-wrap: break-word;
    background-color: #f2f2f2;
}
.student_info
{
    width: 193px;
    height: 25px;
    left: 0px;
    top: 2px;
    position: relative;
    color: #4D4D4D;
    font-size: 16px;
    font-family: nanumgothic;
    font-weight: 500;
    word-wrap: break-word;
    text-decoration: none;
    
} 
  </style>