import {createApp} from 'vue'
import App from './App.vue'
import '/node_modules/primeflex/primeflex.css';
import './style.css';
import PrimeVue from "primevue/config";
import 'primevue/resources/themes/aura-light-green/theme.css'
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';    
import FocusTrap from 'primevue/focustrap';

const app  = createApp(App);
app.use(PrimeVue);
app.directive('focustrap', FocusTrap);


app.mount('#app');
