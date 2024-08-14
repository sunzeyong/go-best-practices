package thirdapi

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRepos(t *testing.T) {
	got, err := GetRepos()
	if err != nil {
		t.Fatalf("fail to get repos, err:%v", err)
	}

	t.Logf("got: %v", got)
}

func TestGetReposWithParams(t *testing.T) {
	Convey("base cases", t, func() {
		cases := []struct {
			Name      string
			Input     RepoParams
			OutputLen int
		}{
			{"长度为1", RepoParams{PerPage: 1, Page: 1}, 1},
			{"长度为2", RepoParams{PerPage: 2, Page: 1}, 2},
		}
		for _, tt := range cases {
			Convey(tt.Name, func() {
				output, err := GetReposWithParams(tt.Input)
				So(err, ShouldBeNil)
				So(len(output), ShouldEqual, tt.OutputLen)
			})
		}
	})
}
