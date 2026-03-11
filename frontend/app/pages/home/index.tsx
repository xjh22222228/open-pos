import styles from './style.module.scss'
import { Outlet } from 'react-router'
import Sidebar from '~/components/sidebar'
import Header from '~/components/header'

export function meta() {
  return [{ title: `首页 - ${import.meta.env.VITE_TITLE}` }]
}

export default function Home() {
  return (
    <div className={styles.home}>
      <Sidebar />

      <div className="flex-1">
        <Header />

        <Outlet />
      </div>
    </div>
  )
}
