# StudyOnline - 在线教育平台后端

一个支持教师、学生、管理员三种角色的在线教育管理平台，提供教案管理、AI 智能生成教案、数据集管理、作业与讨论等功能。

---

## 目录结构

```
studyonline/
├── main.go                 # 入口，注册路由/中间件，启动服务
├── config/                 # 配置读取（init/project.ini）
├── constant/               # 常量定义（身份、教案状态、分类、文件大小限制）
├── dao/
│   ├── entity/             # GORM 数据模型（表结构）
│   ├── mysql/              # MySQL 连接初始化 & 自动建表
│   └── redis/              # Redis 连接初始化
├── handler/                # HTTP Handler 层，参数校验 & 响应格式化
│   └── middleware/         # 中间件（auth、CORS、限流）
├── service/                # 业务逻辑层
├── log/                    # 日志初始化
├── util/                   # 工具函数（token、密码加密、验证码邮件）
├── init/                   # 配置文件 project.ini
├── static/                 # 静态文件目录（文件下载/上传使用）
├── script/                 # 脚本与 Excel 模板
├── test/                   # 测试文件
└── studyonline_AI_Lessons_Plan/  # Python AI 教案生成服务
```

---

## 技术栈

| 组件 | 技术 |
|------|------|
| Web 框架 | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io/) |
| 数据库 | MySQL |
| 缓存 | Redis |
| AI 服务 | Python Flask + DeepSeek API |
| 密码加密 | bcrypt |
| Token | UUID |

---

## 快速启动

### 前置依赖

- Go 1.18+
- MySQL 5.7+
- Redis
- Python 3.8+（AI 教案生成服务）

### 1. 配置

编辑 `init/project.ini`：

```ini
[mysql]
ip = 127.0.0.1
port = 3306
user = root
password = 123456
database = studyonline

[redis]
ip = 127.0.0.1
port = 6379
password =

[admin]
username = admin
password = password

[deepseek]
api_key = <your-api-key>
url = https://api.deepseek.com/v1/chat/completions
model = deepseek-v4-flash
```

### 2. 启动 Go 服务

```bash
go mod tidy
go run main.go
```

### 3. 启动 AI 教案生成服务（可选）

```bash
cd studyonline_AI_Lessons_Plan
pip install -r requirements.txt
python app.py  # 监听 0.0.0.0:12010
```

---

## 用户角色

| 角色 | 常量 | 说明 |
|------|------|------|
| 学生 | `Student` | 查看资源、下载数据集、提交作业、参与讨论 |
| 教师 | `Teacher` | 管理教案、发布数据集、批改作业、管理学生 |
| 管理员 | `Admin` | 管理全部用户、数据集、公告等 |

---

## 数据模型

### 核心实体

| 实体 | 文件 | 说明 |
|------|------|------|
| `Student` | `dao/entity/student.go` | 学生信息（姓名、学号、院系、密码等） |
| `Teacher` | `dao/entity/teacher.go` | 教师信息 |
| `Admin` | `dao/entity/admin.go` | 管理员信息 |
| `LessonPlan` | `dao/entity/LessonPlan.go` | 教案（标题、课时、目标、重难点、内容、思政要点等） |
| `LessonPlanStudent` | `dao/entity/LessonPlanStudent.go` | 教师与学生的教案关联 |
| `Dataset` | `dao/entity/dataset.go` | 数据集（名称、上传者、文件路径、访问权限等） |
| `Resource` | `dao/entity/resource.go` | 教学资源（文件名、上传者、文件路径） |
| `Unit` | `dao/entity/unit.go` | 教学单元 |
| `Homework` | `dao/entity/homework.go` | 作业（标题、内容、截止时间） |
| `Submission` | `dao/entity/submission.go` | 作业提交（学生提交、得分、评语） |
| `Score` | `dao/entity/score.go` | 学生成绩 |
| `Discuss` | `dao/entity/discuss.go` | 教案讨论 |
| `DiscussLike` | `dao/entity/lessonplan_discuss_like.go` | 讨论点赞 |
| `Comment` | `dao/entity/comment.go` | 教案评论 |
| `Permission` | `dao/entity/permission.go` | 数据集访问权限 |
| `Announcement` | `dao/entity/announcement.go` | 系统公告 |
| `DownloadLog` | `dao/entity/downloadlog.go` | 下载日志 |

---

