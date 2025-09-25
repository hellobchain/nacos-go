import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
export default new Vuex.Store({
    state: { token: localStorage.getItem('token') || '' },
    mutations: {
        SET_TOKEN: (s, v) => {
            s.token = v
            v ? localStorage.setItem('token', v) : localStorage.removeItem('token')
        }
    }
})