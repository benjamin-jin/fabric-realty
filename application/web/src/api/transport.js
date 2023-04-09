import request from '@/utils/request'

// 获取登录界面角色选择列表
export function getTransportList(data) {
  return request({
    url: '/getTransportList',
    method: 'post',
    data
  })
}