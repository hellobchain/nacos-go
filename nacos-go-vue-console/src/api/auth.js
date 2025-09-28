import request from '@/utils/request'
export const login = data => request({ url: '/v1/auth/login', method: 'post', data })

export const getUserInfo = () => request({ url: '/v1/auth/user' })
export const changePassword = data => request({ url: '/v1/auth/user', method: 'post', data })

// 新增
export const addUser = data => request({ url: '/v1/auth/user/register', method: 'post', data })
export const updateUser = data => request({ url: '/v1/auth/user/update', method: 'post', data })
export const deleteUser = data => request({ url: '/v1/auth/user/delete?username='+data, method: 'delete' })
export const getUserList = () => request({ url: '/v1/auth/user/list'})
export const logout = () => Promise.resolve()