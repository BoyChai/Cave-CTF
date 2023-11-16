import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// element-plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// element-plus/icons-vue
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

// axios
import axios from 'axios';
axios.defaults.baseURL = 'http://server.boychai.xyz:8002';

// const v = createApp(App).use(router).use(ElementPlus).mount('#app')
const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(router);
app.use(ElementPlus);
app.mount('#app');