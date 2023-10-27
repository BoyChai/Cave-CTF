import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// element-plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// axios
import axios from 'axios';
axios.defaults.baseURL = 'http://127.0.0.1:8080';

// const v = createApp(App).use(router).use(ElementPlus).mount('#app')
const app = createApp(App);
app.use(router);
app.use(ElementPlus);
app.mount('#app');