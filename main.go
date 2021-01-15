package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Robot 机器人
type Robot struct {
	Url    string
	Token  string
	Secret string
}

// sign 签名
func (robot *Robot) sign(request *http.Request) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, robot.Secret)
	hmac256 := hmac.New(sha256.New, []byte(robot.Secret))
	hmac256.Write([]byte(stringToSign))
	query := url.Values{}
	query.Set("access_token", robot.Token)
	query.Set("timestamp", timestamp)
	query.Set("sign", url.QueryEscape(base64.StdEncoding.EncodeToString(hmac256.Sum(nil))))
	request.URL.RawQuery = query.Encode()
}

// SendMessage 发送信息
func (robot *Robot) SendMessage(msg Message) (*Result, error) {
	message, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", robot.Url, &net.Buffers{message})
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json; charset=utf8")
	robot.sign(request)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := Result{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	// example: https://oapi.dingtalk.com/robot/send?access_token=xxxx&timestamp=xxxx&sign=xxxx

	// 实例化机器人
	robot := Robot{
		Url:    "https://oapi.dingtalk.com/robot/send",
		Token:  "f2179c1b79d47d3e0e6734ad7fd337f478b2d6efa1e82392ca1739302b6fae96",
		Secret: "SECd334bd7abf6d241693cf206530cb01c639153e61af12a2b0047ed5d0f67395b1",
	}

	// text 消息
	res, err := robot.SendMessage(Message{
		Msgtype: "text",
		Text: Text{
			Content: "你好，测试",
		},
	})
	fmt.Println(res, err)

	// markdown 消息
	res, err = robot.SendMessage(Message{
		Msgtype: "markdown",
		Markdown: Markdown{
			Title: "markdown 测试",
			Text:  "## 今日天气\n天气晴，有大风，有大雷\n",
		},
	})
	fmt.Println(res, err)

	// link 消息
	res, err = robot.SendMessage(Message{
		Msgtype: "link",
		Link: Link{
			Title:      "link 测试",
			Text:       "水水水水还是手术室护士",
			MessageURL: "https://google.com",
		},
	})
	fmt.Println(res, err)
}
