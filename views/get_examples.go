package views

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//定义新的数据类型
type Spider struct {
	url    string
	header map[string]string
}

//定义 Spider get的方法
func (keyword Spider) get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return string(body)

}
func GetExample(word string) []string {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}

	url := "http://dict.youdao.com/w/eng/" + word
	spider := &Spider{url, header}
	html := spider.get_html_header()

	pattern := `<div class="examples">\s*<p>\s*(.*?)</p>\s*<p>(.*?)</p>\s*</div>`
	rp := regexp.MustCompile(pattern)
	examples := rp.FindAllStringSubmatch(html, -1)
	if examples == nil {
		return nil
	}
	return examples[0][1:3]
}
