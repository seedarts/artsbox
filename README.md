# ArtsBox AI 图片生成器

<div align="center">

![ArtsBox](https://img.shields.io/badge/ArtsBox-v1.0.0-blue)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS-lightgrey)
![License](https://img.shields.io/badge/License-MIT-green)
![Wails](https://img.shields.io/badge/Wails-v2.11.0-purple)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Vue](https://img.shields.io/badge/Vue-3.x-4FC08D?logo=vue.js)

**一款基于 Wails + Vue 的现代化 AI 图片生成桌面应用**

[功能特性](#-功能特性) •
[快速开始](#-快速开始) •
[使用指南](#-使用指南) •
[开发文档](#-开发文档) •
[API 文档](./API.md) •
[Wiki](./WIKI.md)

</div>

---

## 📖 项目简介

ArtsBox 是一款跨平台桌面应用，利用先进的 AI 图像生成技术，让创作者能够通过简单的文字描述或图片输入，快速生成高质量的艺术作品。





### 核心优势

- 🎨 **多种生成模式** - 支持文生图、图文生图、多图融合、组图输出
- 🚀 **原生性能** - 基于 Wails 框架，拥有接近原生应用的性能
- 🎯 **简洁易用** - 现代化 UI 设计，操作直观流畅
- 💾 **本地优先** - 配置和图片均保存在本地，保护隐私
- 🔌 **API 兼容** - 支持多种 AI 图像生成服务（默认 Seedream）

---

## ✨ 功能特性

### 1️⃣ 文生图 (Text-to-Image)
纯文本描述生成单张图片
- 输入创意描述
- 选择图片质量（2K/4K）
- 一键生成高质量图片

### 2️⃣ 图文生图 (Image+Text-to-Image)
基于参考图片 + 文字描述生成新图片
- 上传 1 张参考图片
- 添加修改描述
- 生成符合预期的新图

### 3️⃣ 多图融合 (Multi-Image-Fusion)
融合多张图片风格生成全新作品
- 上传最多 10 张参考图片
- 输入融合描述
- 创造独特的艺术作品

### 4️⃣ 组图输出 (Batch-Generation)
批量生成 1-4 张连贯图片
- **文生组图**: 文字 → 多图（如四季变换）
- **单图生组图**: 单图 + 文字 → 多图（如品牌设计套件）
- **多图生组图**: 多图 + 文字 → 多图（如时间序列）

### 🔧 其他特性

- ⚙️ **灵活配置** - 自定义 API 地址、模型选择
- 📁 **自动保存** - 生成的图片自动保存到本地
- 🖼️ **图片管理** - 缩略图网格预览，快速查看历史作品
- 🔄 **实时刷新** - 生成完成自动加载显示
- 🖱️ **快速操作** - 点击缩略图用系统默认程序打开原图

---

## 🚀 快速开始

### 系统要求

- **Windows**: Windows 10/11
- **macOS**: macOS 10.15+ (Intel/Apple Silicon)
- **WebView2**: Windows 11 自带，Windows 10 需手动安装

### 下载安装

#### 方式一：使用安装程序（推荐）

1. 下载最新版本的安装包
2. 双击运行安装程序
3. 按照向导完成安装
4. 从桌面或开始菜单启动应用

#### 方式二：绿色版

1. 下载 `artsbox.exe`
2. 放置在任意目录
3. 双击运行

### 首次配置

1. **启动应用** - 双击打开 ArtsBox
2. **打开设置** - 点击左侧底部的 ⚙️ 设置图标
3. **配置 API**:
   - **API Key**: 输入您的 AI 服务 API 密钥
   - **Base URL**: 默认 `https://api.artsmcp.com/v1`（可自定义）
   - **模型选择**: 选择 `doubao-seedream-4-5-251128` 或 `doubao-seedream-4-0-250828`
   - **保存目录**: 设置图片保存路径（默认：`%USERPROFILE%\Pictures\ArtsBox`）
4. **保存配置** - 点击保存按钮

---

## 📝 使用指南

### 文生图模式

1. 点击顶部 **"文生图"** 标签
2. 在提示词输入框中输入描述（支持中英文）
   ```
   示例：一只可爱的海獭宝宝，漂浮在碧蓝的海面上，阳光洒在它毛茸茸的身上
   ```
3. 选择图片质量：2K 或 4K
4. 点击 **"开始生成"** 按钮
5. 等待生成完成，图片自动显示在右侧预览区

### 图文生图模式

1. 点击顶部 **"图文生成"** 标签
2. 点击上传区域，选择 1 张本地图片
3. 输入修改描述
   ```
   示例：保持人物姿势不变，将背景改为樱花盛开的公园
   ```
4. 点击 **"开始生成"**

### 多图融合模式

1. 点击顶部 **"多图融合"** 标签
2. 点击上传区域，选择 2-10 张本地图片
3. 输入融合描述
   ```
   示例：融合所有图片的艺术风格，创作一幅超现实主义作品
   ```
4. 点击 **"开始生成"**

### 组图输出模式

1. 点击顶部 **"组图输出"** 标签
2. 选择子模式：
   - **文生组图**: 纯文本输入
   - **单图生组图**: 上传 1 张参考图
   - **多图生组图**: 上传 2-10 张参考图
3. 拖动滑块设置生成数量（1-4 张）
4. 输入提示词
   ```
   示例（文生组图）：生成一组共4张连贯插画，展现同一场景的春夏秋冬
   ```
5. 点击 **"开始生成"**

### 图片管理

- **查看历史**: 右侧自动显示所有已保存的图片
- **刷新列表**: 点击 🔄 刷新按钮
- **打开原图**: 点击缩略图用系统默认图片查看器打开
- **修改目录**: 点击 📁 按钮选择新的保存目录

---

## 🛠️ 开发文档

### 技术栈

- **后端**: Go 1.21+
- **GUI 框架**: Wails v2.11.0
- **前端**: Vue 3 + Vite
- **构建工具**: Wails CLI
- **API**: 兼容 OpenAI 格式的图像生成服务

### 项目结构

```
artsbox/
├── main.go                 # 应用入口
├── app.go                  # 核心业务逻辑
├── pkg/
│   ├── config/            # 配置管理
│   └── openai/            # API 客户端
├── frontend/
│   ├── src/
│   │   ├── components/    # Vue 组件
│   │   ├── App.vue        # 主应用
│   │   └── style.css      # 样式
│   └── wailsjs/           # Wails JS 绑定
└── build/                  # 构建输出
```

### 本地开发

#### 环境准备

1. **安装 Go** (1.21+)
   ```bash
   # 验证安装
   go version
   ```

2. **安装 Node.js** (16+)
   ```bash
   # 验证安装
   node --version
   npm --version
   ```

3. **安装 Wails CLI**
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   
   # 验证安装
   wails version
   ```

#### 克隆项目

```bash
git clone https://github.com/yourusername/artsbox02.git
cd artsbox02/artsbox
```

#### 安装依赖

```bash
# 安装 Go 依赖
go mod tidy

# 安装前端依赖
cd frontend
npm install
cd ..
```

#### 开发模式运行

```bash
wails dev
```

应用会自动启动，支持热重载。修改代码后会自动重新编译。

#### 生产构建

```bash
# 构建可执行文件
wails build

# 输出: build/bin/artsbox.exe (Windows)
#       build/bin/artsbox.app (macOS)
```

#### 构建安装程序

**使用 PowerShell 脚本**（无需额外工具）:
```powershell
.\Install.ps1
```

**使用 NSIS**（需安装 NSIS）:
```bash
# 安装 NSIS 后
wails build -nsis

# 或使用自定义脚本
makensis installer-simple.nsi
```

### 核心 API 方法

查看完整 API 文档：[API.md](./API.md)

| 方法 | 说明 |
|------|------|
| `GetSettings()` | 获取应用配置 |
| `SaveSettings()` | 保存应用配置 |
| `GenerateImage()` | 文生图 |
| `GenerateFromImage()` | 图文生图 |
| `GenerateFromMultiImages()` | 多图融合 |
| `GenerateBatchImages()` | 文生组图 |
| `GenerateBatchFromImage()` | 单图生组图 |
| `GenerateBatchFromMultiImages()` | 多图生组图 |
| `GetSavedImages()` | 获取保存的图片列表 |
| `GetImageAsDataURL()` | 获取图片 base64 数据 |
| `OpenImageFile()` | 打开图片文件 |
| `SelectSaveDirectory()` | 选择保存目录 |

---

## 📚 更多文档

- [Wiki 文档](./WIKI.md) - 详细的使用教程和常见问题
- [API 文档](./API.md) - 完整的 API 接口说明
- [开发计划](./development_plan.md) - 项目设计和开发路线图
- [打包指南](./PACKAGING.md) - 应用打包和分发指南

---

## ⚙️ 配置说明

### 配置文件位置

- **Windows**: `%USERPROFILE%\.artsbox\config.json`
- **macOS**: `~/.artsbox/config.json`

### 配置文件格式

```json
{
  "api_key": "your-api-key",
  "base_url": "https://api.cozex.cn/v1",
  "model": "doubao-seedream-4-5-251128",
  "save_dir": "C:\\Users\\YourName\\Pictures\\ArtsBox"
}
```

### 支持的模型

- `doubao-seedream-4-5-251128` (推荐) - 最新模型
- `doubao-seedream-4-0-250828` - 稳定版本

---

## 🐛 常见问题

### 应用无法启动

**问题**: 双击 `artsbox.exe` 没有反应

**解决方案**:
1. 检查是否安装了 WebView2 运行时（Windows 10）
2. 右键应用 → 属性 → 解除锁定
3. 查看错误日志（如果有）

### 图片无法显示

**问题**: 右侧预览区域空白

**解决方案**:
1. 点击 🔄 刷新按钮
2. 检查保存目录是否有图片文件
3. 确认图片格式（支持：png, jpg, jpeg, jpe, webp）

### API 调用失败

**问题**: 生成图片时报错

**解决方案**:
1. 检查 API Key 是否正确
2. 确认 Base URL 是否正确
3. 检查网络连接
4. 查看 API 服务状态

### 开发模式无法启动

**问题**: `wails dev` 卡住不动

**解决方案**:
1. 确保端口 5173 和 34115 未被占用
2. 删除 `frontend/node_modules` 重新安装
3. 尝试使用生产构建：`wails build`

---

## 🤝 贡献指南

欢迎贡献代码、报告问题或提出建议！

### 报告 Bug

1. 在 [Issues](../../issues) 中搜索是否已存在相关问题
2. 如果没有，创建新 Issue
3. 提供详细的复现步骤和环境信息

### 提交代码

1. Fork 本仓库
2. 创建特性分支：`git checkout -b feature/amazing-feature`
3. 提交更改：`git commit -m 'Add amazing feature'`
4. 推送到分支：`git push origin feature/amazing-feature`
5. 提交 Pull Request

### 开发规范

- 遵循 Go 和 Vue 的最佳实践
- 添加必要的注释
- 确保代码通过测试
- 更新相关文档

---

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](./LICENSE.txt) 文件了解详情

---

## 🙏 致谢

- [Wails](https://wails.io) - 强大的 Go + Web 桌面应用框架
- [Vue.js](https://vuejs.org) - 渐进式 JavaScript 框架
- [豆包 Seedream](https://www.volcengine.com/docs/6791/1209058) - AI 图像生成服务

---

## 📧 联系方式

- **WX**:artsmcp 
- **TG**：suarts
- **Issue**: [GitHub Issues](../../issues)

---

<div align="center">

**如果觉得这个项目有帮助，请给个 ⭐️ Star！**

Made with ❤️ by ArtsBox Team

</div>
