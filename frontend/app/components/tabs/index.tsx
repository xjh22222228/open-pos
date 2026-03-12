import React from 'react'
import { Tabs } from 'antd'
import useTabStore from '~/stores/tabStore'

type TargetKey = React.MouseEvent | React.KeyboardEvent | string

const TabsComponent: React.FC = () => {
  const tabStore = useTabStore((state) => state)

  const onChange = (key: string) => {
    tabStore.setTabActiveKey(key)
  }

  const onEdit = (targetKey: TargetKey, action: 'add' | 'remove') => {
    if (action === 'remove') {
      tabStore.removeTab(targetKey as string)
    }
  }

  return (
    <div>
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
