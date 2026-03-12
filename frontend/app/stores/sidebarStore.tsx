import { create } from 'zustand'
import { getOpenSidebar, setOpenSidebar } from '~/utils/storage'
import type { MenuProps } from 'antd'
import { SettingOutlined, HomeOutlined } from '@ant-design/icons'

type MenuItem = Required<MenuProps>['items'][number]

const items: MenuItem[] = [
  {
    key: '/home/dashboard',
    label: '系统首页',
    icon: <HomeOutlined />,
  },
  {
    key: 'sub2',
    label: '系统管理',
    icon: <SettingOutlined />,
    children: [
      { key: '/home/system-settings', label: '系统设置' },
      { key: '/home/change-password', label: '修改密码' },
    ],
  },
]

interface SidebarState {
  openSidebar: boolean
  sidebarMenus: typeof items
  toggleSidebar: () => void
}

export default create<SidebarState>((set) => ({
  openSidebar: getOpenSidebar(),
  sidebarMenus: items,
  toggleSidebar: () => {
    set((state) => {
      setOpenSidebar(!state.openSidebar)
      return { openSidebar: !state.openSidebar }
    })
  },
}))
