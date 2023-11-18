<!-- 배지 상세 정보 화면 -->

<template>
    <div className="AndroidLarge2">
    <div class="2" style="width: 100%; height: 100%; padding-top: 51px; padding-bottom: 110px; padding-left: 24px; padding-right: 24px; background: white; flex-direction: column; justify-content: flex-start; align-items: center; gap: 81px; display: inline-flex">
        <div style="width: 129px; height: 43px; text-align: center; color: black; font-size: 25px; font-family: Inter; font-weight: 700; word-wrap: break-word">배지 정보</div>
        <div class="Popup" style="width: 312px; height: 515px; position: relative">
            <div class="Bg" style="width: 312px; height: 447px; left: 0px; top: 0px; position: absolute; background: white; box-shadow: 4px 17px 82px #A4BFE9; border-radius: 8px"></div>
            <router-link to="/verify" class="ButtonPrimary" style="width: 312px; height: 50px; padding-left: 32px; padding-right: 32px; padding-top: 16px; padding-bottom: 16px; left: 0px; top: 465px; position: absolute; background: #507BBC; border-radius: 20px; flex-direction: column; justify-content: flex-start; align-items: center; gap: 12px; display: inline-flex">
               <div class="prove" style="text-align: center; color: white; font-size: 16px; font-family: Inter; font-weight: 600; word-wrap: break-word">검증</div>
            </router-link>
            <div style="width: 95.36px; height: 26px; left: 108.32px; top: 187px; position: absolute; text-align: center; color: black; font-size: 20px; font-family: Inter; font-weight: 700; word-wrap: break-word">배지 이름</div>
            <div class="Line2" style="width: 278.67px; height: 0px; left: 13.89px; top: 280px; position: absolute; border: 1px #4D4D4D dotted"></div>
            <div class="Rectangle18" style="width: 97.21px; height: 109px; left: 106.47px; top: 62px; position: absolute; background: #D9D9D9"></div>
            <router-link to="/student/main" style="width: 76.84px; height: 25px; left: 10.18px; top: 14px; position: absolute; text-align: center; color: #507BBC; font-size: 15px; font-family: Inter; font-weight: 700; word-wrap: break-word">돌아가기</router-link>
            <div style="width: 150.91px; height: 20px; left: -15px; top: 303px; position: absolute; color: #4D4D4D; font-size: 14px; font-family: Inter; font-weight: 400; word-wrap: break-word">배지 발행일</div>
            <p style="width: 217.57px; height: 20px; left: 49.07px; top: 229px; position: absolute; text-align: center; color: #4D4D4D; font-size: 14px; font-family: Inter; font-weight: 400; word-wrap: break-word"> {{ $route.params.id }}</p>
            <div class="YyyyMmDd" style="width: 150.91px; height: 20px; left: 143.50px; top: 303px; position: absolute; text-align: right; color: black; font-size: 14px; font-family: Inter; font-weight: 600; word-wrap: break-word">YYYY.MM.DD</div>
            <div style="width: 150.91px; height: 20px; left: -5px; top: 332px; position: absolute; color: #4D4D4D; font-size: 14px; font-family: Inter; font-weight: 400; word-wrap: break-word">배지 유효 기간</div>
            <div class="YyyyMmDd" style="width: 150.91px; height: 20px; left: 142.57px; top: 332px; position: absolute; text-align: right; color: black; font-size: 14px; font-family: Inter; font-weight: 600; word-wrap: break-word">YYYY.MM.DD</div>
            <div style="width: 150.91px; height: 20px; left: -20px; top: 361px; position: absolute; color: #4D4D4D; font-size: 14px; font-family: Inter; font-weight: 400; word-wrap: break-word">배지 등급</div>
            <div class="Lv1" style="width: 150.91px; height: 20px; left: 142.57px; top: 361px; position: absolute; text-align: right; color: black; font-size: 14px; font-family: Inter; font-weight: 600; word-wrap: break-word">Lv.1</div>
            <div style="width: 150.91px; height: 20px; left: -15px; top: 390px; position: absolute; color: #4D4D4D; font-size: 14px; font-family: Inter; font-weight: 400; word-wrap: break-word">발행자 정보</div>
            <div style="width: 150.91px; height: 20px; left: 141.65px; top: 390px; position: absolute; text-align: right; color: black; font-size: 14px; font-family: Inter; font-weight: 600; word-wrap: break-word">부경대학교 도서관 관리자</div>
        </div>
    </div>
</div>
</template>

<script>
import badge_data from '@/assets/badge.js';

export default {
  data() {
    return {
      badgeList : badge_data,
    };    
}}    

</script>

<style>
/* .AndroidLarge2
{width: 360px;
height: 800px;
background: white;
position: relative;
} */
</style>