## API 接口

### 认证相关

| 方法 | 路径 | Handler | 说明 |
|------|------|---------|------|
| POST | `/login/student` | `handler/student/login.go` | 学生登录 |
| POST | `/login/teacher` | `handler/teacher/login.go` | 教师登录 |
| POST | `/login/admin` | `handler/admin/login.go` | 管理员登录 |
| POST | `/register/student` | `handler/student/register.go` | 学生注册 |
| POST | `/register/teacher` | `handler/teacher/register.go` | 教师注册 |
| POST | `/sendCode` | `handler/sendCode.go` | 发送邮件验证码 |
| POST | `/resetPassword` | `handler/resetPassword.go` | 重置密码 |

### 学生相关

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/student/info` | `handler/student/info.go` | 获取个人信息 | auth-student |
| PUT | `/student/update` | `handler/student/update.go` | 修改个人信息 | auth-student |
| GET | `/student/homework` | `handler/student/homework.go` | 查看作业列表 | auth-student |
| PUT | `/student/updatePassword` | `handler/student/updatePassword.go` | 修改密码 | auth-student |

### 教师相关

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/teacher/info` | `handler/teacher/info.go` | 获取个人信息 | auth-teacher |
| PUT | `/teacher/update` | `handler/teacher/update.go` | 修改个人信息 | auth-teacher |
| GET | `/teacher/student` | `handler/teacher/student.go` | 查看学生列表 | auth-teacher |
| PUT | `/teacher/updatePassword` | `handler/teacher/updatePassword.go` | 修改密码 | auth-teacher |

### 教案管理

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/lessonplan/list` | `handler/lessonplan/list.go` | 教案列表 | auth-teacher |
| POST | `/lessonplan/add` | `handler/lessonplan/add.go` | 新增教案 | auth-teacher |
| GET | `/lessonplan/detail/:id` | `handler/lessonplan/detail.go` | 教案详情 | auth-teacher |
| PUT | `/lessonplan/update/:id` | `handler/lessonplan/update.go` | 修改教案 | auth-teacher |
| DELETE | `/lessonplan/delete/:id` | `handler/lessonplan/delete.go` | 删除教案 | auth-teacher |
| GET | `/lessonplan/publish/:id` | `handler/lessonplan/publish.go` | 发布教案 | auth-teacher |
| GET | `/lessonplan/revert/:id` | `handler/lessonplan/revert.go` | 撤回发布 | auth-teacher |
| POST | `/lessonplan/generate` | `handler/lessonplan/generate.go` | AI 生成教案 | auth-teacher |

**教案状态：** `Draft(0)` → `Published(1)` → `Reverted(2)`

### 数据集管理

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/dataset/list` | `handler/dataset/list.go` | 数据集列表 | auth |
| POST | `/dataset/add` | `handler/dataset/add.go` | 上传数据集 | auth-teacher/auth-admin |
| PUT | `/dataset/update/:id` | `handler/dataset/update.go` | 修改数据集 | auth-teacher/auth-admin |
| DELETE | `/dataset/delete/:id` | `handler/dataset/delete.go` | 删除数据集 | auth-teacher/auth-admin |
| GET | `/dataset/download/:id` | `handler/dataset/download.go` | 下载数据集 | auth |

**数据集分类：** 公开(Public)、课程相关(Course)、实验课(Lab)

### 资源管理

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/resource/list` | `handler/resource/list.go` | 资源列表 | auth |
| POST | `/resource/add` | `handler/resource/add.go` | 上传资源 | auth-teacher/auth-admin |
| PUT | `/resource/update/:id` | `handler/resource/update.go` | 修改资源 | auth-teacher/auth-admin |
| DELETE | `/resource/delete/:id` | `handler/resource/delete.go` | 删除资源 | auth-teacher/auth-admin |
| GET | `/resource/download/:id` | `handler/resource/download.go` | 下载资源 | auth |

### 成绩管理

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/score/list` | `handler/score/list.go` | 成绩列表 | auth-teacher |
| POST | `/score/add` | `handler/score/add.go` | 录入成绩 | auth-teacher |
| PUT | `/score/update/:id` | `handler/score/update.go` | 修改成绩 | auth-teacher |
| DELETE | `/score/delete/:id` | `handler/score/delete.go` | 删除成绩 | auth-teacher |
| GET | `/score/student` | `handler/score/student.go` | 学生查看成绩 | auth-student |

