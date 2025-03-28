import { useToast } from '../plugins/toast'
import router from '../router'
import type { ApiResponse } from '../api/api'

const BASE_URL = 'http://localhost:8080/api/v1'

export async function request<T>(
  url: string,
  options: RequestInit = {}
): Promise<ApiResponse<T>> {
  const token = localStorage.getItem('token')
  
  // 合并默认配置和自定义配置
  const config: RequestInit = {
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
      ...(options.headers || {})
    },
    ...options
  }

  try {
    const response = await fetch(`${BASE_URL}${url}`, config)
    const data = await response.json()

    // 处理 401 未授权的情况
    if (data.status === 401) {
      localStorage.removeItem('token')
      router.replace('/login')
      const toast = useToast()
      toast.error('登录已过期，请重新登录')
      return Promise.reject(new Error('Unauthorized'))
    }

    return data
  } catch (error) {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
}

// 导出常用的请求方法
export const http = {
  get: <T>(url: string, options?: RequestInit) => 
    request<T>(url, { ...options, method: 'GET' }),
    
  post: <T>(url: string, data?: any, options?: RequestInit) =>
    request<T>(url, {
      ...options,
      method: 'POST',
      body: JSON.stringify(data)
    }),
    
  put: <T>(url: string, data?: any, options?: RequestInit) =>
    request<T>(url, {
      ...options,
      method: 'PUT',
      body: JSON.stringify(data)
    }),
    
  delete: <T>(url: string, options?: RequestInit) =>
    request<T>(url, { ...options, method: 'DELETE' })
} 