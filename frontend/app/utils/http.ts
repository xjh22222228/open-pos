import axios from 'axios'
import nprogress from 'nprogress'
import 'nprogress/nprogress.css'
import { message } from 'antd'
import { getToken, logout } from './storage'

nprogress.configure({ showSpinner: false })

const http = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 10000,
})

http.interceptors.request.use(
  (config) => {
    nprogress.start()
    const token = getToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    nprogress.done()
    return Promise.reject(error)
  },
)

http.interceptors.response.use(
  (response) => {
    nprogress.done()
    const res = response.data
    // 处理业务错误码
    if (res.errorCode !== 0) {
      // 登录失效判断
      if (res.errorCode === 401) {
        logout()
        window.location.href = '/'
        return Promise.reject(new Error(res.message || '登录失效'))
      }
      message.error(res.message || '请求失败')
      return Promise.reject(new Error(res.message || '请求失败'))
    }
    return res
  },
  (error) => {
    nprogress.done()
    if (error.response) {
      const { status, data } = error.response
      // 兼容判断：无论是 HTTP 状态码还是业务错误码为 401
      if (status === 401 || data?.errorCode === 401) {
        logout()
        window.location.href = '/'
      } else {
        message.error(data?.message || '服务器内部错误')
      }
    } else {
      message.error('网络连接异常，请稍后再试')
    }
    return Promise.reject(error)
  },
)

export default http
