package githubx

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
				output, err := GetReposV2(tt.Input)
				So(err, ShouldBeNil)
				So(len(output), ShouldEqual, tt.OutputLen)
			})
		}
	})
}

func TestGetReposWithRequest(t *testing.T) {
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
				output, err := GetReposV3(tt.Input)
				So(err, ShouldBeNil)
				So(len(output), ShouldEqual, tt.OutputLen)
			})
		}
	})
}

func TestCreateTagName(t *testing.T) {
	tagName := "v0.0.9"
	Convey("base case", t, func() {
		Convey("add"+tagName, func() {
			input := CreateTagInput{
				TagName: tagName,
			}
			output, err := CreateTagName(input)
			So(err, ShouldBeNil)
			So(output.Author.Login, ShouldEqual, "sunzeyong")
			So(output.TagName, ShouldEqual, tagName)
		})
	})
}
