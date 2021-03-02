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
	assert.Equal(t, 200, resp.StatusCode, "应该返回状态码200")
}

func TestAboutPage(t *testing.T) {
	baseURL := "http://localhost:3000"

	var (
		resp *http.Response
		err error
	)

	resp, err = http.Get(baseURL + "/about")

	assert.NoError(t, err, "有错误发生，err不为空！")
	assert.Equal(t, 200, resp.StatusCode, "应该返回状态码200")
}

func TestAllPage(t *testing.T)  {





}
