import type { BusinessDays } from '~/types'

export function getDayText(day: BusinessDays) {
  const days = ['', '周一', '周二', '周三', '周四', '周五', '周六', '周日']
  return days[day]
}
