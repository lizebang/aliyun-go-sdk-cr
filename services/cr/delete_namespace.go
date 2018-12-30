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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteNamespace invokes the cr.DeleteNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/deletenamespace.html
func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
	response = CreateDeleteNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNamespaceWithChan invokes the cr.DeleteNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/deletenamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteNamespaceWithChan(request *DeleteNamespaceRequest) (<-chan *DeleteNamespaceResponse, <-chan error) {
	responseChan := make(chan *DeleteNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNamespace(request)
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

// DeleteNamespaceWithCallback invokes the cr.DeleteNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/deletenamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteNamespaceWithCallback(request *DeleteNamespaceRequest, callback func(response *DeleteNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNamespaceResponse
		var err error
		defer close(result)
		response, err = client.DeleteNamespace(request)
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

// DeleteNamespaceRequest is the request struct for api DeleteNamespace
type DeleteNamespaceRequest struct {
	*requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

// DeleteNamespaceResponse is the response struct for api DeleteNamespace
type DeleteNamespaceResponse struct {
	*responses.BaseResponse
}

// CreateDeleteNamespaceRequest creates a request to invoke DeleteNamespace API
func CreateDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
	request = &DeleteNamespaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "DeleteNamespace", "/namespace/[Namespace]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteNamespaceResponse creates a response to parse from DeleteNamespace response
func CreateDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
	response = &DeleteNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
