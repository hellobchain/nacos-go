import axios from 'axios'
import store from '@/store'
import { Notify } from '@/components/Notify'
import { logout } from '../api/auth'
const service = axios.create({ baseURL: process.env.VUE_APP_BASE_API, timeout: 8000 })
service.interceptors.request.use(cfg => {
    if (store.state.token) cfg.headers.Authorization = 'Bearer ' + store.state.token
    return cfg
})
service.interceptors.response.use(
    res => {
        const { code, message, data } = res.data
        if (code !== 200) {
            if (code === 401) {
                // 登录超时
                Notify.error("登录超时，请重新登录")
                store.commit('SET_TOKEN', '')
                window.location.href = '/login'
                return
            }
            Notify.error(message || 'Error')
            return Promise.reject(new Error(message))
        }
        return data
    },
    err => {
        Notify.error(err.message || 'Network Error')
        return Promise.reject(err)
    }
)
export default service