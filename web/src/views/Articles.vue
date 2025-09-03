<template>
  <div class="page-container">
    <div class="page-title">文章列表</div>
    
    <!-- 搜索和操作区域 -->
    <div class="search-bar">
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input
            v-model="searchForm.title"
            placeholder="请输入文章标题搜索"
            prefix-icon="el-icon-search"
            clearable
            @keyup.enter.native="handleSearch"
          />
        </el-col>
        <el-col :span="6">
          <el-select
            v-model="searchForm.status"
            placeholder="请选择状态"
            clearable
            style="width: 100%"
          >
            <el-option label="已发布" value="published" />
            <el-option label="草稿" value="draft" />
          </el-select>
        </el-col>
        <el-col :span="10">
          <el-button type="primary" @click="handleSearch">
            <i class="el-icon-search"></i>
            搜索
          </el-button>
          <el-button @click="resetSearch">
            <i class="el-icon-refresh"></i>
            重置
          </el-button>
          <el-button type="success" @click="$router.push('/write')">
            <i class="el-icon-plus"></i>
            新建文章
          </el-button>
        </el-col>
      </el-row>
    </div>
    
    <!-- 文章列表表格 -->
    <el-table
      v-loading="loading"
      :data="articleList"
      stripe
      style="width: 100%"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      
      <el-table-column prop="id" label="ID" width="80" />
      
      <el-table-column prop="title" label="标题" min-width="200">
        <template slot-scope="scope">
          <el-link type="primary" @click="viewArticle(scope.row)">
            {{ scope.row.title }}
          </el-link>
        </template>
      </el-table-column>
      
      <el-table-column prop="author" label="作者" width="120" />
      
      <el-table-column prop="status" label="状态" width="100">
        <template slot-scope="scope">
          <el-tag
            :type="scope.row.status === 'published' ? 'success' : 'info'"
            size="small"
          >
            {{ scope.row.status === 'published' ? '已发布' : '草稿' }}
          </el-tag>
        </template>
      </el-table-column>
      
      <el-table-column prop="views" label="阅读量" width="100" />
      
      <el-table-column prop="created_at" label="创建时间" width="180">
        <template slot-scope="scope">
          {{ formatDate(scope.row.created_at) }}
        </template>
      </el-table-column>
      
      <el-table-column label="操作" width="200" fixed="right">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="primary"
            @click="editArticle(scope.row)"
          >
            编辑
          </el-button>
          <el-button
            size="mini"
            :type="scope.row.status === 'published' ? 'warning' : 'success'"
            @click="toggleStatus(scope.row)"
          >
            {{ scope.row.status === 'published' ? '下线' : '发布' }}
          </el-button>
          <el-button
            size="mini"
            type="danger"
            @click="deleteArticle(scope.row)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    
    <!-- 分页组件 -->
    <div class="pagination-container">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="pagination.page"
        :page-sizes="[10, 20, 50, 100]"
        :page-size="pagination.size"
        layout="total, sizes, prev, pager, next, jumper"
        :total="pagination.total"
      />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Articles',
  data() {
    return {
      loading: false,
      articleList: [],
      selectedArticles: [],
      // 搜索表单
      searchForm: {
        title: '',
        status: ''
      },
      // 分页信息
      pagination: {
        page: 1,
        size: 10,
        total: 0
      }
    }
  },
  methods: {
    /**
     * 获取文章列表
     */
    async getArticleList() {
      this.loading = true
      try {
        const params = {
          page: this.pagination.page,
          size: this.pagination.size,
          ...this.searchForm
        }
        
        const response = await this.$http.get('/articles', { params })
        const { data } = response.data
        
        this.articleList = data.list || []
        this.pagination.total = data.total || 0
      } catch (error) {
        this.$message.error(error.response?.data?.message || '获取文章列表失败')
      } finally {
        this.loading = false
      }
    },
    
    /**
     * 搜索文章
     */
    handleSearch() {
      this.pagination.page = 1
      this.getArticleList()
    },
    
    /**
     * 重置搜索条件
     */
    resetSearch() {
      this.searchForm = {
        title: '',
        status: ''
      }
      this.pagination.page = 1
      this.getArticleList()
    },
    
    /**
     * 查看文章详情
     */
    viewArticle(article) {
      // 这里可以跳转到文章详情页或者打开弹窗
      this.$message.info(`查看文章：${article.title}`)
    },
    
    /**
     * 编辑文章
     */
    editArticle(article) {
      this.$router.push(`/write?id=${article.id}`)
    },
    
    /**
     * 切换文章状态
     */
    async toggleStatus(article) {
      const action = article.status === 'published' ? '下线' : '发布'
      
      try {
        await this.$confirm(`确定要${action}这篇文章吗？`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        
        const newStatus = article.status === 'published' ? 'draft' : 'published'
        await this.$http.put(`/articles/${article.id}/status`, { status: newStatus })
        
        this.$message.success(`${action}成功`)
        this.getArticleList()
      } catch (error) {
        if (error !== 'cancel') {
          this.$message.error(error.response?.data?.message || `${action}失败`)
        }
      }
    },
    
    /**
     * 删除文章
     */
    async deleteArticle(article) {
      try {
        await this.$confirm('确定要删除这篇文章吗？删除后无法恢复！', '警告', {
          confirmButtonText: '确定删除',
          cancelButtonText: '取消',
          type: 'error'
        })
        
        await this.$http.delete(`/articles/${article.id}`)
        
        this.$message.success('删除成功')
        this.getArticleList()
      } catch (error) {
        if (error !== 'cancel') {
          this.$message.error(error.response?.data?.message || '删除失败')
        }
      }
    },
    
    /**
     * 处理表格选择变化
     */
    handleSelectionChange(selection) {
      this.selectedArticles = selection
    },
    
    /**
     * 处理分页大小变化
     */
    handleSizeChange(size) {
      this.pagination.size = size
      this.pagination.page = 1
      this.getArticleList()
    },
    
    /**
     * 处理当前页变化
     */
    handleCurrentChange(page) {
      this.pagination.page = page
      this.getArticleList()
    },
    
    /**
     * 格式化日期
     */
    formatDate(dateString) {
      if (!dateString) return '-'
      const date = new Date(dateString)
      return date.toLocaleString('zh-CN')
    }
  },
  
  created() {
    this.getArticleList()
  }
}
</script>

<style scoped>
.search-bar {
  margin-bottom: 20px;
  padding: 20px;
  background: white;
  border-radius: 4px;
}

.pagination-container {
  margin-top: 20px;
  text-align: right;
}

.el-table {
  background: white;
  border-radius: 4px;
}
</style>