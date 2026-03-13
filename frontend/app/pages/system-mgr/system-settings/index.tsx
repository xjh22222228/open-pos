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
import { StoreType, PayMethod } from '~/types/enum'
import { getStoreTypeText, getPayMethodText } from '~/utils/store'
import type { BusinessDays } from '~/types'
import { getDayText } from '~/utils/date'

export function meta() {
  return [{ title: `个人信息 - ${import.meta.env.VITE_TITLE}` }]
}

type FileType = Parameters<GetProp<UploadProps, 'beforeUpload'>>[0]

type FieldType = {
  // 基础设置
  businessLicense: string
  storeAvatar: string
  storeName: string
  phone: string
  landlineNumber: string
  isClosed: boolean
  businessDays: BusinessDays[]
  businessHours: [string, string][]
  // 支付设置
  payMethods: PayMethod[]
}

interface FormCustomProps extends FormItemProps {
  component: React.ReactNode
}

const businessDaysOpts: CheckboxOptionType<number>[] = [
  { label: getDayText(1), value: 1 },
  { label: getDayText(2), value: 2 },
  { label: getDayText(3), value: 3 },
  { label: getDayText(4), value: 4 },
  { label: getDayText(5), value: 5 },
  { label: getDayText(6), value: 6 },
  { label: getDayText(7), value: 7 },
]

const payMethodOpts: CheckboxOptionType<PayMethod>[] = [
  { label: getPayMethodText(PayMethod.Cash), value: PayMethod.Cash },
  {
    label: getPayMethodText(PayMethod.CreditCard),
    value: PayMethod.CreditCard,
  },
  { label: getPayMethodText(PayMethod.Alipay), value: PayMethod.Alipay },
  { label: getPayMethodText(PayMethod.WechatPay), value: PayMethod.WechatPay },
  {
    label: getPayMethodText(PayMethod.BankTransfer),
    value: PayMethod.BankTransfer,
  },
  { label: getPayMethodText(PayMethod.Other), value: PayMethod.Other },
]

const storeTypeOpts: CheckboxOptionType<StoreType>[] = [
  {
    value: StoreType.Minimart,
    label: getStoreTypeText(StoreType.Minimart),
  },
  {
    value: StoreType.Catering,
    label: getStoreTypeText(StoreType.Catering),
  },
  {
    value: StoreType.FreshFood,
    label: getStoreTypeText(StoreType.FreshFood),
  },
  {
    value: StoreType.Takeout,
    label: getStoreTypeText(StoreType.Takeout),
  },
  {
    value: StoreType.Hardware,
    label: getStoreTypeText(StoreType.Hardware),
  },
  {
    value: StoreType.Clothing,
    label: getStoreTypeText(StoreType.Clothing),
  },
  {
    value: StoreType.Bookstore,
    label: getStoreTypeText(StoreType.Bookstore),
  },
  {
    value: StoreType.ElectronicProducts,
    label: getStoreTypeText(StoreType.ElectronicProducts),
  },
  {
    value: StoreType.Other,
    label: getStoreTypeText(StoreType.Other),
  },
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
        label: '店铺LOGO',
        name: 'storeAvatar',
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
        label: '店铺执照',
        name: 'businessLicense',
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
        label: '店铺名称',
        name: 'storeName',
        rules: [{ required: true, message: '请输入店铺名称' }],
        component: <Input />,
      },
      {
        label: '联系电话',
        name: 'phone',
        component: <Input />,
      },
      {
        label: '固定电话',
        name: 'landlineNumber',
        component: <Input />,
      },
      {
        label: '店铺类型',
        name: 'storeType',
        component: <Select options={storeTypeOpts} />,
      },
      {
        label: '营业周期',
        name: 'businessDays',
        component: <Checkbox.Group options={businessDaysOpts} />,
      },
      {
        label: '营业时间',
        name: 'businessHours',
        component: <Input />,
      },
      {
        label: '是否打烊',
        name: 'isClosed',
        component: <Switch />,
      },
      {
        label: '支付方式',
        name: 'payMethods',
        component: <Checkbox.Group options={payMethodOpts} />,
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
