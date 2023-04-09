import request from '@/utils/request'

// 获取登录界面角色选择列表
export function getStorageList(data) {
  return request({
    url: '/getStorageList',
    method: 'post',
    data
  })
}