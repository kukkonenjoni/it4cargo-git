package util

type Config struct {
	ProjectId    string       `json:"projectid"`
	TagName      string       `json:"tagname"`
	ReleaseName  string       `json:"releasename"`
	Base         string       `json:"base"`
	BaseRepo     string       `json:"baserepo"`
	BaseRepoName string       `json:"basereponame"`
	Repositories []Repository `json:"repositories"`
}
type Repository struct {
	ClientName string `json:"clientname"`
	Branch     string `json:"branch"`
	GitUrl     string `json:"giturl"`
	FileName   string `json:"filename"`
}
