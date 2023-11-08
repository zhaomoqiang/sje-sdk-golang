package v1

import (
	"fmt"
	"sje-openapi-for-golang/common"
	"sje-openapi-for-golang/services/virtualman/v1/types"
	"testing"
)

func getClient() *VirtualmanClient {
	info := &common.ServiceInfo{
		Scheme:      "https",
		Host:        "xxx",
		Credentials: &common.Credentials{AccessKeyId: "your ak", AccessKeySecret: "your sk"},
	}
	client, err := NewVirtualmanClient(info)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return client
}

func TestFigureClient_CreateTaskByAudio(t *testing.T) {
	define := &types.TextVirtualmanTaskDefine{
		VirtualmanId: "1214234234234",
		Text:         "这就是一段测试内容而已",
		SpeakerId:    "1111111111111111111",
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

func TestFigureClient_GetVirtualmanList(t *testing.T) {
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

func TestFigureClient_GetSpeakerList(t *testing.T) {
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
