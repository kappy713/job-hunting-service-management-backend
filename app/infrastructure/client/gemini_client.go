package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"

	"job-hunting-service-management-backend/app/internal/entity"
)

// Gemini APIクライアントの構造体
type GeminiClient struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

// Gemini APIのリクエスト構造体
type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

// Gemini APIのレスポンス構造体
type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content ContentResponse `json:"content"`
}

type ContentResponse struct {
	Parts []PartResponse `json:"parts"`
}

type PartResponse struct {
	Text string `json:"text"`
}

// 新しいGeminiクライアントを作成
func NewGeminiClient() *GeminiClient {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY environment variable is required")
	}

	return &GeminiClient{
		APIKey:  apiKey,
		BaseURL: "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent",
		Client: &http.Client{
			Timeout: 120 * time.Second, // 2分に延長
		},
	}
}

// リトライ機能付きHTTPリクエスト送信
func (g *GeminiClient) sendRequestWithRetry(ctx context.Context, req *http.Request, maxRetries int) (*http.Response, error) {
	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			// リトライ前に少し待機（指数バックオフ）
			waitTime := time.Duration(attempt) * 10 * time.Second
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(waitTime):
			}
		}

		resp, err := g.Client.Do(req)
		if err == nil {
			return resp, nil
		}

		lastErr = err
		// タイムアウトエラーの場合はリトライ
		if strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "deadline exceeded") {
			continue
		}
		// その他のエラーの場合は即座に終了
		return nil, err
	}

	return nil, fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}

// プロンプトテンプレートを処理してユーザー情報を埋め込む
func (g *GeminiClient) ProcessPromptTemplate(serviceName string, user *entity.User) (string, error) {
	promptTemplate, exists := entity.ServicePrompts[serviceName]
	if !exists {
		return "", fmt.Errorf("prompt template not found for service: %s", serviceName)
	}

	tmpl, err := template.New("prompt").Parse(promptTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, user); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// Gemini APIにリクエストを送信
func (g *GeminiClient) GenerateContent(ctx context.Context, prompt string) (string, error) {
	request := GeminiRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: prompt,
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s?key=%s", g.BaseURL, g.APIKey)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// リトライ機能付きでリクエスト送信（最大2回リトライ）
	resp, err := g.sendRequestWithRetry(ctx, req, 2)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var geminiResp GeminiResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content generated")
	}

	return geminiResp.Candidates[0].Content.Parts[0].Text, nil
}

// JSONコンテンツを抽出する関数
func extractJSONFromMarkdown(content string) (string, error) {
	// マークダウンコードブロック（```json と ``` で囲まれた部分）を探す（改行対応）
	jsonPattern := regexp.MustCompile("(?s)```(?:json)?\\s*\\n?(.*?)\\n?\\s*```")
	matches := jsonPattern.FindStringSubmatch(content)

	if len(matches) > 1 {
		return strings.TrimSpace(matches[1]), nil
	}

	// マークダウンが見つからない場合は、{ } で囲まれた部分を探す（改行対応）
	bracePattern := regexp.MustCompile(`(?s)\{.*\}`)
	braceMatch := bracePattern.FindString(content)

	if braceMatch != "" {
		return strings.TrimSpace(braceMatch), nil
	}

	// どちらも見つからない場合は、元のコンテンツをそのまま返す
	return strings.TrimSpace(content), nil
}

// サービス用のコンテンツを生成
func (g *GeminiClient) GenerateServiceContent(ctx context.Context, serviceName string, user *entity.User) (map[string]interface{}, error) {
	prompt, err := g.ProcessPromptTemplate(serviceName, user)
	if err != nil {
		return nil, fmt.Errorf("failed to process prompt template: %w", err)
	}

	content, err := g.GenerateContent(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	// マークダウン形式からJSONを抽出
	jsonContent, err := extractJSONFromMarkdown(content)
	if err != nil {
		return nil, fmt.Errorf("failed to extract JSON from response: %w", err)
	}

	// JSONレスポンスをパース
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonContent), &result); err != nil {
		return nil, fmt.Errorf("failed to parse generated JSON content: %w\nRaw content: %s\nExtracted JSON: %s", err, content, jsonContent)
	}

	return result, nil
}
