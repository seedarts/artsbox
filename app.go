package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"artsbox/pkg/config"
	"artsbox/pkg/openai"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	settings *config.Settings
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// 加载配置
	settings, err := config.Load()
	if err != nil {
		runtime.LogError(ctx, "Failed to load config: "+err.Error())
		settings = config.DefaultSettings()
	}
	a.settings = settings
}

// GenerateResult 生成结果
type GenerateResult struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Images  []string `json:"images"`
}

// GetSettings 获取配置
func (a *App) GetSettings() *config.Settings {
	return a.settings
}

// SaveSettings 保存配置
func (a *App) SaveSettings(settings config.Settings) error {
	if err := config.Save(&settings); err != nil {
		return err
	}
	a.settings = &settings
	return nil
}

// GenerateImage 生成图片（文生图）
func (a *App) GenerateImage(prompt string, size string, n int) GenerateResult {
	if a.settings.APIKey == "" {
		return GenerateResult{
			Success: false,
			Message: "请先配置 API Key",
		}
	}

	if prompt == "" {
		return GenerateResult{
			Success: false,
			Message: "请输入提示词",
		}
	}

	client := openai.NewClient(a.settings.APIKey, a.settings.BaseURL)

	req := openai.TextToImageRequest{
		Model:          a.settings.Model,
		Prompt:         prompt,
		Size:           size,
		Stream:         false,
		ResponseFormat: "url",
		Watermark:      false,
	}

	resp, err := client.GenerateFromText(req)
	if err != nil {
		return GenerateResult{
			Success: false,
			Message: "生成失败: " + err.Error(),
		}
	}

	var imageURLs []string
	for _, img := range resp.Data {
		imageURLs = append(imageURLs, img.URL)
	}

	return GenerateResult{
		Success: true,
		Message: "生成成功",
		Images:  imageURLs,
	}
}

// GenerateBatchImages 文生组图
func (a *App) GenerateBatchImages(prompt string, size string, maxImages int) GenerateResult {
	if a.settings.APIKey == "" {
		return GenerateResult{
			Success: false,
			Message: "请先配置 API Key",
		}
	}

	if prompt == "" {
		return GenerateResult{
			Success: false,
			Message: "请输入提示词",
		}
	}

	if maxImages < 1 || maxImages > 4 {
		return GenerateResult{
			Success: false,
			Message: "生成数量必须在 1-4 范围内",
		}
	}

	client := openai.NewClient(a.settings.APIKey, a.settings.BaseURL)

	req := openai.ImageToImageRequest{
		Model:                     a.settings.Model,
		Prompt:                    prompt,
		Size:                      size,
		SequentialImageGeneration: "auto",
		SequentialImageGenerationOptions: &openai.BatchGenerationOptions{
			MaxImages: maxImages,
		},
		Stream:         false,
		ResponseFormat: "url",
		Watermark:      false,
	}

	resp, err := client.GenerateFromImage(req)
	if err != nil {
		return GenerateResult{
			Success: false,
			Message: "生成失败: " + err.Error(),
		}
	}

	var imageURLs []string
	for _, img := range resp.Data {
		imageURLs = append(imageURLs, img.URL)
	}

	return GenerateResult{
		Success: true,
		Message: "生成成功",
		Images:  imageURLs,
	}
}

// GenerateFromImage 图文生图（从 base64 数据）
func (a *App) GenerateFromImage(prompt string, imageBase64 string) GenerateResult {
	if a.settings.APIKey == "" {
		return GenerateResult{
			Success: false,
			Message: "请先配置 API Key",
		}
	}

	if imageBase64 == "" {
		return GenerateResult{
			Success: false,
			Message: "未提供图片数据",
		}
	}

	client := openai.NewClient(a.settings.APIKey, a.settings.BaseURL)

	// 构造图生图请求
	req := openai.ImageToImageRequest{
		Model:          a.settings.Model,
		Prompt:         prompt,
		Image:          imageBase64, // 单图输入，字符串
		Size:           "2K",
		Stream:         false,
		ResponseFormat: "url",
		Watermark:      false,
	}

	resp, err := client.GenerateFromImage(req)
	if err != nil {
		return GenerateResult{
			Success: false,
			Message: "生成失败: " + err.Error(),
		}
	}

	var imageURLs []string
	for _, img := range resp.Data {
		imageURLs = append(imageURLs, img.URL)
	}

	return GenerateResult{
		Success: true,
		Message: "生成成功",
		Images:  imageURLs,
	}
}

