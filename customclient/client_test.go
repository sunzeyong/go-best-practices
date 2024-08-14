package customclient

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
	Domain = "https://api.github.com"
)
