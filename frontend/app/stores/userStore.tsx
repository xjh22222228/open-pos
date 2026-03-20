import { create } from 'zustand'
import type { UserInfo } from '~/types'
import { logout, setToken } from '~/utils/storage'
import { login, type LoginParams } from '~/services/login'

interface UserState {
  userInfo: Partial<UserInfo>

  userLogin: (params: LoginParams) => Promise<void>
  userLogout: () => void
}

export default create<UserState>((set) => ({
  userInfo: {},

  async userLogin(params) {
    const res = await login(params)
    setToken(res.data.token)
    set({ userInfo: res.data.user })
  },

  userLogout() {
    logout()
  },
}))
