<!-- 보고서 작성 페이지 -->
<template>
<div class="1" >
    <router-link to="/about" style="width: 83px; height: 25px; left: 7px; top: 92px; position: absolute; text-align: center; color: #507BBC; font-size: 15px; font-family: Helvetica; font-weight: 700; word-wrap: break-word">돌아가기</router-link>
    <div style="width: 129px; height: 43px; left: 115px; top: 54px; position: absolute; text-align: center; color: black; font-size: 25px; font-family: Inter; font-weight: 700; word-wrap: break-word">활동 보고</div>
    <div style="width: 83px; height: 20px; left: 24px; top: 157px; position: absolute; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">제목</div>
    <div style="width: 83px; height: 20px; left: 24px; top: 200px; position: absolute; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">저자</div>
    <div style="width: 83px; height: 20px; left: 32px; top: 243px; position: absolute; color: black; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">출판사</div>
    <div style="width: 163px; height: 20px; left: 164px; top: 157px; position: absolute; text-align: right; color: #B4B4B4; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">제목을 입력하세요</div>
    <div style="width: 152px; height: 20px; left: 175px; top: 200px; position: absolute; text-align: right; color: #B4B4B4; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">저자를 입력하세요</div>
    <div style="width: 186px; height: 20px; left: 141px; top: 243px; position: absolute; text-align: right; color: #B4B4B4; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">출판사를 입력하세요</div>
    <div class="Rectangle11" style="width: 317px; height: 49px; left: 21px; top: 720px; position: absolute; background: #507BBC; border-radius: 11px"></div>
    <div style="width: 83px; height: 25px; left: 141px; top: 732px; position: absolute; text-align: center; color: white; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">제출</div>
    <div class="Line1"></div>
    <!-- </div> -->
    <textarea v-model="message" style="width: 317px; height: 409px; left: 21px; top: 294px; position: absolute; background: white; box-shadow: 1px 2px 12px rgba(0, 0, 0, 0.25) inset; border-radius: 19px" placeholder="여러줄을 입력해보세요"></textarea>
    <!-- <div style="width: 186px; height: 20px; left: 41px; top: 313px; position: absolute; color: #B4B4B4; font-size: 18px; font-family: Inter; font-weight: 700; word-wrap: break-word">내용을 입력하세요</div> -->
</div>
</template>

<script>
</script>

<style>
.class1
{
 width: 360px;
 height: 800px;
 position: relative;
 background: white;
}
</style>