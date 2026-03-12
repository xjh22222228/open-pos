import { create } from 'zustand'
import { getTabActiveKey, setTabActiveKey, setTabs, getTabs } from '~/utils/storage'
import type { TabProps } from '~/types'
import sidebarStore from './sidebarStore'

interface TabState {
  // Route path
  tabActiveKey: TabProps['key']
  tabs: TabProps[]
  addTab: (key: TabProps['key']) => void
  removeTab: (key: TabProps['key']) => void
  setTabActiveKey: (key: TabProps['key']) => void
}

const tabStore = create<TabState>((set, get) => ({
  tabActiveKey: getTabActiveKey(),
  tabs: getTabs(),
  addTab: (key) => {
    const state = get()
    const hasTab = state.tabs.some((item) => item.key === key)
    if (hasTab) {
      set(() => ({ tabActiveKey: key }))
    } else {
      const sidebarMenus = sidebarStore.getState().sidebarMenus
      let data: any = null
      const tabData = (menus: typeof sidebarMenus) => {
        for (let i = 0; i < menus.length; i++) {
          const item: any = menus[i]
          if (item.key === key) {
            data = item
            break
          }
          if (item.children) {
            data = tabData(item.children)
            if (data) {
              break
            }
          }
        }
        return data
      }
      tabData(sidebarMenus)
      set(() => ({
        tabs: data
          ? [
              ...state.tabs,
              {
                key: data.key,
                label: data.label,
              },
            ]
          : state.tabs,
        tabActiveKey: key,
      }))
    }
  },
  removeTab(key) {
    set((state) => {
      const tabs = state.tabs.filter((tab) => tab.key !== key)
      const idx = state.tabs.findIndex((tab) => tab.key === state.tabActiveKey)
      const tabActiveKey = state.tabs[idx === 0 ? idx + 1 : idx - 1]?.key || ''
      return {
        tabs,
        tabActiveKey,
      }
    })
  },
  setTabActiveKey(tabActiveKey) {
    set(() => ({
      tabActiveKey,
    }))
  },
}))

tabStore.subscribe((state) => {
  setTabActiveKey(state.tabActiveKey)
  setTabs(state.tabs)
})

export default tabStore
