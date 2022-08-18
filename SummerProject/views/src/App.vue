<template>
    <div>
       <FirstPage></FirstPage>
       <router-view></router-view>
       <BottomPage id="bottom"></BottomPage>
      
    </div>
</template>
<script>
    import FirstPage from './FirstPage.vue';

    import BottomPage from './BottomPage.vue'


    export default {
        name: 'app',
        data() {
            return {}
        },
        components: {
            FirstPage,

            BottomPage,


        },

    }
</script>
<style>
    body {
        width: 1250px;
        height: 600px;
        background-image: url(../pic/背景.jpeg);
    }
    
    #bottom {
        margin-top: 580px;
    }
</style>