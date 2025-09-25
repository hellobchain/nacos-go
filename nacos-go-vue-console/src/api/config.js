import request from '@/utils/request'
export const getConfigs = () => request({ url: '/v1/cs/configs/list' })
export const publish = data => request({ url: '/v1/cs/configs', method: 'post', data })
export const delConfig = p => request({ url: '/v1/cs/configs?dataId=' + p.dataId + '&group=' + p.group + '&tenant=' + (p.tenant || ''), method: 'delete' })