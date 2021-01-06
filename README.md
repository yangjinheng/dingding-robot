~~~go
func main() {
	// 实例化机器人
	robot := Robot{
		Url:    "https://oapi.dingtalk.com/robot/send",
		Token:  "20257dae9a0ae7227755261a4033938a662bbee36bd31935fd10547ba5a27bd4",
		Secret: "SEC560a4f53b9720d5113fe4c2a2076071b5f1f4f9acdb9625add2f984f0dd1e6e5",
	}

	// text 消息
	res, err := robot.SendMessage(Message{
		Msgtype: "text",
		Text: Text{
			Content: "text 测试",
		},
	})
	fmt.Println(res, err)
}
~~~

