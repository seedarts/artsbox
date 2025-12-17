package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Client OpenAI API 客户端
type Client struct {
	APIKey  string
	BaseURL string
	client  *http.Client
}

// NewClient 创建新的 API 客户端
func NewClient(apiKey, baseURL string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: baseURL,
		client: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

// TextToImageRequest 文生图请求参数
type TextToImageRequest struct {
	Model          string `json:"model"`
	Prompt         string `json:"prompt"`
	Size           string `json:"size,omitempty"`
	Stream         bool   `json:"stream"`
	ResponseFormat string `json:"response_format"`
	Watermark      bool   `json:"watermark"`
}

// BatchGenerationOptions 组图生成选项
type BatchGenerationOptions struct {
	MaxImages int `json:"max_images"`
}

// ImageToImageRequest 图生图/多图融合/组图请求参数
type ImageToImageRequest struct {
	Model                            string                  `json:"model"`
	Prompt                           string                  `json:"prompt"`
	Image                            interface{}             `json:"image,omitempty"` // 字符串(单图)或数组(多图)
	Size                             string                  `json:"size,omitempty"`
	SequentialImageGeneration        string                  `json:"sequential_image_generation,omitempty"`
	SequentialImageGenerationOptions *BatchGenerationOptions `json:"sequential_image_generation_options,omitempty"`
	Stream                           bool                    `json:"stream"`
	ResponseFormat                   string                  `json:"response_format"`
	Watermark                        bool                    `json:"watermark"`
}

// ImageResponse API 响应
type ImageResponse struct {
	Created int64 `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

// ErrorResponse API 错误响应
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    string `json:"code"`
	} `json:"error"`
}

// GenerateFromText 文生图
func (c *Client) GenerateFromText(req TextToImageRequest) (*ImageResponse, error) {
	url := c.BaseURL + "/images/generations"

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("API error: %s", string(body))
		}
		return nil, fmt.Errorf("API error: %s", errResp.Error.Message)
	}

	var imgResp ImageResponse
	if err := json.Unmarshal(body, &imgResp); err != nil {
		return nil, err
	}

	return &imgResp, nil
}

// GenerateFromImage 图生图（使用JSON格式）
func (c *Client) GenerateFromImage(req ImageToImageRequest) (*ImageResponse, error) {
	url := c.BaseURL + "/images/generations"

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("API error: %s", string(body))
		}
		return nil, fmt.Errorf("API error: %s", errResp.Error.Message)
	}

	var imgResp ImageResponse
	if err := json.Unmarshal(body, &imgResp); err != nil {
		return nil, err
	}

	return &imgResp, nil
}

// DownloadImage 下载图片
func DownloadImage(url, savePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: status %d", resp.StatusCode)
	}

	// 确保目录存在
	dir := filepath.Dir(savePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
