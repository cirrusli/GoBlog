<template>
    <div style="display: flex;">
        
    <div id="pic" style="flex: left;">
      <img src="../pic/企鹅.png" style="flex: left;">
      <img src="../pic/企鹅.png" style="flex: left;">
      <img src="../pic/企鹅.png" style="flex: left;">
      </div>
      <div style="margin-top: 27px;margin-left: 280px;flex: right; " id="search" >
        <input type="text" v-model="search">
        <a id="search1" @click='into'><img src="../pic/搜索 (8).png"></a>
      </div>
        <div style="position: absolute; margin-top: 15px;left: 65%;">
            <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal"  background-color="#d4e9f0" router>
                
                <el-menu-item><router-link to="/main">主页</router-link></el-menu-item>
                <el-submenu >
                  <template slot="title" >登录/注册</template>
<el-menu-item>
    <router-link to="/login">登录</router-link>


</el-menu-item>
<el-menu-item>
    <router-link to="/set">注册</router-link>
</el-menu-item>

</el-submenu>
<el-menu-item index="3">访问历史</el-menu-item>
<el-menu-item index="4"><a href="https://www.ele.me" target="_blank">个人中心</a></el-menu-item>
</el-menu>
</div>
<div style="position: absolute;margin-top: 67px;margin-left: 145px;">
    <span style="font-family:Arial,Helvetica,sans-serif;font-size:1em;color: rgba(9, 244, 166, 0.763);font-style: oblique;">三只松鼠Vblog</span></div>

</div>

</template>
<script>
    export default {
        name: 'FirstPage',
        data() {
            return {
                activeIndex: '1',
                activeIndex2: '1',
                search: ''
            }
        },
        created() {},
        computed: {},
        methods: {
            handleSelect(key, keyPath) {
                console.log(key, keyPath);
            },
            // into: function() {
            //   $.get('http')
            // }
        },
    }
</script>
<style>
    #pic {
        margin-left: 100px;
        margin-top: 5px;
    }
    
    input {
        background: none;
        border: black solid 3px;
        float: left;
        border-radius: 28px;
        font-size: 8px;
        text-align: center;
        padding: 8px;
        color: rgb(44, 44, 44);
    }
    
    #search1 {
        vertical-align: middle;
        margin-left: 10px;
        float: right;
        background-color: rgb(44, 44, 44);
        height: 35px;
        width: 40px;
        border-radius: 50%;
        text-align: center;
        line-height: 25px;
        color: #fff;
    }
    
    #search:hover>a {
        background: rgb(5, 94, 177);
        cursor: pointer;
    }
</style>