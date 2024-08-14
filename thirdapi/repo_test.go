package thirdapi

import "testing"

func TestGetRepos(t *testing.T) {
	got, err := GetRepos()
	if err != nil {
		t.Fatalf("fail to get repos, err:%v", err)
	}

	t.Logf("got: %v", got)
}
