import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons'
import styles from './style.module.scss'
import useSidebarStore from '~/stores/sidebarStore'
import { Dropdown, type MenuProps } from 'antd'
import useUserStore from '@/stores/userStore'
import { useNavigate } from 'react-router'

export default function Header() {
  const toggleSidebar = useSidebarStore((state) => state.toggleSidebar)
  const openSidebar = useSidebarStore((state) => state.openSidebar)
  const userLogout = useUserStore((state) => state.userLogout)
  const navigate = useNavigate()

  const items: MenuProps['items'] = [
    {
      key: '1',
      label: '系统设置',
      onClick: () => {
        navigate('/home/system-settings')
      },
    },
    {
      key: '2',
      label: '修改密码',
      onClick: () => {
        navigate('/home/change-password')
      },
    },
    {
      key: '3',
      label: '退出登录',
      onClick: async () => {
        userLogout()
        await navigate('/')
        location.reload()
      },
    },
  ]

  return (
    <header className={styles.header}>
      <div className="flex items-center">
        <span className="text-[16px] cursor-pointer" onClick={toggleSidebar}>
          {openSidebar ? <MenuFoldOutlined /> : <MenuUnfoldOutlined />}
        </span>

        <span className="ml-2.5">OPEN POS</span>
      </div>

      <div>
        <Dropdown menu={{ items }} placement="bottomRight">
          <img
            src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg"
            alt=""
            className="w-[30px] h-[30px] object-cover cursor-pointer"
            draggable="false"
          />
        </Dropdown>
      </div>
    </header>
  )
}
