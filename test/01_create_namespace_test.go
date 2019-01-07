package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_CreateNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateCreateNamespaceRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	cnrb := cr.NewCreateNamespaceRequestBody(os.Getenv("NAMESPACE") + "-delete")
	body, err := cnrb.Marshal()
	assert.Nil(t, err)
	request.SetContent(body)

	response, err := client.CreateNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}

func Test_CreateNamespaceForTest(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateCreateNamespaceRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	cnrb := cr.NewCreateNamespaceRequestBody(os.Getenv("NAMESPACE"))
	body, err := cnrb.Marshal()
	assert.Nil(t, err)
	request.SetContent(body)

	response, err := client.CreateNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
