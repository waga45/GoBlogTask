# GoBlog 后台管理系统

基于 Vue2 + Element UI 开发的博客后台管理系统。

## 项目特性

- 🔐 完整的登录/注册功能
- 🛡️ 路由守卫，未登录无法访问内部页面
- 📝 文章管理（列表、新建、编辑、删除）
- 🎨 基于 Element UI 的现代化界面
- 📱 响应式设计

## 技术栈

- **前端框架**: Vue 2.6.14
- **UI 组件库**: Element UI 2.15.9
- **路由管理**: Vue Router 3.5.4
- **状态管理**: Vuex 3.6.2
- **HTTP 客户端**: Axios 0.27.2
- **构建工具**: Vue CLI 5.0.8

## 项目结构

```
src/
├── assets/          # 静态资源
│   └── css/         # 全局样式
├── components/      # 公共组件
├── router/          # 路由配置
├── store/           # Vuex 状态管理
├── views/           # 页面组件
│   ├── Login.vue    # 登录页面
│   ├── Layout.vue   # 主布局
│   ├── Articles.vue # 文章列表
│   └── WriteArticle.vue # 写文章
├── App.vue          # 根组件
└── main.js          # 入口文件
```

## 功能说明

### 1. 登录系统
- 支持用户登录和注册
- 表单验证
- Token 认证
- 自动跳转

### 2. 路由守卫
- 未登录用户无法访问内部页面
- 自动重定向到登录页
- 登录后跳转到原始页面

### 3. 文章管理
- 文章列表展示
- 搜索和筛选
- 新建/编辑文章
- 文章状态管理（草稿/发布）
- 文章删除

### 4. 用户界面
- 顶部导航栏显示用户信息
- 左侧菜单导航
- 退出登录功能

## API 接口

项目配置的 API 基础地址为：`http://127.0.0.1:8888/v1/`

### 主要接口：

- `POST /auth/login` - 用户登录
- `POST /auth/register` - 用户注册
- `GET /articles` - 获取文章列表
- `POST /articles` - 创建文章
- `PUT /articles/:id` - 更新文章
- `DELETE /articles/:id` - 删除文章
- `PUT /articles/:id/status` - 更新文章状态

## 安装和运行

### 1. 安装依赖
```bash
npm install --legacy-peer-deps
```

### 2. 启动开发服务器
```bash
npm run serve
```

### 3. 构建生产版本
```bash
npm run build
```

## 开发说明

### 环境要求
- Node.js >= 14.0.0
- npm >= 6.0.0

### 开发端口
- 前端开发服务器：http://localhost:8080
- 后端 API 服务器：http://127.0.0.1:8888

### 注意事项
1. 确保后端 API 服务正常运行
2. 检查 API 接口地址配置
3. 登录后的 Token 会保存在 localStorage 中
4. 路由使用 history 模式，需要服务器配置支持

## 待完善功能

- [ ] 文章分类管理
- [ ] 文件上传功能
- [ ] 富文本编辑器
- [ ] 用户权限管理
- [ ] 数据统计面板
- [ ] 评论管理

## 许可证

MIT License