package main

import (
	"encoding/json"
	"fmt"
	"plugin"
	// "time"
)

type config struct {
	Timeout         int      `json:"timeout"`
	Aider           string   `json:"aider"`
	HTTPProxy       string   `json:"http_proxy"`
	PassList        []string `json:"pass_list"`
	ExtraPluginPath string   `json:"extra_plugin_path"`
}

type Meta struct {
	System   string   `json:"system"`
	PathList []string `json:"pathlist"`
	FileList []string `json:"filelist"`
	PassList []string `json:"passlist"`
}

type Task struct {
	Type   string `json:"type"`
	Netloc string `json:"netloc"`
	Target string `json:"target"`
	Meta   Meta   `json:"meta"`
}

type Greeter interface {
	Check(taskJSON string) []map[string]interface{}
	GetPlugins() []map[string]interface{}
	SetConfig(configJSON string)
	ShowLog()
}

func main() {
	// 加载go plugin
	plug, err := plugin.Open("./kunpeng_go.so")
	if err != nil {
		fmt.Println(err)
		return
	}
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		return
	}
	kunpeng, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		return
	}
	// 开启日志打印
	kunpeng.ShowLog()

	// 获取插件信息
	pocs := kunpeng.GetPlugins()
	fmt.Println(len(pocs))

	// 修改配置
	c := &config{
		Timeout: 15,
		// Aider:     "",
		// HTTPProxy: "",
		// PassList:  []string{"ptest"},
		ExtraPluginPath: "/home/sachinly/goproject/src/kunpeng/plugin/",
	}
	configJSONBytes, _ := json.Marshal(c)
	kunpeng.SetConfig(string(configJSONBytes))

	// 扫描目标
	task := Task{
		Type:   "web",
		Netloc: "http://192.168.7.127:9080",
		Target: "struts2",
		Meta: Meta{
			System:   "",
			PathList: []string{},
			FileList: []string{},
			PassList: []string{""},
		},
	}
	task2 := Task{
		Type:   "web",
		Netloc: "http://192.168.7.127:8808",
		Target: "struts2",
		Meta: Meta{
			System:   "",
			PathList: []string{},
			FileList: []string{"http://192.168.7.127:8808/memoindex.action","http://192.168.7.127:8808/index.action"},
			PassList: []string{},
		},
	}
	jsonBytes, _ := json.Marshal(task)
	result := kunpeng.Check(string(jsonBytes))
	fmt.Println(result)
	// time.Sleep(time.Second * 21)
	jsonBytes, _ = json.Marshal(task2)
	result = kunpeng.Check(string(jsonBytes))
	fmt.Println(result)
}
