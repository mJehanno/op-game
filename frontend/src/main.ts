import {createApp} from 'vue'
import Home from '@/pages/Home.vue';
import Game from '@/pages/Game.vue';
import NotFound from '@/pages/NotFound.vue';
import App from './App.vue';
import {createRouter, createWebHashHistory} from 'vue-router';
import '/node_modules/primeflex/primeflex.css';
import './style.css';
import PrimeVue from "primevue/config";
import ToastService from 'primevue/toastservice';
import Tooltip from 'primevue/tooltip';
import 'primevue/resources/themes/aura-light-green/theme.css'
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';    
import { ScoreManagerService } from './plugins/ScoreManager';
import { createPinia } from 'pinia'
import Ladder from '@/pages/Ladder.vue';

const routes =  [
  { path: "/", component: Home },
  { 
    path: "/game", 
    component: Game, 
  },
  { path: "/scores", component: Ladder},
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

const pinia = createPinia();
const app  = createApp(App);

//plugins
app.use(PrimeVue);
app.use(router);
app.use(ToastService);
app.use(ScoreManagerService);
app.use(pinia);

// directives
app.directive('tooltip', Tooltip);

app.mount('#app');
