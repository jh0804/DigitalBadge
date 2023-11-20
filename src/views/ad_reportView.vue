<!-- 관리자 보고서 화면 -->
<template>
  <div class="container" style="width: 100%; height: 100%; position: relative; background: white;">
    <!-- <router-link to="/admin/main" style="width: 83px; height: 25px; left: 7px; top: 22px; position: absolute; text-align: center;  word-wrap: break-word; color: #507BBC; font-size: 15px; font-family: Helvetica; font-weight: 700; text-decoration: none;"> 돌아가기 </router-link> -->
    <router-link to="/admin/main" style="position: absolute; left: 10px; top: 60px;">
      <img :src="require('@/assets/pknu_logo.png')" alt="돌아가기" style="width: 40px; height: 40px;">
    </router-link>

  <div style="width: 190px; height: 43px; left: 90px; top: 60px; position: absolute; text-align: center; color: black; font-size: 25px; font-family: Inter; font-weight: 700; word-wrap: break-word">어린왕자를 읽고</div>
    <div style="width: 83px; height: 20px; left: 22px; top: 157px; position: absolute; text-align: left; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">제목</div>
    <div style="width: 83px; height: 20px; left: 22px; top: 200px; position: absolute; text-align: left; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">저자</div>
    <div style="width: 83px; height: 20px; left: 22px; top: 243px; position: absolute; text-align: left; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">출판사</div>
  <!-- 이 밑에 세개 각자 axios로 입력받아서 출력해야함 -->
  <div style="width: 163px; height: 20px; left: 164px; top: 157px; position: absolute; text-align: right; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">어린 왕자</div>
  <div style="width: 172px; height: 20px; left: 155px; top: 200px; position: absolute; text-align: right; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">앙투안 드 생텍쥐페리</div>
  <div style="width: 186px; height: 20px; left: 141px; top: 243px; position: absolute; text-align: right; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">열린 책들</div>


  <!-- 승인 버튼 -->
  <button @click="approve" class="btn_confirm" style="width: 116px; height: 45px; left: 53px; top: 722px; position: absolute">
      <div class="Rectangle11" style="width: 116px; height: 45px; left: 0px; top: 0px; position: absolute; background: white; border-radius: 50px; border: 1px #507BBC solid"></div>
      <div style="width: 58px; height: 22.96px; left: 29px; top: 15px; position: absolute; text-align: center; color: #507BBC; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">승인</div>
  </button>

<!-- 배지 발급 버튼 -->
  <button class="btn_issue" style="width: 116px; height: 45px; left: 190px; top: 722px; position: absolute">
      <div class="Rectangle12" style="width: 116px; height: 45px; left: 0px; top: 0px; position: absolute; background: #507BBC; border-radius: 50px"></div>
      <div style="width: 73px; height: 23px; left: 21px; top: 15px; position: absolute; text-align: center; color: white; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">배지 발급</div>
  </button>
  
  <div class="Line1" style="width: 360px; height: 0px; left: 0px; top: 126px; position: absolute; border: 1px #D9D9D9 solid"></div>
  <div class="Rectangle11" style="width: 317px; height: 409px; left: 21px; top: 294px; position: absolute; background: white; border-radius: 19px; border: 1px #BEBEBE solid"></div>
  <!-- 이 밑에 내용을 데이터 받아서 출력해야 함-->
  <div style="width: 186px; height: 20px; left: 41px; top: 313px; position: absolute; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">이 책은 재밌습니다.</div>
</div>
</template>

<script>
  export default {
    name: 'bogoseoForm',
    data() {
      return {
      // 모달을 숨겨 주는 변수선언
      Report_title:'',
      Book_title: '',
      author: '',
      contents: '',
      approval: false,
      };
  },
  methods: {
    approve() {
      this.approval('dd');
    },
  },
  };
  
</script>


<style>
.btn_proval {
    width: 116px;
    height: 45px;
    left: 53px;
    top: 722px;
    position: absolute;
    border: none;
    border-radius: 50px;
  }

  .Rectangle11 {
    width: 116px;
    height: 45px;
    left: 0px;
    top: 0px;
    position: absolute;
    background: white;
    border-radius: 50px;
    border: 1px #507BBC solid;
  }

  .btn_text_p {
    width: 58px;
    height: 22.96px;
    left: 29px;
    top: 12px;
    position: absolute;
    text-align: center;
    color: #507BBC;
    font-size: 16px;
    font-family: Inter;
    font-weight: 700;
    word-wrap: break-word;
  }

  .btn_issue {
    width: 116px;
    height: 45px;
    left: 190px;
    top: 722px;
    position: absolute;
    border: none;
    border-radius: 50px;
  }

  .Rectangle12 {
    width: 116px;
    height: 45px;
    left: 0px;
    top: 0px;
    position: absolute;
    background: #507BBC;
    border-radius: 50px;
  }

  .btn_text_i {
    width: 73px;
    height: 23px;
    left: 21px;
    top: 12px;
    position: absolute;
    text-align: center;
    color: white;
    font-size: 16px;
    font-family: Inter;
    font-weight: 700;
    word-wrap: break-word;
  }
</style>