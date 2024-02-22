import {createApp} from 'vue'
import Home from '@/components/Home.vue';
import Game from '@/components/Game.vue';
import ScoreBoard from '@/components/ScoreBoard.vue';
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
  { path: "/game", component: Game},
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
