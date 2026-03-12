import { Typography, Button, Form, Input, message } from 'antd'
import type { FormProps } from 'antd'

export function meta() {
  return [{ title: `系统设置 - ${import.meta.env.VITE_TITLE}` }]
}

type FieldType = {
  password: string
  newPassword: string
  confirmNewPassword?: string
}

export default function ChangePassword() {
  const onFinish: FormProps<FieldType>['onFinish'] = (values) => {}

  return <div className="p-2.5">33333</div>
}
