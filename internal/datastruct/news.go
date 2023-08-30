package datastruct

type NewsRequestBody struct {
	Action                 string   `json:"action"`
	Keyword                string   `json:"keyword"`
	ArticlesPage           int      `json:"articlesPage"`
	ArticlesCount          int      `json:"articlesCount"`
	ArticlesSortBy         string   `json:"articlesSortBy"`
	ArticlesSortByAsc      bool     `json:"articlesSortByAsc"`
	ArticlesArticleBodyLen int      `json:"articlesArticleBodyLen"`
	ResultType             string   `json:"resultType"`
	DataType               []string `json:"dataType"`
	ApiKey                 string   `json:"apiKey"`
	ForceMaxDataTimeWindow int      `json:"forceMaxDataTimeWindow"`
}
