<!-- 정보 변경 화면 -->
<!-- 학번이랑 (학생 / 관리자)만 입력 받으면 됨 & 이름이랑 이메일은 구글에서 데이터 받아와서 출력하는 걸 할 수 있으면 좋을 듯 -->

<template>
  <div className="AndroidLarge2">
    <router-link to="/" style=" left: -115px; top: -15px; position: relative; text-align: center; background: #ffffff; border:none;" class="modal-exit-btn">
          <h4 style="color: #507BBC; font-weight: 300; word-wrap: break-word; font-Family: Inter;" > 돌아가기 </h4>
        </router-link>
        <h3 style="position: relative; top: -68px;"> 정보 수정 </h3>

    <div style="position: relative; top:-50px">
        <label><input type="radio" v-bind:value="radioValue1" v-model="picked"> 학생 </label>   
        <label><input type="radio" v-bind:value="radioValue2" v-model="picked"> 관리자 </label>

        </div>
        <form v-if="picked === '학생'" @submit.prevent="submitForm1">
        <div>
            <!-- 여기는 학생 폼 -->
          <label for="name" style="width: 100%; height: 100%; left: -18px; top: -40px; font-weight: 700; position: relative;" >학생 이름</label>
          <input type="text" id="name" style=" width: 194px; height: 38px; left: 15px; top: -40px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none; text-align: right;" v-model="name" />
        </div>
        <div>
          <label for="email" style="width: 100%; height: 100%; left: -18px; top: -30px; font-weight: 700; position: relative;">이메일 </label>
          <input type="email" id="email" style="width: 194px; height: 38px; left: 15px; top: -30px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="email" />
        </div>
        <div>
          <label for="roleId" style="width: 100%; height: 100%; left: -18px; top: -20px; font-weight: 700; position: relative;">학번 </label>
          <input type="number" id="roleId" style="width: 194px; height: 38px; left: 15px; top: -20px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="roleId" />
        </div>
        <div>
          <label for="major" style="width: 100%; height: 100%; left: 10px; top: -25px; font-weight: 700; position: relative;">학과 </label>
          <input type="text" id="major" style="width: 255px; height: 38px; left: -10px; top: 10px; position: relative; margin-top:10px; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="major" />
        </div>
          <button type="submit" class="submit_btn"> 저장 </button>
          <router-link to="/student/main" class="ok_btn"> 확인 </router-link>     
        </form>

        <!-- 관리자 폼 -->
        <form v-else-if="picked === '관리자'" @submit.prevent="submitForm2">
        <div>       
          <label for="name" style="width: 100%; height: 100%; left: -18px; top: -40px; font-weight: 700; position: relative;" >관리자 이름</label>
          <input type="text" id="name" style=" width: 194px; height: 38px; left: 15px; top: -40px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none; text-align: right;" v-model="name" />
        </div>
        <div>
          <label for="email" style="width: 100%; height: 100%; left: -18px; top: -30px; font-weight: 700; position: relative;">이메일 </label>
          <input type="email" id="email" style="width: 194px; height: 38px; left: 15px; top: -30px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="email" />
        </div>
        <div>
          <label for="roleId" style="width: 100%; height: 100%; left: -18px; top: -20px; font-weight: 700; position: relative;">사번 </label>
          <input type="number" id="roleId" style="width: 194px; height: 38px; left: 15px; top: -20px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="roleId" />
        </div>
          <button type="submit" class="submit_btn"> 저장 </button> 
          <div>
          <router-link to="/admin/main" class="ok_btn"> 확인 </router-link>     
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
.AndroidLarge2
{width: 360px;
height: 800px;
background: white;
}
.submit_btn
{ width: 70px;
  height:40px;
  top: 20px;
  left: 0px;
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
  top: 20px;
  left: 10px;
  position: relative;
  background: rgb(248, 248, 248);
  color: #507BBC;
  border-radius: 20px; 
  border-color: #507BBC;
  border:3px;
  text-align: center;
  text-decoration: none;
}
</style>