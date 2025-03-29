<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useToast } from '../plugins/toast'
import type { ApiResponse, License, Product, PaginationData } from '../api/types'
import { api, API_ROUTES } from '../api/config'
import * as XLSX from 'xlsx'

const toast = useToast()
const licenses = ref<License[]>([])
const products = ref<Product[]>([])
const loading = ref(false)
const exportLoading = ref(false)
const showGenerateModal = ref(false)
const showEditModal = ref(false)
const showFilters = ref(false)
const currentLicense = ref<License | null>(null)
const showQuickFilter = ref(false)
const showActivationFilter = ref(false)

// 筛选条件
const filters = ref({
  activation_status: '',
  enable_status: '',
  product_id: '',
  license_key: '',
  activated_start_date: '',
  activated_end_date: '',
  expires_start_date: '',
  expires_end_date: '',
  remark: ''
})

const formData = ref({
  product_id: 0,
  duration_days: 365,
  count: 1,
  remark: ''
})

const editFormData = ref({
  remark: '',
  expires_at: ''
})

// 分页相关
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const showAll = ref(false)

const pageSizeOptions = [10, 20, 50, 100, -1] // -1 表示显示全部

// 快捷筛选状态
const quickFilter = ref<'enabled' | 'disabled' | 'all'>('enabled')
const activationFilter = ref<'activated' | 'inactive' | 'expired' | 'all'>('all')

// 获取当前激活状态筛选文本
const getActivationFilterText = () => {
  switch (activationFilter.value) {
    case 'activated':
      return '已激活'
    case 'inactive':
      return '未激活'
    case 'expired':
      return '已过期'
    case 'all':
      return '激活状态'
    default:
      return '激活状态'
  }
}

// 获取当前激活状态筛选图标
const getActivationFilterIcon = () => {
  switch (activationFilter.value) {
    case 'activated':
      return 'i-ri-checkbox-circle-line'
    case 'inactive':
      return 'i-ri-time-line'
    case 'expired':
      return 'i-ri-close-circle-line'
    case 'all':
      return 'i-ri-filter-3-line'
    default:
      return 'i-ri-filter-3-line'
  }
}

// 获取当前激活状态筛选颜色
const getActivationFilterColor = () => {
  switch (activationFilter.value) {
    case 'activated':
      return 'text-green-600'
    case 'inactive':
      return 'text-yellow-600'
    case 'expired':
      return 'text-red-600'
    case 'all':
      return 'text-blue-600'
    default:
      return 'text-gray-600'
  }
}

// 激活状态快捷筛选处理
const handleActivationFilter = (status: 'activated' | 'inactive' | 'expired' | 'all') => {
  activationFilter.value = status
  filters.value.activation_status = status === 'all' ? '' : status
  showActivationFilter.value = false
  handleFilter()
}

// 获取当前启用状态筛选文本
const getFilterText = () => {
  switch (quickFilter.value) {
    case 'enabled':
      return '已启用'
    case 'disabled':
      return '已禁用'
    case 'all':
      return '启用状态'
    default:
      return '启用状态'
  }
}

// 获取当前启用状态筛选图标
const getFilterIcon = () => {
  switch (quickFilter.value) {
    case 'enabled':
      return 'i-ri-checkbox-circle-line'
    case 'disabled':
      return 'i-ri-close-circle-line'
    case 'all':
      return 'i-ri-list-check'
    default:
      return 'i-ri-filter-3-line'
  }
}

// 获取当前启用状态筛选颜色
const getFilterColor = () => {
  switch (quickFilter.value) {
    case 'enabled':
      return 'text-green-600'
    case 'disabled':
      return 'text-red-600'
    case 'all':
      return 'text-blue-600'
    default:
      return 'text-gray-600'
  }
}

// 启用状态快捷筛选处理
const handleQuickFilter = (status: 'enabled' | 'disabled' | 'all') => {
  quickFilter.value = status
  filters.value.enable_status = status === 'all' ? '' : status
  showQuickFilter.value = false
  handleFilter()
}

