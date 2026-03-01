# StudyOnline-在线教育平台后端

## 启动流程

1. init里配置mysql和redis
2. go mod tidy
3. 运行main.go

## 技术选型

登录模块：用Redis存储token，有效时长12小时

其他模块：用Mysql做后端数据持久化

服务限流：按IP限速（每个IP独立桶），桶大小为50，每秒40个token，预计100人使用，理论qps上限=4000

## 部署到远端内网服务器

1. 先登录相关平台网站，ssh方式脸上服务器。
2. 内网访问不了github，所以有一个gitee网站同步此项目：https://gitee.com/evan_yin/studyonline-main
3. 云端git pull最新代码，用nohup方式后台启动相应进程即可。