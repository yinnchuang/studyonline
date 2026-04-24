#!/bin/bash

# 获取脚本所在目录，然后设置为项目根目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
WORK_DIR="$SCRIPT_DIR/.."
cd "$WORK_DIR" || exit 1

# 函数：记录日志
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

log "开始停止 StudyOnline 服务..."

# 停止 main.go
if [ -f "main.pid" ]; then
    MAIN_PID=$(cat main.pid)
    if kill -0 "$MAIN_PID" 2>/dev/null; then
        log "停止 main.go (PID: $MAIN_PID)..."
        kill "$MAIN_PID"
        sleep 1
        if kill -0 "$MAIN_PID" 2>/dev/null; then
            kill -9 "$MAIN_PID"
        fi
    fi
    rm -f main.pid
fi

# 停止 app.py
cd "$WORK_DIR/studyonline_AI_Lessons_Plan" || exit 1
if [ -f "app.pid" ]; then
    APP_PID=$(cat app.pid)
    if kill -0 "$APP_PID" 2>/dev/null; then
        log "停止 app.py (PID: $APP_PID)..."
        kill "$APP_PID"
        sleep 1
        if kill -0 "$APP_PID" 2>/dev/null; then
            kill -9 "$APP_PID"
        fi
    fi
    rm -f app.pid
fi

cd "$WORK_DIR"

# 清理编译的 main 文件（可选）
# rm -f main

log "所有服务已停止！"
