import { defineConfig, presetIcons, presetTypography, presetMini } from 'unocss'

export default defineConfig({
  presets: [
    presetIcons({
      scale: 1.2,
      cdn: 'https://esm.sh/',
      collections: {
        ri: () => import('@iconify-json/ri/icons.json').then(i => i.default),
      },
      extraProperties: {
        'display': 'inline-block',
        'vertical-align': 'middle'
      },
    }),
    presetTypography(),
    presetMini(),
  ],
  theme: {
    colors: {
      primary: '#42b883',
      secondary: '#35495e',
      accent: '#3eaf7c',
      dark: '#1a1a1a',
      light: '#ffffff',
    },
    breakpoints: {
      'sm': '640px',
      'md': '768px',
      'lg': '1024px',
      'xl': '1280px',
      '2xl': '1536px',
    },
  },
  rules: [
    ['text-balance', { 'text-wrap': 'balance' }],
    ['text-pretty', { 'text-wrap': 'pretty' }],
    ['shadow-soft', { 'box-shadow': '0 2px 12px 0 rgba(0, 0, 0, 0.1)' }],
  ],
  shortcuts: [
    // 基础按钮
    ['btn', 'px-4 py-2 rounded-lg inline-block transition-all duration-200 ease-in-out cursor-pointer bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 active:scale-95'],
    ['btn-primary', 'btn bg-primary text-white hover:bg-[#3aa876] active:scale-95'],
    ['btn-secondary', 'btn bg-secondary text-white hover:bg-opacity-90 active:scale-95'],
    ['btn-outline', 'btn border-2 border-primary text-primary hover:bg-primary hover:text-white active:scale-95'],
    
    // 布局
    ['container', 'max-w-7xl mx-auto px-4 sm:px-6 lg:px-8'],
    ['flex-center', 'flex items-center justify-center'],
    ['flex-between', 'flex items-center justify-between'],
    ['grid-auto-fit', 'grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3'],
    
    // 卡片
    ['card', 'p-6 rounded-xl border border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900 shadow-soft hover:shadow-lg transition-all duration-200'],
    
    // 导航
    ['nav-link', 'text-gray-600 dark:text-gray-300 hover:text-primary dark:hover:text-primary transition-colors duration-200'],
    
    // 响应式文本
    ['h1-responsive', 'text-3xl sm:text-4xl lg:text-5xl font-bold'],
    ['h2-responsive', 'text-2xl sm:text-3xl lg:text-4xl font-bold'],
    ['h3-responsive', 'text-xl sm:text-2xl lg:text-3xl font-bold'],
    
    // 间距
    ['section', 'py-12 sm:py-16 lg:py-20'],
    ['section-sm', 'py-8 sm:py-12 lg:py-16'],
    
  ],
}) 