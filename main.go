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
	// webHook在群设置-智能群助手-添加机器人-自定义
	webHook := ""
	client := &http.Client{}
	data.Msgtype = "markdown"
	data.Markdown = make(map[string]string)
	// 这里填入要发送的标题
	data.Markdown["title"] = "柚子"
	// 这里填入要发送的内容
	data.Markdown["text"] = "小萝莉"
	r, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", webHook, strings.NewReader(string(r)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Charset", "UTF-8")
	resp, _ := client.Do(req)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
