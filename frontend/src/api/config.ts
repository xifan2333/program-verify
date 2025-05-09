import { globalToastMethods } from "../plugins/toast";
const toast = globalToastMethods;


// API配置
export const API_CONFIG = {
  // API基础URL
  BASE_URL: import.meta.env.DEV ? 'http://localhost:8080/api/v1' : `${window.location.origin}/api/v1`,
  
  // 超时时间（毫秒）
  TIMEOUT: 10000,
  
  // 请求头
  HEADERS: {
    'Content-Type': 'application/json'
  }
}

// API路径配置
export const API_ROUTES = {
  // 认证相关
  AUTH: {
    LOGIN: '/auth/login',
    VERIFY: '/auth/verify'  // 添加验证接口
  },
  
  // 用户相关
  USER: {
    UPDATE: '/user/update'
  },
  
  // 产品相关
  PRODUCTS: {
    LIST: '/products',
    DETAIL: (id: number) => `/products/${id}`,
    CREATE: '/products',
    UPDATE: (id: number) => `/products/${id}`,
    STATS: '/products/stats'
  },
  
  // 许可证相关
  LICENSES: {
    LIST: '/licenses',
    DETAIL: (id: number) => `/licenses/${id}`,
    CREATE: '/licenses',
    UPDATE: (id: number) => `/licenses/${id}`,
    VERIFY: '/licenses/verify',
    STATS: '/licenses/stats'
  },

  // 数据分析相关
  ANALYTICS: {
    STATS: '/analytics/stats',
    REVENUE_TREND: '/analytics/revenue-trend',
    PRODUCT_ACTIVATION: '/analytics/product-activation'
  }
}

// 统一处理未授权
const handleUnauthorized = () => {
  localStorage.removeItem('token')
 
  // 如果在路由守卫中，不需要重复跳转
  if (!window.location.pathname.includes('/login')) {
    window.location.href = '/login'
  }
}

// API请求工具函数
export const api = {
  // 获取认证头
  getAuthHeader: (): HeadersInit => {
    const token = localStorage.getItem('token')
    return token ? { 'Authorization': `Bearer ${token}` } : {}
  },
  
  // 构建完整URL
  buildUrl: (path: string) => {
    const baseUrl = API_CONFIG.BASE_URL
    // 开发环境使用相对路径，生产环境使用完整URL
   
    return `${baseUrl}${path}`
  },
  
  // 处理响应
  handleResponse: async <T>(response: Response): Promise<T> => {
    const data = await response.json()
    
    // 处理未授权
    if (response.status === 401 || data.status === 401) {
      toast.error(data.message || '无效的令牌认证，请重新登录')
      handleUnauthorized()
      
    }
    
    // 处理其他错误
    if (!response.ok) {
      const errorMessage = data.message || '请求失败'
      throw new Error(errorMessage)
    }
    return data
  },
  
  // GET请求
  get: async <T>(path: string, params?: Record<string, string>): Promise<T> => {
    const url = api.buildUrl(path)
    const finalUrl = params ? `${url}?${new URLSearchParams(params).toString()}` : url
    
    try {
      const response = await fetch(finalUrl, {
        headers: {
          ...API_CONFIG.HEADERS,
          ...api.getAuthHeader()
        }
      })
      return api.handleResponse<T>(response)
    } catch (error) {
      throw error
    }
  },
  
  // POST请求
  post: async <T>(path: string, data?: any): Promise<T> => {
    try {
      const response = await fetch(api.buildUrl(path), {
        method: 'POST',
        headers: {
          ...API_CONFIG.HEADERS,
          ...api.getAuthHeader()
        },
        body: data ? JSON.stringify(data) : undefined
      })
      return api.handleResponse<T>(response)
    } catch (error) {
      throw error
    }
  },
  
  // PUT请求
  put: async <T>(path: string, data?: any): Promise<T> => {
    try {
      const response = await fetch(api.buildUrl(path), {
        method: 'PUT',
        headers: {
          ...API_CONFIG.HEADERS,
          ...api.getAuthHeader()
        },
        body: data ? JSON.stringify(data) : undefined
      })
      return api.handleResponse<T>(response)
    } catch (error) {
      throw error
    }
  },
  
  // DELETE请求
  delete: async <T>(path: string): Promise<T> => {
    try {
      const response = await fetch(api.buildUrl(path), {
        method: 'DELETE',
        headers: {
          ...API_CONFIG.HEADERS,
          ...api.getAuthHeader()
        }
      })
      return api.handleResponse<T>(response)
    } catch (error) {
      throw error
    }
  }
}



