package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lizebang/aliyun-go-sdk-cr/services/cr"
)

func Test_UpdateNamespace(t *testing.T) {
	client, err := cr.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	assert.Nil(t, err)

	request := cr.CreateUpdateNamespaceRequest()
	request.SetDomain(os.Getenv("DOMAIN"))
	request.Namespace = os.Getenv("NAMESPACE")
	unrb := cr.UpdateNamespaceRequestBody{}
	unrb.Namespace.AutoCreate = false
	unrb.Namespace.DefaultVisibility = cr.NamespaceVisibilityPublic
	body, err := json.Marshal(unrb)
	assert.Nil(t, err)
	request.SetContent(body)

	response, err := client.UpdateNamespace(request)
	assert.Nil(t, err)
	assert.True(t, response.IsSuccess())
}
