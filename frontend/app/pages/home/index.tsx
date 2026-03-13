import styles from './style.module.scss'
import { Outlet } from 'react-router'
import Sidebar from '~/components/sidebar'
import Header from '~/components/header'
import Tabs from '~/components/tabs'
import { Suspense, useEffect } from 'react'
import { Spin } from 'antd'
import { useLocation } from 'react-router'
import useTabStore from '~/stores/tabStore'

export function meta() {
  return [{ title: `系统主页 - ${import.meta.env.VITE_TITLE}` }]
}

export const handle = {
  name: '123123',
}

export default function Home() {
  const location = useLocation()
  const tabStore = useTabStore()

  useEffect(() => {
    tabStore.addTab(location.pathname)
  }, [location])

  return (
    <div className={styles.home}>
      <Sidebar />

      <div className={styles.wrapper}>
        <Header />
        <Tabs />

        <div className={styles.homeContainer}>
          <Suspense fallback={<Spin />}>
            <Outlet />
          </Suspense>
        </div>
      </div>
    </div>
  )
}