### 讨论与评论

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/discuss/list/:lessonplan_id` | `handler/discuss/list.go` | 讨论列表 | auth |
| POST | `/discuss/add` | `handler/discuss/add.go` | 发起讨论 | auth |
| DELETE | `/discuss/delete/:id` | `handler/discuss/delete.go` | 删除讨论 | auth |
| POST | `/discuss/like/:id` | `handler/discuss/like.go` | 点赞讨论 | auth |
| GET | `/comment/list/:discuss_id` | `handler/comment/list.go` | 评论列表 | auth |
| POST | `/comment/add` | `handler/comment/add.go` | 发表评论 | auth |
| DELETE | `/comment/delete/:id` | `handler/comment/delete.go` | 删除评论 | auth |

### 作业与提交

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/homework/list` | `handler/homework/list.go` | 作业列表 | auth |
| POST | `/homework/add` | `handler/homework/add.go` | 发布作业 | auth-teacher |
| DELETE | `/homework/delete/:id` | `handler/homework/delete.go` | 删除作业 | auth-teacher |
| POST | `/submission/add` | `handler/submission/add.go` | 提交作业 | auth-student |

### 其他

| 方法 | 路径 | Handler | 说明 | 中间件 |
|------|------|---------|------|--------|
| GET | `/downloadlog/list` | `handler/downloadlog/downloadlog.go` | 下载日志 | auth-admin |
| GET | `/announcement/list` | `handler/announcement/list.go` | 公告列表 | auth |
| POST | `/announcement/add` | `handler/announcement/add.go` | 发布公告 | auth-admin |
| POST | `/permission/grant` | `handler/permission/grant.go` | 授予数据集权限 | auth-teacher/auth-admin |
| GET | `/permission/students/:dataset_id` | `handler/permission/list.go` | 已授权学生列表 | auth-teacher/auth-admin |

---

## 架构分层

```
Handler 层  ← 参数校验、响应封装、调用 Service
   ↓
Service 层  ← 业务逻辑编排
   ↓
DAO 层      ← 数据库操作（entity / mysql / redis）
```

- **Token 认证**：登录后服务端生成 UUID Token，存入 Redis，有效期 12 小时。请求需携带 `Authorization` Header。
- **密码加密**：bcrypt 加盐哈希存储，密码强度要求 ≥8 位且包含字母+数字。
- **限流策略**：基于 IP 的令牌桶限流，桶容量 50，每秒 40 token 补充，理论 QPS 上限 4000。
- **文件上传限制**：最大 32MB。

---

## AI 教案生成服务

独立的 Python Flask 服务，集成 DeepSeek API 自动生成教案的六个部分：

| 步骤 | API 函数 | 说明 |
|------|----------|------|
| 1 | `tongyi_generate_objectives` | 生成教学目标 |
| 2 | `tongyi_generate_key` | 生成教学重点 |
| 3 | `tongyi_generate_difficult` | 生成教学难点 |
| 4 | `tongyi_generate_content` | 生成完整教学过程（Markdown 格式） |
| 5 | `tongyi_generate_ideological` | 生成思政要点 |
| 6 | `tongyi_generate_reflection` | 生成教学反思 |

### Python 服务 API

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/lesson/plan/generate` | AI 生成完整教案并保存到数据库 |
| GET | `/lessonplan/<id>/evaluate` | 评价教案质量 |
| POST | `/lessonplan/<id>/reflect` | 生成教学反思 |

**DeepSeek 模型：** `deepseek-v4-flash`

---

## 部署

> 内网部署，无法直接访问 GitHub，需通过 Gitee 同步：
> https://gitee.com/evan_yin/studyonline-main

```bash
# 1. 登录内网服务器
ssh user@server

# 2. 拉取最新代码
git clone https://gitee.com/evan_yin/studyonline-main.git
cd studyonline-main

# 3. 配置 init/project.ini

# 4. 编译运行
go build -o studyonline
nohup ./studyonline &

# 5. AI 服务（可选）
cd studyonline_AI_Lessons_Plan
nohup python app.py &
```

---

## 开发规范

- Go 代码遵循项目编码规范
- 数据库变更需同步更新 GORM 模型
- 敏感信息（密码、API Key）不硬编码，统一从配置文件读取
- 日志使用结构化格式，关键操作记录审计日志
