<template>
    <div >
        <div id="circle"></div>
        <div><img src="../pic/程序员.jpeg" id="pic1"></div>
        <div id="main1" style="width: 500px;height: 120px;font-style: normal;font-size: larger;font-weight: 550;color: #07e4fc;"> 
           <div style="text-align: center;margin-top: 18px;">  Study make the world better </div>
           <div style="text-align: center;padding-top: 8px;">  学习让世界变得更加美好 </div>
           <div style="text-align: center;padding-top: 8px;">  Apprendre à rendre le monde meilleur </div>
        </div>
        <div style="display: flex;" id="icon">
            <a href="https://im.qq.com/index" target="_blank"><img src="../pic/icon_QQ.png"></a>
            <a href="https://weixin.qq.com/"><img src="../pic/微信.png"></a>
            <a href="https://gitee.com/"><img src="../pic/gitee.png"></a>
            <a href="https://weibo.com/"><img src="../pic/微博 (1).png"></a>
        </div>

    </div>
</template>
<script>
    export default {
        name: 'MainPage',
        data() {
            return {}
        },
        created() {},
        computed: {},
        methods: {},
    }
</script>
<style>
    #main1 {
        position: absolute;
        border: #4599ab solid 5px;
        left: 26%;
        top: 43%;
        border-radius: 30px;
    }
    
    #pic1 {
        height: 80px;
        width: 80px;
        margin-top: 70px;
        margin-left: 540px;
        border-radius: 50%;
        border: #07e4fc solid 3px;
    }
    
    #icon {
        margin-top: 160px;
        margin-left: 496px;
    }
    
    #icon>a {
        padding-left: 10px;
    }
</style>