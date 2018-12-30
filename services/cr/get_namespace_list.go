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

// GetNamespaceList invokes the cr.GetNamespaceList API synchronously
// api document: https://help.aliyun.com/api/cr/getnamespacelist.html
func (client *Client) GetNamespaceList(request *GetNamespaceListRequest) (response *GetNamespaceListResponse, err error) {
	response = CreateGetNamespaceListResponse()
	err = client.DoAction(request, response)
	return
}

// GetNamespaceListWithChan invokes the cr.GetNamespaceList API asynchronously
// api document: https://help.aliyun.com/api/cr/getnamespacelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNamespaceListWithChan(request *GetNamespaceListRequest) (<-chan *GetNamespaceListResponse, <-chan error) {
	responseChan := make(chan *GetNamespaceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNamespaceList(request)
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

// GetNamespaceListWithCallback invokes the cr.GetNamespaceList API asynchronously
// api document: https://help.aliyun.com/api/cr/getnamespacelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNamespaceListWithCallback(request *GetNamespaceListRequest, callback func(response *GetNamespaceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNamespaceListResponse
		var err error
		defer close(result)
		response, err = client.GetNamespaceList(request)
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

// GetNamespaceListRequest is the request struct for api GetNamespaceList
type GetNamespaceListRequest struct {
	*requests.RoaRequest
}

// GetNamespaceListResponse is the response struct for api GetNamespaceList
type GetNamespaceListResponse struct {
	*responses.BaseResponse
}

// CreateGetNamespaceListRequest creates a request to invoke GetNamespaceList API
func CreateGetNamespaceListRequest() (request *GetNamespaceListRequest) {
	request = &GetNamespaceListRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "GetNamespaceList", "/namespace", "", "")
	request.Method = requests.GET
	return
}

// CreateGetNamespaceListResponse creates a response to parse from GetNamespaceList response
func CreateGetNamespaceListResponse() (response *GetNamespaceListResponse) {
	response = &GetNamespaceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
