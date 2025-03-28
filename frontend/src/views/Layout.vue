<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isCollapsed = ref(false)

const handleLogout = () => {
  localStorage.removeItem('token')
  router.push('/login')
}

const menuItems = [
  { path: '/', icon: 'i-ri-home-4-line', label: '首页' },
  { path: '/products', icon: 'i-ri-apps-2-line', label: '产品管理' },
  { path: '/licenses', icon: 'i-ri-key-2-line', label: '许可证管理' }
]
</script>

<template>
  <div class="min-h-screen bg-gray-100 dark:bg-gray-900">
    <!-- 侧边栏 -->
    <aside
      class="fixed left-0 top-0 h-screen bg-white dark:bg-gray-800 shadow-lg transition-all duration-300"
      :class="isCollapsed ? 'w-16' : 'w-64'"
    >
      <div class="flex items-center justify-between  border-b dark:border-gray-700">
        <div class="overflow-hidden transition-all duration-300" :style="{ width: isCollapsed ? '0' : '160px' }">
          <h1 class="text-xl font-bold text-primary whitespace-nowrap py-4 pl-4">软件授权</h1>
        </div>
        <button class="icon-btn" @click="isCollapsed = !isCollapsed">
          <div :class="isCollapsed ? 'i-ri-menu-unfold-line mr-6' : 'i-ri-menu-fold-line mr-4'"></div>
        </button>
      </div>
      <nav :class="isCollapsed ? 'px-2' : 'px-4'">
        <ul class="space-y-2">
          <li v-for="item in menuItems" :key="item.path">
            <router-link
              :to="item.path"
              class="flex items-center rounded-lg transition-all duration-300 hover:bg-gray-100 my-2 p-3"
              :class="[
                { 'bg-primary/10 text-primary': $route.path === item.path },
                isCollapsed ? 'justify-center' : 'gap-3'
              ]"
            >
              <div 
                :class="[
                  item.icon,
                  'text-xl'
                ]"
              ></div>
              <div class="overflow-hidden transition-all duration-300" :style="{ width: isCollapsed ? '0' : '120px' }">
                <span class="whitespace-nowrap">{{ item.label }}</span>
              </div>
            </router-link>
          </li>
        </ul>
      </nav>
    </aside>

    <!-- 主内容区 -->
    <main
      class="transition-all duration-300 min-h-screen"
      :class="isCollapsed ? 'ml-16' : 'ml-64'"
    >
      <!-- 顶部导航栏 -->
      <header class="h-16 bg-white dark:bg-gray-800 shadow-sm flex items-center justify-between px-6">
        <h2 class="text-lg font-semibold">{{ $route.name }}</h2>
        <button class="icon-btn" @click="handleLogout">
          <div class="i-ri-logout-box-line"></div>
        </button>
      </header>

      <!-- 页面内容 -->
      <div class="p-6">
        <router-view></router-view>
      </div>
    </main>
  </div>
</template> 