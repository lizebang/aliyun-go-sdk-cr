package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_GetRepo(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetRepoRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	request.RepoNamespace = os.Getenv("NAMESPACE")
	request.RepoName = os.Getenv("REPO")

	response, err := client.GetRepo(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
