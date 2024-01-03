<template>
<div class="2" style="width: 100%; height: 100%; position: relative; background: white">
  <div class="Verified" style="width: 109px; height: 37px; left: 125px; top: 426px; position: absolute; text-align: center; color: black; font-size: 25px; font-family: Inter; font-weight: 700; word-wrap: break-word">verified</div>
  <div class="Valid">Your badge is valid !</div>
  <img class="check" :src="require('@/assets/check.png')"/>
</div>
</template>


<style>
.valid{
width: 100%; 
text-align: center; 
color: #27214D; 
font-size: 20px; 
font-family: Alata; 
font-weight: 400; 
line-height: 30px;
word-wrap: break-word;
left:100px;
top:255px;
position: absolute;


}
.check{
    position: absolute;
    left:112px;
    top:265px;
}
</style>