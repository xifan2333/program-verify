import { createVNode, render, type App } from 'vue'
import Toast from '../components/Toast.vue'

interface ToastOptions {
  message: string
  type?: 'success' | 'error' | 'warning' | 'info'
  duration?: number
}

class ToastManager {
  private container: HTMLElement | null = null
  private vnode: any = null

  constructor() {
    this.init()
  }

  private init() {
    if (!this.container) {
      this.container = document.createElement('div')
      document.body.appendChild(this.container)
    }
  }

  create(options: ToastOptions) {
    // 如果已经有实例，先销毁它
    if (this.vnode) {
      render(null, this.container!)
      this.vnode = null
    }

    // 创建新的实例
    this.vnode = createVNode(Toast, {
      message: options.message,
      type: options.type || 'info',
      duration: options.duration
    })

    render(this.vnode, this.container!)
    this.vnode.component?.exposed?.show()
  }
}

const toastManager = new ToastManager()

export const useToast = () => {
  return {
    success: (message: string, duration?: number) => toastManager.create({ message, type: 'success', duration }),
    error: (message: string, duration?: number) => toastManager.create({ message, type: 'error', duration }),
    warning: (message: string, duration?: number) => toastManager.create({ message, type: 'warning', duration }),
    info: (message: string, duration?: number) => toastManager.create({ message, type: 'info', duration })
  }
}

export default {
  install: (app: App) => {
    app.config.globalProperties.$toast = useToast()
  }
} 