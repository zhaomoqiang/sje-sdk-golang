package types

import "fmt"

type GetVirtualmanListRes struct {
	CommonResponse
	/**
	 * 数据
	 */
	Data *GetVirtualmanListData `json:"data,omitempty"`
}

type GetVirtualmanListData struct {
	/**
	 * 数据结果集列表
	 */
	Results []*GetVirtualmanListDefine `json:"results,omitempty"`
	/**
	 * 总条数
	 */
	Total int `json:"total,omitempty"`
}

type GetVirtualmanListDefine struct {
	/**
	 * 数字人id
	 */
	Id string `json:"id"`
	/**
	 * 数字人名称
	 */
	Name string `json:"name"`
	/**
	 * 封面图地址
	 */
	ThumbnailUrl string `json:"thumbnailUrl"`
}

type AudioVirtualmanTaskDefine struct {
	/**
	 * 数字人id
	 */
	VirtualmanId string `json:"virtualmanId"`
	/**
	 * 音频url
	 */
	AudioUrl string `json:"audioUrl"`
	/**
	 * 结果回调地址
	 */
	CallbackUrl string `json:"callbackUrl"`
}

type TextVirtualmanTaskDefine struct {
	/**
	 * 数字人id
	 */
	VirtualmanId string `json:"virtualmanId"`
	/**
	 * 文本
	 */
	Text string `json:"text"`
	/**
	 * 音色id
	 */
	SpeakerId string `json:"speakerId"`
	/**
	 * 音量 由低到高依次为：0.5、1、1.5、2
	 */
	Volume float32 `json:"volume"`
	/**
	 * 语速 由低到高依次为：-2、-1、0、1、2
	 */
	Speed float32 `json:"speed"`
	/**
	 * 结果回调地址
	 */
	CallbackUrl string `json:"callbackUrl"`
}

func (define *GetVirtualmanListDefine) String() string {
	return fmt.Sprintf("Id=%s,Name=%s,ThumbnailUrl=%s", define.Id, define.Name, define.ThumbnailUrl)
}

func (define *AudioVirtualmanTaskDefine) String() string {
	return fmt.Sprintf("VirtualmanId=%s,AudioUrl=%s,CallbackUrl=%s", define.VirtualmanId, define.AudioUrl, define.CallbackUrl)
}

func (define *TextVirtualmanTaskDefine) String() string {
	return fmt.Sprintf("VirtualmanId=%s,Text=%s,SpeakerId=%s,Volume=%f,Speed=%f,CallbackUrl=%s", define.VirtualmanId, define.Text, define.SpeakerId, define.Volume, define.Speed, define.CallbackUrl)
}
