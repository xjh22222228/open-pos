import { useState } from 'react'
import type { Route } from '../../+types/root'
import styles from './style.module.scss'
import type { FormProps } from 'antd'
import { Button, Checkbox, Form, Input } from 'antd'
import Footer from '~/components/footer'
import { useNavigate } from 'react-router'

export function meta({}: Route.MetaArgs) {
  return [{ title: `登陆 - ${import.meta.env.VITE_TITLE}` }]
}

type FieldType = {
  username?: string
  password?: string
  remember?: string
}

export default function Login() {
  const navigate = useNavigate()

  const onFinish: FormProps<FieldType>['onFinish'] = (values) => {
    navigate('/home/dashboard')
    console.log('Success:', values)
  }

  return (
    <div className={styles.login}>
      <Form
        className="flex-1 w-[500px] !max-w-full px-[20px] flex aliitems-center justify-center flex-col"
        name="form"
        labelCol={{ span: 8 }}
        wrapperCol={{ span: 16 }}
        initialValues={{ remember: true }}
        onFinish={onFinish}
        autoComplete="off"
      >
        <Form.Item<FieldType>
          label="用户名"
          name="username"
          rules={[{ required: true, message: '请输入您的用户名!' }]}
        >
          <Input />
        </Form.Item>

        <Form.Item<FieldType>
          label="密码"
          name="password"
          rules={[{ required: true, message: '请输入您的密码!' }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item<FieldType> name="remember" valuePropName="checked" label={null}>
          <Checkbox>记住我</Checkbox>
        </Form.Item>

        <Form.Item label={null}>
          <Button type="primary" htmlType="submit" block>
            登陆
          </Button>
        </Form.Item>
      </Form>

      <Footer />
    </div>
  )
}
