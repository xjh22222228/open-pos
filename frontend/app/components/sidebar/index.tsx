import { useState, useEffect } from 'react'
import type { MenuProps } from 'antd'
import { Menu } from 'antd'
import useSidebarStore from '~/stores/sidebarStore'
import classNames from 'classnames'
import useTabStore from '~/stores/tabStore'
import { useNavigate } from 'react-router'

export default function Sidebar() {
  const openSidebar = useSidebarStore((state) => state.openSidebar)
  const sidebarMenus = useSidebarStore((state) => state.sidebarMenus)
  const tabStore = useTabStore()
  const navigate = useNavigate()
  const [openKeys, setOpenKeys] = useState<string[]>([])

  const onClick: MenuProps['onClick'] = (e) => {
    if (tabStore.tabActiveKey !== e.key) {
      navigate(e.key)
    }
  }

  // 获取当前激活菜单项的父菜单 keys
  const getOpenKeys = (menus: any, activeKey: string): string[] => {
    const openKeys: string[] = []
    const findKeys = (items: any, parents: string[] = []): boolean => {
      for (const item of items) {
        if (item && typeof item === 'object' && 'key' in item) {
          const currentParents = [...parents, item.key as string]
          if (item.key === activeKey) {
            openKeys.push(...parents)
            return true
          }
          if (item.children) {
            if (findKeys(item.children, currentParents)) {
              openKeys.push(item.key as string)
              return true
            }
          }
        }
      }
      return false
    }
    findKeys(menus)
    return openKeys
  }

  useEffect(() => {
    const keys = getOpenKeys(sidebarMenus, tabStore.tabActiveKey)
    setOpenKeys([...openKeys, ...keys])
  }, [tabStore.tabActiveKey])

  const onOpenChange = (openKeys: string[]) => {
    setOpenKeys(openKeys)
  }

  return (
    <div
      className={classNames('overflow-hidden relative flex flex-col w-[200px]', {
        '!w-[80px]': !openSidebar,
      })}
    >
      <div
        className={classNames('p-2 flex items-center text-center', {
          'justify-center': !openSidebar,
        })}
      >
        <img
          src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg"
          alt=""
          className="w-[35px] h-[35px]"
        />
        {openSidebar ? <span className="ml-2">OPEN POS</span> : null}
      </div>

      <Menu
        className="h-full flex-1 overflow-hidden overflow-y-auto select-none"
        onClick={onClick}
        selectedKeys={[tabStore.tabActiveKey]}
        openKeys={openKeys}
        mode="inline"
        inlineCollapsed={!openSidebar}
        items={sidebarMenus}
        onOpenChange={onOpenChange}
      />
    </div>
  )
}
