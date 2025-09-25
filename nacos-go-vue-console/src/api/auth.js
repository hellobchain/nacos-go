import request from '@/utils/request'
export const login = data => request({ url: '/v1/auth/login', method: 'post', data })
export const logout = () => Promise.resolve()