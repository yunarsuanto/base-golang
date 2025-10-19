package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yunarsuanto/base-go/constants"
)

func client() *http.Client {
	return &http.Client{
		// Timeout: time.Second * constants.HttpClientDefaultTimeout,
		Transport: &http.Transport{
			DisableCompression: true,
			DisableKeepAlives:  true,
			// Dial: (&net.Dialer{
			// 	Timeout: 60 * time.Second,
			// }).Dial,
			// TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

func do(method string, endpoint string, params map[string]string, headers map[string]string, body io.Reader) ([]byte, *constants.ErrorResponse) {
	var result []byte
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return result, errs
	}
	req.Close = true

	q := req.URL.Query()
	for i, v := range params {
		q.Add(i, v)
	}
	req.URL.RawQuery = q.Encode()

	for i, v := range headers {
		req.Header.Add(i, v)
	}

	netClient := client()
	resp, err := netClient.Do(req)
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return result, errs
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Errorln(err)
		}
	}()

	result, err = io.ReadAll(resp.Body)
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return result, errs
	}
	if resp.StatusCode != 200 {
		log.Error(err, ConvertBytesToString(result))
		return nil, ErrHttpClient(endpoint)
	}

	return result, nil
}

func HttpClientDoMultipart(method string, endpoint string, params map[string]string, headers map[string]string, body map[string]io.Reader, result interface{}) *constants.ErrorResponse {
	var b bytes.Buffer
	var err error

	w := multipart.NewWriter(&b)

	for key, r := range body {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer func() {
				err := x.Close()
				if err != nil {
					log.Errorln(err)
				}
			}()
		}

		if x, ok := r.(*os.File); ok {
			fw, err = w.CreateFormFile(key, x.Name())
			if err != nil {
				errs := ErrorInternalServer(err.Error())
				return errs
			}
		} else {
			fw, err = w.CreateFormField(key)
			if err != nil {
				errs := ErrorInternalServer(err.Error())
				return errs
			}
		}
		_, err = io.Copy(fw, r)
		if err != nil {
			errs := ErrorInternalServer(err.Error())
			return errs
		}
	}

	err = w.Close()
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return errs
	}

	if headers != nil {
		headers["Content-Type"] = w.FormDataContentType()
	} else {
		headers = map[string]string{
			"Content-Type": w.FormDataContentType(),
			"Accept":       "application/json",
		}
	}

	buf, errs := do(method, endpoint, params, headers, &b)
	if errs != nil {
		return errs
	}

	err = json.Unmarshal(buf, &result)
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return errs
	}

	return nil
}

func HttpClientDoJson(method string, endpoint string, params map[string]string, headers map[string]string, body interface{}, result interface{}) *constants.ErrorResponse {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return errs
	}

	reqBody := bytes.NewBuffer(jsonBody)
	buf, errs := do(method, endpoint, params, headers, reqBody)
	if errs != nil {
		return errs
	}

	err = json.Unmarshal(buf, &result)
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return errs
	}

	return nil
}

func HttpClientDo(method string, endpoint string, params map[string]string, headers map[string]string, result interface{}) *constants.ErrorResponse {
	buf, errs := do(method, endpoint, params, headers, nil)
	if errs != nil {
		return errs
	}

	err := json.Unmarshal(buf, result)
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return errs
	}

	return nil
}

func HttpClientDoUrlEncoded(method string, endpoint string, params map[string]string, headers map[string]string, body url.Values, result interface{}) *constants.ErrorResponse {
	reqBody := strings.NewReader(body.Encode())
	buf, errs := do(method, endpoint, params, headers, reqBody)
	if errs != nil {
		return errs
	}

	err := json.Unmarshal(buf, &result)
	if err != nil {
		errs := ErrorInternalServer(err.Error())
		return errs
	}

	return nil
}
