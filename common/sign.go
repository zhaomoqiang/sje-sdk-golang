package common

import (
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

const (
	API_TIMESTAMP string = "api-timestamp"
	AUTHORIZATION string = "authorization"
)

var HEADER_KEYS_TO_IGNORE = []string{
	"authorization",
	"content-type",
	"content-length",
	"user-agent",
	"presigned-expires",
	"expect",
	"host"}

func Sign(signParameters *SignParameters, credentials *Credentials) string {
	metaData := getMetaDate(signParameters.Date)
	bodySha := HashSHA256(signParameters.Body)
	hashedCanonReq := hashedCanonicalRequest(signParameters, metaData, bodySha)
	stringToSign := concat("\n", metaData.Algorithm, metaData.Date, metaData.CredentialScope, hashedCanonReq)
	signature := hex.EncodeToString(HmacSHA256([]byte(credentials.AccessKeySecret), stringToSign))
	return buildAuthHeader(signature, metaData, credentials)
}

func buildAuthHeader(signature string, meta *MetaData, credentials *Credentials) string {
	credential := credentials.AccessKeyId + "/" + meta.CredentialScope
	return meta.Algorithm +
		" Credential=" + credential +
		" SignedHeaders=" + meta.SignedHeaders +
		" Signature=" + signature
}

func getMetaDate(timestamp string) *MetaData {
	return &MetaData{Date: timestamp, Algorithm: "HMAC-SHA256", CredentialScope: concat("/", timestamp[:8], "request")}
}

func hashedCanonicalRequest(signParameters *SignParameters, meta *MetaData, bodyHash string) string {
	getCanonicalHeaders(signParameters, meta)
	canonicalRequest := concat("\n", signParameters.Method, signQueryEncoder(signParameters.Query), meta.CanonicalHeaders, meta.SignedHeaders, bodyHash)
	return HashSHA256([]byte(canonicalRequest))
}

func getCanonicalHeaders(signParameters *SignParameters, meta *MetaData) {
	var needSignHeaders []string
	needSignHeaders = append(needSignHeaders, API_TIMESTAMP)
	var signedHeaders []string
	originHeaders := signParameters.Headers
	for k, _ := range originHeaders {
		signedHeaders = append(signedHeaders, strings.ToLower(k))
	}
	if signParameters.NeedSignHeaderKeys != nil && len(signParameters.NeedSignHeaderKeys) > 0 {
		for _, k := range signParameters.NeedSignHeaderKeys {
			needSignHeaders = append(needSignHeaders, strings.ToLower(k))
		}
		signedHeaders = filter(signedHeaders, needSignHeaders, true)
	}
	signedHeaders = filter(signedHeaders, HEADER_KEYS_TO_IGNORE, false)

	sort.Strings(signedHeaders)
	signedHeaderKeys := strings.Join(signedHeaders, ";")

	var signedHeadersToSign string
	for _, k := range signedHeaders {
		value := strings.TrimSpace(originHeaders[k])
		signedHeadersToSign += k + ":" + value + "\n"
	}
	meta.SignedHeaders = signedHeaderKeys
	meta.CanonicalHeaders = signedHeadersToSign
}

func signQueryEncoder(params url.Values) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var result []string
	for _, k := range keys {
		v := params[k]
		if v != nil {
			result = append(result, concat("=", k, v[0]))
		} else {
			result = append(result, k+"=")
		}
	}
	return strings.Join(result, "&")
}

func concat(sep string, str ...string) string {
	return strings.Join(str, sep)
}

func filter(signedHeaders []string, needSignHeaders []string, op bool) []string {
	var headers []string
	var valid bool
	for _, k := range signedHeaders {
		for _, kk := range needSignHeaders {
			if op {
				if k == kk {
					valid = true
					break
				}
			} else {
				if k != kk {
					valid = true
					break
				}
			}
		}
		if valid {
			headers = append(headers, k)
		}
	}
	return headers
}
