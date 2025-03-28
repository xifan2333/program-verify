<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch, nextTick } from 'vue'
import { useToast } from '../plugins/toast'
import type { StatsSummary, ProductActivation, RevenueTrendData } from '../api/api'
import { api, API_ROUTES } from '../api/config'
import * as echarts from 'echarts'

const toast = useToast()
const loading = ref(false)
const revenueChartRef = ref<HTMLElement>()
const activationChartRef = ref<HTMLElement>()
let revenueChart: echarts.ECharts | null = null
let activationChart: echarts.ECharts | null = null

// 统计数据
const stats = ref<StatsSummary>({
  products: {
    total: 0,
    enabled: 0,
    disabled: 0
  },
  licenses: {
    total: 0,
    activated: 0,
    inactive: 0,
    expired: 0,
    disabled: 0
  },
  revenue: {
    total: 0,
    today: 0,
    this_month: 0
  }
})

// 产品激活数据
const productActivation = ref<ProductActivation>({
  products: [],
  summary: {
    total_licenses: 0,
    total_activated: 0,
    total_revenue: 0,
    avg_activation_rate: 0
  }
})
// 收益趋势数据
const revenueTrend = ref<RevenueTrendData>({
  dates: [],
  revenue: [],
  total: 0,
  average: 0
})

// 获取统计数据
const fetchStats = async () => {
  loading.value = true
  try {
    const [statsRes, trendRes, activationRes] = await Promise.all([
      api.get(API_ROUTES.ANALYTICS.STATS),
      api.get(API_ROUTES.ANALYTICS.REVENUE_TREND),
      api.get(API_ROUTES.ANALYTICS.PRODUCT_ACTIVATION)
    ])

    if (statsRes.status === 200) {
      stats.value = statsRes.data
    }
    if (trendRes.status === 200) {
      revenueTrend.value = {
        dates: trendRes.data.dates,
        revenue: trendRes.data.revenue,
        total: trendRes.data.total,
        average: trendRes.data.average
      }
    }
    if (activationRes.status === 200) {
      productActivation.value = activationRes.data
    }
  } catch (error) {
    toast.error('获取统计数据失败')
    console.error('Failed to fetch stats:', error)
  } finally {
    loading.value = false
  }
}

// 初始化收益趋势图表
const initRevenueChart = () => {
  if (!revenueChartRef.value) return
  
  revenueChart = echarts.init(revenueChartRef.value)
  const option = {
    title: {
      text: '收益趋势',
      left: 'center'
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'line'
      },
      formatter: function(params: any[]) {
        return `${params[0].name}<br/>收益: ¥${params[0].value}`
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: revenueTrend.value.dates,
      axisLabel: {
        interval: 0,
        rotate: 30
      }
    },
    yAxis: {
      type: 'value',
      name: '收益 (¥)'
    },
    series: [
      {
        name: '日收益',
        type: 'line',
        data: revenueTrend.value.revenue,
        itemStyle: {
          color: '#10b981'
        },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: 'rgba(16, 185, 129, 0.3)'
            },
            {
              offset: 1,
              color: 'rgba(16, 185, 129, 0.1)'
            }
          ])
        }
      }
    ]
  }
  revenueChart.setOption(option)
}

// 初始化产品激活图表
const initActivationChart = () => {
  if (!activationChartRef.value) return
  
  activationChart = echarts.init(activationChartRef.value)
  const option = {
    title: {
      text: '产品激活量',
      left: 'center'
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      },
      formatter: function(params: any[]) {
        const data = params[0].data
        return `${params[0].name}<br/>
                激活量: ${data.activated_licenses}<br/>
                总许可: ${data.total_licenses}<br/>
                激活率: ${data.activation_rate.toFixed(2)}%<br/>
                收益: ¥${data.revenue}`
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: productActivation.value.products.map(item => item.name),
      axisLabel: {
        interval: 0,
        rotate: 30
      }
    },
    yAxis: [
      {
        type: 'value',
        name: '许可证数量',
        position: 'left'
      },
      {
        type: 'value',
        name: '激活率 (%)',
        position: 'right',
        max: 100,
        axisLabel: {
          formatter: '{value}%'
        }
      }
    ],
    series: [
      {
        name: '许可证数量',
        type: 'bar',
        barWidth: '40%',
        data: productActivation.value.products.map(item => ({
          value: item.total_licenses,
          activated_licenses: item.activated_licenses,
          total_licenses: item.total_licenses,
          activation_rate: item.activation_rate,
          revenue: item.revenue
        })),
        itemStyle: {
          color: '#10b981'
        }
      },
      {
        name: '激活率',
        type: 'line',
        yAxisIndex: 1,
        data: productActivation.value.products.map(item => item.activation_rate),
        itemStyle: {
          color: '#3b82f6'
        }
      }
    ]
  }
  activationChart.setOption(option)
}

