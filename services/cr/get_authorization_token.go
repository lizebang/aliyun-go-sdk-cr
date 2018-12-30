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

// GetAuthorizationToken invokes the cr.GetAuthorizationToken API synchronously
// api document: https://help.aliyun.com/api/cr/getauthorizationtoken.html
func (client *Client) GetAuthorizationToken(request *GetAuthorizationTokenRequest) (response *GetAuthorizationTokenResponse, err error) {
	response = CreateGetAuthorizationTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetAuthorizationTokenWithChan invokes the cr.GetAuthorizationToken API asynchronously
// api document: https://help.aliyun.com/api/cr/getauthorizationtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAuthorizationTokenWithChan(request *GetAuthorizationTokenRequest) (<-chan *GetAuthorizationTokenResponse, <-chan error) {
	responseChan := make(chan *GetAuthorizationTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAuthorizationToken(request)
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

// GetAuthorizationTokenWithCallback invokes the cr.GetAuthorizationToken API asynchronously
// api document: https://help.aliyun.com/api/cr/getauthorizationtoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAuthorizationTokenWithCallback(request *GetAuthorizationTokenRequest, callback func(response *GetAuthorizationTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAuthorizationTokenResponse
		var err error
		defer close(result)
		response, err = client.GetAuthorizationToken(request)
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

// GetAuthorizationTokenRequest is the request struct for api GetAuthorizationToken
type GetAuthorizationTokenRequest struct {
	*requests.RoaRequest
}

// GetAuthorizationTokenResponse is the response struct for api GetAuthorizationToken
type GetAuthorizationTokenResponse struct {
	*responses.BaseResponse
}

// CreateGetAuthorizationTokenRequest creates a request to invoke GetAuthorizationToken API
func CreateGetAuthorizationTokenRequest() (request *GetAuthorizationTokenRequest) {
	request = &GetAuthorizationTokenRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "GetAuthorizationToken", "/tokens", "", "")
	request.Method = requests.GET
	return
}

// CreateGetAuthorizationTokenResponse creates a response to parse from GetAuthorizationToken response
func CreateGetAuthorizationTokenResponse() (response *GetAuthorizationTokenResponse) {
	response = &GetAuthorizationTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