const handlePageSizeChange = (size: number) => {
  pageSize.value = size
  showAll.value = size === -1
  currentPage.value = 1 // 切换每页数量时重置到第一页
  fetchLicenses()
}

const handlePageChange = (page: number) => {
  currentPage.value = page
  fetchLicenses()
}

const fetchLicenses = async () => {
  loading.value = true
  try {
    const params = {
      ...(showAll.value ? {} : {
        page: currentPage.value.toString(),
        page_size: pageSize.value.toString(),
      }),
      ...(filters.value.activation_status && { activation_status: filters.value.activation_status }),
      ...(filters.value.enable_status && { enable_status: filters.value.enable_status }),
      ...(filters.value.product_id && { product_id: filters.value.product_id }),
      ...(filters.value.license_key && { license_key: filters.value.license_key }),
      ...(filters.value.activated_start_date && { activated_start_date: filters.value.activated_start_date }),
      ...(filters.value.activated_end_date && { activated_end_date: filters.value.activated_end_date }),
      ...(filters.value.expires_start_date && { expires_start_date: filters.value.expires_start_date }),
      ...(filters.value.expires_end_date && { expires_end_date: filters.value.expires_end_date }),
      ...(filters.value.remark && { remark: filters.value.remark })
    }

    const response = await api.get<ApiResponse<PaginationData<License>>>(API_ROUTES.LICENSES.LIST, params)
    if (response.status === 200) {
      licenses.value = response.data.items
      total.value = response.data.total
    } else {
      toast.error(response.message || '获取许可证列表失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '获取许可证列表失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}

const fetchProducts = async () => {
  try {
    const response = await api.get<ApiResponse<PaginationData<Product>>>(API_ROUTES.PRODUCTS.LIST, { status: 'enabled' })
    if (response.status === 200) {
      products.value = response.data.items
    } else {
      toast.error(response.message || '获取产品列表失败')  
    }
  } catch (error) {
    
    toast.error(error instanceof Error ? error.message : '获取产品列表失败')
    console.error(error)
    
  }
}

const handleGenerate = async () => {
  if (!formData.value.product_id) return
  
  try {
    const response = await api.post<ApiResponse<{ count: number }>>(API_ROUTES.LICENSES.CREATE, formData.value)
    if (response.status === 200) {
      showGenerateModal.value = false
      formData.value = { product_id: 0, duration_days: 365, count: 1, remark: '' }
      toast.success(response.message)
      fetchLicenses()
    } else {
      toast.error(response.message || '生成失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '生成失败')
    console.error(error)
  }
}

const handleDisable = async (id: number) => {
  try {
    const response = await api.put<ApiResponse<null>>(API_ROUTES.LICENSES.UPDATE(id), {
      enable_status: 'disabled'
    })
    if (response.status === 200) {
      toast.success(response.message)
      fetchLicenses()
    } else {
      toast.error(response.message || '操作失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '操作失败')
    console.error(error)
  }
}

const handleEnable = async (id: number) => {
  try {
    const response = await api.put<ApiResponse<null>>(API_ROUTES.LICENSES.UPDATE(id), {
      enable_status: 'enabled'
    })
    if (response.status === 200) {
      toast.success(response.message)
      fetchLicenses()
    } else {
      toast.error(response.message || '操作失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '操作失败')
    console.error(error)
  }
}

const openEditModal = (license: License) => {
  currentLicense.value = license
  editFormData.value = {
    remark: license.remark || '',
    expires_at: license.expires_at ? new Date(license.expires_at).toISOString().slice(0, 16) : ''
  }
  showEditModal.value = true
}

const handleEdit = async () => {
  if (!currentLicense.value) return
  
  try {
    const updateData: any = {}
    
    if (editFormData.value.remark !== currentLicense.value.remark) {
      updateData.remark = editFormData.value.remark
    }
    
    if (editFormData.value.expires_at) {
      updateData.expires_at = new Date(editFormData.value.expires_at).toISOString()
    }
    
    if (Object.keys(updateData).length === 0) {
      showEditModal.value = false
      return
    }
    
    const response = await api.put<ApiResponse<License>>(API_ROUTES.LICENSES.UPDATE(currentLicense.value.id), updateData)
    if (response.status === 200) {
      showEditModal.value = false
      currentLicense.value = null
      editFormData.value = { remark: '', expires_at: '' }
      toast.success(response.message)
      fetchLicenses()
    } else {
      toast.error(response.message || '更新失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '更新失败')
    console.error(error)
  }
}

const handleFilter = () => {
  currentPage.value = 1 // 重置到第一页
  fetchLicenses()
}

const resetFilters = () => {
  filters.value = {
    activation_status: '',
    enable_status: '',
    product_id: '',
    license_key: '',
    activated_start_date: '',
    activated_end_date: '',
    expires_start_date: '',
    expires_end_date: '',
    remark: ''
  }
  handleFilter()
}

// 导出功能
const handleExport = async () => {
  exportLoading.value = true
  try {
    // 构建导出数据
    const headers = ['ID', '产品', '许可证密钥', '有效期(天)', '状态', '激活时间', '过期时间', '备注']
    const rows = licenses.value.map(license => [
      license.id,
      license.product.name,
      license.license_key,
      license.duration_days,
      license.enable_status === 'disabled' ? '已禁用' :
      license.activation_status === 'expired' ? '已过期' :
      license.activation_status === 'activated' ? '已激活' : '未激活',
      license.activated_at ? new Date(license.activated_at).toLocaleString() : '-',
      license.expires_at ? new Date(license.expires_at).toLocaleString() : '-',
      license.remark || '-'
    ])

    // 创建工作簿
    const wb = XLSX.utils.book_new()
    
    // 创建工作表
    const ws = XLSX.utils.aoa_to_sheet([headers, ...rows])
    
    // 设置列宽
    const colWidths = [
      { wch: 8 },   // ID
      { wch: 20 },  // 产品
      { wch: 32 },  // 许可证密钥
      { wch: 10 },  // 有效期
      { wch: 8 },   // 状态
      { wch: 20 },  // 激活时间
      { wch: 20 },  // 过期时间
      { wch: 20 }   // 备注
    ]
    ws['!cols'] = colWidths

    // 添加工作表到工作簿
    XLSX.utils.book_append_sheet(wb, ws, '许可证列表')

    // 生成 Excel 文件
    const excelBuffer = XLSX.write(wb, { bookType: 'xlsx', type: 'array' })
    const blob = new Blob([excelBuffer], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' })
    
    // 创建下载链接
    const link = document.createElement('a')
    const url = URL.createObjectURL(blob)
    link.setAttribute('href', url)
    link.setAttribute('download', `licenses_${new Date().toISOString().slice(0,10)}.xlsx`)
    link.style.visibility = 'hidden'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    URL.revokeObjectURL(url)

    toast.success('导出成功')
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '导出失败，请稍后重试')
  } finally {
    exportLoading.value = false
  }
}

onMounted(() => {
  fetchLicenses()
  fetchProducts()
})
</script>

<template>
  <div class="space-y-6">
    <div class="flex-between mb-4">
      <h1 class="text-2xl font-bold">许可证管理</h1>
      <div class="flex gap-2">
        <!-- 激活状态快捷筛选按钮 -->
        <div class="relative">
          <button
            class="flex items-center gap-2 px-3 py-1.5 rounded-lg border dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
            @click="showActivationFilter = !showActivationFilter"
          >
            <div :class="[getActivationFilterIcon(), getActivationFilterColor()]"></div>
            <span>{{ getActivationFilterText() }}</span>
            <div class="i-ri-arrow-down-s-line text-sm"></div>
          </button>
          
          <!-- 激活状态下拉菜单 -->
          <div
            v-if="showActivationFilter"
            class="absolute top-full right-0 mt-1 w-32 bg-white dark:bg-gray-800 rounded-lg shadow-lg border dark:border-gray-700 z-10"
          >
            <div class="py-1">
              <button
                class="w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
                :class="activationFilter === 'activated' ? 'text-green-600' : 'text-gray-600 dark:text-gray-300'"
                @click="handleActivationFilter('activated')"
              >
                <div class="i-ri-checkbox-circle-line"></div>
                已激活
              </button>
              <button
                class="w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
                :class="activationFilter === 'inactive' ? 'text-yellow-600' : 'text-gray-600 dark:text-gray-300'"
                @click="handleActivationFilter('inactive')"
              >
                <div class="i-ri-time-line"></div>
                未激活
              </button>
              <button
                class="w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
                :class="activationFilter === 'expired' ? 'text-red-600' : 'text-gray-600 dark:text-gray-300'"
                @click="handleActivationFilter('expired')"
              >
                <div class="i-ri-close-circle-line"></div>
                已过期
              </button>
              <button
                class="w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
                :class="activationFilter === 'all' ? 'text-blue-600' : 'text-gray-600 dark:text-gray-300'"
                @click="handleActivationFilter('all')"
              >
                <div class="i-ri-list-check"></div>
                全部
              </button>
            </div>
          </div>
        </div>
        <!-- 启用状态快捷筛选按钮 -->
        <div class="relative">
          <button
            class="flex items-center gap-2 px-3 py-1.5 rounded-lg border dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
            @click="showQuickFilter = !showQuickFilter"
          >
            <div :class="[getFilterIcon(), getFilterColor()]"></div>
            <span>{{ getFilterText() }}</span>
            <div class="i-ri-arrow-down-s-line text-sm"></div>
          </button>
          
          <!-- 启用状态下拉菜单 -->
          <div
            v-if="showQuickFilter"
            class="absolute top-full right-0 mt-1 w-32 bg-white dark:bg-gray-800 rounded-lg shadow-lg border dark:border-gray-700 z-10"
          >
            <div class="py-1">
              <button
                class="w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
                :class="quickFilter === 'enabled' ? 'text-green-600' : 'text-gray-600 dark:text-gray-300'"
                @click="handleQuickFilter('enabled')"
              >
                <div class="i-ri-checkbox-circle-line"></div>
                已启用
              </button>
              <button
                class="w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
                :class="quickFilter === 'disabled' ? 'text-red-600' : 'text-gray-600 dark:text-gray-300'"
                @click="handleQuickFilter('disabled')"
              >
                <div class="i-ri-close-circle-line"></div>
                已禁用
              </button>
              <button
                class="w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
                :class="quickFilter === 'all' ? 'text-blue-600' : 'text-gray-600 dark:text-gray-300'"
                @click="handleQuickFilter('all')"
              >
                <div class="i-ri-list-check"></div>
                全部
              </button>
            </div>
          </div>
        </div>
        <button 
          class="btn bg-blue-400 text-white hover:bg-blue-500"
          @click="handleExport"
          :disabled="exportLoading"
        >
          <div class="i-ri-download-line mr-1"></div>
          {{ exportLoading ? '导出中...' : '导出' }}
        </button>
        <button class="btn btn-primary" @click="showGenerateModal = true">
          <i class="i-ri-add-line"></i>
          生成许可证
        </button>
      </div>
    </div>

    <!-- 筛选条件部分 -->
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm mb-4 overflow-hidden">
      <div class="flex items-center justify-between p-4 border-b dark:border-gray-700">
        <h3 class="text-lg font-medium text-gray-900 dark:text-gray-100">筛选条件</h3>
        <button 
          class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 transition-colors"
          @click="showFilters = !showFilters"
        >
          <div class="i-ri-arrow-up-s-line text-xl" :class="{ 'rotate-180': !showFilters }"></div>
        </button>
      </div>
      
      <div v-show="showFilters" class="p-4 space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          <!-- 产品筛选 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">产品</label>
            <select
              v-model="filters.product_id"
              class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
            >
              <option value="">全部</option>
              <option v-for="product in products" :key="product.id" :value="product.id">
                {{ product.name }}
              </option>
            </select>
          </div>

          <!-- 许可证密钥 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">许可证密钥</label>
            <input
              v-model="filters.license_key"
              type="text"
              placeholder="输入许可证密钥"
              class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
            />
          </div>

          <!-- 激活时间范围 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">激活时间</label>
            <div class="flex items-center space-x-2">
              <input
                v-model="filters.activated_start_date"
                type="date"
                class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
              />
              <span class="text-gray-500 dark:text-gray-400">-</span>
              <input
                v-model="filters.activated_end_date"
                type="date"
                class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
          </div>

          <!-- 过期时间范围 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">过期时间</label>
            <div class="flex items-center space-x-2">
              <input
                v-model="filters.expires_start_date"
                type="date"
                class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
              />
              <span class="text-gray-500 dark:text-gray-400">-</span>
              <input
                v-model="filters.expires_end_date"
                type="date"
                class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
          </div>

          <!-- 备注 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">备注</label>
            <input
              v-model="filters.remark"
              type="text"
              placeholder="输入备注信息"
              class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
            />
          </div>
        </div>

        <!-- 操作按钮 -->
        <div class="flex justify-end space-x-3 mt-4 pt-4 border-t dark:border-gray-700">
          <button
            @click="resetFilters"
            class="btn mr-4"
          >
            <div class="i-ri-refresh-line mr-1"></div>
            重置
          </button>
          <button
            @click="handleFilter"
            class="btn btn-primary"
          > 
            <div class="i-ri-filter-3-line mr-1"></div>
            筛选
          </button>
        </div>
      </div>
    </div>

    <!-- 许可证列表 -->
    <div class="card">
      <div v-if="loading" class="flex-center py-8">
        <div class="i-ri-loader-4-line animate-spin text-2xl text-primary"></div>
      </div>
      <div v-else class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b dark:border-gray-700">
              <th class="px-4 py-2 text-left">ID</th>
              <th class="px-4 py-2 text-left">产品</th>
              <th class="px-4 py-2 text-left">许可证密钥</th>
              <th class="px-4 py-2 text-left">有效期(天)</th>
              <th class="px-4 py-2 text-left">状态</th>
              <th class="px-4 py-2 text-left">激活时间</th>
              <th class="px-4 py-2 text-left">过期时间</th>
              <th class="px-4 py-2 text-left">备注</th>
              <th class="px-4 py-2 text-left">操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="license in licenses" :key="license.id" class="border-b dark:border-gray-700">
              <td class="px-4 py-2">{{ license.id }}</td>
              <td class="px-4 py-2">{{ license.product.name }}</td>
              <td class="px-4 py-2 font-mono">{{ license.license_key }}</td>
              <td class="px-4 py-2">{{ license.duration_days }}</td>
              <td class="px-4 py-2">
                <span
                  class="px-2 py-1 rounded-full text-xs"
                  :class="{
                    'bg-green-100 text-green-800': license.activation_status === 'activated' && license.enable_status === 'enabled',
                    'bg-yellow-100 text-yellow-800': license.activation_status === 'inactive' && license.enable_status === 'enabled',
                    'bg-red-100 text-red-800': license.enable_status === 'disabled',
                    'bg-gray-100 text-gray-800': license.activation_status === 'expired'
                  }"
                >
                  {{ 
                    license.enable_status === 'disabled' ? '已禁用' :
                    license.activation_status === 'expired' ? '已过期' :
                    license.activation_status === 'activated' ? '已激活' : '未激活'
                  }}
                </span>
              </td>
              <td class="px-4 py-2">{{ license.activated_at ? new Date(license.activated_at).toLocaleString() : '-' }}</td>
              <td class="px-4 py-2">{{ license.expires_at ? new Date(license.expires_at).toLocaleString() : '-' }}</td>
              <td class="px-4 py-2">{{ license.remark || '-' }}</td>
              <td class="px-4 py-2">
                <div class="flex gap-2">
                  <button
                    class="icon-btn"
                    @click="openEditModal(license)"
                    :disabled="license.enable_status === 'disabled'"
                  >
                    <div class="i-ri-edit-line"></div>
                  </button>
                  <button
                    class="icon-btn"
                    @click="license.enable_status === 'enabled' ? handleDisable(license.id) : handleEnable(license.id)"
                    :class="license.enable_status === 'enabled' ? 'text-red-500 hover:text-red-600' : 'text-green-500 hover:text-green-600'"
                  >
                    <div :class="license.enable_status === 'enabled' ? 'i-ri-close-circle-line' : 'i-ri-checkbox-circle-line'"></div>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- 分页 -->
      <div class="flex justify-between items-center mt-4 px-4">
        <div class="flex items-center gap-2">
          <span class="text-sm text-gray-600 dark:text-gray-400">每页显示</span>
          <select
            v-model="pageSize"
            class="px-2 py-1 text-sm rounded border dark:border-gray-700 bg-white dark:bg-gray-800"
            @change="(e) => handlePageSizeChange(Number((e.target as HTMLSelectElement).value))"
          >
            <option v-for="size in pageSizeOptions" :key="size" :value="size">
              {{ size === -1 ? '全部' : size }}
            </option>
          </select>
          <span class="text-sm text-gray-600 dark:text-gray-400">条</span>
        </div>
        
        <div v-if="!showAll" class="flex gap-2">
          <button
            class="px-3 py-1 rounded border dark:border-gray-700"
            :disabled="currentPage === 1"
            @click="handlePageChange(currentPage - 1)"
          >
            上一页
          </button>
          <span class="px-3 py-1">
            第 {{ currentPage }} 页 / 共 {{ Math.ceil(total / pageSize) }} 页
          </span>
          <button
            class="px-3 py-1 rounded border dark:border-gray-700"
            :disabled="currentPage >= Math.ceil(total / pageSize)"
            @click="handlePageChange(currentPage + 1)"
          >
            下一页
          </button>
        </div>
        <div v-else class="text-sm text-gray-600 dark:text-gray-400">
          共 {{ total }} 条数据
        </div>
      </div>
    </div>

    <!-- 生成许可证弹窗 -->
    <div v-if="showGenerateModal" class="fixed inset-0 bg-black/50 flex-center">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 w-96">
        <h2 class="text-xl font-bold mb-4">生成许可证</h2>
        <form @submit.prevent="handleGenerate" class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1">选择产品</label>
            <select
              v-model="formData.product_id"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
              required
            >
              <option value="0">请选择产品</option>
              <option v-for="product in products" :key="product.id" :value="product.id">
                {{ product.name }}
              </option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">有效期(天)</label>
            <input
              v-model="formData.duration_days"
              type="number"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
              required
            >
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">生成数量</label>
            <input
              v-model="formData.count"
              type="number"
              min="1"
              max="100"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
              required
            >
          </div>
          <div class="flex justify-end gap-2 my-4">
            <button
              type="button"
              class="btn"
              @click="showGenerateModal = false"
            >
              取消
            </button>
            <button type="submit" class="btn btn-primary">
              生成
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- 编辑许可证弹窗 -->
    <div v-if="showEditModal" class="fixed inset-0 bg-black/50 flex-center">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 w-96">
        <h2 class="text-xl font-bold mb-4">编辑许可证</h2>
        <form @submit.prevent="handleEdit" class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1">过期时间</label>
            <input
              v-model="editFormData.expires_at"
              type="datetime-local"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
            >
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">备注</label>
            <input
              v-model="editFormData.remark"
              type="text"
              placeholder="输入备注信息"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
            >
          </div>
          <div class="flex justify-end gap-2 my-4">
            <button
              type="button"
              class="btn"
              @click="showEditModal = false"
            >
              取消
            </button>
            <button type="submit" class="btn btn-primary">
              保存
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template> 