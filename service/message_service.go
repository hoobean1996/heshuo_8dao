package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TextMessage 定义接收的XML消息结构
type TextMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int64    `xml:"MsgId"`
	MsgDataId    string   `xml:"MsgDataId"`
	Idx          string   `xml:"Idx"`
}

type TextResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`   // 接收方帐号(收到的OpenID)
	FromUserName string   `xml:"FromUserName"` // 开发者微信号
	CreateTime   int64    `xml:"CreateTime"`   // 消息创建时间
	MsgType      string   `xml:"MsgType"`      // 消息类型
	Content      string   `xml:"Content"`      // 回复的消息内容
}

// HandleTextMessage 处理接收到的文本消息
func HandleTextMessage(c *gin.Context) {
	// 读取请求体
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// 解析XML消息
	var message TextMessage
	if err := xml.Unmarshal(body, &message); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Failed to parse XML message",
		})
		return
	}

	// 验证消息类型
	if message.MsgType != "text" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid message type",
		})
		return
	}

	response := TextResponse{
		ToUserName:   message.FromUserName,         // 原发送者作为接收者
		FromUserName: message.ToUserName,           // 原接收者作为发送者
		CreateTime:   time.Now().Unix(),            // 当前时间戳
		MsgType:      "text",                       // 消息类型为文本
		Content:      "收到你的消息: " + message.Content, // 回复的内容
	}

	fmt.Println("request", message)
	fmt.Println("response", response)

	c.XML(http.StatusOK, response)
}
