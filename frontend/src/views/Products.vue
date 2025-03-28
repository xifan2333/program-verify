<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useToast } from '../plugins/toast'
import type { Product } from '../api/api'
import { api, API_ROUTES } from '../api/config'
import * as XLSX from 'xlsx'

const toast = useToast()
const products = ref<Product[]>([])
const loading = ref(false)
const exportLoading = ref(false)
const showCreateModal = ref(false)
const showEditModal = ref(false)
const currentProduct = ref<Product | null>(null)
const showFilters = ref(false)
const showQuickFilter = ref(false)

// 筛选条件
const filters = ref({
  status: 'enabled', // 默认显示启用状态
  name: '',
  min_price: '',
  max_price: '',
  start_date: '',
  end_date: ''
})

// 快捷筛选状态
const quickFilter = ref<'enabled' | 'disabled' | 'all'>('enabled')

const formData = ref({
  name: '',
  price: 0
})

// 分页相关
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const showAll = ref(false)

const pageSizeOptions = [10, 20, 50, 100, -1] // -1 表示显示全部

// 从 URL 获取初始查询参数
const initFromQuery = () => {
  const urlParams = new URLSearchParams(window.location.search)
  
  filters.value = {
    status: urlParams.get('status') || '',
    name: urlParams.get('name') || '',
    min_price: urlParams.get('min_price') || '',
    max_price: urlParams.get('max_price') || '',
    start_date: urlParams.get('start_date') || '',
    end_date: urlParams.get('end_date') || ''
  }
  
  currentPage.value = parseInt(urlParams.get('page') || '1')
  pageSize.value = parseInt(urlParams.get('page_size') || '10')
}

// 更新 URL 参数
const updateQueryParams = () => {
  const params = new URLSearchParams()
  
  // 添加筛选条件
  if (filters.value.status) params.set('status', filters.value.status)
  if (filters.value.name) params.set('name', filters.value.name)
  if (filters.value.min_price) params.set('min_price', filters.value.min_price)
  if (filters.value.max_price) params.set('max_price', filters.value.max_price)
  if (filters.value.start_date) params.set('start_date', filters.value.start_date)
  if (filters.value.end_date) params.set('end_date', filters.value.end_date)
  
  // 添加分页参数
  params.set('page', currentPage.value.toString())
  params.set('page_size', pageSize.value.toString())
  
  // 更新 URL，不刷新页面
  const newUrl = `${window.location.pathname}?${params.toString()}`
  window.history.pushState({}, '', newUrl)
}

const handlePageSizeChange = (size: number) => {
  pageSize.value = size
  showAll.value = size === -1
  currentPage.value = 1 // 切换每页数量时重置到第一页
  updateQueryParams()
  fetchProducts()
}

const handleFilter = () => {
  currentPage.value = 1 // 重置到第一页
  updateQueryParams()
  fetchProducts()
}

const handlePageChange = (page: number) => {
  currentPage.value = page
  updateQueryParams()
  fetchProducts()
}

const fetchProducts = async () => {
  loading.value = true
  try {
    const params = {
      ...(showAll.value ? {} : {
        page: currentPage.value.toString(),
        page_size: pageSize.value.toString(),
      }),
      ...(filters.value.status && { status: filters.value.status }),
      ...(filters.value.name && { name: filters.value.name }),
      ...(filters.value.min_price && { min_price: filters.value.min_price }),
      ...(filters.value.max_price && { max_price: filters.value.max_price }),
      ...(filters.value.start_date && { start_date: filters.value.start_date }),
      ...(filters.value.end_date && { end_date: filters.value.end_date })
    }

    const data = await api.get(API_ROUTES.PRODUCTS.LIST, params)
    products.value = data.data.items
    total.value = data.data.total
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '获取产品列表失败，请稍后重试')
  } finally {
    loading.value = false
  }
}

// 表单验证
const validateProductForm = () => {
  if (!formData.value.name || formData.value.price <= 0) {
    toast.warning('请填写完整的产品信息')
    return false
  }
  return true
}

const handleCreate = async () => {
  if (!validateProductForm()) return
  
  try {
    await api.post(API_ROUTES.PRODUCTS.CREATE, formData.value)
    showCreateModal.value = false
    formData.value = { name: '', price: 0 }
    fetchProducts()
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '创建失败，请稍后重试')
  }
}

