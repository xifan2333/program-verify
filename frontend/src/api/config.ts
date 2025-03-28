// API配置
export const API_CONFIG = {
  // API基础URL
  BASE_URL: import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080/api/v1',
  
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
    LOGIN: '/auth/login'
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
    // 如果baseUrl是相对路径，直接拼接
    if (baseUrl.startsWith('/')) {
      return `${baseUrl}${path}`
    }
    // 否则使用URL构造函数
    return `${baseUrl}${path}`
  },
  
  // 处理响应
  handleResponse: async (response: Response) => {
    const data = await response.json()
    if (!response.ok) {
      throw new Error(data.message || '请求失败')
    }
    return data
  },
  
  // GET请求
  get: async (path: string, params?: Record<string, string>) => {
    const url = api.buildUrl(path)
    const finalUrl = params ? `${url}?${new URLSearchParams(params).toString()}` : url
    
    const response = await fetch(finalUrl, {
      headers: {
        ...API_CONFIG.HEADERS,
        ...api.getAuthHeader()
      }
    })
    return api.handleResponse(response)
  },
  
  // POST请求
  post: async (path: string, data?: any) => {
    const response = await fetch(api.buildUrl(path), {
      method: 'POST',
      headers: {
        ...API_CONFIG.HEADERS,
        ...api.getAuthHeader()
      },
      body: data ? JSON.stringify(data) : undefined
    })
    return api.handleResponse(response)
  },
  
  // PUT请求
  put: async (path: string, data?: any) => {
    const response = await fetch(api.buildUrl(path), {
      method: 'PUT',
      headers: {
        ...API_CONFIG.HEADERS,
        ...api.getAuthHeader()
      },
      body: data ? JSON.stringify(data) : undefined
    })
    return api.handleResponse(response)
  },
  
  // DELETE请求
  delete: async (path: string) => {
    const response = await fetch(api.buildUrl(path), {
      method: 'DELETE',
      headers: {
        ...API_CONFIG.HEADERS,
        ...api.getAuthHeader()
      }
    })
    return api.handleResponse(response)
  }
} 