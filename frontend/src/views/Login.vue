<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from '../plugins/toast'
import { API_ROUTES, api } from '../api/config'
import type { ApiResponse, LoginResponse } from '../api/types'

const router = useRouter()
const toast = useToast()
const username = ref('')
const password = ref('')
const loading = ref(false)

const handleLogin = async () => {
  if (!username.value || !password.value) {
    toast.warning('请输入用户名和密码')
    return
  }
  
  loading.value = true
  try {
    const response = await api.post<ApiResponse<LoginResponse>>(
      API_ROUTES.AUTH.LOGIN,
      {
        username: username.value,
        password: password.value
      }
    )
    
    if (response.status === 200) {
      localStorage.setItem('token', response.data.token)
      toast.success(response.message)
      router.replace('/')
    } else {
      toast.error(response.message || '登录失败')
    }
  } catch (error) {
    toast.error(error instanceof Error ? error.message : '登录失败')
    console.error(error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex-center bg-gray-100 dark:bg-gray-900">
    <div class="w-96 p-8 bg-white dark:bg-gray-800 rounded-lg shadow-lg">
      <div class="text-center mb-8">
        <div class="flex-center gap-3 mb-4">
          <div class="i-ri-key-2-line text-4xl text-primary"></div>
          <h1 class="text-2xl font-bold">软件授权管理系统</h1>
        </div>
      </div>
      
      <form @submit.prevent="handleLogin" class="space-y-4">
        <div class="mt-4">
          <label class="block text-sm font-medium mb-2">用户名</label>
          <input
            v-model="username"
            type="text"
            class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
            placeholder="请输入用户名"
            required
          >
        </div>
        
        <div class="mt-4">
          <label class="block text-sm font-medium mb-2">密码</label>
          <input
            v-model="password"
            type="password"
            class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
            placeholder="请输入密码"
            required
          >
        </div>
        
        <button
          type="submit"
          class="w-full btn-primary mt-4"
          :disabled="loading"
        >
          {{ loading ? '登录中...' : '登录' }}
        </button>
      </form>
    </div>
  </div>
</template> 