const handleEdit = async () => {
  if (!currentProduct.value) return
  
  if (!validateProductForm()) return
  
  try {
    const data = await api.put(API_ROUTES.PRODUCTS.UPDATE(currentProduct.value.id), formData.value)
    if (data.status === 200) {
      showEditModal.value = false
      currentProduct.value = null
      formData.value = { name: '', price: 0 }
      fetchProducts()
    } else {
      toast.error(data.message || '更新失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '更新失败，请稍后重试')
  }
}

const handleDisable = async (id: number) => {
  try {
    const data = await api.put(API_ROUTES.PRODUCTS.UPDATE(id), { status: 'disabled' })
    if (data.status === 200) {
      toast.success(data.message)
      fetchProducts()
    } else {
      toast.error(data.message || '禁用产品失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '操作失败，请稍后重试')
  }
}

const handleEnable = async (id: number) => {
  try {
    const data = await api.put(API_ROUTES.PRODUCTS.UPDATE(id), { status: 'enabled' })
    if (data.status === 200) {
      toast.success(data.message)
      fetchProducts()
    } else {
      toast.error(data.message || '重新启用产品失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '操作失败，请稍后重试')
  }
}

const openEditModal = (product: Product) => {
  currentProduct.value = product
  formData.value = {
    name: product.name,
    price: product.price
  }
  showEditModal.value = true
}

const resetFilters = () => {
  filters.value = {
    status: 'enabled',
    name: '',
    min_price: '',
    max_price: '',
    start_date: '',
    end_date: ''
  }
  handleFilter()
}

// 导出功能
const handleExport = async () => {
  exportLoading.value = true
  try {
    // 构建导出数据
    const headers = ['ID', '名称', '价格', '状态', '创建时间']
    const rows = products.value.map(product => [
      product.id,
      product.name,
      product.price,
      product.status === 'enabled' ? '启用' : '禁用',
      new Date(product.created_at).toLocaleString()
    ])

    // 创建工作簿
    const wb = XLSX.utils.book_new()
    
    // 创建工作表
    const ws = XLSX.utils.aoa_to_sheet([headers, ...rows])
    
    // 设置列宽
    const colWidths = [
      { wch: 8 },  // ID
      { wch: 20 }, // 名称
      { wch: 10 }, // 价格
      { wch: 8 },  // 状态
      { wch: 20 }  // 创建时间
    ]
    ws['!cols'] = colWidths

    // 添加工作表到工作簿
    XLSX.utils.book_append_sheet(wb, ws, '产品列表')

    // 生成 Excel 文件
    const excelBuffer = XLSX.write(wb, { bookType: 'xlsx', type: 'array' })
    const blob = new Blob([excelBuffer], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' })
    
    // 创建下载链接
    const link = document.createElement('a')
    const url = URL.createObjectURL(blob)
    link.setAttribute('href', url)
    link.setAttribute('download', `products_${new Date().toISOString().slice(0,10)}.xlsx`)
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

// 快捷筛选处理
const handleQuickFilter = (status: 'enabled' | 'disabled' | 'all') => {
  quickFilter.value = status
  filters.value.status = status === 'all' ? '' : status
  showQuickFilter.value = false
  handleFilter()
}

// 获取当前筛选状态文本
const getFilterText = () => {
  switch (quickFilter.value) {
    case 'enabled':
      return '启用'
    case 'disabled':
      return '禁用'
    case 'all':
      return '启用状态'
    default:
      return '启用状态'
  }
}

// 获取当前筛选状态图标
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

// 获取当前筛选状态颜色
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

onMounted(() => {
  initFromQuery()
  fetchProducts()
})
</script>

<template>
  <div class="space-y-6">
    <div class="flex-between mb-4">
      <h1 class="text-2xl font-bold">产品管理</h1>
      <div class="flex gap-2">
        <!-- 快捷筛选按钮 -->
        <div class="relative">
          <button
            class="flex items-center gap-2 px-3 py-1.5 rounded-lg border dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
            @click="showQuickFilter = !showQuickFilter"
          >
            <div :class="[getFilterIcon(), getFilterColor()]"></div>
            <span>{{ getFilterText() }}</span>
            <div class="i-ri-arrow-down-s-line text-sm"></div>
          </button>
          
          <!-- 下拉菜单 -->
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
                启用
              </button>
              <button
                class="w-full px-4 py-2 text-left hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
                :class="quickFilter === 'disabled' ? 'text-red-600' : 'text-gray-600 dark:text-gray-300'"
                @click="handleQuickFilter('disabled')"
              >
                <div class="i-ri-close-circle-line"></div>
                禁用
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
        <button class="btn btn-primary" @click="showCreateModal = true">
          <i class="i-ri-add-line"></i>
          创建产品
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
          <!-- 状态筛选 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">状态</label>
            <select
              v-model="filters.status"
              class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
            >
              <option value="">全部</option>
              <option value="enabled">启用</option>
              <option value="disabled">禁用</option>
            </select>
          </div>

          <!-- 名称筛选 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">名称</label>
            <input
              v-model="filters.name"
              type="text"
              placeholder="输入产品名称"
              class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
            />
          </div>

          <!-- 价格范围 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">价格范围</label>
            <div class="flex items-center space-x-2">
              <input
                v-model="filters.min_price"
                type="number"
                placeholder="最低价"
                class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
              />
              <span class="text-gray-500 dark:text-gray-400">-</span>
              <input
                v-model="filters.max_price"
                type="number"
                placeholder="最高价"
                class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
          </div>

          <!-- 创建时间范围 -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">创建时间</label>
            <div class="flex items-center space-x-2">
              <input
                v-model="filters.start_date"
                type="date"
                class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
              />
              <span class="text-gray-500 dark:text-gray-400">-</span>
              <input
                v-model="filters.end_date"
                type="date"
                class="block w-full px-3 py-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
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

    <!-- 产品列表 -->
    <div class="card">
      <div v-if="loading" class="flex-center py-8">
        <div class="i-ri-loader-4-line animate-spin text-2xl text-primary"></div>
      </div>
      <div v-else class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b dark:border-gray-700">
              <th class="px-4 py-2 text-left">ID</th>
              <th class="px-4 py-2 text-left">名称</th>
              <th class="px-4 py-2 text-left">价格</th>
              <th class="px-4 py-2 text-left">状态</th>
              <th class="px-4 py-2 text-left">创建时间</th>
              <th class="px-4 py-2 text-left">操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="product in products" :key="product.id" class="border-b dark:border-gray-700">
              <td class="px-4 py-2">{{ product.id }}</td>
              <td class="px-4 py-2">{{ product.name }}</td>
              <td class="px-4 py-2">¥{{ product.price }}</td>
              <td class="px-4 py-2">
                <span
                  class="px-2 py-1 rounded-full text-xs"
                  :class="product.status === 'enabled' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'"
                >
                  {{ product.status === 'enabled' ? '启用' : '禁用' }}
                </span>
              </td>
              <td class="px-4 py-2">{{ new Date(product.created_at).toLocaleString() }}</td>
              <td class="px-4 py-2">
                <div class="flex gap-2">
                  <button
                    class="icon-btn"
                    @click="openEditModal(product)"
                    :disabled="product.status !== 'enabled'"
                  >
                    <div class="i-ri-edit-line"></div>
                  </button>
                  <button
                    class="icon-btn"
                    @click="product.status === 'enabled' ? handleDisable(product.id) : handleEnable(product.id)"
                    :class="product.status === 'enabled' ? 'text-red-500 hover:text-red-600' : 'text-green-500 hover:text-green-600'"
                  >
                    <div :class="product.status === 'enabled' ? 'i-ri-close-circle-line' : 'i-ri-checkbox-circle-line'"></div>
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

    <!-- 创建产品弹窗 -->
    <div v-if="showCreateModal" class="fixed inset-0 bg-black/50 flex-center">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 w-96">
        <h2 class="text-xl font-bold mb-4">创建产品</h2>
        <form @submit.prevent="handleCreate" class="space-y-4">
          <div>
            <label class="block text-sm font-medium mb-1">产品名称</label>
            <input
              v-model="formData.name"
              type="text"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
              required
            >
          </div>
          <div>
            <label class="block text-sm font-medium mb-1">价格</label>
            <input
              v-model="formData.price"
              type="number"
              step="0.01"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
              required
            >
          </div>
          <div class="flex justify-end gap-2 my-4">
            <button
              type="button"
              class="btn"
              @click="showCreateModal = false"
            >
              取消
            </button>
            <button type="submit" class="btn btn-primary">
              创建
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- 编辑产品弹窗 -->
    <div v-if="showEditModal" class="fixed inset-0 bg-black/50 flex-center">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 w-96">
        <h2 class="text-xl font-bold mb-4">编辑产品</h2>
        <form @submit.prevent="handleEdit" class="space-y-4">
          <div>
            <label class="block text-sm font-medium mx-2">产品名称</label>
            <input
              v-model="formData.name"
              type="text"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
              required
            >
          </div>
          <div>
            <label class="block text-sm font-medium mx-2">价格</label>
            <input
              v-model="formData.price"
              type="number"
              step="0.01"
              class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
              required
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