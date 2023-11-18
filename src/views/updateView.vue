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
            <!-- 이름은 구글에서 가져오도록 수정해야 함 -->
          <label for="name" style="width: 100%; height: 100%; left: -18px; top: -40px; font-weight: 700; position: relative;" >학생 이름</label>
          <input type="text" id="name" style=" width: 194px; height: 38px; left: 15px; top: -40px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none; text-align: right;" v-model="name" />
        </div>
        <div>
          <label for="email" style="width: 100%; height: 100%; left: -18px; top: -30px; font-weight: 700; position: relative;">이메일 </label>
          <input type="email" id="email" style="width: 194px; height: 38px; left: 15px; top: -30px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="email" />
        </div>
        <div>
          <label for="student_num" style="width: 100%; height: 100%; left: -18px; top: -20px; font-weight: 700; position: relative;">학번 </label>
          <input type="number" id="student_num" style="width: 194px; height: 38px; left: 15px; top: -20px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="student_num" />
        </div>
        <div>
          <label for="major" style="width: 100%; height: 100%; left: 10px; top: -25px; font-weight: 700; position: relative;">학과 </label>
          <input type="text" id="major" style="width: 255px; height: 38px; left: -10px; top: 10px; position: relative; margin-top:10px; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="major" />
        </div>
          <!-- 밑에 버튼 누르면 데이터 베이스로 데이터 보내지게 해야 함 & 메인 화면으로 감 -->
          <router-link to="/student/main" style=" width: 70px; height:40px; top: 20px; position: relative; background: #507BBC; border-radius: 20px; border:none; text-align: center; ">
          <h4 style=" font-color : white; text-align: center; ">확인</h4>
        </router-link>
        </form>
        <form v-else-if="picked === '관리자'" @submit.prevent="submitForm2">
        <div>
            <!-- 이름은 구글에서 가져오도록 수정해야 함 & 관리자 입력 정보 변경 -->
          <label for="name" style="width: 100%; height: 100%; left: -18px; top: -40px; font-weight: 700; position: relative;" >관리자 이름</label>
          <input type="text" id="name" style=" width: 194px; height: 38px; left: 15px; top: -40px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none; text-align: right;" v-model="name" />
        </div>
        <div>
          <label for="email" style="width: 100%; height: 100%; left: -18px; top: -30px; font-weight: 700; position: relative;">이메일 </label>
          <input type="email" id="email" style="width: 194px; height: 38px; left: 15px; top: -30px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="email" />
        </div>
        <div>
          <label for="student_num" style="width: 100%; height: 100%; left: -18px; top: -20px; font-weight: 700; position: relative;">사번 </label>
          <input type="number" id="student_num" style="width: 194px; height: 38px; left: 15px; top: -20px; position: relative; background: #EFEFEF; box-shadow: 0px 3px 7px rgba(0, 0, 0, 0.25) inset; border-radius: 11px; border: none;" v-model="student_num" />
        </div>
          <!-- 밑에 버튼 누르면 데이터 베이스로 데이터 보내지게 해야 함 & 메인 화면으로 감 -->
          <router-link to="/admin/main" style=" width: 70px; height:40px; top: 20px; position: relative; background: #507BBC; border-radius: 20px; border:none; text-align: center; ">
          <h4 style=" font-color : white; text-align: center; ">확인</h4>
        </router-link>

        </form>


    </div>
</template>

<script>
    export default {
        data() {
            return {
                picked: '',
                radioValue1: '학생',
                radioValue2: '관리자'
            };
        },
    methods: {
    submitForm1() {
      // 학생 폼 제출 동작
      // console.log('dd'); 이런게 필요한가본데
    },
    submitForm2() {
      // 관리자 폼 제출 동작
    }
  }

}
</script>

<style>
.AndroidLarge2
{width: 360px;
height: 800px;
background: white;
}
</style>