<!-- 관리자 메인화면 -->

<template>
    <div class="main">  
       <div className="Rectangle3 ">
        <div className="Rectangle4"></div>
        <div className="Rectangle5"></div>
        <div style="width: 135px; height: 28px; left: 133px; top: 200px; position: absolute; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">이름</div>
        <div style="width: 115px; height: 20px; left: 172px; top: 234px; position: absolute; color: #686868; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">도서관 관리자</div>
        <div style="width: 115px; height: 20px; left: 140px; top: 264px; position: absolute; color: #686868; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">사번</div>
        <button @click="modalOpen = true" style="width: 83px; height: 25px; left: 247px; top: 131px; position: absolute; text-align: center; color: #507BBC; font-size: 13px; font-family: Inter; font-weight: 700; word-wrap: break-word">정보 수정</button>
        <div class="black-bg" v-if="modalOpen === true">
        <div class="white-bg">
            <div class="container">
            <form @submit.prevent="submitForm">
          <div>
            <label for="name">이름  </label>
            <input type="text" id="name" v-model="name" />
          </div>
          <div>
            <label for="admin_number">사번 </label>
            <input type="number" id="admin_number" v-model="grade" />
          </div>
        </form>
  </div>
          <button @click="modalOpen = false" style="  text-align: center, color: #507BBC, fontSize: 15, fontFamily: Inter, fontWeight: 700; wordWrap: break-word" class="modal-exit-btn">
            취소
          </button>
          <button @click="modalOpen = false" style=" text-align: center, color: #507BBC, fontSize: 15, fontFamily: Inter, fontWeight: 700; wordWrap: break-word" class="modal-exit-btn">
            제출하기
          </button>
        </div>
      </div>
      </div>
      <img class="PknuLogo1" :src="require('@/assets/pknu_logo.png')"/>
      <div class="Group1" style="width: 100%; height: 100%; position: relative">
    <div class="Rectangle10" style="width: 317px; height: 442px; left: 0px; top: 0px; position: absolute; background: white; box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.25); border-radius: 19px"></div>
    <div class="Rectangle17" style="width: 260px; height: 50px; left: 20px; top: 58px; position: absolute; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px"></div>
    <div class="Oo20xxxxxxx" style="width: 193px; height: 25px; left: 0px; top: 33px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 500; word-wrap: break-word">김OO | 20XXXXXXX</div>
    <div style="width: 170px; height: 25px; left: 20px; top: 75px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">어린왕자를 읽고</div>
    <div class="Rectangle18" style="width: 260px; height: 50px; left: 20px; top: 167px; position: absolute; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px"></div>
    <div class="Oo20xxxxxxx" style="width: 193px; height: 25px; left: 0px; top: 142px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 500; word-wrap: break-word">이OO | 20XXXXXXX</div>
    <div style="width: 170px; height: 25px; left: 20px; top: 184px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">백설공주를 읽고</div>
    <div class="Rectangle19" style="width: 260px; height: 50px; left: 20px; top: 276px; position: absolute; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25); border-radius: 11px"></div>
    <div class="Oo20xxxxxxx" style="width: 193px; height: 25px; left: 0px; top: 251px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 500; word-wrap: break-word">박OO | 20XXXXXXX</div>
    <div style="width: 170px; height: 25px; left: 20px; top: 292px; position: absolute; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">1984를 읽고</div>
    <div class="Rectangle20" style="width: 13px; height: 355px; left: 304px; top: 40px; position: absolute; background: #EEEEEE; box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25) inset; border-radius: 15px"></div>
    <div class="Rectangle21" style="width: 13px; height: 44px; left: 304px; top: 61px; position: absolute; background: #507BBC; border-radius: 17px"></div>
</div>
     </div>
  </template>
  
  
  <script>
  export default {
    name: 'SignupForm',
    data() {
      return {
      // 모달을 숨겨 주는 변수선언
      modalOpen: false,
      name: '',
      password: '',
      passwordConfirm: '',
      position: '',
      teamName: '',
      };
  },
  methods: {
    submitForm() {
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
    top: -20px; 
    position: relative; 
    background: #507BBC;
  }
  .Rectangle4{
    width: 308px; 
    height: 221px; 
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
  
  .white-bg {
    width: 100%;
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
  .group1{
    width: 317px;
    height: 442px;
    left: 20px;
    top: 321px;
    position: relative;
  }

  
  </style>