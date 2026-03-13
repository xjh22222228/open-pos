import { Typography, Button, Form, Input, message } from 'antd'
import type { FormProps } from 'antd'

const { Title } = Typography

export function meta() {
  return [{ title: `修改密码 - ${import.meta.env.VITE_TITLE}` }]
}

type FieldType = {
  password: string
  newPassword: string
  confirmNewPassword?: string
}

export default function ChangePassword() {
  const onFinish: FormProps<FieldType>['onFinish'] = (values) => {
    if (values.newPassword !== values.confirmNewPassword) {
      message.error('新密码和确认密码不一致')
      return
    }
  }

  return (
    <div className="p-2.5 h-full bg-white">
      <Title level={3}>修改密码</Title>

      <Form
        className="!mt-6"
        name="form"
        labelCol={{ span: 6 }}
        wrapperCol={{ span: 18 }}
        style={{ maxWidth: 600 }}
        onFinish={onFinish}
        autoComplete="off"
      >
        <Form.Item<FieldType>
          label="原密码"
          name="password"
          rules={[{ required: true, message: '请输入原密码' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item<FieldType>
          label="新密码"
          name="newPassword"
          rules={[{ required: true, message: '请输入新密码' }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item<FieldType>
          label="再次输入新密码"
          name="confirmNewPassword"
          rules={[{ required: true, message: '请再次输入新密码' }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item label={null}>
          <Button type="primary" htmlType="submit">
            提交
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}