// GenerateFromMultiImages 多图融合（从 base64 数据数组）
func (a *App) GenerateFromMultiImages(prompt string, imageBase64List []string) GenerateResult {
	if a.settings.APIKey == "" {
		return GenerateResult{
			Success: false,
			Message: "请先配置 API Key",
		}
	}

	if len(imageBase64List) == 0 {
		return GenerateResult{
			Success: false,
			Message: "未提供图片数据",
		}
	}

	if len(imageBase64List) > 10 {
		return GenerateResult{
			Success: false,
			Message: "最多支持 10 张图片",
		}
	}

	client := openai.NewClient(a.settings.APIKey, a.settings.BaseURL)

	// 构造多图融合请求
	req := openai.ImageToImageRequest{
		Model:          a.settings.Model,
		Prompt:         prompt,
		Image:          imageBase64List, // 多图输入，数组
		Size:           "2K",
		Stream:         false,
		ResponseFormat: "url",
		Watermark:      false,
	}

	resp, err := client.GenerateFromImage(req)
	if err != nil {
		return GenerateResult{
			Success: false,
			Message: "生成失败: " + err.Error(),
		}
	}

	var imageURLs []string
	for _, img := range resp.Data {
		imageURLs = append(imageURLs, img.URL)
	}

	return GenerateResult{
		Success: true,
		Message: "生成成功",
		Images:  imageURLs,
	}
}

// GenerateBatchFromImage 单图生组图
func (a *App) GenerateBatchFromImage(prompt string, imageBase64 string, maxImages int) GenerateResult {
	if a.settings.APIKey == "" {
		return GenerateResult{
			Success: false,
			Message: "请先配置 API Key",
		}
	}

	if imageBase64 == "" {
		return GenerateResult{
			Success: false,
			Message: "未提供图片数据",
		}
	}

	if maxImages < 1 || maxImages > 4 {
		return GenerateResult{
			Success: false,
			Message: "生成数量必须在 1-4 范围内",
		}
	}

	client := openai.NewClient(a.settings.APIKey, a.settings.BaseURL)

	req := openai.ImageToImageRequest{
		Model:                     a.settings.Model,
		Prompt:                    prompt,
		Image:                     imageBase64,
		Size:                      "2K",
		SequentialImageGeneration: "auto",
		SequentialImageGenerationOptions: &openai.BatchGenerationOptions{
			MaxImages: maxImages,
		},
		Stream:         false,
		ResponseFormat: "url",
		Watermark:      false,
	}

	resp, err := client.GenerateFromImage(req)
	if err != nil {
		return GenerateResult{
			Success: false,
			Message: "生成失败: " + err.Error(),
		}
	}

	var imageURLs []string
	for _, img := range resp.Data {
		imageURLs = append(imageURLs, img.URL)
	}

	return GenerateResult{
		Success: true,
		Message: "生成成功",
		Images:  imageURLs,
	}
}

// GenerateBatchFromMultiImages 多图生组图
func (a *App) GenerateBatchFromMultiImages(prompt string, imageBase64List []string, maxImages int) GenerateResult {
	if a.settings.APIKey == "" {
		return GenerateResult{
			Success: false,
			Message: "请先配置 API Key",
		}
	}

	if len(imageBase64List) == 0 {
		return GenerateResult{
			Success: false,
			Message: "未提供图片数据",
		}
	}

	if len(imageBase64List) > 10 {
		return GenerateResult{
			Success: false,
			Message: "最多支持 10 张图片",
		}
	}

	if maxImages < 1 || maxImages > 4 {
		return GenerateResult{
			Success: false,
			Message: "生成数量必须在 1-4 范围内",
		}
	}

	client := openai.NewClient(a.settings.APIKey, a.settings.BaseURL)

	req := openai.ImageToImageRequest{
		Model:                     a.settings.Model,
		Prompt:                    prompt,
		Image:                     imageBase64List,
		Size:                      "2K",
		SequentialImageGeneration: "auto",
		SequentialImageGenerationOptions: &openai.BatchGenerationOptions{
			MaxImages: maxImages,
		},
		Stream:         false,
		ResponseFormat: "url",
		Watermark:      false,
	}

	resp, err := client.GenerateFromImage(req)
	if err != nil {
		return GenerateResult{
			Success: false,
			Message: "生成失败: " + err.Error(),
		}
	}

	var imageURLs []string
	for _, img := range resp.Data {
		imageURLs = append(imageURLs, img.URL)
	}

	return GenerateResult{
		Success: true,
		Message: "生成成功",
		Images:  imageURLs,
	}
}