// 监听窗口大小变化，调整图表大小
const handleResize = () => {
  revenueChart?.resize()
  activationChart?.resize()
}

// 监听数据变化，更新图表
watch([() => revenueTrend.value, () => productActivation.value], ([newRevenueTrend, newProductActivation]) => {
  if (newRevenueTrend) {
    nextTick(() => {
      initRevenueChart()
    })
  }
  if (newProductActivation?.products?.length > 0) {
    nextTick(() => {
      initActivationChart()
    })
  }
}, { deep: true })

// 在组件挂载时初始化
onMounted(async () => {
  await fetchStats()
  nextTick(() => {
    if (revenueTrend.value?.dates?.length > 0) {
      initRevenueChart()
    }
    if (productActivation.value?.products?.length > 0) {
      initActivationChart()
    }
  })
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  revenueChart?.dispose()
  activationChart?.dispose()
})
</script>

<template>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold">系统概览</h1>
    
    <!-- 加载状态 -->
    <div v-if="loading" class="flex-center py-8">
      <div class="i-ri-loader-4-line animate-spin text-2xl text-primary"></div>
    </div>
    
    <div v-else>
      <!-- 许可证统计 -->
      <div class="mb-8">
        <h2 class="text-xl font-bold mb-4">许可证统计</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <!-- 已激活许可证 -->
          <div class="card">
            <div class="flex items-center gap-4">
              <div class="i-ri-shield-check-line text-3xl text-green-500"></div>
              <div>
                <h3 class="text-lg font-medium">已激活</h3>
                <p class="text-2xl font-bold">{{ stats.licenses.activated }}</p>
              </div>
            </div>
          </div>
          
          <!-- 未激活许可证 -->
          <div class="card">
            <div class="flex items-center gap-4">
              <div class="i-ri-time-line text-3xl text-yellow-500"></div>
              <div>
                <h3 class="text-lg font-medium">未激活</h3>
                <p class="text-2xl font-bold">{{ stats.licenses.inactive }}</p>
              </div>
            </div>
          </div>

          <!-- 已过期许可证 -->
          <div class="card">
            <div class="flex items-center gap-4">
              <div class="i-ri-close-circle-line text-3xl text-red-500"></div>
              <div>
                <h3 class="text-lg font-medium">已过期</h3>
                <p class="text-2xl font-bold">{{ stats.licenses.expired }}</p>
              </div>
            </div>
          </div>

          <!-- 已禁用许可证 -->
          <div class="card">
            <div class="flex items-center gap-4">
              <div class="i-ri-forbid-line text-3xl text-gray-500"></div>
              <div>
                <h3 class="text-lg font-medium">已禁用</h3>
                <p class="text-2xl font-bold">{{ stats.licenses.disabled }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 收益统计 -->
      <div class="mb-8">
        <h2 class="text-xl font-bold mb-4">收益统计</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <!-- 总收益 -->
          <div class="card">
            <div class="flex items-center gap-4">
              <div class="i-ri-money-cny-circle-line text-3xl text-blue-500"></div>
              <div>
                <h3 class="text-lg font-medium">总收益</h3>
                <p class="text-2xl font-bold">¥{{ stats.revenue.total.toFixed(2) }}</p>
              </div>
            </div>
          </div>

          <!-- 今日收益 -->
          <div class="card">
            <div class="flex items-center gap-4">
              <div class="i-ri-calendar-todo-line text-3xl text-green-500"></div>
              <div>
                <h3 class="text-lg font-medium">今日收益</h3>
                <p class="text-2xl font-bold">¥{{ stats.revenue.today.toFixed(2) }}</p>
              </div>
            </div>
          </div>

          <!-- 本月收益 -->
          <div class="card">
            <div class="flex items-center gap-4">
              <div class="i-ri-calendar-line text-3xl text-purple-500"></div>
              <div>
                <h3 class="text-lg font-medium">本月收益</h3>
                <p class="text-2xl font-bold">¥{{ stats.revenue.this_month.toFixed(2) }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 图表展示 -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- 收益趋势图 -->
        <div class="card">
          <div ref="revenueChartRef" class="w-full h-80"></div>
        </div>
        
        <!-- 产品激活量图 -->
        <div class="card">
          <div ref="activationChartRef" class="w-full h-80"></div>
        </div>
      </div>
    </div>
  </div>
</template> 