package common

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	_GlobalClient *http.Client
	emptyBytes    []byte
)

const (
	POST string = "POST"
	GET  string = "GET"
)

func init() {
	_GlobalClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     10 * time.Second,
		},
	}
}

// Client
type Client struct {
	Client      *http.Client
	ServiceInfo *ServiceInfo
}

func IsLegalUrl(url string) bool {
	re := regexp.MustCompile(`(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
	result := re.FindAllStringSubmatch(url, -1)
	if result == nil {
		return false
	}
	return true
}

// NewClient
func NewClient(info *ServiceInfo) (*Client, error) {
	if !IsLegalUrl(info.Scheme + "://" + info.Host) {
		return nil, fmt.Errorf("host illegality, please check the input parameters %s", info.Host)
	}
	if info.Credentials.AccessKeyId == "" || info.Credentials.AccessKeySecret == "" {
		return nil, errors.New("AccessKeyId or AccessKeySecret missing, Please check the input parameters")
	}
	client := &Client{Client: _GlobalClient, ServiceInfo: info.Clone()}
	return client, nil
}

func getStandardHeader() (header http.Header) {
	header = http.Header{}
	header.Add("api-sdk-name", SDKName)
	header.Add("api-sdk-version", SDKVersion)
	return header
}

func header2Map(header http.Header) map[string]string {
	result := make(map[string]string, len(header))
	for k, vv := range header {
		result[strings.ToLower(k)] = vv[0]
	}
	return result
}

func (client *Client) Get(ctx context.Context, api string, query url.Values) ([]byte, string, error) {
	return client.request(ctx, api, GET, query, emptyBytes, "")
}

func (client *Client) Post(ctx context.Context, api string, query url.Values, body []byte) ([]byte, string, error) {
	return client.request(ctx, api, POST, query, body, "application/json")
}

func (client *Client) request(ctx context.Context, api string, method string, query url.Values, body []byte, ct string) ([]byte, string, error) {
	u := url.URL{
		Scheme:   client.ServiceInfo.Scheme,
		Host:     client.ServiceInfo.Host,
		Path:     api,
		RawQuery: query.Encode(),
	}
	requestBody := bytes.NewReader(body)
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return []byte(""), "", errors.New("failed to build request")
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	header := getStandardHeader()
	hm := header2Map(header)
	hm[API_TIMESTAMP] = timestamp
	signParameters := &SignParameters{Method: method, Date: timestamp, Query: query, Body: body, Headers: hm}
	authorization := Sign(signParameters, client.ServiceInfo.Credentials)
	header.Add(API_TIMESTAMP, timestamp)
	header.Add(AUTHORIZATION, authorization)
	req.Header = header
	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}

	req.Body = ioutil.NopCloser(requestBody)
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithTimeout(ctx, client.ServiceInfo.Timeout)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := client.Client.Do(req)
	if err != nil {
		return []byte(""), "", err
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	traceId := resp.Header.Get("trace-id")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), traceId, err
	}
	if code < 200 || code > 299 {
		return data, traceId, fmt.Errorf("api %s http code %d traceId %s body %s", api, code, traceId, string(data))
	}
	return data, traceId, nil
}
