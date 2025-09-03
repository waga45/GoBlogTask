<template>
  <div class="layout-container">
    <!-- 顶部栏 -->
    <div class="layout-header">
      <div class="logo">
        <h2 style="margin: 0; color: #1890ff;">GoBlog 管理后台</h2>
      </div>
      
      <!-- 用户信息区域 -->
      <div class="user-info">
        <span class="username">{{ userInfo.username || '用户' }}</span>
        <el-button 
          type="danger" 
          size="small" 
          @click="handleLogout"
        >
          退出
        </el-button>
      </div>
    </div>
    
    <!-- 主体内容区域 -->
    <div class="layout-body">
      <!-- 左侧菜单栏 -->
      <div class="layout-sidebar">
        <el-menu
          :default-active="activeMenu"
          class="sidebar-menu"
          background-color="#001529"
          text-color="rgba(255, 255, 255, 0.65)"
          active-text-color="#fff"
          router
        >
          <el-menu-item index="/articles">
            <i class="el-icon-document"></i>
            <span slot="title">文章列表</span>
          </el-menu-item>
          
          <el-menu-item index="/write">
            <i class="el-icon-edit"></i>
            <span slot="title">写文章</span>
          </el-menu-item>
        </el-menu>
      </div>
      
      <!-- 右侧内容区域 -->
      <div class="layout-content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Layout',
  computed: {
    ...mapGetters({
      userInfo: 'getUserInfo',
      activeMenu: 'getActiveMenu'
    })
  },
  methods: {
    /**
     * 处理用户退出登录
     */
    handleLogout() {
      this.$confirm('确定要退出登录吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 清除用户信息
        this.$store.dispatch('logout')
        
        // 跳转到登录页
        this.$router.push('/login')
        
        this.$message.success('已退出登录')
      }).catch(() => {
        // 用户取消退出
      })
    }
  },
  
  created() {
    // 检查用户信息是否存在
    if (!this.userInfo) {
      // 如果没有用户信息，尝试从localStorage获取
      const userInfo = localStorage.getItem('userInfo')
      if (userInfo) {
        this.$store.commit('SET_USER_INFO', JSON.parse(userInfo))
      } else {
        // 如果没有用户信息，跳转到登录页
        this.$router.push('/login')
      }
    }
  }
}
</script>

<style scoped>
.layout-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.layout-header {
  height: 60px;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  z-index: 1000;
}

.logo h2 {
  margin: 0;
  color: #1890ff;
  font-weight: 600;
}

.layout-body {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.layout-sidebar {
  width: 200px;
  background: #001529;
  overflow-y: auto;
}

.sidebar-menu {
  border-right: none;
  height: 100%;
}

.layout-content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  background: #f0f2f5;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 15px;
}

.username {
  color: #333;
  font-weight: 500;
  font-size: 14px;
}

/* 菜单项样式 */
.sidebar-menu .el-menu-item {
  color: rgba(255, 255, 255, 0.65);
  border-right: none;
}

.sidebar-menu .el-menu-item:hover {
  background-color: #1890ff !important;
  color: #fff;
}

.sidebar-menu .el-menu-item.is-active {
  background-color: #1890ff !important;
  color: #fff;
}

.sidebar-menu .el-menu-item i {
  margin-right: 8px;
}
</style>