#!/bin/bash

log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

log "拉取最新代码..."
if git pull; then
    log "代码更新成功"
else
    log "代码更新失败！"
    exit 1
fi
