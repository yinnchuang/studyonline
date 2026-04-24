#!/bin/bash

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$SCRIPT_DIR/.."

# 获取当前用户和用户组
CURRENT_USER=$(whoami)
CURRENT_GROUP=$(id -gn)

echo "=========================================="
echo "  StudyOnline Service 安装脚本"
echo "=========================================="
echo ""
echo "项目目录: $PROJECT_DIR"
echo "当前用户: $CURRENT_USER"
echo "当前用户组: $CURRENT_GROUP"
echo ""

# 备份原始 service 文件
if [ -f "$SCRIPT_DIR/studyonline.service" ]; then
    cp "$SCRIPT_DIR/studyonline.service" "$SCRIPT_DIR/studyonline.service.bak"
fi

# 生成 service 文件
cat > "$SCRIPT_DIR/studyonline.service" << EOF
[Unit]
Description=StudyOnline Service
After=network.target

[Service]
Type=forking
User=$CURRENT_USER
Group=$CURRENT_GROUP
WorkingDirectory=$PROJECT_DIR
ExecStart=$SCRIPT_DIR/start.sh
ExecStop=$SCRIPT_DIR/stop.sh
Restart=always
RestartSec=5
StandardOutput=journal
StandardError=journal
SyslogIdentifier=studyonline

[Install]
WantedBy=multi-user.target
EOF

echo "已生成 studyonline.service 文件"
echo ""
echo "=========================================="
echo "  下一步操作："
echo "=========================================="
echo ""
echo "1. 查看生成的 service 文件："
echo "   cat $SCRIPT_DIR/studyonline.service"
echo ""
echo "2. 如果确认无误，运行以下命令安装服务："
echo "   sudo cp $SCRIPT_DIR/studyonline.service /etc/systemd/system/"
echo "   sudo systemctl daemon-reload"
echo "   sudo systemctl start studyonline"
echo "   sudo systemctl enable studyonline"
echo ""
echo "3. 查看服务状态："
echo "   sudo systemctl status studyonline"
echo ""
