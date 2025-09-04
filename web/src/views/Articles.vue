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
            v-model="searchForm.state"
            placeholder="请选择状态"
            clearable
            style="width: 100%"
          >
            <el-option label="全部" :value="-1" />
            <el-option label="有效" :value="1" />
            <el-option label="无效" :value="0" />
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

      <el-table-column prop="title" label="标题" min-width="300">
        <template slot-scope="scope">
          <div>
            <el-link type="primary" @click="viewArticle(scope.row)">
              {{ scope.row.title }}
            </el-link>
            <div class="article-content-preview">
              {{ scope.row.content }}
            </div>
          </div>
        </template>
      </el-table-column>

      <el-table-column prop="state" label="状态" width="100">
        <template slot-scope="scope">
          <el-tag
            :type="scope.row.state === 1 ? 'success' : 'danger'"
            size="small"
          >
            {{ scope.row.state === 1 ? '有效' : '无效' }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column prop="createTime" label="创建时间" width="180">
        <template slot-scope="scope">
          {{ formatDate(scope.row.createTime) }}
        </template>
      </el-table-column>

      <el-table-column label="操作" width="150" fixed="right">
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
        state: -1
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
        const requestData = {
          title: this.searchForm.title || '',
          state: this.searchForm.state,
          pageIndex: this.pagination.page,
          pageSize: this.pagination.size
        }

        const response = await this.$http.post('/posts/list', requestData)

        if (response.data.code === 100) {
          const { data } = response.data
          this.articleList = data.list || []
          this.pagination.total = data.totalCount || 0
        } else {
          this.$message.error(response.data.message || '获取文章列表失败')
        }
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
        state: -1
      }
      this.pagination.page = 1
      this.getArticleList()
    },

    /**
     * 查看文章详情
     */
    async viewArticle(article) {
      try {
        const response = await this.$http.get('/posts/detail', {
          params: { id: String(article.id) } // 确保ID作为字符串传递
        })

        if (response.data.code === 100) {
          const { data } = response.data
          const title = data.Title || data.title || '无标题'
          const state = (data.State || data.state) === 1 ? '有效' : '无效'
          const createTime = this.formatDate(data.CreateTime || data.createTime)
          const content = data.Content || data.content || '无内容'
          
          // 构建自定义HTML内容
          const htmlContent = `
            <div style="padding: 20px; font-family: Arial, sans-serif;">
              <!-- 标题部分 -->
              <h2 style="margin: 0 0 20px 0; font-weight: bold; font-size: 20px; color: #333; border-bottom: 2px solid #409EFF; padding-bottom: 10px;">
                ${title}
              </h2>
              
              <!-- 时间和状态行 -->
              <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 25px; padding: 10px 0; background-color: #f8f9fa; border-radius: 5px; padding: 15px;">
                <span style="color: #666; font-size: 14px;">
                  <i class="el-icon-time" style="margin-right: 5px;"></i>
                  创建时间：${createTime}
                </span>
                <span style="padding: 4px 12px; border-radius: 15px; font-size: 12px; font-weight: bold; ${
                  state === '有效' 
                    ? 'background-color: #67C23A; color: white;' 
                    : 'background-color: #F56C6C; color: white;'
                }">
                  ${state}
                </span>
              </div>
              
              <!-- 内容部分 -->
              <div style="margin-top: 25px; padding-top: 20px; border-top: 1px solid #eee;">
                <h4 style="margin: 0 0 15px 0; color: #333; font-size: 16px;">文章内容：</h4>
                <div style="background-color: #fafafa; padding: 20px; border-radius: 8px; border-left: 4px solid #409EFF; line-height: 1.6; color: #555; max-height: 300px; overflow-y: auto; white-space: pre-wrap; word-wrap: break-word;">
                  ${content}
                </div>
              </div>
            </div>
          `

          this.$alert(htmlContent, '文章详情', {
            dangerouslyUseHTMLString: true,
            customClass: 'article-detail-dialog',
            showClose: true,
            closeOnClickModal: true,
            closeOnPressEscape: true
          })
        } else {
          this.$message.error(response.data.message || '获取文章详情失败')
        }
      } catch (error) {
        this.$message.error(error.response?.data?.message || '获取文章详情失败')
      }
    },

    /**
     * 编辑文章
     */
    editArticle(article) {
      // 确保ID作为字符串传递，避免long类型精度丢失
      this.$router.push(`/write?id=${article.id}`)
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

        const response = await this.$http.post('/posts/remove', {
          id: String(article.id) // 确保ID作为字符串传递，避免long类型精度丢失
        })

        if (response.data.code === 100) {
          this.$message.success('删除成功')
          this.getArticleList()
        } else {
          this.$message.error(response.data.message || '删除失败')
        }
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

.article-content-preview {
  color: #999;
  font-size: 12px;
  margin-top: 4px;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 280px;
}
</style>

<style>
/* 文章详情弹框样式 */
.article-detail-dialog {
  width: 60% !important;
  min-height: 500px !important;
  max-width: 800px !important;
  min-width: 600px !important;
}

.article-detail-dialog .el-message-box {
  width: 60% !important;
  min-height: 500px !important;
  max-width: 800px !important;
  min-width: 600px !important;
}

.article-detail-dialog .el-message-box__content {
  padding: 0 !important;
  max-height: 70vh !important;
  overflow-y: auto !important;
}

.article-detail-dialog .el-message-box__message {
  margin: 0 !important;
  padding: 0 !important;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  .article-detail-dialog {
    width: 80% !important;
    min-width: 500px !important;
  }
  
  .article-detail-dialog .el-message-box {
    width: 80% !important;
    min-width: 500px !important;
  }
}

@media (max-width: 768px) {
  .article-detail-dialog {
    width: 95% !important;
    min-width: 320px !important;
  }
  
  .article-detail-dialog .el-message-box {
    width: 95% !important;
    min-width: 320px !important;
  }
}
</style>
