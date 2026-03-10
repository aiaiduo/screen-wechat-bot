package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// 解析命令行参数
	var (
		url     string
		element string
		width   string
		height  string
		bot     string
	)

	flag.StringVar(&url, "url", "https://baidu.com", "给我一个你想要截图的URL")
	flag.StringVar(&url, "u", "https://baidu.com", "给我一个你想要截图的URL")
	flag.StringVar(&element, "element", "#s_lg_img", "给我你关心的页面元素")
	flag.StringVar(&element, "e", "#s_lg_img", "给我你关心的页面元素")
	flag.StringVar(&width, "kuan", "1200", "页面宽度")
	flag.StringVar(&width, "k", "1200", "页面宽度")
	flag.StringVar(&height, "gao", "800", "页面高度")
	flag.StringVar(&height, "g", "800", "页面高度")
	flag.StringVar(&bot, "bot", "d63e3f22-3a88-43fb-a2ad-ad78ba5b43b5", "机器人地址")
	flag.StringVar(&bot, "b", "d63e3f22-3a88-43fb-a2ad-ad78ba5b43b5", "机器人地址")
	flag.Parse()

	// 启动浏览器
	path, _ := launcher.LookPath()
	browser := rod.New().ControlURL(launcher.New().Bin(path).MustLaunch()).MustConnect()
	defer browser.MustClose()

	// 设置浏览器窗口大小
	page := browser.MustPage(url)
	page.MustSetViewport(1200, 800, 1.0, false)

	// 等待页面加载完成
	page.MustWaitLoad()

	// 等待2秒确保页面完全加载
	time.Sleep(2 * time.Second)

	// 截图
	var img []byte
	// 截取整个页面
	img = page.MustScreenshot()

	// 发送到企业微信机器人
	sendToWeChatBot(bot, img)
}

func sendToWeChatBot(bot string, img []byte) {
	// 企业微信机器人API地址
	apiURL := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", bot)
	fmt.Printf("正在发送截图到企业微信机器人，API地址: %s\n", apiURL)

	// 构建请求体
	data := map[string]interface{}{
		"msgtype": "image",
		"image": map[string]string{
			"base64":  base64.StdEncoding.EncodeToString(img),
			"md5":     fmt.Sprintf("%x", md5.Sum(img)),
		},
	}

	// 转换为JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("JSON编码失败: %v\n", err)
		return
	}

	// 创建HTTP客户端，设置超时时间
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// 发送请求
	fmt.Println("正在发送请求...")
	resp, err := client.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("发送失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 打印响应
	fmt.Printf("发送成功，响应: %s\n", body)
}
