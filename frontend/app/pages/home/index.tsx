import styles from './style.module.scss'
import { Outlet } from 'react-router'
import Sidebar from '~/components/sidebar'
import Header from '~/components/header'
import Tabs from '~/components/tabs'

export function meta() {
  return [{ title: `系统首页 - ${import.meta.env.VITE_TITLE}` }]
}

export const handle = {
  name: '123123',
}

export default function Home() {
  return (
    <div className={styles.home}>
      <Sidebar />

      <div className="flex-1">
        <Header />
        <Tabs />

        <Outlet />
      </div>
    </div>
  )
}
