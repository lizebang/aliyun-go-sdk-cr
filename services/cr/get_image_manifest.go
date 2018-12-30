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

// GetImageManifest invokes the cr.GetImageManifest API synchronously
// api document: https://help.aliyun.com/api/cr/getimagemanifest.html
func (client *Client) GetImageManifest(request *GetImageManifestRequest) (response *GetImageManifestResponse, err error) {
	response = CreateGetImageManifestResponse()
	err = client.DoAction(request, response)
	return
}

// GetImageManifestWithChan invokes the cr.GetImageManifest API asynchronously
// api document: https://help.aliyun.com/api/cr/getimagemanifest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetImageManifestWithChan(request *GetImageManifestRequest) (<-chan *GetImageManifestResponse, <-chan error) {
	responseChan := make(chan *GetImageManifestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImageManifest(request)
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

// GetImageManifestWithCallback invokes the cr.GetImageManifest API asynchronously
// api document: https://help.aliyun.com/api/cr/getimagemanifest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetImageManifestWithCallback(request *GetImageManifestRequest, callback func(response *GetImageManifestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImageManifestResponse
		var err error
		defer close(result)
		response, err = client.GetImageManifest(request)
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

// GetImageManifestRequest is the request struct for api GetImageManifest
type GetImageManifestRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Tag           string `position:"Path" name:"Tag"`
}

// GetImageManifestResponse is the response struct for api GetImageManifest
type GetImageManifestResponse struct {
	*responses.BaseResponse
}

// CreateGetImageManifestRequest creates a request to invoke GetImageManifest API
func CreateGetImageManifestRequest() (request *GetImageManifestRequest) {
	request = &GetImageManifestRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "GetImageManifest", "/repos/[RepoNamespace]/[RepoName]/tags/[Tag]/manifest", "", "")
	request.Method = requests.GET
	return
}

// CreateGetImageManifestResponse creates a response to parse from GetImageManifest response
func CreateGetImageManifestResponse() (response *GetImageManifestResponse) {
	response = &GetImageManifestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
