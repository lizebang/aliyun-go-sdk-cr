package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_CreateRepo(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateCreateRepoRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	crrb, _ := cr.NewCreateRepoRequestBody(os.Getenv("NAMESPACE"), os.Getenv("REPO")+"-without-repo-source",
		os.Getenv("REPO")+"-without-repo-source", os.Getenv("REPO")+"-without-repo-source", cr.RepoTypePrivate)
	body, err := crrb.Marshal()
	assert.Nil(t, err)
	request.SetContent(body)

	response, err := client.CreateRepo(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CreateRepoWithRepoSource(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateCreateRepoRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	crrb, _ := cr.NewCreateRepoRequestBody(os.Getenv("NAMESPACE"), os.Getenv("REPO")+"-with-repo-source",
		os.Getenv("REPO")+"-with-repo-source", os.Getenv("REPO")+"-with-repo-source", cr.RepoTypePrivate)
	crrb.SetRepoSource(cr.SourceRepoTypeGitHub, os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_REPO"))
	body, err := crrb.Marshal()
	assert.Nil(t, err)
	request.SetContent(body)

	response, err := client.CreateRepo(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CreateRepoForTest(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateCreateRepoRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	crrb, _ := cr.NewCreateRepoRequestBody(os.Getenv("NAMESPACE"), os.Getenv("REPO"), os.Getenv("REPO"),
		os.Getenv("REPO"), cr.RepoTypePrivate)
	crrb.SetRepoSource(cr.SourceRepoTypeGitHub, os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_REPO"))
	body, err := crrb.Marshal()
	assert.Nil(t, err)
	request.SetContent(body)

	response, err := client.CreateRepo(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
