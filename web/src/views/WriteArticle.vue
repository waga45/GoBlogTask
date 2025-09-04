<template>
  <div class="page-container">
    <div class="page-title">{{ isEdit ? '编辑文章' : '写文章' }}</div>
    
    <el-form
      ref="articleForm"
      :model="articleForm"
      :rules="articleRules"
      label-width="100px"
    >
      <!-- 文章标题 -->
      <el-form-item label="文章标题" prop="title">
        <el-input
          v-model="articleForm.title"
          placeholder="请输入文章标题"
          maxlength="100"
          show-word-limit
        />
      </el-form-item>
      
      <!-- 文章内容 -->
      <el-form-item label="文章内容" prop="content">
        <el-input
          v-model="articleForm.content"
          type="textarea"
          :rows="15"
          placeholder="请输入文章内容（支持Markdown格式）"
        />
      </el-form-item>
      
      <!-- 文章状态 -->
      <el-form-item label="文章状态">
        <el-radio-group v-model="articleForm.state">
          <el-radio :label="0">草稿(无效)</el-radio>
          <el-radio :label="1">发布(有效)</el-radio>
        </el-radio-group>
      </el-form-item>
      
      <!-- 操作按钮 -->
      <el-form-item>
        <el-button
          type="primary"
          :loading="submitLoading"
          @click="handleSubmit"
        >
          {{ isEdit ? '更新文章' : '保存文章' }}
        </el-button>
        <el-button @click="handleCancel">
          取消
        </el-button>
        <el-button
          v-if="!isEdit"
          type="info"
          @click="handlePreview"
        >
          预览
        </el-button>
      </el-form-item>
    </el-form>
    
    <!-- 预览对话框 -->
    <el-dialog
      title="文章预览"
      :visible.sync="previewVisible"
      width="80%"
      top="5vh"
    >
      <div class="preview-content">
        <h2>{{ articleForm.title }}</h2>
        <div class="article-meta">
          <span>状态：{{ articleForm.state === 1 ? '有效' : '无效' }}</span>
        </div>
        <div class="article-content">
          <pre>{{ articleForm.content }}</pre>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'WriteArticle',
  data() {
    return {
      isEdit: false,
      articleId: null,
      submitLoading: false,
      previewVisible: false,
      // 文章表单数据
      articleForm: {
        title: '',
        content: '',
        state: 1
      },
      // 表单验证规则
      articleRules: {
        title: [
          { required: true, message: '请输入文章标题', trigger: 'blur' },
          { min: 1, max: 100, message: '标题长度在 1 到 100 个字符', trigger: 'blur' }
        ],
        content: [
          { required: true, message: '请输入文章内容', trigger: 'blur' },
          { min: 1, message: '文章内容不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    /**
     * 获取文章详情（编辑模式）
     */
    async getArticleDetail() {
      try {
        const response = await this.$http.get('/posts/detail', {
          params: { id: String(this.articleId) } // 确保ID作为字符串传递
        })
        
        if (response.data.code === 100) {
          const { data } = response.data
          this.articleForm = {
            title: data.title || '',
            content: data.content || '',
            state: data.state || 1
          }
        } else {
          this.$message.error(response.data.message || '获取文章详情失败')
          this.$router.push('/articles')
        }
      } catch (error) {
        this.$message.error(error.response?.data?.message || '获取文章详情失败')
        this.$router.push('/articles')
      }
    },
    
    /**
     * 提交文章
     */
    handleSubmit() {
      this.$refs.articleForm.validate(async (valid) => {
        if (valid) {
          this.submitLoading = true
          try {
            let response
            if (this.isEdit) {
              // 更新文章
              const updateData = {
                id: this.articleId, // 保持字符串格式，避免long类型精度丢失
                title: this.articleForm.title,
                content: this.articleForm.content,
                state: this.articleForm.state
              }
              response = await this.$http.post('/posts/update', updateData)
            } else {
              // 创建文章
              response = await this.$http.post('/posts/new', this.articleForm)
            }
            
            if (response.data.code === 100) {
              this.$message.success(this.isEdit ? '文章更新成功' : '文章创建成功')
              this.$router.push('/articles')
            } else {
              this.$message.error(response.data.message || '操作失败')
            }
          } catch (error) {
            this.$message.error(error.response?.data?.message || '操作失败')
          } finally {
            this.submitLoading = false
          }
        }
      })
    },
    
    /**
     * 取消操作
     */
    handleCancel() {
      this.$confirm('确定要取消吗？未保存的内容将丢失！', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '继续编辑',
        type: 'warning'
      }).then(() => {
        this.$router.push('/articles')
      }).catch(() => {
        // 用户选择继续编辑
      })
    },
    
    /**
     * 预览文章
     */
    handlePreview() {
      if (!this.articleForm.title || !this.articleForm.content) {
        this.$message.warning('请先填写标题和内容')
        return
      }
      this.previewVisible = true
    }
  },
  
  created() {
    // 检查是否为编辑模式
    this.articleId = this.$route.query.id
    if (this.articleId) {
      this.isEdit = true
      this.getArticleDetail()
    }
  }
}
</script>

<style scoped>
.preview-content {
  padding: 20px;
}

.preview-content h2 {
  color: #333;
  margin-bottom: 15px;
  padding-bottom: 10px;
  border-bottom: 1px solid #e8e8e8;
}

.article-meta {
  margin-bottom: 15px;
  color: #666;
  font-size: 14px;
}

.article-meta span {
  margin-right: 20px;
}

.article-summary {
  margin-bottom: 20px;
  padding: 15px;
  background: #f5f5f5;
  border-radius: 4px;
  color: #666;
}

.article-content {
  line-height: 1.8;
  color: #333;
}

.article-content pre {
  white-space: pre-wrap;
  word-wrap: break-word;
  font-family: inherit;
  margin: 0;
}

.el-tag {
  margin-right: 10px;
  margin-bottom: 10px;
}
</style>