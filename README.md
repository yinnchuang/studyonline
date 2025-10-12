# StudyOnline-在线教育平台后端

## 启动流程

1.init里配置mysql和redis 

2.go mod tidy 

3.运行main.go

## 技术选型

登录模块：用Redis存储token，有效时长12小时

其他模块：用Mysql做后端数据持久化