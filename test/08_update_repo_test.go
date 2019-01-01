package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_UpdateRepo(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateUpdateRepoRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	request.RepoNamespace = os.Getenv("NAMESPACE")
	request.RepoName = os.Getenv("REPO")
	urrb := cr.UpdateRepoRequestBody{}
	urrb.Repo.Summary = os.Getenv("REPO")
	urrb.Repo.Detail = os.Getenv("REPO")
	urrb.Repo.RepoType = cr.RepoTypePublic
	body, err := json.Marshal(urrb)
	assert.Nil(t, err)
	request.SetContent(body)

	response, err := client.UpdateRepo(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
