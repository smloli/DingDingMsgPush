package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strings"
)

type Data struct {
	Msgtype string `json:"msgtype"`
	Markdown map[string]string `json:"markdown"`
}

func main() {
	var data Data
	webHook := "https://api.dingtalk.com/robot/send?access_token=1670b5c7b7b673f537d1f392070c702f466f4243fc518ae354d9cc25f8f2606d"
	client := &http.Client{}
	data.Msgtype = "markdown"
	data.Markdown = make(map[string]string, 1)
	// 这里填入要发送的标题
	data.Markdown["title"] = "柚子"
	// 这里填入要发送的内容
	data.Markdown["text"] = "小萝莉"
	r, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", webHook, strings.NewReader(string(r)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Charset", "UTF-8")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}