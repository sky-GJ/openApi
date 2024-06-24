// import './assets/main.css'
//import './assets/base.css'
import './assets/theme.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ArcoVue from '@arco-design/web-vue';
import App from './App.vue'
import router from './router'
import '@arco-design/web-vue/dist/arco.css';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(ArcoVue);
app.use(ArcoVueIcon);

app.mount('#app');