package thirdapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-querystring/query"
)

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type RepoParams struct {
	PerPage int `url:"per_page"`
	Page    int `url:"page"`
}

var (
	ErrFailGetResp = errors.New("fail to get resp")
	ErrFailSendReq = errors.New("fail to send request")
)

// 直接GET请求
func GetRepos() ([]Repo, error) {
	resp, err := http.Get("https://api.github.com/users/sunzeyong/repos")
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("%w, err: %v", ErrFailGetResp, "resp is nil")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, http code is not ok, cur: %v", ErrFailGetResp, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to read resp.Body, err: %v", ErrFailGetResp, err)
	}
	var output []Repo
	err = json.Unmarshal(data, &output)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to unmarshal, err: %w", ErrFailGetResp, err)
	}
	return output, nil
}

func GetReposWithParams(p RepoParams) ([]Repo, error) {
	// 结构体转url.Values
	queryValues, err := query.Values(p)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to convert query, err:%v", ErrFailSendReq, err)
	}
	u, err := url.Parse("https://api.github.com/users/sunzeyong/repos")
	if err != nil {
		return nil, fmt.Errorf("%w, fail to parse url, err:%v", ErrFailSendReq, err)
	}
	u.RawQuery = queryValues.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("%w, err: %v", ErrFailGetResp, "resp is nil")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, http code is not ok, cur: %v", ErrFailGetResp, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to read resp.Body, err: %v", ErrFailGetResp, err)
	}
	var output []Repo
	err = json.Unmarshal(data, &output)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to unmarshal, err: %w", ErrFailGetResp, err)
	}
	return output, nil
}

func GetWithRequest(p RepoParams) ([]Repo, error) {
	// prepare request
	u, err := url.Parse("https://api.github.com/users/sunzeyong/repos")
	if err != nil {
		return nil, fmt.Errorf("%w, fail to parse url, err: %v", ErrFailSendReq, err)
	}
	queryValue, err := query.Values(p)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to convert query, err:%v", ErrFailSendReq, err)
	}
	u.RawQuery = queryValue.Encode()

	// new request
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to new request, err:%v", ErrFailSendReq, err)
	}

	// send request and process resp
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to send request, err:%v", ErrFailSendReq, err)
	}

	if resp == nil {
		return nil, fmt.Errorf("%w, err: %v", ErrFailGetResp, "resp is nil")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, http code is not ok, cur: %v", ErrFailGetResp, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to read resp.Body, err: %v", ErrFailGetResp, err)
	}
	var output []Repo
	err = json.Unmarshal(data, &output)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to unmarshal, err: %w", ErrFailGetResp, err)
	}
	return output, nil
}

type CreateTagInput struct {
	TagName string `json:"tag_name"`
}

type CreateTagOutput struct {
	URL    string `json:"url"`
	Id     int    `json:"id"`
	Author struct {
		Login string `json:"login"`
	} `json:"author"`
	TagName string `json:"tag_name"`
}

func CreateTagName(input CreateTagInput) (*CreateTagOutput, error) {
	// prepare params
	owner, repo := "sunzeyong", "go-best-practices"
	path := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)

	inputByte, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to marshal data, err: %v", ErrFailSendReq, err)
	}

	// prepare request, add header
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(inputByte))
	if err != nil {
		return nil, fmt.Errorf("%w, fail to new request, err:%v", ErrFailSendReq, err)
	}
	key := os.Getenv("GITHUBKEY")
	req.Header.Add("Authorization", "Bearer "+key)

	// send request and process resp
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to send request, err:%v", ErrFailSendReq, err)
	}

	if resp == nil {
		return nil, fmt.Errorf("%w, err: %v", ErrFailGetResp, "resp is nil")
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("%w, http code is not ok, cur: %v", ErrFailGetResp, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to read resp.Body, err: %v", ErrFailGetResp, err)
	}
	output := &CreateTagOutput{}
	err = json.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("%w, fail to unmarshal, err: %w", ErrFailGetResp, err)
	}
	return output, nil
}
