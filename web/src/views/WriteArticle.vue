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
      
      <!-- 文章摘要 -->
      <el-form-item label="文章摘要" prop="summary">
        <el-input
          v-model="articleForm.summary"
          type="textarea"
          :rows="3"
          placeholder="请输入文章摘要"
          maxlength="200"
          show-word-limit
        />
      </el-form-item>
      
      <!-- 文章分类 -->
      <el-form-item label="文章分类" prop="category">
        <el-select
          v-model="articleForm.category"
          placeholder="请选择文章分类"
          style="width: 200px"
        >
          <el-option label="技术分享" value="tech" />
          <el-option label="生活随笔" value="life" />
          <el-option label="学习笔记" value="study" />
          <el-option label="项目总结" value="project" />
        </el-select>
      </el-form-item>
      
      <!-- 文章标签 -->
      <el-form-item label="文章标签">
        <el-tag
          v-for="tag in articleForm.tags"
          :key="tag"
          closable
          @close="removeTag(tag)"
          style="margin-right: 10px"
        >
          {{ tag }}
        </el-tag>
        <el-input
          v-if="inputVisible"
          ref="saveTagInput"
          v-model="inputValue"
          size="small"
          style="width: 100px"
          @keyup.enter.native="handleInputConfirm"
          @blur="handleInputConfirm"
        />
        <el-button
          v-else
          size="small"
          @click="showInput"
        >
          + 添加标签
        </el-button>
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
        <el-radio-group v-model="articleForm.status">
          <el-radio label="draft">保存为草稿</el-radio>
          <el-radio label="published">立即发布</el-radio>
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
          <span>分类：{{ getCategoryName(articleForm.category) }}</span>
          <span v-if="articleForm.tags.length">标签：{{ articleForm.tags.join(', ') }}</span>
        </div>
        <div class="article-summary">
          <strong>摘要：</strong>{{ articleForm.summary }}
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
      inputVisible: false,
      inputValue: '',
      // 文章表单数据
      articleForm: {
        title: '',
        summary: '',
        content: '',
        category: '',
        tags: [],
        status: 'draft'
      },
      // 表单验证规则
      articleRules: {
        title: [
          { required: true, message: '请输入文章标题', trigger: 'blur' },
          { min: 5, max: 100, message: '标题长度在 5 到 100 个字符', trigger: 'blur' }
        ],
        summary: [
          { required: true, message: '请输入文章摘要', trigger: 'blur' },
          { min: 10, max: 200, message: '摘要长度在 10 到 200 个字符', trigger: 'blur' }
        ],
        content: [
          { required: true, message: '请输入文章内容', trigger: 'blur' },
          { min: 50, message: '文章内容至少50个字符', trigger: 'blur' }
        ],
        category: [
          { required: true, message: '请选择文章分类', trigger: 'change' }
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
        const response = await this.$http.get(`/articles/${this.articleId}`)
        const { data } = response.data
        
        this.articleForm = {
          title: data.title || '',
          summary: data.summary || '',
          content: data.content || '',
          category: data.category || '',
          tags: data.tags || [],
          status: data.status || 'draft'
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
            if (this.isEdit) {
              // 更新文章
              await this.$http.put(`/articles/${this.articleId}`, this.articleForm)
              this.$message.success('文章更新成功')
            } else {
              // 创建文章
              await this.$http.post('/articles', this.articleForm)
              this.$message.success('文章保存成功')
            }
            
            this.$router.push('/articles')
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
    },
    
    /**
     * 显示标签输入框
     */
    showInput() {
      this.inputVisible = true
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus()
      })
    },
    
    /**
     * 确认添加标签
     */
    handleInputConfirm() {
      const inputValue = this.inputValue.trim()
      if (inputValue && !this.articleForm.tags.includes(inputValue)) {
        this.articleForm.tags.push(inputValue)
      }
      this.inputVisible = false
      this.inputValue = ''
    },
    
    /**
     * 移除标签
     */
    removeTag(tag) {
      const index = this.articleForm.tags.indexOf(tag)
      if (index > -1) {
        this.articleForm.tags.splice(index, 1)
      }
    },
    
    /**
     * 获取分类名称
     */
    getCategoryName(category) {
      const categoryMap = {
        tech: '技术分享',
        life: '生活随笔',
        study: '学习笔记',
        project: '项目总结'
      }
      return categoryMap[category] || category
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