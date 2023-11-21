<!-- 정보 변경 화면 -->
<!-- 학번이랑 (학생 / 관리자)만 입력 받으면 됨 & 이름이랑 이메일은 구글에서 데이터 받아와서 출력하는 걸 할 수 있으면 좋을 듯 -->

<template>
  <div className="AndroidLarge2">
    <h3 style="position: relative; top: 50px;"> 정보 수정 </h3>
    <router-link to="/" style=" left: -122px; top: 40px; position: relative; text-align: center; background: #ffffff; border:none; text-decoration: none;" class="modal-exit-btn">
          <h4 style="color: #507BBC; font-weight: 700; word-wrap: break-word; font-Family: nanumgothic;" > 돌아가기 </h4>          
        </router-link>
    <div class="divider"></div>
    <div style="position: relative; top:70px">
        <label><input type="radio" v-bind:value="radioValue1" v-model="picked"> 학생 </label>   
        <label><input type="radio" v-bind:value="radioValue2" v-model="picked"> 관리자 </label>

        </div>
        <form v-if="picked === '학생'" @submit.prevent="submitForm1">
        <div>
            <!-- 여기는 학생 폼 -->
          <label for="name" style="width: 100%; height: 100%; left: -22px; top: 80px; font-weight: 700; position: relative; color:#173970;" >학생 이름</label>
          <input type="text" id="name" style=" width: 194px; height: 38px; left: 21px; top: 80px; position: relative; background: #FFFFFF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 30px; border: none; text-align: right;" v-model="name" />
        </div>
        <div>
          <label for="email" style="width: 100%; height: 100%; left: -30px; top: 95px; font-weight: 700; position: relative; color:#173970;">이메일 </label>
          <input type="email" id="email" style="width: 194px; height: 38px; left: 26px; top: 95px; position: relative; background: FFFFFF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 30px; border: none;" v-model="email" />
        </div>
        <div>
          <label for="roleId" style="width: 100%; height: 100%; left: -35px; top: 110px; font-weight: 700; position: relative; color:#173970;">학번 </label>
          <input type="number" id="roleId" style="width: 194px; height: 38px; left: 33px; top: 110px; position: relative; background: #FFFFFF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 30px; border: none;" v-model="roleId" />
        </div>
          <button type="submit" class="submit_btn"> 저장 </button>
          <div class="ok_btn" >
          <router-link to="/student/main" style="word-wrap: break-word; color: #507BBC; text-decoration: none;"> 확인 </router-link>
        </div>  
        </form>

        <!-- 관리자 폼 -->
        <form v-else-if="picked === '관리자'" @submit.prevent="submitForm2">
        <div>       
          <label for="name" style="width: 100%; height: 100%; left: -13px; top: 80px; font-weight: 700; position: relative; color:#173970;" >관리자 이름</label>
          <input type="text" id="name" style=" width: 194px; height: 38px; left: 13px; top: 80px; position: relative; background: #ffffff; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 30px; border: none; text-align: right;" v-model="name" />
        </div>
        <div>
          <label for="email" style="width: 100%; height: 100%; left: -30px; top: 95px; font-weight: 700; position: relative; color:#173970;">이메일 </label>
          <input type="email" id="email" style="width: 194px; height: 38px; left: 26px; top: 95px; position: relative; background: #ffffff; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 30px; border: none;" v-model="email" />
        </div>
        <div>
          <label for="roleId" style="width: 100%; height: 100%; left: -35px; top: 110px; font-weight: 700; position: relative; color:#173970;">사번 </label>
          <input type="number" id="roleId" style="width: 194px; height: 38px; left: 33px; top: 110px; position: relative; background: #ffffff; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 30px; border: none;" v-model="roleId" />
        </div>
          <button type="submit" class="submit_btn"> 저장 </button> 
          <div class="ok_btn" >
          <router-link to="/admin/main" style="word-wrap: break-word; color: #507BBC; text-decoration: none;"> 확인 </router-link>     
          </div>
        </form>
    </div>
</template>

<script>
import axios from 'axios';

    export default {
        data() {
            return {
                picked: '',
                radioValue1: '학생',
                radioValue2: '관리자',
                name: '',
                email: '',
                roleId: '',
                major: '',
            };
        },
    created() {
    // Vue 컴포넌트가 생성될 때 서버로부터 데이터를 불러옴
    this.loadData();
  },
    methods: {
      loadData() {
      // 서버로부터 데이터를 불러오는 로직
      axios.get('/api/member') // 예시 경로, 실제로 사용하는 경로에 맞게 변경
        .then(response => {
          const userData = response.data;
          this.name = userData.name;
          this.email = userData.email;
          // 다른 필요한 데이터들도 동일하게 처리
        })
        .catch(error => {
          console.error('Error loading data:', error);
        });
    },
    submitForm1() {
       // 학생 폼 제출 동작
      const formData = {
        role: this.picked,
        name: this.name,
        email: this.email,
        roleId: this.roleId,
        major: this.major,
      };
      this.sendFormData(formData, '/api/student/update'); // 여기 엔드포인트 입력훼
    },
    submitForm2() {
      // 관리자 폼 제출 동작
      const formData = {
        role: this.picked,
        name: this.name,
        email: this.email,
        roleId: this.roleId,
        // Add other fields for the admin form as needed
      };
      this.sendFormData(formData, '/api/admin/update'); // 여기 엔드포인트 입력훼
    },
    sendFormData(formData, endpoint) {
      axios.post(endpoint, formData)
        .then(response => {
          console.log(response.data); // 성공적으로 전송되면 콘솔에 로그 출력
          // 이동 또는 다른 동작 수행
        })
        .catch(error => {
          console.error('Error submitting form data:', error);
          // 오류 처리 또는 사용자에게 오류 메시지 표시
        });
    }
  }
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
.submit_btn
{ width: 70px;
  height:40px;
  top: 140px;
  left: -30px;
  position: relative;
  background: #507BBC;
  color: white;
  border-radius: 20px;
  border:none;
  text-align: center;
}
.ok_btn
{ width: 70px;
  height:40px;
  top: 100px;
  left: 190px;
  position: relative;
  background: rgb(237, 237, 237);
  color: #507BBC;
  border-radius: 20px; 
  border:1px #507BBC;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: center;
}
.divider {
  width: 100%;
  height: 1px;
  background-color: #D9D9D9; /* 선의 색상 설정 */
  position: absolute;
  top: 125px; /* 선의 위치 조정 */
}
</style>