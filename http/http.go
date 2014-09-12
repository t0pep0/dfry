package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type httpError string

func (h httpError) Error() string {
	return string(h)
}

const (
	_POST          = "POST"
	_CON_TYPE      = "Content-Type"
	_X_FORM        = "application/x-www-form-urlencoded"
	_INVALID_PARAM = httpError("Invalid param type")
)

func clientDo(req *http.Request) (answer io.Reader, err error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return answer, err
	}
	answer = resp.Body
	return answer, err
}

func Get(uri string) (answer io.Reader, err error) {
	resp, err := http.Get(uri)
	if err != nil {
		return answer, err
	}
	answer = resp.Body
	return answer, err
}

func Post(uri string, params map[string][]string) (answer io.Reader, err error) {
	data := url.Values(params)
	req, err := http.NewRequest(_POST, uri, strings.NewReader(data.Encode()))
	if err != nil {
		return answer, err
	}
	req.Header.Set(_CON_TYPE, _X_FORM)
	return clientDo(req)
}

func PostMiltipart(uri string, params map[string]interface{}) (answer io.Reader, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, value := range params {
		switch val := value.(type) {
		case string:
			writer.WriteField(key, val)
		case *os.File:
			part, _ := writer.CreateFormFile(key, val.Name())
			io.Copy(part, val)
		default:
			return answer, _INVALID_PARAM
		}
	}
	writer.Close()
	req, err := http.NewRequest(_POST, uri, body)
	req.Header.Add(_CON_TYPE, writer.FormDataContentType())
	return clientDo(req)
}
