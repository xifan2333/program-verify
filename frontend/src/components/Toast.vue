<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'

interface ToastProps {
  message: string
  type?: 'success' | 'error' | 'warning' | 'info'
  duration?: number
}

const props = withDefaults(defineProps<ToastProps>(), {
  type: 'info',
  duration: 3000
})

const visible = ref(false)
let timer: number | null = null

const show = () => {
  visible.value = true
  if (timer) clearTimeout(timer)
  timer = window.setTimeout(() => {
    visible.value = false
  }, props.duration)
}

onMounted(() => {
  show()
})

onBeforeUnmount(() => {
  if (timer) {
    clearTimeout(timer)
  }
})

defineExpose({
  show
})
</script>

<template>
  <div class="fixed top-4 right-4 z-50">
    <Transition
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="transform translate-x-full opacity-0"
      enter-to-class="transform translate-x-0 opacity-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="transform translate-x-0 opacity-100"
      leave-to-class="transform translate-x-full opacity-0"
    >
      <div
        v-show="visible"
        class="min-w-[300px] p-4 rounded-lg shadow-lg flex items-center gap-2"
        :class="{
          'bg-green-50 dark:bg-green-900/30 text-green-800 dark:text-green-200 border border-green-200 dark:border-green-800': type === 'success',
          'bg-red-50 dark:bg-red-900/30 text-red-800 dark:text-red-200 border border-red-200 dark:border-red-800': type === 'error',
          'bg-yellow-50 dark:bg-yellow-900/30 text-yellow-800 dark:text-yellow-200 border border-yellow-200 dark:border-yellow-800': type === 'warning',
          'bg-blue-50 dark:bg-blue-900/30 text-blue-800 dark:text-blue-200 border border-blue-200 dark:border-blue-800': type === 'info'
        }"
      >
        <div
          class="text-xl"
          :class="{
            'i-ri-checkbox-circle-line text-green-500 dark:text-green-400': type === 'success',
            'i-ri-error-warning-line text-red-500 dark:text-red-400': type === 'error',
            'i-ri-alert-line text-yellow-500 dark:text-yellow-400': type === 'warning',
            'i-ri-information-line text-blue-500 dark:text-blue-400': type === 'info'
          }"
        ></div>
        <p class="flex-1">{{ message }}</p>
      </div>
    </Transition>
  </div>
</template> 