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

// CreateNamespace invokes the cr.CreateNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/createnamespace.html
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
	response = CreateCreateNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNamespaceWithChan invokes the cr.CreateNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/createnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateNamespaceWithChan(request *CreateNamespaceRequest) (<-chan *CreateNamespaceResponse, <-chan error) {
	responseChan := make(chan *CreateNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNamespace(request)
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

// CreateNamespaceWithCallback invokes the cr.CreateNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/createnamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateNamespaceWithCallback(request *CreateNamespaceRequest, callback func(response *CreateNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNamespaceResponse
		var err error
		defer close(result)
		response, err = client.CreateNamespace(request)
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

// CreateNamespaceRequest is the request struct for api CreateNamespace
type CreateNamespaceRequest struct {
	*requests.RoaRequest
}

// CreateNamespaceResponse is the response struct for api CreateNamespace
type CreateNamespaceResponse struct {
	*responses.BaseResponse
}

// CreateCreateNamespaceRequest creates a request to invoke CreateNamespace API
func CreateCreateNamespaceRequest() (request *CreateNamespaceRequest) {
	request = &CreateNamespaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "CreateNamespace", "/namespace", "", "")
	request.Method = requests.PUT
	return
}

// CreateCreateNamespaceResponse creates a response to parse from CreateNamespace response
func CreateCreateNamespaceResponse() (response *CreateNamespaceResponse) {
	response = &CreateNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// CreateNamespaceRequestBody is the request body struct for api CreateNamespace
type CreateNamespaceRequestBody struct {
	Namespace struct {
		Namespace string `json:"Namespace"`
	} `json:"Namespace"`
}

// NewCreateNamespaceRequestBody creates a request body
func NewCreateNamespaceRequestBody(namespace string) (body *CreateNamespaceRequestBody, err error) {
	cnrb := &CreateNamespaceRequestBody{}
	cnrb.Namespace.Namespace = namespace
	return cnrb, nil
}

// Marshal returns the JSON encoding of request body
func (cnrb *CreateNamespaceRequestBody) Marshal() (body []byte, err error) {
	return json.Marshal(cnrb)
}
