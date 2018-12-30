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

// DeleteImage invokes the cr.DeleteImage API synchronously
// api document: https://help.aliyun.com/api/cr/delete-image-version.html
func (client *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
	response = CreateDeleteImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteImageWithChan invokes the cr.DeleteImage API asynchronously
// api document: https://help.aliyun.com/api/cr/delete-image-version.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteImageWithChan(request *DeleteImageRequest) (<-chan *DeleteImageResponse, <-chan error) {
	responseChan := make(chan *DeleteImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteImage(request)
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

// DeleteImageWithCallback invokes the cr.DeleteImage API asynchronously
// api document: https://help.aliyun.com/api/cr/delete-image-version.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteImageWithCallback(request *DeleteImageRequest, callback func(response *DeleteImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteImageResponse
		var err error
		defer close(result)
		response, err = client.DeleteImage(request)
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

// DeleteImageRequest is the request struct for api DeleteImage
type DeleteImageRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Tag           string `position:"Path" name:"Tag"`
}

// DeleteImageResponse is the response struct for api DeleteImage
type DeleteImageResponse struct {
	*responses.BaseResponse
}

// CreateDeleteImageRequest creates a request to invoke DeleteImage API
func CreateDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "DeleteImage", "/repos/[RepoNamespace]/[RepoName]/tags/[Tag]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteImageResponse creates a response to parse from DeleteImage response
func CreateDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
