import { StoreType, PayMethod } from '~/types/enum'

export function getStoreTypeText(type: StoreType): string {
  const storeTypeText = {
    [StoreType.Catering]: '餐饮',
    [StoreType.FreshFood]: '生鲜',
    [StoreType.Minimart]: '超市',
    [StoreType.Takeout]: '外卖',
    [StoreType.Hardware]: '五金',
    [StoreType.Clothing]: '服装',
    [StoreType.Bookstore]: '书店',
    [StoreType.ElectronicProducts]: '电子产品',
    [StoreType.Other]: '其他',
  }
  return storeTypeText[type]
}

export function getPayMethodText(type: PayMethod) {
  const payMethodText = {
    [PayMethod.Cash]: '现金',
    [PayMethod.CreditCard]: '信用卡',
    [PayMethod.Alipay]: '支付宝',
    [PayMethod.WechatPay]: '微信',
    [PayMethod.BankTransfer]: '银行转账',
    [PayMethod.Other]: '其他',
  }
  return payMethodText[type]
}
