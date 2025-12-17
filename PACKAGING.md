# ArtsBox 打包发布指南

本文档介绍如何将 ArtsBox 打包成可分发的安装程序。

## 方法一：PowerShell 安装脚本（推荐，无需额外工具）

### 优点
- ✅ 无需安装 NSIS 或其他工具
- ✅ 自动创建桌面和开始菜单快捷方式
- ✅ 包含卸载功能
- ✅ 适合快速分发

### 使用步骤

1. **构建应用程序**
   ```powershell
   wails build
   ```

2. **运行安装脚本**
   ```powershell
   .\Install.ps1
   ```
   脚本会自动：
   - 请求管理员权限
   - 将应用安装到 `C:\Program Files\ArtsBox`
   - 创建桌面快捷方式
   - 创建开始菜单项
   - 生成卸载脚本

3. **分发方式**
   
   **选项 A: 打包成压缩包**
   ```
   ArtsBox-v1.0.0.zip
   ├── Install.ps1
   └── build/
       └── bin/
           └── artsbox.exe
   ```
   
   用户下载后解压，右键运行 `Install.ps1` → 选择"以管理员身份运行"

   **选项 B: 仅分发 EXE（绿色版）**
   直接分发 `build\bin\artsbox.exe`，用户可以放在任意位置运行

## 方法二：NSIS 安装程序（专业级）

### 前提条件
需要安装 NSIS (Nullsoft Scriptable Install System)

### 安装 NSIS

1. 下载 NSIS: https://nsis.sourceforge.io/Download
2. 安装到默认路径: `C:\Program Files (x86)\NSIS`
3. 将 NSIS 添加到系统 PATH

### 使用步骤

1. **构建应用程序**
   ```powershell
   wails build
   ```

2. **编译安装程序（简化版）**
   ```powershell
   makensis installer-simple.nsi
   ```
   生成: `ArtsBox-Setup-v1.0.0.exe` (约 10 MB)

3. **编译安装程序（完整版）**
   ```powershell
   makensis installer.nsi
   ```
   生成: `ArtsBox-Setup-v1.0.0.exe` (包含许可证页面)

4. **分发**
   直接分发生成的 `ArtsBox-Setup-v1.0.0.exe`
   用户双击即可安装

### NSIS 安装程序特性

- ✅ 专业的安装向导界面
- ✅ 许可协议页面
- ✅ 自定义安装目录
- ✅ 注册到"程序和功能"
- ✅ 标准的卸载程序
- ✅ 安装进度显示

## 方法三：使用 Wails 内置 NSIS（如果已安装）

如果系统已安装 NSIS，可以直接使用 Wails 命令：

```powershell
wails build -nsis
```

这将自动生成安装程序到 `build/bin/` 目录。

## 安装程序功能对比

| 功能 | PowerShell 脚本 | NSIS 简化版 | NSIS 完整版 | Wails -nsis |
|------|----------------|------------|------------|-------------|
| 无需额外工具 | ✅ | ❌ | ❌ | ❌ |
| 安装向导界面 | ❌ | ✅ | ✅ | ✅ |
| 许可协议页面 | ❌ | ❌ | ✅ | ✅ |
| 自定义安装路径 | ❌ | ✅ | ✅ | ✅ |
| 桌面快捷方式 | ✅ | ✅ | ✅ | ✅ |
| 开始菜单项 | ✅ | ✅ | ✅ | ✅ |
| 卸载功能 | ✅ | ✅ | ✅ | ✅ |
| 注册表集成 | ❌ | ✅ | ✅ | ✅ |
| 程序和功能显示 | ❌ | ✅ | ✅ | ✅ |
| 制作难度 | ⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐ |

## 推荐方案

### 快速测试/内部使用
→ **方法一**（PowerShell 脚本）
- 命令: `.\Install.ps1`
- 最简单，适合开发测试

### 个人分发/小范围使用
→ **绿色版**（仅 EXE）
- 直接分发 `build\bin\artsbox.exe`
- 无需安装，放哪都能用

### 正式发布/公开分发
→ **方法二**（NSIS 安装程序）
- 命令: `makensis installer-simple.nsi`
- 专业、标准、易用

## 文件清单

项目中包含以下打包相关文件：

```
artsbox/
├── Install.ps1              # PowerShell 安装脚本
├── installer-simple.nsi     # NSIS 简化安装脚本
├── installer.nsi            # NSIS 完整安装脚本
├── LICENSE.txt             # 许可证文件
├── wails.json              # Wails 配置（已配置 NSIS）
└── PACKAGING.md            # 本文档
```

## 版本号管理

修改版本号需要同步更新：
1. `wails.json` → `info.productVersion`
2. `installer-simple.nsi` → `OutFile` 行
3. `installer.nsi` → `!define APP_VERSION`
4. `Install.ps1` → 标题行

## 发布检查清单

打包前请确认：
- [ ] 已运行 `wails build` 构建最新版本
- [ ] 已测试应用程序功能正常
- [ ] 已更新版本号
- [ ] 已测试安装程序
- [ ] 已测试卸载程序
- [ ] 已准备 README 和使用说明

## 安装后文件位置

**使用安装程序安装后**:
- 应用程序: `C:\Program Files\ArtsBox\artsbox.exe`
- 配置文件: `%USERPROFILE%\.artsbox\config.json`
- 图片保存: `%USERPROFILE%\Pictures\ArtsBox` (默认)
- 桌面快捷方式: `桌面\ArtsBox AI 图片生成器.lnk`
- 开始菜单: `开始菜单\ArtsBox\`

**绿色版（直接运行）**:
- 应用程序: 任意位置
- 配置文件: `%USERPROFILE%\.artsbox\config.json`
- 图片保存: `%USERPROFILE%\Pictures\ArtsBox` (默认)

## 技术支持

如有问题，请检查：
1. WebView2 运行时是否已安装（Windows 11 自带）
2. 是否有网络连接（调用 API 需要）
3. API Key 是否配置正确
4. 保存目录是否有写入权限
