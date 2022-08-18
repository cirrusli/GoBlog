import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import jquery from 'jquery';

import router from '../src/router/index.js'
Vue.prototype.$ = jquery //全局引入jquery;


Vue.use(router)




Vue.use(ElementUI);

new Vue({
    el: '#app',
    render: h => h(App),
    router,
});

Vue.config.productionTip = false

new Vue({
    render: h => h(App),
}).$mount('#app')