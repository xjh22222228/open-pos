import { useState, useMemo } from 'react'
import { Button, Form, Input, Upload, Select, Switch, Checkbox } from 'antd'
import type {
  FormProps,
  GetProp,
  UploadFile,
  UploadProps,
  FormItemProps,
  CheckboxOptionType,
} from 'antd'
import ImgCrop from 'antd-img-crop'
import type { BusinessDays } from '~/types'
import { getDayText } from '~/utils/date'

export function meta() {
  return [{ title: `系统设置 - ${import.meta.env.VITE_TITLE}` }]
}

type FileType = Parameters<GetProp<UploadProps, 'beforeUpload'>>[0]

type FieldType = {
  avatar: string
  name: string
  phone: string
  workDays: BusinessDays[]
  workHours: [string, string][]
}

interface FormCustomProps extends FormItemProps {
  component: React.ReactNode
}

const workDaysOpts: CheckboxOptionType<number>[] = [
  { label: getDayText(1), value: 1 },
  { label: getDayText(2), value: 2 },
  { label: getDayText(3), value: 3 },
  { label: getDayText(4), value: 4 },
  { label: getDayText(5), value: 5 },
  { label: getDayText(6), value: 6 },
  { label: getDayText(7), value: 7 },
]

export default function SystemSettings() {
  const [form] = Form.useForm()
  const [fileList, setFileList] = useState<UploadFile[]>([
    {
      uid: '-1',
      name: 'image.png',
      status: 'done',
      url: 'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png',
    },
  ])

  const onFinish: FormProps<FieldType>['onFinish'] = (values) => {
    console.log(values)
  }

  const onChange: UploadProps['onChange'] = ({ fileList: newFileList }) => {
    setFileList(newFileList)
  }

  const onPreview = async (file: UploadFile) => {
    let src = file.url as string
    if (!src) {
      src = await new Promise((resolve) => {
        const reader = new FileReader()
        reader.readAsDataURL(file.originFileObj as FileType)
        reader.onload = () => resolve(reader.result as string)
      })
    }
    form.setFieldValue('storeAvatar', src)
  }

  const baseFormItems: FormCustomProps[] = useMemo(
    () => [
      {
        label: '头像',
        name: 'avatar',
        component: (
          <ImgCrop rotationSlider>
            <Upload
              action="https://660d2bd96ddfa2943b33731c.mockapi.io/api/upload"
              listType="picture-card"
              fileList={fileList}
              onChange={onChange}
              onPreview={onPreview}
            >
              {fileList.length < 1 && '+ Upload'}
            </Upload>
          </ImgCrop>
        ),
      },
      {
        label: '姓名',
        name: 'name',
        rules: [{ required: true, message: '请输入姓名' }],
        component: <Input />,
      },
      {
        label: '联系电话',
        name: 'phone',
        component: <Input />,
      },
      {
        label: '上班周期',
        name: 'workDays',
        component: <Checkbox.Group options={workDaysOpts} />,
      },
      {
        label: '上班时间',
        name: 'workHours',
        component: <Input />,
      },
    ],
    [],
  )

  return (
    <div className="p-2.5 h-full bg-white">
      <Form
        className="!mt-6"
        name="form"
        form={form}
        labelCol={{ span: 4 }}
        wrapperCol={{ span: 20 }}
        style={{ maxWidth: 600 }}
        onFinish={onFinish}
        autoComplete="off"
      >
        {baseFormItems.map(({ component, ...item }, idx) => (
          <Form.Item<FieldType> {...item} key={idx}>
            {component}
          </Form.Item>
        ))}

        <Form.Item label={null}>
          <Button type="primary" htmlType="submit">
            保存
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}
