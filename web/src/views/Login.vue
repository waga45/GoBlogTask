<template>
  <div class="login-container">
    <div class="login-box">
      <h1 class="login-title">GoBlog</h1>
      
      <el-tabs v-model="activeTab" class="login-tabs">
        <!-- 登录表单 -->
        <el-tab-pane label="登录" name="login">
          <el-form
            ref="loginForm"
            :model="loginForm"
            :rules="loginRules"
            label-width="0px"
          >
            <el-form-item prop="username">
              <el-input
                v-model="loginForm.username"
                placeholder="请输入用户名"
                prefix-icon="el-icon-user"
                size="large"
              />
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                v-model="loginForm.password"
                type="password"
                placeholder="请输入密码"
                prefix-icon="el-icon-lock"
                size="large"
              />
            </el-form-item>
            <el-form-item prop="code">
              <el-row :gutter="10">
                <el-col :span="14">
                  <el-input
                    v-model="loginForm.code"
                    placeholder="请输入验证码"
                    prefix-icon="el-icon-picture"
                    size="large"
                    @keyup.enter.native="handleLogin"
                  />
                </el-col>
                <el-col :span="10">
                  <img
                    v-if="captchaImage"
                    :src="captchaImage"
                   
                    alt="验证码"
                    class="captcha-image"
                    @click="getCaptcha"
                    title="点击刷新验证码"
                  />
                </el-col>
              </el-row>
            </el-form-item>
            <el-form-item>
              <el-button
                type="primary"
                size="large"
                style="width: 100%"
                :loading="loginLoading"
                @click="handleLogin"
              >
                登录
              </el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        
        <!-- 注册表单 -->
        <el-tab-pane label="注册" name="register">
          <el-form
            ref="registerForm"
            :model="registerForm"
            :rules="registerRules"
            label-width="0px"
          >
            <el-form-item prop="username">
              <el-input
                v-model="registerForm.username"
                placeholder="请输入用户名"
                prefix-icon="el-icon-user"
                size="large"
              />
            </el-form-item>
            <el-form-item prop="email">
              <el-input
                v-model="registerForm.email"
                placeholder="请输入邮箱"
                prefix-icon="el-icon-message"
                size="large"
              />
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                v-model="registerForm.password"
                type="password"
                placeholder="请输入密码"
                prefix-icon="el-icon-lock"
                size="large"
              />
            </el-form-item>
            <el-form-item prop="confirmPassword">
              <el-input
                v-model="registerForm.confirmPassword"
                type="password"
                placeholder="请确认密码"
                prefix-icon="el-icon-lock"
                size="large"
              />
            </el-form-item>
            <el-form-item prop="code">
              <el-row :gutter="10">
                <el-col :span="14">
                  <el-input
                    v-model="registerForm.code"
                    placeholder="请输入验证码"
                    prefix-icon="el-icon-picture"
                    size="large"
                    @keyup.enter.native="handleRegister"
                  />
                </el-col>
                <el-col :span="10">
                  <img
                    v-if="captchaImage"
                    :src="captchaImage"
                    alt="验证码"
                    class="captcha-image"
                    @click="getCaptcha"
                    title="点击刷新验证码"
                  />
                </el-col>
              </el-row>
            </el-form-item>
            <el-form-item>
              <el-button
                type="primary"
                size="large"
                style="width: 100%"
                :loading="registerLoading"
                @click="handleRegister"
              >
                注册
              </el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    // 确认密码验证规则
    const validateConfirmPassword = (rule, value, callback) => {
      if (value !== this.registerForm.password) {
        callback(new Error('两次输入密码不一致'))
      } else {
        callback()
      }
    }
    
    return {
      activeTab: 'login',
      loginLoading: false,
      registerLoading: false,
      // 验证码相关
      captchaImage: '',
      captchaId: '',
      // 登录表单数据
      loginForm: {
        username: '',
        password: '',
        code: ''
      },
      // 注册表单数据
      registerForm: {
        username: '',
        email: '',
        password: '',
        confirmPassword: '',
        code: ''
      },
      // 登录表单验证规则
      loginRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 20, message: '用户名长度在 3 到 20 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 5, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
        ],
        code: [
          { required: true, message: '请输入验证码', trigger: 'blur' },
          { len: 4, message: '验证码长度为4位', trigger: 'blur' }
        ]
      },
      // 注册表单验证规则
      registerRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 20, message: '用户名长度在 3 到 20 个字符', trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请确认密码', trigger: 'blur' },
          { validator: validateConfirmPassword, trigger: 'blur' }
        ],
        code: [
          { required: true, message: '请输入验证码', trigger: 'blur' },
          { len: 4, message: '验证码长度为4位', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    // 获取验证码
    getCaptcha() {
      console.log('开始获取验证码...')
      this.$http.get('/user/captcha')
        .then(response => {
          console.log('验证码API响应:', response)
          console.log('响应状态:', response.status)
          console.log('响应数据:', response.data)
          
          if (response.data && response.data.code === 100) {
            this.captchaImage = response.data.data.image64
            this.captchaId = response.data.data.captchaId
            console.log('验证码获取成功')
          } else {
            console.error('验证码API返回错误:', response.data)
            this.$message.error(response.data?.message || '获取验证码失败')
          }
        })
        .catch(error => {
          console.error('获取验证码网络错误:', error)
          console.error('错误详情:', error.response)
          if (error.response) {
            console.error('响应状态:', error.response.status)
            console.error('响应数据:', error.response.data)
          }
          this.$message.error('获取验证码失败，请检查网络连接')
        })
    },
    
    // 刷新验证码
    refreshCaptcha() {
      this.getCaptcha()
    },
    
    /**
     * 处理登录操作
     */
    handleLogin() {
      this.$refs.loginForm.validate(async (valid) => {
        if (valid) {
          this.loginLoading = true
          try {
            const loginData = {
              username: this.loginForm.username,
              password: this.loginForm.password,
              code: this.loginForm.code,
              captchaId: this.captchaId
            }
            const response = await this.$http.post('/user/login', loginData)
            
            if (response.data.code === 100) {
              // 保存用户信息和token到store
              this.$store.dispatch('login', {
                userInfo: { username: response.data.data.username },
                token: response.data.data.token
              })
              
              this.$message.success('登录成功')
              this.$router.push('/dashboard')
            } else {
              this.$message.error(response.data.message || '登录失败')
              // 登录失败后刷新验证码
              this.getCaptcha()
              this.loginForm.code = ''
            }
          } catch (error) {
            this.$message.error(error.response?.data?.message || '登录失败')
            // 登录失败后刷新验证码
            this.getCaptcha()
            this.loginForm.code = ''
          } finally {
            this.loginLoading = false
          }
        }
      })
    },
    
    /**
     * 处理注册操作
     */
    handleRegister() {
      this.$refs.registerForm.validate(async (valid) => {
        if (valid) {
          this.registerLoading = true
          try {
            const registerData = {
              username: this.registerForm.username,
              email: this.registerForm.email,
              password: this.registerForm.password,
              code: this.registerForm.code,
              captchaId: this.captchaId
            }
            const response = await this.$http.post('/user/register', registerData)
            
            if (response.data.code === 100) {
              // 保存token和用户信息
              this.$store.dispatch('login', {
                token: response.data.data.token,
                userInfo: { username: response.data.data.username }
              })
              this.$message.success('注册成功')
              this.$router.push('/dashboard')
            } else {
              this.$message.error(response.data.message || '注册失败')
              // 注册失败后刷新验证码
              this.getCaptcha()
              this.registerForm.code = ''
            }
          } catch (error) {
            this.$message.error(error.response?.data?.message || '注册失败')
            // 注册失败后刷新验证码
            this.getCaptcha()
            this.registerForm.code = ''
          } finally {
            this.registerLoading = false
          }
        }
      })
    },
    
    /**
     * 重置表单
     */
    resetForms() {
      this.loginForm = {
        username: '',
        password: '',
        code: ''
      }
      this.registerForm = {
        username: '',
        email: '',
        password: '',
        confirmPassword: '',
        code: ''
      }
      this.$refs.loginForm && this.$refs.loginForm.resetFields()
      this.$refs.registerForm && this.$refs.registerForm.resetFields()
    }
  },
  
  created() {
    // 如果已经登录，直接跳转到仪表盘
    if (this.$store.getters.getLoginStatus) {
      this.$router.push('/dashboard')
    } else {
      // 获取验证码
      this.getCaptcha()
    }
  }
}
</script>

<style scoped>
/* 验证码图片样式 */
.captcha-image {
  width: 100%;
  height: 40px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  cursor: pointer;
  object-fit: contain;
  background-color: #f5f7fa;
}

.captcha-image:hover {
  border-color: #409eff;
}
</style>