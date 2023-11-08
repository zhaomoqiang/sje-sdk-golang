package types

type CommonResponse struct {
	/**
	 * 提示码
	 */
	Code string `json:"code"`
	/**
	 * 提示消息
	 */
	Message string `json:"message,omitempty"`

	TraceId string `json:"traceId,omitempty"`
}

type CreateTaskRes struct {
	CommonResponse
	/**
	 * 数据
	 */
	Data *TaskInfo `json:"data,omitempty"`
}

type TaskInfo struct {
	TaskId string `json:"taskId,omitempty"`
}

type PageDefine struct {
	/**
	 * 页条数
	 */
	PageSize int8 `json:"pageSize" query:"pageSize"`
	/**
	 * 页码
	 */
	Page int8 `json:"page" query:"page"`
}
