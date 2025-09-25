import request from '@/utils/request'
export const getTenants = () => request({ url: '/v1/cs/tenants' })
export const addTenant = t => request({ url: '/v1/cs/tenants', method: 'post', data: { tenant: t } })
export const delTenant = t => request({ url: '/v1/cs/tenants?tenant=' + t, method: 'delete' })