import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons'
import styles from './style.module.scss'
import useSidebarStore from '~/stores/sidebarStore'
import { Dropdown, type MenuProps } from 'antd'

export default function Header() {
  const toggleSidebar = useSidebarStore((state) => state.toggleSidebar)
  const openSidebar = useSidebarStore((state) => state.openSidebar)

  const items: MenuProps['items'] = [
    {
      key: '1',
      label: '系统设置',
    },
    {
      key: '2',
      label: '修改密码',
    },
    {
      key: '3',
      label: '退出登录',
    },
  ]

  return (
    <header className={styles.header}>
      <div>
        <span className="text-[20px] cursor-pointer" onClick={toggleSidebar}>
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
