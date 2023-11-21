<!-- 관리자 보고서 화면 -->
<template>
<div class="AndroidLarge2" >
  <div class="container" style="width: 100%; height: 100%; position: relative; background: #f2f2f2;">
    <!-- <router-link to="/admin/main" style="width: 83px; height: 25px; left: 7px; top: 22px; position: absolute; text-align: center;  word-wrap: break-word; color: #507BBC; font-size: 15px; font-family: Helvetica; font-weight: 700; text-decoration: none;"> 돌아가기 </router-link> -->
    <router-link to="/admin/main" style="position: absolute; left: 10px; top: 60px;">
      <img :src="require('@/assets/pknu_logo.png')" alt="돌아가기" style="width: 40px; height: 40px;">
    </router-link>

  <div style="width: 190px; height: 43px; left: 90px; top: 60px; position: absolute; text-align: center; color: black; font-size: 25px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">{{ reportList[$route.params.id-1].title}}</div>
  <div style="width: 83px; height: 20px; left: 30px; top: 157px; position: absolute; text-align: left; color: #173970; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">작성자</div>
  <div style="width: 83px; height: 20px; left: 30px; top: 200px; position: absolute; text-align: left; color: #173970; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">저자</div>
  <div style="width: 83px; height: 20px; left: 30px; top: 243px; position: absolute; text-align: left; color: #173970; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">출판사</div>
  <!-- 이 밑에 세개 각자 axios로 입력받아서 출력해야함 -->
  <div style="width: 163px; height: 20px; left: 164px; top: 157px; position: absolute; text-align: right; color: black; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">{{ reportList[$route.params.id-1].name}}</div>
  <div style="width: 190px; height: 20px; right: 33px; top: 200px; position: absolute; text-align: right; color: black; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">{{ reportList[$route.params.id-1].author}}</div>
  <div style="width: 186px; height: 20px; left: 141px; top: 243px; position: absolute; text-align: right; color: black; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">{{ reportList[$route.params.id-1].publisher}}</div>


  <!-- 승인 버튼 -->
  <button class="btn_proval" @clink="submitForm">
    <router-link to = "/admin/main">
    <div class="Rectangle11"></div>
    <div class="btn_text_p">승인</div>
    </router-link>
  </button>  
  <div class="Line1" style="width: 360px; height: 0px; left: 0px; top: 126px; position: absolute; border: 1px #D9D9D9 solid"></div>
  <div class="Rectangle11" style="width: 317px; height: 409px; left: 21px; top: 294px; position: absolute; background: white; border-radius: 19px; border: 1px #BEBEBE solid"></div>
  <div style="width: 280px; height: 20px; left: 41px; top: 313px; position: absolute; color: black; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">{{ reportList[$route.params.id-1].content}}</div>
</div>

<!-- 배지 발급 버튼 -->
<button class="btn_issue" @click="modalOpen = true" style="width: 120px; height: 49px; left: 180px; top: 720px; position: absolute; background: #507BBC; border-radius: 30px; border:none;" >
      <div style="width: 73px; height: 23px; left: 21px; top: 15px; position: absolute; text-align: center; color: white; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">배지 발급</div>
    </button>
      <div class="black-bg" v-if="modalOpen === true">
          <div class="white-bg">
            <h3> 배지가 발급되었습니다.</h3>
            <router-link to="/admin/main" class="ok_btn">
              확인
            </router-link>
          </div>          
      </div>
</div>
</template>

<script>
import router from '@/router';
import report_data from '@/assets/js/report.js';

  export default {
    name: 'bogoseoForm',
    data() {
        return {
            reportList: report_data,
            modalOpen: false,
            Report_title: '',
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
    components: { router }
};
  
</script>


<style>
@font-face{
  font-family: 'nanumgothic';
  src: url('@/assets/fonts/NanumGothic.ttf');
}
.AndroidLarge2
{width: 360px;
height: 800px;
background: #f2f2f2;
position: relative;
}
.btn_proval {
    width: 116px;
    height: 45px;
    left: 53px;
    top: 722px;
    position: absolute;
    border: none;
    border-radius: 50px;
  }
  .black-bg {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 320px;
  height: 800px;
  top: 0px;
  left: 0px;
  background-color: rgba(0, 0, 0, 0.432);
  position: absolute;
  padding: 20px;
}
.white-bg {
  /* width: 100%; */
  width: 300px;
  height: 130px;
  background-color: white;
  padding: 20px;
  border-radius: 18px;
  align-items: center;
}
  .Rectangle11 {
    width: 116px;
    height: 45px;
    left: 0px;
    top: 0px;
    position: absolute;
    background: #f2f2f2;
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
    /* style="width: 116px; height: 45px; left: 0px; top: 0px; position: absolute; background: #507BBC; border-radius: 50px" */
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
    font-family: nanumgothic;
    font-weight: 700;
    word-wrap: break-word;
  }
  .ok_btn
  {
    left: 80px;
    top: 5px;
    width: 40%;
    height: 40px;
    text-align:center;
    background-color: #507BBC;
    color: #f2f2f2;
    font-Size: 15px;
    font-Family: nanumgothic;
    position: relative;
    border:none;
    border-radius: 30px;
    text-decoration: none;
  }
</style>