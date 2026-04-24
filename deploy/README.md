# StudyOnline 部署说明

## 文件说明

- `start.sh` - 启动脚本
- `stop.sh` - 停止脚本
- `install.sh` - 自动安装脚本（推荐使用）
- `studyonline.service` - Systemd 服务文件

## 部署步骤（推荐方式）

### 1. 使用自动安装脚本

```bash
cd deploy
./install.sh
```

这个脚本会：
- 自动获取项目路径
- 自动获取当前用户和用户组
- 自动生成 studyonline.service 文件

### 2. 手动安装（如果不使用 install.sh）

编辑 `studyonline.service`，修改以下内容：
```ini
User=your_username       # 修改为实际的用户名
Group=your_group         # 修改为实际的用户组
WorkingDirectory=/path/to/studyonline  # 修改为实际的项目路径
ExecStart=/path/to/studyonline/deploy/start.sh
ExecStop=/path/to/studyonline/deploy/stop.sh
```

### 2. 设置执行权限

```bash
chmod +x deploy/start.sh
chmod +x deploy/stop.sh
```

### 3. 创建 Python 虚拟环境（可选，脚本会自动创建）

```bash
cd studyonline_AI_Lessons_Plan
python3 -m venv venv
source venv/bin/activate
# 如果有 requirements.txt
pip install -r requirements.txt
```

### 4. 安装 Systemd 服务

```bash
# 复制 service 文件到 systemd 目录
sudo cp deploy/studyonline.service /etc/systemd/system/

# 重载 systemd 配置
sudo systemctl daemon-reload

# 启动服务
sudo systemctl start studyonline

# 设置开机自启
sudo systemctl enable studyonline
```

## 常用命令

```bash
# 查看服务状态
sudo systemctl status studyonline

# 查看服务日志
sudo journalctl -u studyonline -f

# 停止服务
sudo systemctl stop studyonline

# 重启服务
sudo systemctl restart studyonline

# 重新加载配置（修改了 service 文件后）
sudo systemctl daemon-reload
sudo systemctl restart studyonline
```

## 日志查看

- Systemd 日志：`sudo journalctl -u studyonline -f`
- main.go 日志：`/path/to/studyonline/main.log`
- app.py 日志：`/path/to/studyonline/studyonline_AI_Lessons_Plan/app.log`

## 手动启动/停止（不使用 Systemd）

```bash
# 手动启动
./deploy/start.sh

# 手动停止
./deploy/stop.sh
```
