package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHomePage(t *testing.T)  {
	baseURL := "http://localhost:3000"

	var (
		resp *http.Response
		errors error
	)

	resp, errors = http.Get(baseURL+ "/")

	assert.NoError(t, errors, "有错误发生，errors不为空")
	assert.Equal(t, resp.StatusCode,200, "应该返回状态码200")
}
