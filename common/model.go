package common

import (
	"net/url"
	"time"
)

type ServiceInfo struct {
	Timeout     time.Duration
	Scheme      string
	Host        string
	Credentials *Credentials
}

type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
}

func (serviceInfo *ServiceInfo) Clone() *ServiceInfo {
	ret := new(ServiceInfo)
	//base info
	if serviceInfo.Timeout != time.Duration(0) {
		ret.Timeout = serviceInfo.Timeout
	} else {
		ret.Timeout = 5 * time.Second
	}
	ret.Host = serviceInfo.Host
	ret.Scheme = serviceInfo.Scheme
	//credential
	ret.Credentials = serviceInfo.Credentials.Clone()
	return ret
}

func (cred *Credentials) Clone() *Credentials {
	return &Credentials{
		AccessKeyId:     cred.AccessKeyId,
		AccessKeySecret: cred.AccessKeySecret,
	}
}

type SignParameters struct {
	Method             string
	Date               string
	Query              url.Values
	Body               []byte
	Headers            map[string]string
	NeedSignHeaderKeys []string
}

type MetaData struct {
	Algorithm        string
	CredentialScope  string
	SignedHeaders    string
	CanonicalHeaders string
	Date             string
}
