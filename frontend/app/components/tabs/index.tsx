import React from 'react'
import { Tabs } from 'antd'
import useTabStore from '~/stores/tabStore'
import { useNavigate } from 'react-router'

type TargetKey = React.MouseEvent | React.KeyboardEvent | string

const TabsComponent: React.FC = () => {
  const navigate = useNavigate()
  const tabStore = useTabStore((state) => state)

  const onChange = (key: string) => {
    navigate(key)
    tabStore.setTabActiveKey(key)
  }

  const onEdit = (targetKey: TargetKey, action: 'add' | 'remove') => {
    if (action === 'remove') {
      tabStore.removeTab(targetKey as string)
    }
  }

  return (
    <div className="bg-white select-none">
      <Tabs
        hideAdd
        onChange={onChange}
        activeKey={tabStore.tabActiveKey}
        type="editable-card"
        onEdit={onEdit}
        items={tabStore.tabs}
        size="small"
      />
    </div>
  )
}

export default TabsComponent
