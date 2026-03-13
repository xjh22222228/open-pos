import React, { useState, useEffect } from 'react'
import styles from './number.module.scss'
import NumberFlow from '@number-flow/react'
import { Link } from 'react-router'
import classNames from 'classnames'

const Number: React.FC = () => {
  const [state, setState] = useState([
    {
      title: '今日销售额',
      total: 0,
      Icon: <i className={classNames('iconfont icon-xiaoshoue', styles.icon)}></i>,
      prefix: '￥',
      path: '/home/bill',
    },
    {
      title: '今日净收入',
      total: 0,
      Icon: <i className={classNames('iconfont icon-xiaoshoue1', styles.icon)}></i>,
      prefix: '￥',
      path: '/home/reminder',
    },
    {
      title: '今日订单笔数',
      total: 0,
      Icon: <i className={classNames('iconfont icon-dingdan', styles.icon)}></i>,
      path: '/home/todayTask',
    },
    {
      title: '今日销售数量',
      total: 0,
      Icon: <i className={classNames('iconfont icon-dingdanshu', styles.icon)}></i>,
      path: '/home/todoList',
    },
  ])

  useEffect(() => {}, [])

  return (
    <div
      className={classNames(
        'grid gap-3 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4',
        styles.number,
      )}
    >
      {state.map((item) => (
        <div className={styles.item} key={item.title}>
          <Link to={item.path} className={styles['block-item']}>
            {item.Icon}
            <div className={styles.data}>
              <div className={styles.title}>{item.title}</div>
              <NumberFlow value={item.total} prefix={item.prefix} />
            </div>
          </Link>
        </div>
      ))}
    </div>
  )
}

export default Number
