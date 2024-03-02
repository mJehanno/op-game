import {createApp} from 'vue'
import Home from '@/pages/Home.vue';
import Game from '@/pages/Game.vue';
import ScoreBoard from '@/pages/Ladder.vue';
import Multer from '@/pages/Multer.vue';
import App from './App.vue';
import {createRouter, createWebHashHistory} from 'vue-router';
import '/node_modules/primeflex/primeflex.css';
import './style.css';
import PrimeVue from "primevue/config";
import ToastService from 'primevue/toastservice';
import 'primevue/resources/themes/aura-light-green/theme.css'
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';    
const routes =  [
  { path: "/", component: Home},
  { 
    path: "/game", 
    component: Game, 
    children: [{ path: "mult/:level", component: Multer}]
  },
  { path: "/scores", component: ScoreBoard},
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

const app  = createApp(App);
app.use(PrimeVue);
app.use(router)
app.use(ToastService);

app.mount('#app');
