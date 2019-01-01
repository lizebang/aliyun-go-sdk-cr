package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_DeleteRepo(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateDeleteRepoRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	request.RepoNamespace = os.Getenv("NAMESPACE")

	request.RepoName = os.Getenv("REPO") + "-without-repo-source"
	response, err := client.DeleteRepo(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())

	request.RepoName = os.Getenv("REPO") + "-with-repo-source"
	response, err = client.DeleteRepo(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
