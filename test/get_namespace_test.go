package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_GetNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateGetNamespaceRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	request.Namespace = os.Getenv("NAMESPACE")

	response, err := client.GetNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
