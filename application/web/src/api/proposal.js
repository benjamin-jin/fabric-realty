import request from '@/utils/request'


export function CreateProposal() {
  return request({
    url: '/createProposal',
    method: 'post'
  })
}

export function PostProposal(data) {
  return request({
    url: '/postProposal',
    method: 'post',
    data
  })
}

export function DeleteProposal(data) {
  return request({
    url: '/deleteProposal',
    method: 'post',
    data
  })
}
export function GetProposal(data) {
  return request({
    url: '/getProposal',
    method: 'post',
    data
  })
}
export function GetProposalList(data) {
  return request({
    url: '/getProposalList',
    method: 'post',
    data
  })
}