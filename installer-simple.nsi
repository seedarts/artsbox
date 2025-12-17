; ArtsBox AI 图片生成器 - 简化安装脚本
; 不需要外部依赖，直接使用 NSIS 基础功能

Name "ArtsBox AI 图片生成器"
OutFile "ArtsBox-Setup-v1.0.0.exe"
InstallDir "$PROGRAMFILES64\ArtsBox"
RequestExecutionLevel admin

; 页面
Page directory
Page instfiles

; 安装
Section ""
    SetOutPath "$INSTDIR"
    File "build\bin\artsbox.exe"
    
    ; 创建开始菜单快捷方式
    CreateDirectory "$SMPROGRAMS\ArtsBox"
    CreateShortcut "$SMPROGRAMS\ArtsBox\ArtsBox AI 图片生成器.lnk" "$INSTDIR\artsbox.exe"
    CreateShortcut "$SMPROGRAMS\ArtsBox\卸载.lnk" "$INSTDIR\Uninstall.exe"
    
    ; 创建桌面快捷方式
    CreateShortcut "$DESKTOP\ArtsBox AI 图片生成器.lnk" "$INSTDIR\artsbox.exe"
    
    ; 创建卸载程序
    WriteUninstaller "$INSTDIR\Uninstall.exe"
    
    ; 注册表信息
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\ArtsBox" "DisplayName" "ArtsBox AI 图片生成器"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\ArtsBox" "UninstallString" "$INSTDIR\Uninstall.exe"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\ArtsBox" "DisplayIcon" "$INSTDIR\artsbox.exe"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\ArtsBox" "Publisher" "ArtsBox"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\ArtsBox" "DisplayVersion" "1.0.0"
SectionEnd

; 卸载
Section "Uninstall"
    Delete "$INSTDIR\artsbox.exe"
    Delete "$INSTDIR\Uninstall.exe"
    Delete "$DESKTOP\ArtsBox AI 图片生成器.lnk"
    Delete "$SMPROGRAMS\ArtsBox\ArtsBox AI 图片生成器.lnk"
    Delete "$SMPROGRAMS\ArtsBox\卸载.lnk"
    RMDir "$SMPROGRAMS\ArtsBox"
    RMDir "$INSTDIR"
    DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\ArtsBox"
SectionEnd
