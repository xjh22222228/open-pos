import styles from './index.module.scss'
import Number from './number'

export function meta() {
  return [{ title: `系统主页 - ${import.meta.env.VITE_TITLE}` }]
}

export default function Dashboard() {
  return (
    <div>
      <Number />
    </div>
  )
}
