import { useState } from 'react'
import type { Route } from '../../+types/root'
import styles from './style.module.scss'
import type { FormProps } from 'antd'
import { Button, Checkbox, Form, Input, message } from 'antd'
import Footer from '~/components/footer'
import { useNavigate } from 'react-router'
import useUserStore from '~/stores/userStore'

export function meta({}: Route.MetaArgs) {
  return [{ title: `登陆 - ${import.meta.env.VITE_TITLE}` }]
}

type FieldType = {
  tenantCode: string
  username: string
  password: string
  remember: string
}

export default function Login() {
  const navigate = useNavigate()
  const { userLogin } = useUserStore()
  const [loading, setLoading] = useState(false)

  const onFinish: FormProps<FieldType>['onFinish'] = async (values) => {
    try {
      setLoading(true)
      await userLogin({
        tenantCode: values.tenantCode,
        username: values.username,
        password: values.password,
      })
      message.success('登录成功')
      navigate('/home/dashboard')
    } catch (error) {
      console.error(error)
    } finally {
      setLoading(false)
    }
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
          label="租户编码"
          name="tenantCode"
          rules={[{ required: true, message: '请输入租户编码!' }]}
        >
          <Input />
        </Form.Item>

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
          <Button type="primary" htmlType="submit" block loading={loading}>
            登陆
          </Button>
        </Form.Item>
      </Form>

      <Footer />
    </div>
  )
}
