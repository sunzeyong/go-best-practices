package thirdapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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


