import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)
const routes = [
    { path: '/login', name: 'Login', component: () => import('@/components/Login.vue') },
    {
        path: '/',
        component: () => import('@/components/Layout.vue'),
        redirect: '/tenant',
        children: [
            { path: 'tenant', component: () => import('@/components/Tenant.vue') },
            { path: 'service', component: () => import('@/components/Service.vue') },
            { path: 'config', component: () => import('@/components/Config.vue') },
            { path: 'publish', component: () => import('@/components/Publish.vue') }
        ]
    }
]
export default new VueRouter({ mode: 'history', routes })