package thirdapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 1. 全局自定义client，可以做统一出口逻辑处理
// 2. 利用范型统一处理反序列化，直接获取到目标结构体和error
// 3. resp的http状态和业务状态做统一判断

var (
	client *http.Client
	once   sync.Once

	ErrSendReq = errors.New("fail to send request")
)

func InitClient() {
	tr := &http.Transport{
		MaxIdleConns:        1000, // 最大空间连接总数
		MaxIdleConnsPerHost: 10,   // 每个host最大空闲连接数，不设置有系统默认值
		MaxConnsPerHost:     1000, // 每个host最大连接数 不设置就是不限制
		IdleConnTimeout:     10 * time.Minute,
		DisableKeepAlives:   false,
	}
	client = &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
}

func Send[T any](req *http.Request, output T) error {
	once.Do(InitClient)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%w, fail to exec client.Do, err: %v", ErrSendReq, err)
	}
	if resp == nil {
		return fmt.Errorf("%w, resp is nil", ErrSendReq)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%w, fail to read resp.Body, err: %v", ErrSendReq, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("%w, http status is abnormal, curr code: %v, body data: %v",
			ErrSendReq, resp.StatusCode, string(data))
	}

	err = json.Unmarshal(data, output)
	if err != nil {
		return fmt.Errorf("%w, fail to unmarshal data to output, err: %v", ErrSendReq, err)
	}

	return nil
}

func SendV2[T any](req *http.Request) (T, error) {
	once.Do(InitClient)
	var output T

	resp, err := client.Do(req)
	if err != nil {
		return output, fmt.Errorf("%w, fail to exec client.Do, err: %v", ErrSendReq, err)
	}
	if resp == nil {
		return output, fmt.Errorf("%w, resp is nil", ErrSendReq)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return output, fmt.Errorf("%w, fail to read resp.Body, err: %v", ErrSendReq, err)
	}
	defer func() {
		resp.Body.Close()
	}()

	if resp.StatusCode >= http.StatusBadRequest {
		return output, fmt.Errorf("%w, http status is abnormal, curr code: %v, body data: %v",
			ErrSendReq, resp.StatusCode, string(data))
	}

	err = json.Unmarshal(data, &output)
	if err != nil {
		return output, fmt.Errorf("%w, fail to unmarshal data to output, err: %v", ErrSendReq, err)
	}

	return output, nil
}
