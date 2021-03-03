package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)


func TestAllPage(t *testing.T)  {
	baseURL := "http://localhost:3000"

	tests := []struct {
		url      string
		method   string
		expected int
	}{
	 	{"/", "GET", 200},
	 	{"/about", "GET", 200},
	 	{"/notfound", "GET", 404},
	 	{"/articles/1", "GET", 200},
	 	{"/articles", "GET", 200},
	 	{"/articles/create", "GET", 200},
	 	{"/articles", "POST", 200},
	 	{"/articles/2/edit", "GET", 200},
	 	{"/articles/2", "GET", 200},
	 	{"/articles/-1", "POST", 404},
	}

	for _, test := range tests {
		var (
			resp *http.Response
			err error
		)

		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseURL+test.url, data)
		default:
			resp, err = http.Get(baseURL+test.url)
		}

		assert.NoError(t, err, "请求"+test.method+test.url+"有错误发生，err不为空！")
		assert.Equal(t, test.expected, resp.StatusCode, test.url+"应该返回状态码"+strconv.Itoa(test.expected))
	}

}
