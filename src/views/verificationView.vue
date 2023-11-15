<template>
    <div class="main">
      <div className="badge_box">
        <div className="badge" style="width: 78px; height: 90px; left: 52px; top: 424px; position: absolute; background: #D9D9D9; border-radius: 6px"></div>
        <div style="width: 83px; height: 25px; left: 48px; top: 523px; position: absolute; text-align: center; color: #4D4D4D; font-size: 16px; font-family: Inter; font-weight: 700; word-wrap: break-word">배지 이름</div>
      </div>
       <div className="Rectangle3">
        <div className="Rectangle4"></div>
        <div className="Rectangle5"></div>
        <div style="width: 135px; height: 28px; left: 133px; top: 187px; position: absolute; color: black; font-size: 22px; font-family: Inter; font-weight: 700; word-wrap: break-word">이름</div>
        <div style="width: 115px; height: 20px; left: 141px; top: 226px; position: absolute; color: #686868; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">학번</div>
        <div style="width: 115px; height: 20px; left: 141px; top: 259px; position: absolute; color: #686868; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">학과</div>
        <button @click="modalOpen = true" style="width: 83px; height: 25px; left: 247px; top: 131px; position: absolute; text-align: center; color: #507BBC; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">정보 수정</button>
        <div class="black-bg" v-if="modalOpen === true">
        <div class="white-bg">
            <div class="container">
	        	<form @submit.prevent="submitForm">
			    <div>
			    	<label for="email">email</label>
				    <input type="text" id="email" v-model="email" />
			    </div>
			    <div>
				    <label for="name">성함</label>
				    <input type="text" id="name" v-model="name" />
			    </div>
			    <div>
				    <label for="password">비밀번호</label>
				    <input type="password" id="password" v-model="password" />
			    </div>
			    <div>
				    <label for="passwordConfirm">비밀번호 확인</label>
				    <input type="password" id="passwordConfirm" v-model="passwordConfirm" />
			    </div>
			    <div>
				    <label for="position">직책</label>
				    <input type="text" id="position" v-model="position" />
			    </div>
			    <div>
				    <label for="teamName">팀 이름</label>
				    <input type="text" id="teamName" v-model="teamName" />
			    </div>
			        <!-- <button type="submit">회원가입</button> -->
	    	</form>
	</div>
          <button @click="modalOpen = false" style=" text-align: center, color: #507BBC, fontSize: 15, fontFamily: Inter, fontWeight: 700; wordWrap: break-word" class="modal-exit-btn">
            돌아가기
          </button>
        </div>
      </div>
      </div>
      <img class="PknuLogo1" :src="require('@/assets/pknu_logo.png')"/>
      <router-link to="/bogoseo" class="bogoseobtn" >
       <div style="width: 83px; height: 25px; left: 141px; top: 738px; position: absolute; color: white; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">활동 보고</div>
      </router-link>
     </div>
  </template>
  
  <script>
  export default {
  data() {
    return {
      // 모달을 숨겨 주는 변수선언
      modalOpen: false
    };
  },
    name: 'SignupForm',
	data() {
		return {
			email: '',
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
    width: 146px;
    height: 28px;
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
  .bogoseobtn{
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