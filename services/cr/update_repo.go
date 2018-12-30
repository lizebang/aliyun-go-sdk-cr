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

// UpdateRepo invokes the cr.UpdateRepo API synchronously
// api document: https://help.aliyun.com/api/cr/updaterepo.html
func (client *Client) UpdateRepo(request *UpdateRepoRequest) (response *UpdateRepoResponse, err error) {
	response = CreateUpdateRepoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRepoWithChan invokes the cr.UpdateRepo API asynchronously
// api document: https://help.aliyun.com/api/cr/updaterepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRepoWithChan(request *UpdateRepoRequest) (<-chan *UpdateRepoResponse, <-chan error) {
	responseChan := make(chan *UpdateRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRepo(request)
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

// UpdateRepoWithCallback invokes the cr.UpdateRepo API asynchronously
// api document: https://help.aliyun.com/api/cr/updaterepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRepoWithCallback(request *UpdateRepoRequest, callback func(response *UpdateRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRepoResponse
		var err error
		defer close(result)
		response, err = client.UpdateRepo(request)
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

// UpdateRepoRequest is the request struct for api UpdateRepo
type UpdateRepoRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

// UpdateRepoResponse is the response struct for api UpdateRepo
type UpdateRepoResponse struct {
	*responses.BaseResponse
}

// CreateUpdateRepoRequest creates a request to invoke UpdateRepo API
func CreateUpdateRepoRequest() (request *UpdateRepoRequest) {
	request = &UpdateRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "UpdateRepo", "/repos/[RepoNamespace]/[RepoName]", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateRepoResponse creates a response to parse from UpdateRepo response
func CreateUpdateRepoResponse() (response *UpdateRepoResponse) {
	response = &UpdateRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// UpdateRepoRequestBody is the request body struct for api UpdateRepo
type UpdateRepoRequestBody struct {
	Repo struct {
		Summary  string   `json:"Summary"`
		Detail   string   `json:"Detail"`
		RepoType repoType `json:"RepoType"`
	} `json:"Repo"`
}

// NewUpdateRepoRequestBody creates a request body
func NewUpdateRepoRequestBody(summary, detail string, rtype repoType) (body *UpdateRepoRequestBody, err error) {
	urrb := &UpdateRepoRequestBody{}
	urrb.Repo.Summary = summary
	urrb.Repo.Detail = detail
	urrb.Repo.RepoType = rtype
	return urrb, nil
}

// Marshal returns the JSON encoding of the request body
func (urrb *UpdateRepoRequestBody) Marshal() (body []byte, err error) {
	return json.Marshal(urrb)
}
