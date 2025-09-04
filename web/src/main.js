import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import './assets/css/global.css'
import axios from 'axios'

// 配置axios
axios.defaults.baseURL = 'http://127.0.0.1:8888/v1/'
axios.defaults.timeout = 10000

// 自定义JSON解析器，处理long类型ID精度丢失
function parseJSONWithLongIds(jsonString) {
  if (typeof jsonString !== 'string') {
    return jsonString
  }
  
  try {
    // 使用正则表达式将所有"id":数字 替换为 "id":"数字"
    const processedString = jsonString.replace(/"id":\s*(\d+)/g, '"id":"$1"')
    return JSON.parse(processedString)
  } catch (e) {
    // 如果解析失败，尝试原始解析
    try {
      return JSON.parse(jsonString)
    } catch (e2) {
      return jsonString
    }
  }
}

// 配置自定义响应转换器
axios.defaults.transformResponse = [
  function (data) {
    return parseJSONWithLongIds(data)
  }
]

// 请求拦截器
axios.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器
axios.interceptors.response.use(
  response => {
    return response
  },
  error => {
    if (error.response && error.response.status === 401) {
      // 清除token并跳转到登录页
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
      router.push('/login')
    }
    return Promise.reject(error)
  }
)

Vue.prototype.$http = axios

Vue.use(ElementUI)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