// SaveFile 保存图片到本地
func (a *App) SaveFile(url string, filename string) string {
	if filename == "" {
		filename = fmt.Sprintf("image_%d.png", time.Now().Unix())
	}

	savePath := filepath.Join(a.settings.SaveDir, filename)

	if err := openai.DownloadImage(url, savePath); err != nil {
		return "保存失败: " + err.Error()
	}

	return savePath
}

// SelectSaveDirectory 选择保存目录
func (a *App) SelectSaveDirectory() string {
	dirPath, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "选择保存目录",
		DefaultDirectory: a.settings.SaveDir,
	})

	if err != nil || dirPath == "" {
		return a.settings.SaveDir
	}

	return dirPath
}

// GetSavedImages 获取保存目录中的所有图片
func (a *App) GetSavedImages() []string {
	if a.settings.SaveDir == "" {
		return []string{}
	}

	var images []string
	files, err := os.ReadDir(a.settings.SaveDir)
	if err != nil {
		runtime.LogError(a.ctx, "Failed to read directory: "+err.Error())
		return []string{}
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		ext := strings.ToLower(filepath.Ext(filename)) // 转为小写比较
		// 支持的图片格式: png, jpg, jpeg, jpe, webp
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".jpe" || ext == ".webp" {
			// 返回文件的绝对路径
			images = append(images, filepath.Join(a.settings.SaveDir, filename))
		}
	}

	// 按修改时间倒序排序（最新的在前面）
	type ImageInfo struct {
		Path    string
		ModTime time.Time
	}

	imageInfos := make([]ImageInfo, 0, len(images))
	for _, img := range images {
		info, err := os.Stat(img)
		if err == nil {
			imageInfos = append(imageInfos, ImageInfo{
				Path:    img,
				ModTime: info.ModTime(),
			})
		}
	}

	// 排序：最新的在前
	for i := 0; i < len(imageInfos)-1; i++ {
		for j := i + 1; j < len(imageInfos); j++ {
			if imageInfos[i].ModTime.Before(imageInfos[j].ModTime) {
				imageInfos[i], imageInfos[j] = imageInfos[j], imageInfos[i]
			}
		}
	}

	result := make([]string, len(imageInfos))
	for i, info := range imageInfos {
		result[i] = info.Path
	}

	runtime.LogInfo(a.ctx, fmt.Sprintf("Loaded %d images from %s", len(result), a.settings.SaveDir))
	return result
}

// OpenImageFile 打开图片文件
func (a *App) OpenImageFile(imagePath string) error {
	// 使用系统默认程序打开图片
	runtime.BrowserOpenURL(a.ctx, "file:///"+filepath.ToSlash(imagePath))
	return nil
}

// GetImageAsDataURL 获取图片的 base64 Data URL
func (a *App) GetImageAsDataURL(imagePath string) string {
	data, err := os.ReadFile(imagePath)
	if err != nil {
		runtime.LogError(a.ctx, "Failed to read image: "+err.Error())
		return ""
	}

	// 根据文件扩展名确定 MIME 类型
	ext := strings.ToLower(filepath.Ext(imagePath))
	mimeType := "image/jpeg" // 默认
	switch ext {
	case ".png":
		mimeType = "image/png"
	case ".jpg", ".jpeg", ".jpe":
		mimeType = "image/jpeg"
	case ".webp":
		mimeType = "image/webp"
	}

	// 转换为 base64
	encoded := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", mimeType, encoded)
}
