<template>
    <div>
        <div id="login">
            <div class="login-box">
                <h2>注册系统</h2>
                <form>
                  <div class="user-box">
                    <input type="text" name="" required="" v-model="username">
                    <label>邮箱号</label>
                  </div>
                  <div class="user-box">
                    <input type="password" name="" required="" v-model="password">
                    <label>密码</label>
                  </div>
                  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                  <a @click="handleSubmit">
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                    注册
                  </a>
                 
                </form>
              </div>
         
        </div>
    </div>
</template>
<script>
    export default {
        name: 'SetPage',
        data() {
            return {}
        },
        created() {},
        computed: {},
        methods: {},
    }
</script>
<style>

</style>