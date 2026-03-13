import { create } from 'zustand'
import { getOpenSidebar } from '~/utils/storage'
import type { UserInfo } from '~/types'
import { logout } from '~/utils/storage'

interface UserState {
  userInfo: Partial<UserInfo>

  userLogout: () => void
}

export default create<UserState>((set) => ({
  userInfo: {},

  userLogout() {
    logout()
  },
}))
