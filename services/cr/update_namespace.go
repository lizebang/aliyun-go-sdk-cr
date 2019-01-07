package cr

// MIT License

// Copyright (c) 2018 Li Zebang

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateNamespace invokes the cr.UpdateNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/updatenamespace.html
func (client *Client) UpdateNamespace(request *UpdateNamespaceRequest) (response *UpdateNamespaceResponse, err error) {
	response = CreateUpdateNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNamespaceWithChan invokes the cr.UpdateNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/updatenamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateNamespaceWithChan(request *UpdateNamespaceRequest) (<-chan *UpdateNamespaceResponse, <-chan error) {
	responseChan := make(chan *UpdateNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNamespace(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateNamespaceWithCallback invokes the cr.UpdateNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/updatenamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateNamespaceWithCallback(request *UpdateNamespaceRequest, callback func(response *UpdateNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNamespaceResponse
		var err error
		defer close(result)
		response, err = client.UpdateNamespace(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateNamespaceRequest is the request struct for api UpdateNamespace
type UpdateNamespaceRequest struct {
	*requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

// UpdateNamespaceResponse is the response struct for api UpdateNamespace
type UpdateNamespaceResponse struct {
	*responses.BaseResponse
}

// CreateUpdateNamespaceRequest creates a request to invoke UpdateNamespace API
func CreateUpdateNamespaceRequest() (request *UpdateNamespaceRequest) {
	request = &UpdateNamespaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "UpdateNamespace", "/namespace/[Namespace]", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateNamespaceResponse creates a response to parse from UpdateNamespace response
func CreateUpdateNamespaceResponse() (response *UpdateNamespaceResponse) {
	response = &UpdateNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// UpdateNamespaceRequestBody is the request body struct for api UpdateNamespace
type UpdateNamespaceRequestBody struct {
	Namespace struct {
		AutoCreate        bool                `json:"AutoCreate"`
		DefaultVisibility namespaceVisibility `json:"DefaultVisibility"`
	} `json:"Namespace"`
}

// NewUpdateNamespaceRequestBody creates a request body
func NewUpdateNamespaceRequestBody(auto bool, visibility namespaceVisibility) *UpdateNamespaceRequestBody {
	unrb := &UpdateNamespaceRequestBody{}
	unrb.Namespace.AutoCreate = auto
	unrb.Namespace.DefaultVisibility = visibility
	return unrb
}

// Marshal returns the JSON encoding of the request body
func (unrb *UpdateNamespaceRequestBody) Marshal() (body []byte, err error) {
	return json.Marshal(unrb)
}
