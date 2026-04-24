#!/bin/bash

# 获取脚本所在目录，然后设置为项目根目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
WORK_DIR="$SCRIPT_DIR/.."
cd "$WORK_DIR" || exit 1

# 函数：记录日志
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

log "开始启动 StudyOnline 服务..."

# 1. 编译 main.go
log "编译 main.go..."
if go build -o main main.go; then
    log "编译成功"
else
    log "编译失败！"
    exit 1
fi

# 2. 启动 main.go
log "启动 main.go 服务..."
./main > main.log 2>&1 &
MAIN_PID=$!
echo "$MAIN_PID" > main.pid
log "main.go 已启动，PID: $MAIN_PID"

# 等待一下确保 main.go 启动
sleep 2

# 3. 启动 Python 服务
log "启动 studyonline_AI_Lessons_Plan/app.py..."
cd "$WORK_DIR/studyonline_AI_Lessons_Plan" || exit 1

# 检查虚拟环境是否存在，如果不存在则创建
if [ ! -d "venv" ]; then
    log "创建 Python 虚拟环境..."
    python3 -m venv venv
fi

# 激活虚拟环境
log "激活虚拟环境..."
source venv/bin/activate

# 安装依赖（如果需要）
if [ -f "requirements.txt" ]; then
    log "安装 Python 依赖..."
    pip install -r requirements.txt
fi

# 启动 Python 服务
log "启动 app.py..."
python app.py > app.log 2>&1 &
APP_PID=$!
echo "$APP_PID" > app.pid
log "app.py 已启动，PID: $APP_PID"

cd "$WORK_DIR"

log "所有服务启动完成！"
log "main PID: $MAIN_PID"
log "app PID: $APP_PID"
