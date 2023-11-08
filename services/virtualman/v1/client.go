package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sje-openapi-for-golang/common"
	"sje-openapi-for-golang/services/virtualman/v1/types"
)

const (
	CreateTask     string = "/virtualman/v1/createTask"
	GetSpeakerList string = "/virtualman/v1/getSpeakerList"
	GetFigureList  string = "/virtualman/v1/getVirtualmanList"
)

type VirtualmanClient struct {
	*common.Client
}

func NewVirtualmanClient(info *common.ServiceInfo) (*VirtualmanClient, error) {
	client, err := common.NewClient(info)
	if err != nil {
		return nil, err
	}
	return &VirtualmanClient{client}, nil
}

type queryMarshalFilter func(key string, value []string) (accept bool)

func skipEmptyValue() queryMarshalFilter {
	return func(_ string, value []string) (accept bool) {
		if len(value) == 0 {
			return false
		}

		for _, item := range value {
			if item == "" {
				return false
			}
		}

		return true
	}
}

func marshalToQuery(model interface{}, filters ...queryMarshalFilter) (url.Values, error) {
	ret := url.Values{}
	ret.Add("action", "virtualman")
	if model == nil {
		return ret, nil
	}

	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)
	if modelValue.IsNil() {
		return ret, nil
	}

	if modelType.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
		modelType = modelValue.Type()
	} else {
		return ret, fmt.Errorf("not struct")
	}
	fieldCount := modelType.NumField()

	for i := 0; i < fieldCount; i++ {
		field := modelType.Field(i)
		queryKey := field.Tag.Get("query")
		if queryKey == "" {
			continue
		}

		queryRawValue := modelValue.FieldByName(field.Name)
		queryValues := make([]string, 0)
		if field.Type.Kind() != reflect.Array && field.Type.Kind() != reflect.Slice {
			value := resolveQueryValue(queryRawValue)
			if value == nil {
				continue
			}
			queryValues = append(queryValues, fmt.Sprintf("%v", value))
		} else {
			length := queryRawValue.Len()
			for idx := 0; idx < length; idx++ {
				value := resolveQueryValue(queryRawValue.Index(idx))
				if value == nil {
					continue
				}
				queryValues = append(queryValues, fmt.Sprintf("%v", value))
			}
		}

		for _, fun := range filters {
			if !fun(queryKey, queryValues) {
				goto afterAddQuery
			}
		}

		for _, t := range queryValues {
			ret.Add(queryKey, t)
		}
	afterAddQuery:
	}
	return ret, nil
}

func resolveQueryValue(queryRawValue reflect.Value) interface{} {
	if queryRawValue.Kind() == reflect.Ptr {
		if queryRawValue.IsNil() {
			return nil
		}
		decayedQueryRawValue := queryRawValue.Elem()
		decayedReflectValue := decayedQueryRawValue.Interface()
		return fmt.Sprintf("%v", decayedReflectValue)
	} else {
		queryReflectValue := queryRawValue.Interface()
		return fmt.Sprintf("%v", queryReflectValue)
	}
}

func marshalToJson(model interface{}) ([]byte, error) {
	if model == nil {
		return make([]byte, 0), nil
	}
	result, err := json.Marshal(model)
	if err != nil {
		return []byte{}, fmt.Errorf("can not marshal model to json, %v", err)
	}
	return result, nil
}

func unmarshalResultInto(data []byte, result interface{}) error {
	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

func (fc *VirtualmanClient) CreateTaskByText(define *types.TextVirtualmanTaskDefine) (*types.CreateTaskRes, error) {
	if define.VirtualmanId == "" {
		return nil, errors.New("the virtualmanId must not be empty")
	}
	if define.Text == "" {
		return nil, errors.New("the text must not be empty")
	}
	if define.SpeakerId == "" {
		return nil, errors.New("the speakerId must not be empty")
	}
	if define.CallbackUrl != "" && !common.IsLegalUrl(define.CallbackUrl) {
		return nil, errors.New("this callbackUrl input is illegal")
	}
	query, err := marshalToQuery(nil)
	body, err := marshalToJson(define)
	if err != nil {
		return nil, fmt.Errorf("json serialization failed %s", define.String())
	}
	data, traceId, err := fc.Post(nil, CreateTask, query, body)
	if err != nil {
		return nil, err
	}
	result := new(types.CreateTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}
	result.TraceId = traceId
	return result, nil
}

func (fc *VirtualmanClient) CreateTaskByAudio(define *types.AudioVirtualmanTaskDefine) (*types.CreateTaskRes, error) {
	if define.VirtualmanId == "" {
		return nil, errors.New("the virtualmanId must not be empty")
	}
	if define.AudioUrl == "" || !common.IsLegalUrl(define.AudioUrl) {
		return nil, errors.New("the audioUrl must not be empty or illegal")
	}
	if define.CallbackUrl != "" && !common.IsLegalUrl(define.CallbackUrl) {
		return nil, errors.New("this callbackUrl input is illegal")
	}
	query, err := marshalToQuery(nil)
	body, err := marshalToJson(define)
	if err != nil {
		return nil, fmt.Errorf("json serialization failed %s", define.String())
	}
	data, traceId, err := fc.Post(nil, CreateTask, query, body)
	if err != nil {
		return nil, err
	}
	result := new(types.CreateTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}
	result.TraceId = traceId
	return result, nil
}

func (fc *VirtualmanClient) GetSpeakerList(define *types.PageDefine) (*types.GetSpeakerListRes, error) {
	query, err := marshalToQuery(define)
	if err != nil {
		return nil, err
	}
	data, traceId, err := fc.Get(nil, GetSpeakerList, query)
	if err != nil {
		return nil, err
	}
	result := new(types.GetSpeakerListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}
	result.TraceId = traceId
	return result, nil
}

func (fc *VirtualmanClient) GetVirtualmanList(define *types.PageDefine) (*types.GetVirtualmanListRes, error) {
	query, err := marshalToQuery(define)
	if err != nil {
		return nil, err
	}
	data, traceId, err := fc.Get(nil, GetFigureList, query)
	if err != nil {
		return nil, err
	}
	result := new(types.GetVirtualmanListRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}
	result.TraceId = traceId
	return result, nil
}
