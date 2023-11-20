<!-- 보고서 작성 페이지 -->
<template>
<div class="1" >
    <router-link to="/student/main" style="width: 83px; height: 25px; left: 7px; top: 92px; position: absolute; text-align: center;  word-wrap: break-word; color: #507BBC; font-size: 15px; font-family: nanumgothic; font-weight: 700; text-decoration: none;"> 돌아가기 </router-link>
    <div style="width: 129px; height: 43px; left: 115px; top: 54px; position: absolute; text-align: center; color: black; font-size: 25px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">활동 보고</div>
    <div style="width: 83px; height: 20px; left: 24px; top: 157px; position: absolute; color: black; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">제목</div>
    <div style="width: 83px; height: 20px; left: 24px; top: 200px; position: absolute; color: black; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">저자</div>
    <div style="width: 83px; height: 20px; left: 32px; top: 243px; position: absolute; color: black; font-size: 18px; font-family: nanumgothic; font-weight: 700; word-wrap: break-word">출판사</div>
    <input v-model="title" style="width: 186px; height: 20px; left: 141px; top: 157px; position: absolute; text-align: right; font-size: 18px; font-weight: 600; word-wrap: break-word; border: none;" placeholder=" 제목을 입력하세요">
    <input v-model="author" style="width: 186px; height: 20px; left: 141px; top: 200px; position: absolute; text-align: right; font-size: 18px; font-weight: 600; word-wrap: break-word; border: none;" placeholder=" 저자를 입력하세요">
    <input v-model="publisher" style="width: 186px; height: 20px; left: 141px; top: 243px; position: absolute; text-align: right; font-size: 18px; font-family: Inter; font-weight: 600; word-wrap: break-word; border: none;" placeholder=" 출판사를 입력하세요">
    <div class="Line1"></div>
    <div style="width: 317px; height: 409px; left: 21px; top: 294px; position: absolute; background: white; box-shadow: 1px 2px 12px rgba(0, 0, 0, 0.25) inset; border-radius: 19px">
        <textarea v-model="message" style="width: 290px; height: 380px; left: 10px; top:10px; position: absolute; background: white; border: none; resize: none;" placeholder=" 내용을 입력해주세요"></textarea>
    </div>
    <button @click="modalOpen = true" style="width: 317px; height: 49px; left: 21px; top: 720px; position: absolute; background: #507BBC; border-radius: 30px; border:none;"></button>
    <div style="width: 83px; height: 25px; left: 141px; top: 732px; position: absolute; text-align: center; color: white; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">제출</div>
        <div class="black-bg" v-if="modalOpen === true">
      <div class="white-bg">
        <h3>제출하시겠습니까?</h3>
        <button @click="modalOpen = false" style=" width: 40%; height: 40px; text-align: center; color: #507BBC; font-Size: 15px; fontFamily: Inter; position: relative; border:none; border-radius: 30px; " class="modal-exit-btn">
          취소
        </button>
        <button @click="submitReport" style="width: 40%; height: 40px; text-align: center; font-Size: 15px; font-Family: Inter; word-Wrap: break-word; position: relative; background: #507BBC; border-radius: 30px; border:none; text-decoration: none; color: white;"  > 확인 </button>
        </div>
        </div>
      
</div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
    // 모달을 숨겨 주는 변수선언
    modalOpen: false,
    title: '',
    writter: '',
    publisher: '',
    message: '',
    };
},
methods: {
    // 보고서 제출 함수
    submitReport() {
      // 데이터를 서버로 전송
      const reportData = {
        title: this.title,
        author: this.author,
        publisher: this.publisher,
        message: this.message,
        approval: false,         
      };
            // Axios를 사용한 POST 요청
        axios.post('/api/submit-report', reportData) //여기에 백엔드 api 넣으면 된다
        .then(response => {
          console.log(response.data); // 성공적으로 전송되면 콘솔에 로그 출력
          this.$router.push('/student/main'); // 모달 열기 또는 다른 동작 수행
        })
        .catch(error => {
          console.error('Error submitting report:', error);
          // 오류 처리 또는 사용자에게 오류 메시지 표시
        });
    },
  },
};

</script>

<style>
@font-face{
  font-family: 'nanumgothic';
  src: url('@/assets/fonts/NanumGothic.ttf');
}
.class1
{
 width: 360px;
 height: 400px;
 position: relative;
 background: white;
}
.placeholder
{
color: #B4B4B4;
} 
.black-bg {
  display: flex;
  align-items: center;
  width: 360px;
  height: 800px;
  top: -10px;
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
}

</style>