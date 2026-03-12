import type { TabProps } from '~/types'

export function setOpenSidebar(isOpen: boolean) {
  localStorage.setItem('openSidebar', isOpen ? '1' : '0')
}

export function getOpenSidebar() {
  const open = localStorage.getItem('openSidebar')
  return open === '1' || !open
}

export function setTabActiveKey(key: string) {
  localStorage.setItem('tabActiveKey', key)
}

export function getTabActiveKey() {
  return localStorage.getItem('tabActiveKey') || ''
}

export function setTabs(tabs: TabProps[]) {
  localStorage.setItem('tabs', JSON.stringify(tabs))
}

export function getTabs(): TabProps[] {
  try {
    return JSON.parse(localStorage.getItem('tabs') || '')
  } catch {
    return []
  }
}
