import request from '@/utils/request'

// 获取需求列表
export function getDemandList(data) {
  return request({
    url: '/getDemandList',
    method: 'post',
    data
  })
}