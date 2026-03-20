import http from '~/utils/http'
import type { UserInfo } from '~/types'

export interface LoginParams {
  tenantCode: string
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  user: UserInfo
}

/**
 * 登陆
 */
export function login(data: LoginParams) {
  return http.post<LoginResponse>('/login', data)
}
