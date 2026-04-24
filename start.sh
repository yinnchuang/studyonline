#!/bin/bash

log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

log "开始执行部署流程..."

# 1. Git pull 拉取最新代码
log "拉取最新代码..."
if git pull; then
    log "代码更新成功"
else
    log "代码更新失败！"
    exit 1
fi

# 2. 编译 main.go
log "编译 main.go..."
if go build -o main main.go; then
    log "编译成功"
else
    log "编译失败！"
    exit 1
fi

# 3. 检查并杀掉占用 8080 和 12010 端口的进程
log "检查并清理端口占用..."
for port in 8080 12010; do
    log "检查端口 $port..."
    pids=$(lsof -t -i:$port 2>/dev/null)
    if [ -n "$pids" ]; then
        log "发现端口 $port 被占用，PID: $pids，正在终止..."
        kill -9 $pids
        sleep 1
        log "端口 $port 已清理"
    else
        log "端口 $port 未被占用"
    fi
done

# 4. 后台运行 Go 程序
log "启动 Go 服务..."
nohup ./main > app.log 2>&1 &
log "Go 服务已启动，日志: app.log"

# 5. 后台运行 Python 服务
log "启动 Python 服务..."
cd studyonline_AI_Lessons_Plan || exit 1
source venv/bin/activate
nohup python app.py > llm.log 2>&1 &
log "Python 服务已启动，日志: llm.log"

log "所有服务部署完成！"
