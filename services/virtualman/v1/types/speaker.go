package types

import "fmt"

type GetSpeakerListRes struct {
	CommonResponse
	/**
	 * 数据
	 */
	Data *GetSpeakerListData `json:"data,omitempty"`
}

type GetSpeakerListData struct {
	/**
	 * 数据结果集列表
	 */
	Results []*GetSpeakerListDefine `json:"results,omitempty"`
	/**
	 * 总条数
	 */
	Total int `json:"total,omitempty"`
}

type GetSpeakerListDefine struct {
	/**
	 * 音色id
	 */
	Id string `json:"id"`
	/**
	 * 音色名称
	 */
	Name string `json:"name"`
	/**
	 * 试听音频地址
	 */
	AudioUrl string `json:"audioUrl"`
	/**
	 * 封面图地址
	 */
	ThumbnailUrl string `json:"thumbnailUrl"`
}

func (define *GetSpeakerListDefine) String() string {
	return fmt.Sprintf("Id=%s,Name=%s,AudioUrl=%s,ThumbnailUrl=%s", define.Id, define.Name, define.AudioUrl, define.ThumbnailUrl)
}
