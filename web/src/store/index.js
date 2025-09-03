import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    // 用户信息
    userInfo: JSON.parse(localStorage.getItem('userInfo')) || null,
    // 登录状态
    isLoggedIn: !!localStorage.getItem('token'),
    // 当前激活的菜单
    activeMenu: '/dashboard'
  },
  getters: {
    // 获取用户信息
    getUserInfo: state => state.userInfo,
    // 获取登录状态
    getLoginStatus: state => state.isLoggedIn,
    // 获取登录状态（别名）
    isLoggedIn: state => state.isLoggedIn,
    // 获取当前激活菜单
    getActiveMenu: state => state.activeMenu
  },
  mutations: {
    // 设置用户信息
    SET_USER_INFO(state, userInfo) {
      state.userInfo = userInfo
      localStorage.setItem('userInfo', JSON.stringify(userInfo))
    },
    // 设置登录状态
    SET_LOGIN_STATUS(state, status) {
      state.isLoggedIn = status
    },
    // 设置token
    SET_TOKEN(state, token) {
      localStorage.setItem('token', token)
      state.isLoggedIn = true
    },
    // 清除用户信息
    CLEAR_USER_INFO(state) {
      state.userInfo = null
      state.isLoggedIn = false
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
    },
    // 设置当前激活菜单
    SET_ACTIVE_MENU(state, menu) {
      state.activeMenu = menu
    }
  },
  actions: {
    // 登录操作
    login({ commit }, { userInfo, token }) {
      commit('SET_USER_INFO', userInfo)
      commit('SET_TOKEN', token)
    },
    // 退出登录
    logout({ commit }) {
      commit('CLEAR_USER_INFO')
    },
    // 设置激活菜单
    setActiveMenu({ commit }, menu) {
      commit('SET_ACTIVE_MENU', menu)
    }
  },
  modules: {
  }
})