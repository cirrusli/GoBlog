import Vue from 'vue';
import Router from 'vue-router';
import LoginPage from '../login.vue';
import SetPage from '../SetPage.vue';
import MainPage from '../MainPage.vue'
Vue.use(Router);
export default new Router({

    routes: [ //用来配置路由的相关信息
        {

            path: '/login', //虚拟路径用来设置浏览器中的路径
            title: '登录',
            component: LoginPage,
        },
        {
            path: '/set',
            title: '注册',
            component: SetPage,
        },
        {
            path: '/main',
            title: '主页',
            component: MainPage
        }
    ]
})