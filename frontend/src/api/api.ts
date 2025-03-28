export interface PaginationData<T> {
  total: number
  total_pages: number
  page: number
  page_size: number
  items: T[]
}

// 产品接口
export interface Product {
  id: number
  name: string
  price: number
  status: 'enabled' | 'disabled'
  created_at: string
}

// 登录响应接口
export interface LoginResponse {
  token: string
}

// 许可证接口
export interface License {
  id: number
  product_id: number
  license_key: string
  duration_days: number
  created_at: string
  activated_at: string | null
  expires_at: string | null
  activation_status: 'inactive' | 'activated' | 'expired'
  enable_status: 'enabled' | 'disabled'
  remark: string
  product: Product
}

// API 响应接口
export interface ApiResponse<T> {
  status: number
  message: string
  data: T
}

export interface VerifyLicenseResponse {
  valid: boolean
  message?: string
  license?: {
    product_name: string
    activated_at: string
    expires_at: string
  }
}

// 产品统计信息
export interface ProductStats {
  total: number           // 产品总数
  enabled: number        // 已启用产品数
  disabled: number       // 已禁用产品数
}

// 许可证统计信息
export interface LicenseStats {
  total: number          // 许可证总数
  activated: number      // 已激活数
  inactive: number       // 未激活数
  expired: number        // 已过期数
  disabled: number       // 已禁用数
}

// 收益统计信息
export interface RevenueStats {
  total: number          // 总收益
  today: number         // 今日收益
  this_month: number    // 本月收益
}

// 统计数据汇总
export interface StatsSummary {
  products: ProductStats
  licenses: LicenseStats
  revenue: RevenueStats
}

// 许可证分析数据
export interface LicenseAnalytics {
  // 产品转化率
  product_conversion: Array<{
    id: number
    name: string
    total_licenses: number
    activated_licenses: number
    revenue: number
    conversion_rate: number
  }>

  // 时间趋势
  time_trend: Array<{
    date: string
    new_licenses: number
    activated_licenses: number
    daily_revenue: number
  }>

  // 有效期分布
  duration_distribution: Array<{
    duration_days: number
    count: number
  }>

  // 过期预警
  expiration_warning: {
    expiring_soon: number
  }
}

// 收益趋势数据
export interface RevenueTrend {
  dates: string[]
  revenue: number[]
  total_revenue: number
  avg_daily_revenue: number
}

// 产品激活数据
export interface ProductActivationData {
  id: number
  name: string
  price: number
  total_licenses: number
  activated_licenses: number
  revenue: number
  activation_rate: number
}

// 产品激活分析
export interface ProductActivation {
  products: ProductActivationData[]
  summary: {
    total_licenses: number
    total_activated: number
    total_revenue: number
    avg_activation_rate: number
  }
}

export interface RevenueTrendData {
  dates: string[]
  revenue: number[]
  total: number
  average: number
} 