package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_GetRepoTags(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetRepoTagsRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	request.RepoNamespace = os.Getenv("NAMESPACE")
	request.RepoName = os.Getenv("REPO")

	response, err := client.GetRepoTags(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_GetRepoTagsWithData(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetRepoTagsRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	request.RepoNamespace = os.Getenv("NAMESPACE")
	request.RepoName = os.Getenv("REPO")
	request.Page = "1"
	request.PageSize = "5"

	response, err := client.GetRepoTags(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
