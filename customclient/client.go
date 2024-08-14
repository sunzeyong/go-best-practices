package customclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// 1. resp的http状态和业务状态做统一判断
// 2. 利用范性统一处理反序列化，直接获取到目标结构体和error
// 3. 提供生成Request的便捷函数，如header添加，结构体转url参数

var (
	ErrRequest           = errors.New("err happend at client do request")
	ErrRespNil           = errors.New("resp is nil")
	ErrRespStatusIsNotOk = errors.New("resp status is not ok")
)

func Send[T any](req *http.Request, output T) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("%w, err: %v", ErrRequest, err)
	}

	if resp == nil {
		return ErrRespNil
	}
	if resp.StatusCode != http.StatusOK {
		return ErrRespStatusIsNotOk
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, output)
	if err != nil {
		return err
	}
	return nil
}

func DirectGet() {
	

}
