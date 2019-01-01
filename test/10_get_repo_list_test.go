package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_GetRepoList(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetRepoListRequest()
	request.SetDomain(os.Getenv("DOMAIN"))

	response, err := client.GetRepoList(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_GetRepoListWithData(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetRepoListRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	request.Status = "NORMAL"
	request.RepoNamePrefix = "a"
	request.Page = "1"
	request.PageSize = "5"

	response, err := client.GetRepoList(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
