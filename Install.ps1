# ArtsBox AI 图片生成器 - PowerShell 安装脚本
# 将此脚本和 artsbox.exe 放在同一目录运行

$ErrorActionPreference = "Stop"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  ArtsBox AI 图片生成器 v1.0.0" -ForegroundColor Cyan
Write-Host "  安装程序" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 检查管理员权限
$isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
if (-not $isAdmin) {
    Write-Host "需要管理员权限。正在请求提升权限..." -ForegroundColor Yellow
    Start-Process powershell.exe -ArgumentList "-NoProfile -ExecutionPolicy Bypass -File `"$PSCommandPath`"" -Verb RunAs
    exit
}

# 安装路径
$installDir = "$env:ProgramFiles\ArtsBox"
$exeName = "artsbox.exe"
$sourceExe = Join-Path $PSScriptRoot "build\bin\$exeName"

# 检查源文件
if (-not (Test-Path $sourceExe)) {
    Write-Host "错误: 找不到 $sourceExe" -ForegroundColor Red
    Write-Host "请确保已经运行 'wails build' 命令" -ForegroundColor Red
    Read-Host "按任意键退出"
    exit 1
}

Write-Host "安装位置: $installDir" -ForegroundColor Green
Write-Host ""

# 创建安装目录
if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir -Force | Out-Null
    Write-Host "✓ 创建安装目录" -ForegroundColor Green
}

# 复制文件
Copy-Item $sourceExe -Destination $installDir -Force
Write-Host "✓ 复制应用程序文件" -ForegroundColor Green

# 创建桌面快捷方式
$desktopPath = [Environment]::GetFolderPath("Desktop")
$shortcutPath = Join-Path $desktopPath "ArtsBox AI 图片生成器.lnk"
$WScriptShell = New-Object -ComObject WScript.Shell
$shortcut = $WScriptShell.CreateShortcut($shortcutPath)
$shortcut.TargetPath = Join-Path $installDir $exeName
$shortcut.WorkingDirectory = $installDir
$shortcut.Description = "ArtsBox AI 图片生成器"
$shortcut.Save()
Write-Host "✓ 创建桌面快捷方式" -ForegroundColor Green

# 创建开始菜单快捷方式
$startMenuPath = "$env:ProgramData\Microsoft\Windows\Start Menu\Programs\ArtsBox"
if (-not (Test-Path $startMenuPath)) {
    New-Item -ItemType Directory -Path $startMenuPath -Force | Out-Null
}
$startMenuShortcut = Join-Path $startMenuPath "ArtsBox AI 图片生成器.lnk"
$shortcut2 = $WScriptShell.CreateShortcut($startMenuShortcut)
$shortcut2.TargetPath = Join-Path $installDir $exeName
$shortcut2.WorkingDirectory = $installDir
$shortcut2.Description = "ArtsBox AI 图片生成器"
$shortcut2.Save()
Write-Host "✓ 创建开始菜单快捷方式" -ForegroundColor Green

# 创建卸载脚本
$uninstallScript = @"
# ArtsBox 卸载脚本
`$ErrorActionPreference = "Stop"

Write-Host "正在卸载 ArtsBox AI 图片生成器..." -ForegroundColor Yellow

# 删除安装目录
if (Test-Path "$installDir") {
    Remove-Item "$installDir" -Recurse -Force
    Write-Host "✓ 删除安装目录" -ForegroundColor Green
}

# 删除桌面快捷方式
`$desktopShortcut = Join-Path ([Environment]::GetFolderPath("Desktop")) "ArtsBox AI 图片生成器.lnk"
if (Test-Path `$desktopShortcut) {
    Remove-Item `$desktopShortcut -Force
    Write-Host "✓ 删除桌面快捷方式" -ForegroundColor Green
}

# 删除开始菜单快捷方式
if (Test-Path "$startMenuPath") {
    Remove-Item "$startMenuPath" -Recurse -Force
    Write-Host "✓ 删除开始菜单快捷方式" -ForegroundColor Green
}

Write-Host ""
Write-Host "卸载完成！" -ForegroundColor Green
Write-Host "注意: 用户配置文件未被删除，位于: `$env:USERPROFILE\.artsbox" -ForegroundColor Cyan
Read-Host "按任意键退出"
"@

$uninstallScriptPath = Join-Path $installDir "Uninstall.ps1"
$uninstallScript | Out-File -FilePath $uninstallScriptPath -Encoding UTF8
Write-Host "✓ 创建卸载脚本" -ForegroundColor Green

# 创建卸载快捷方式
$uninstallShortcut = Join-Path $startMenuPath "卸载 ArtsBox.lnk"
$shortcut3 = $WScriptShell.CreateShortcut($uninstallShortcut)
$shortcut3.TargetPath = "powershell.exe"
$shortcut3.Arguments = "-NoProfile -ExecutionPolicy Bypass -File `"$uninstallScriptPath`""
$shortcut3.WorkingDirectory = $installDir
$shortcut3.Description = "卸载 ArtsBox AI 图片生成器"
$shortcut3.Save()
Write-Host "✓ 创建卸载快捷方式" -ForegroundColor Green

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  安装完成！" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "安装位置: $installDir" -ForegroundColor White
Write-Host "配置文件: $env:USERPROFILE\.artsbox\config.json" -ForegroundColor White
Write-Host ""
Write-Host "您可以从以下位置启动应用:" -ForegroundColor White
Write-Host "  • 桌面快捷方式" -ForegroundColor White
Write-Host "  • 开始菜单 > ArtsBox" -ForegroundColor White
Write-Host ""
Write-Host "首次使用请配置豆包 API Key" -ForegroundColor Yellow
Write-Host ""

Read-Host "按任意键退出"
