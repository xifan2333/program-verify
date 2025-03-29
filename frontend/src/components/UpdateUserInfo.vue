<script setup lang="ts">
import { ref } from 'vue'
import { useToast } from '../plugins/toast'
import { API_ROUTES, api } from '../api/config'
import type { ApiResponse, UpdateUserRequest, UpdateUserResponse } from '../api/types'

const props = defineProps<{
  show: boolean
}>()

const emit = defineEmits<{
  (e: 'update:show', value: boolean): void
  (e: 'update:success', data: UpdateUserResponse): void
}>()

const toast = useToast()
const currentPassword = ref('')
const newUsername = ref('')
const newPassword = ref('')
const loading = ref(false)

const handleUpdate = async () => {
  if (!currentPassword.value || !newUsername.value || !newPassword.value) {
    toast.warning('请填写所有字段')
    return
  }
  
  loading.value = true
  try {
    const requestData: UpdateUserRequest = {
      current_password: currentPassword.value,
      new_username: newUsername.value,
      new_password: newPassword.value
    }

    const response = await api.put<ApiResponse<UpdateUserResponse>>(
      API_ROUTES.USER.UPDATE,
      requestData
    )
    
    if (response.status === 200 && response.data) {
      emit('update:success', response.data)
      emit('update:show', false)
      toast.success(response.message)
    } else {
      toast.error(response.message || '更新失败')
    }
  } catch (error) {
    console.error('更新错误:', error)
    toast.error('更新失败，请稍后重试')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div v-if="show" class="fixed inset-0 bg-black bg-opacity-50 flex-center z-50">
    <div class="w-96 p-8 bg-white dark:bg-gray-800 rounded-lg shadow-lg">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-xl font-bold">修改用户信息</h2>
        <button @click="emit('update:show', false)" class="text-gray-500 hover:text-gray-700">
          <div class="i-ri-close-line text-xl"></div>
        </button>
      </div>
      
      <form @submit.prevent="handleUpdate" class="space-y-4">
        <div>
          <label class="block text-sm font-medium mb-2">当前密码</label>
          <input
            v-model="currentPassword"
            type="password"
            class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
            placeholder="请输入当前密码"
            required
          >
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-2">新用户名</label>
          <input
            v-model="newUsername"
            type="text"
            class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
            placeholder="请输入新用户名"
            required
          >
        </div>
        
        <div>
          <label class="block text-sm font-medium mb-2">新密码</label>
          <input
            v-model="newPassword"
            type="password"
            class="w-full px-4 py-2 rounded-lg border dark:border-gray-700 bg-white dark:bg-gray-900"
            placeholder="请输入新密码"
            required
          >
        </div>
        
        <button
          type="submit"
          class="w-full btn-primary mt-4"
          :disabled="loading"
        >
          {{ loading ? '更新中...' : '更新' }}
        </button>
      </form>
    </div>
  </div>
</template> 