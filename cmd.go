package main

import (
	"fmt"
	"sje-openapi-for-golang/common"
	v1 "sje-openapi-for-golang/services/virtualman/v1"
	"sje-openapi-for-golang/services/virtualman/v1/types"
	"sync"
)

func main() {
	k := 0
	for ; k < 1; k++ {
		waitGroup := sync.WaitGroup{}
		waitGroup.Add(1)
		i := 0
		for ; i < 1; i++ {
			go func() {
				//GetSpeakerList()
				GetVirtualmanList()
				//CreateTaskByText()
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()
	}
}

func getClient() *v1.VirtualmanClient {
	info := &common.ServiceInfo{
		Scheme:      "https",
		Host:        "sje-openapi-proxy.test.bhbapp.cn",
		Credentials: &common.Credentials{AccessKeyId: "MBoSiWJ2EHWYBkfU", AccessKeySecret: "89o7Yjd8S4nn2N2B6kQJpwPI"},
	}
	client, err := v1.NewVirtualmanClient(info)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return client
}

func CreateTaskByText() {
	define := &types.TextVirtualmanTaskDefine{
		FigureId:  "6535dec449baf5003096648c",
		Text:      "这就是一段测试内容而已，在实际运用中应该替换为真实的业务文本内容",
		SpeakerId: "652def4ea2767b00302c30e2",
	}
	client := getClient()
	if client != nil {
		result, err := client.CreateTaskByText(define)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(result.Code)
		fmt.Println(result.Message)
		fmt.Println(result.TraceId)
		fmt.Println(result.Data)
	}
}

func GetSpeakerList() {
	define := types.PageDefine{Page: 1, PageSize: 10}
	client := getClient()
	if client != nil {
		result, err := client.GetSpeakerList(&define)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(result.Code)
		fmt.Println(result.Message)
		fmt.Println(result.TraceId)
		fmt.Println(result.Data.Total)
		for _, v := range result.Data.Results {
			fmt.Println(v.String())
		}
	}
}

func GetVirtualmanList() {
	define := types.PageDefine{Page: 1, PageSize: 10}
	client := getClient()
	if client != nil {
		result, err := client.GetVirtualmanList(&define)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(result.Code)
		fmt.Println(result.Message)
		fmt.Println(result.TraceId)
		fmt.Println(result.Data.Total)
		for _, v := range result.Data.Results {
			fmt.Println(v.String())
		}
	}
}
