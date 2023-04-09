import request from '@/utils/request'

// 获取登录界面角色选择列表
export function getTransportList() {
  return request({
    url: '/queryAccountList',
    method: 'post'
  })
}