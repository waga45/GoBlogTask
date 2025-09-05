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

      <el-table-column label="操作" width="240" fixed="right">
        <template slot-scope="scope">
          <div class="action-buttons">
            <el-button
              size="mini"
              type="info"
              icon="el-icon-view"
              @click="viewArticle(scope.row)"
              plain
            >
              查看
            </el-button>
            <el-button
              size="mini"
              type="primary"
              icon="el-icon-edit"
              @click="editArticle(scope.row)"
              plain
            >
              编辑
            </el-button>
            <el-button
              size="mini"
              type="success"
              icon="el-icon-chat-dot-round"
              @click="showComments(scope.row)"
              plain
            >
              评论
            </el-button>
            <el-button
              size="mini"
              type="danger"
              icon="el-icon-delete"
              @click="deleteArticle(scope.row)"
              plain
            >
              删除
            </el-button>
          </div>
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

    <!-- 评论弹框 -->
    <el-dialog
      :title="commentDialog.title"
      :visible.sync="commentDialog.visible"
      width="70%"
      :before-close="closeCommentDialog"
      class="comment-dialog"
    >
      <div v-loading="commentDialog.loading">
        <!-- 添加评论表单 -->
        <div class="add-comment-section">
          <h4>添加评论</h4>
          <el-input
            v-model="newComment.content"
            type="textarea"
            :rows="3"
            placeholder="请输入评论内容..."
            maxlength="500"
            show-word-limit
          ></el-input>
          <div class="comment-actions">
            <el-button type="primary" @click="addComment()">发表评论</el-button>
          </div>
        </div>

        <!-- 评论列表 -->
        <div class="comment-list-section">
          <h4>评论列表 ({{ commentPagination.total }}条)</h4>
          <div v-if="commentList.length === 0" class="no-comments">
            <el-empty description="暂无评论"></el-empty>
          </div>
          <div v-else>
            <div
              v-for="comment in commentList"
              :key="comment.id"
              class="comment-item"
            >
              <div class="comment-header">
                <span class="comment-user">{{ comment.userName }}</span>
                <span class="comment-time">{{ formatDate(comment.createTime) }}</span>
                <el-button
                  type="text"
                  size="mini"
                  class="delete-btn"
                  @click="deleteComment(comment)"
                >
                  删除
                </el-button>
              </div>
              <div class="comment-content">{{ comment.content }}</div>
            </div>
          </div>

          <!-- 评论分页 -->
          <div v-if="commentPagination.total > 0" class="comment-pagination">
            <el-pagination
              @size-change="handleCommentSizeChange"
              @current-change="handleCommentCurrentChange"
              :current-page="commentPagination.page"
              :page-sizes="[5, 10, 20]"
              :page-size="commentPagination.size"
              layout="total, sizes, prev, pager, next"
              :total="commentPagination.total"
              small
            />
          </div>
        </div>
      </div>
     </el-dialog>

     <!-- 文章详情弹框 -->
     <el-dialog
       title="文章详情"
       :visible.sync="articleDetail.visible"
       width="80%"
       :before-close="closeArticleDetailDialog"
       class="article-detail-dialog"
     >
       <div v-loading="articleDetail.loading">
         <!-- 文章内容部分 -->
         <div class="article-detail-content">
           <div class="article-header">
             <h2 class="article-title">{{ articleDetail.data.Title || articleDetail.data.title || '无标题' }}</h2>
             <div class="article-meta">
               <span class="article-time">
                 <i class="el-icon-time"></i>
                 创建时间：{{ formatDate(articleDetail.data.CreateTime || articleDetail.data.createTime) }}
               </span>
               <span 
                 class="article-status"
                 :class="{
                   'status-active': (articleDetail.data.State || articleDetail.data.state) === 1,
                   'status-inactive': (articleDetail.data.State || articleDetail.data.state) !== 1
                 }"
               >
                 {{ (articleDetail.data.State || articleDetail.data.state) === 1 ? '有效' : '无效' }}
               </span>
             </div>
           </div>
           
           <div class="article-content-body">
             <h4>文章内容：</h4>
             <div class="content-text">
               {{ articleDetail.data.Content || articleDetail.data.content || '无内容' }}
             </div>
           </div>
         </div>

         <!-- 评论部分 -->
         <div class="article-comments-section">
           <!-- 添加评论表单 -->
           <div class="add-comment-section">
             <h4>添加评论</h4>
             <el-input
               v-model="detailNewComment.content"
               type="textarea"
               :rows="3"
               placeholder="请输入评论内容..."
               maxlength="500"
               show-word-limit
             ></el-input>
             <div class="comment-actions">
               <el-button type="primary" @click="addDetailComment()">发表评论</el-button>
             </div>
           </div>

           <!-- 评论列表 -->
           <div class="comment-list-section">
             <h4>评论列表 ({{ detailCommentPagination.total }}条)</h4>
             <div v-if="detailCommentList.length === 0" class="no-comments">
               <el-empty description="暂无评论"></el-empty>
             </div>
             <div v-else>
               <div
                 v-for="comment in detailCommentList"
                 :key="comment.id"
                 class="comment-item"
               >
                 <div class="comment-header">
                   <span class="comment-user">{{ comment.userName }}</span>
                   <span class="comment-time">{{ formatDate(comment.createTime) }}</span>
                   <el-button
                     type="text"
                     size="mini"
                     class="delete-btn"
                     @click="deleteDetailComment(comment)"
                   >
                     删除
                   </el-button>
                 </div>
                 <div class="comment-content">{{ comment.content }}</div>
               </div>
             </div>

             <!-- 评论分页 -->
             <div v-if="detailCommentPagination.total > 0" class="comment-pagination">
               <el-pagination
                 @size-change="handleDetailCommentSizeChange"
                 @current-change="handleDetailCommentCurrentChange"
                 :current-page="detailCommentPagination.page"
                 :page-sizes="[5, 10, 20]"
                 :page-size="detailCommentPagination.size"
                 layout="total, sizes, prev, pager, next"
                 :total="detailCommentPagination.total"
                 small
               />
             </div>
           </div>
         </div>
       </div>
     </el-dialog>
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
      },
      // 评论相关数据
      commentDialog: {
        visible: false,
        title: '',
        articleId: '',
        loading: false
      },
      commentList: [],
      commentPagination: {
        page: 1,
        size: 10,
        total: 0
      },
      newComment: {
        content: ''
      },
      // 文章详情弹框相关数据
      articleDetail: {
        visible: false,
        loading: false,
        data: {},
        articleId: ''
      },
      detailCommentList: [],
      detailCommentPagination: {
        page: 1,
        size: 10,
        total: 0
      },
      detailNewComment: {
        content: ''
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
      // 设置文章详情数据
      this.articleDetail = {
        visible: true,
        loading: true,
        data: {},
        articleId: String(article.id || article.Id)
      }
      this.detailCommentList = []
      this.detailCommentPagination.page = 1
      this.detailNewComment.content = ''

      try {
        const response = await this.$http.get('/posts/detail', {
          params: { id: this.articleDetail.articleId }
        })

        if (response.data.code === 100) {
          this.articleDetail.data = response.data.data
          // 同时加载评论列表
          await this.getDetailCommentList()
        } else {
          this.$message.error(response.data.message || '获取文章详情失败')
        }
      } catch (error) {
        this.$message.error(error.response?.data?.message || '获取文章详情失败')
      } finally {
        this.articleDetail.loading = false
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
    },

    /**
     * 显示评论弹框
     */
    async showComments(article) {
      this.commentDialog.visible = true
      this.commentDialog.title = article.title || article.Title || '文章评论'
      this.commentDialog.articleId = String(article.id || article.Id)
      this.commentPagination.page = 1
      this.newComment.content = ''
      await this.getCommentList()
    },

    /**
     * 获取评论列表
     */
    async getCommentList() {
      this.commentDialog.loading = true
      try {
        const requestData = {
          postId: this.commentDialog.articleId,
          pageIndex: this.commentPagination.page,
          pageSize: this.commentPagination.size
        }

        const response = await this.$http.post('comments/list', requestData)

        if (response.data.code === 100) {
          const { data } = response.data
          this.commentList = data.list || []
          this.commentPagination.total = data.totalCount || 0
        } else {
          this.$message.error(response.data.message || '获取评论列表失败')
        }
      } catch (error) {
        this.$message.error(error.response?.data?.message || '获取评论列表失败')
      } finally {
        this.commentDialog.loading = false
      }
    },

    /**
     * 添加评论
     */
    async addComment() {
      if (!this.newComment.content.trim()) {
        this.$message.warning('请输入评论内容')
        return
      }

      try {
        const requestData = {
          postId: this.commentDialog.articleId,
          content: this.newComment.content.trim()
        }

        const response = await this.$http.post('comments/new', requestData)

        if (response.data.code === 100) {
          this.$message.success('评论添加成功')
          this.newComment.content = ''
          this.commentPagination.page = 1
          await this.getCommentList()
        } else {
          this.$message.error(response.data.message || '添加评论失败')
        }
      } catch (error) {
        this.$message.error(error.response?.data?.message || '添加评论失败')
      }
    },

    /**
     * 删除评论
     */
    async deleteComment(comment) {
      try {
        await this.$confirm('确定要删除这条评论吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })

        const requestData = {
          postId: this.commentDialog.articleId,
          id: Number(comment.id)
        }

        const response = await this.$http.post('comments/disable', requestData)

        if (response.data.code === 100) {
          this.$message.success('删除成功')
          await this.getCommentList()
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
     * 处理评论分页大小变化
     */
    handleCommentSizeChange(size) {
      this.commentPagination.size = size
      this.commentPagination.page = 1
      this.getCommentList()
    },

    /**
     * 处理评论当前页变化
     */
    handleCommentCurrentChange(page) {
      this.commentPagination.page = page
      this.getCommentList()
    },

    /**
     * 关闭评论弹框
     */
    closeCommentDialog() {
      this.commentDialog.visible = false
      this.commentList = []
      this.newComment.content = ''
    },

    /**
     * 获取文章详情页面的评论列表
     */
    async getDetailCommentList() {
      try {
        const requestData = {
          postId: this.articleDetail.articleId,
          pageIndex: this.detailCommentPagination.page,
          pageSize: this.detailCommentPagination.size
        }

        const response = await this.$http.post('comments/list', requestData)

        if (response.data.code === 100) {
          const { data } = response.data
          this.detailCommentList = data.list || []
          this.detailCommentPagination.total = data.totalCount || 0
        } else {
          this.$message.error(response.data.message || '获取评论列表失败')
        }
      } catch (error) {
        this.$message.error(error.response?.data?.message || '获取评论列表失败')
      }
    },

    /**
     * 在文章详情页面添加评论
     */
    async addDetailComment() {
      if (!this.detailNewComment.content.trim()) {
        this.$message.warning('请输入评论内容')
        return
      }

      try {
        const requestData = {
          postId: this.articleDetail.articleId,
          content: this.detailNewComment.content.trim()
        }

        const response = await this.$http.post('comments/new', requestData)

        if (response.data.code === 100) {
          this.$message.success('评论添加成功')
          this.detailNewComment.content = ''
          this.detailCommentPagination.page = 1
          await this.getDetailCommentList()
        } else {
          this.$message.error(response.data.message || '添加评论失败')
        }
      } catch (error) {
        this.$message.error(error.response?.data?.message || '添加评论失败')
      }
    },

    /**
     * 在文章详情页面删除评论
     */
    async deleteDetailComment(comment) {
      try {
        await this.$confirm('确定要删除这条评论吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })

        const requestData = {
          postId: this.articleDetail.articleId,
          id: Number(comment.id)
        }

        const response = await this.$http.post('comments/disable', requestData)

        if (response.data.code === 100) {
          this.$message.success('删除成功')
          await this.getDetailCommentList()
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
     * 处理文章详情评论分页大小变化
     */
    handleDetailCommentSizeChange(size) {
      this.detailCommentPagination.size = size
      this.detailCommentPagination.page = 1
      this.getDetailCommentList()
    },

    /**
     * 处理文章详情评论当前页变化
     */
    handleDetailCommentCurrentChange(page) {
      this.detailCommentPagination.page = page
      this.getDetailCommentList()
    },

    /**
     * 关闭文章详情弹框
     */
    closeArticleDetailDialog() {
      this.articleDetail.visible = false
      this.detailCommentList = []
      this.detailNewComment.content = ''
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

.action-buttons {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
  justify-content: flex-start;
}

.action-buttons .el-button {
  margin: 0;
  padding: 5px 8px;
  font-size: 12px;
  border-radius: 3px;
  transition: all 0.3s ease;
}

.action-buttons .el-button:hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>

<style>
/* 评论弹框样式 */
.comment-dialog .el-dialog__body {
  padding: 20px;
  max-height: 70vh;
  overflow-y: auto;
}

.add-comment-section {
  margin-bottom: 30px;
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
}

.add-comment-section h4 {
  margin: 0 0 15px 0;
  color: #333;
  font-size: 16px;
}

.comment-actions {
  margin-top: 15px;
  text-align: right;
}

.comment-list-section h4 {
  margin: 0 0 20px 0;
  color: #333;
  font-size: 16px;
  border-bottom: 2px solid #409EFF;
  padding-bottom: 10px;
}

.comment-item {
  padding: 15px;
  margin-bottom: 15px;
  background-color: #fff;
  border: 1px solid #ebeef5;
  border-radius: 8px;
  transition: box-shadow 0.2s;
}

.comment-item:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.comment-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  padding-bottom: 8px;
  border-bottom: 1px solid #f0f0f0;
}

.comment-user {
  font-weight: bold;
  color: #409EFF;
  font-size: 14px;
}

.comment-time {
  color: #999;
  font-size: 12px;
}

.delete-btn {
  color: #f56c6c !important;
}

.delete-btn:hover {
  color: #f56c6c !important;
  background-color: #fef0f0 !important;
}

.comment-content {
  color: #555;
  line-height: 1.6;
  word-wrap: break-word;
  white-space: pre-wrap;
}

.no-comments {
  text-align: center;
  padding: 40px 0;
}

.comment-pagination {
  margin-top: 20px;
  text-align: center;
}

/* 文章详情弹框样式 */
.article-detail-dialog .el-dialog {
  width: 80% !important;
  max-width: 1200px !important;
  min-width: 800px !important;
  margin: 5vh auto !important;
  display: flex;
  flex-direction: column;
  max-height: 90vh;
  transform: none !important;
  left: 50% !important;
  top: 5vh !important;
  margin-left: -40% !important;
}

.article-detail-dialog .el-dialog__body {
  max-height: 70vh;
  overflow-y: auto;
  padding: 20px;
  flex: 1;
}

.article-detail-content {
  margin-bottom: 30px;
  padding: 20px;
  background-color: #fff;
  border-radius: 8px;
  border: 1px solid #ebeef5;
}

.article-header {
  margin-bottom: 20px;
}

.article-title {
  margin: 0 0 15px 0;
  font-weight: bold;
  font-size: 24px;
  color: #333;
  border-bottom: 2px solid #409EFF;
  padding-bottom: 10px;
}

.article-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  background-color: #f8f9fa;
  border-radius: 5px;
}

.article-time {
  color: #666;
  font-size: 14px;
}

.article-time i {
  margin-right: 5px;
}

.article-status {
  padding: 4px 12px;
  border-radius: 15px;
  font-size: 12px;
  font-weight: bold;
}

.status-active {
  background-color: #67C23A;
  color: white;
}

.status-inactive {
  background-color: #F56C6C;
  color: white;
}

.article-content-body h4 {
  margin: 20px 0 15px 0;
  color: #333;
  font-size: 16px;
}

.content-text {
  background-color: #fafafa;
  padding: 20px;
  border-radius: 8px;
  border-left: 4px solid #409EFF;
  line-height: 1.6;
  color: #555;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.article-comments-section {
  border-top: 2px solid #eee;
  padding-top: 20px;
}

.add-comment-section {
  margin-bottom: 30px;
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
  border: 1px solid #e9ecef;
}

.add-comment-section h4 {
  margin: 0 0 15px 0;
  color: #333;
  font-size: 16px;
}

.comment-actions {
  margin-top: 15px;
  text-align: right;
}

.comment-list-section h4 {
  margin: 0 0 20px 0;
  color: #333;
  font-size: 16px;
  border-bottom: 1px solid #eee;
  padding-bottom: 10px;
}

.comment-item {
  padding: 15px;
  margin-bottom: 15px;
  background-color: #fff;
  border: 1px solid #ebeef5;
  border-radius: 8px;
  transition: box-shadow 0.3s ease;
}

.comment-item:hover {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.comment-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  padding-bottom: 8px;
  border-bottom: 1px solid #f0f0f0;
}

.comment-user {
  font-weight: bold;
  color: #409EFF;
  font-size: 14px;
}

.comment-time {
  color: #999;
  font-size: 12px;
}

.delete-btn {
  color: #F56C6C !important;
  padding: 0 !important;
}

.delete-btn:hover {
  color: #F56C6C !important;
  background-color: transparent !important;
}

.comment-content {
  color: #555;
  line-height: 1.6;
  font-size: 14px;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.no-comments {
  text-align: center;
  padding: 40px 20px;
  color: #999;
}

.comment-pagination {
  margin-top: 20px;
  text-align: center;
  padding-top: 20px;
  border-top: 1px solid #eee;
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
