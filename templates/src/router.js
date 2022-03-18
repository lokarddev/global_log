import Vue from 'vue'
import VueRouter from 'vue-router'
import Dashboard from './pages/Dashboard.vue'
import Logs from "@/pages/Logs";
import Analytics from "@/pages/Analytics";

Vue.use(VueRouter)
const routes = [
    {path: '/', name: 'Home', component: Dashboard},
    {path: '/logs', name: 'Logs', component: Logs},
    {path: '/analytics', name: 'Analytics', component: Analytics},
]
const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})
export default router;
