import request from '@/utils/request'
export const getInstances = () => request({ url: '/v1/ns/instance/list' })
export const deregister = (ip, port) => request({ url: `/v1/ns/instance?ip=${ip}&port=${port}`, method: 'delete' })