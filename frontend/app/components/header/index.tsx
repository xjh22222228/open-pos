import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons'
import styles from './style.module.scss'

export default function Header() {
  return (
    <header className={styles.header}>
      <div>
        <span className="text-[20px] cursor-pointer">
          <MenuFoldOutlined />
        </span>
      </div>
    </header>
  )
}
