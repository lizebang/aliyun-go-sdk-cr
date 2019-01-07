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

// CreateRepo invokes the cr.CreateRepo API synchronously
// api document: https://help.aliyun.com/api/cr/createrepo.html
func (client *Client) CreateRepo(request *CreateRepoRequest) (response *CreateRepoResponse, err error) {
	response = CreateCreateRepoResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRepoWithChan invokes the cr.CreateRepo API asynchronously
// api document: https://help.aliyun.com/api/cr/createrepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepoWithChan(request *CreateRepoRequest) (<-chan *CreateRepoResponse, <-chan error) {
	responseChan := make(chan *CreateRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRepo(request)
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

// CreateRepoWithCallback invokes the cr.CreateRepo API asynchronously
// api document: https://help.aliyun.com/api/cr/createrepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepoWithCallback(request *CreateRepoRequest, callback func(response *CreateRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRepoResponse
		var err error
		defer close(result)
		response, err = client.CreateRepo(request)
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

// CreateRepoRequest is the request struct for api CreateRepo
type CreateRepoRequest struct {
	*requests.RoaRequest
}

// CreateRepoResponse is the response struct for api CreateRepo
type CreateRepoResponse struct {
	*responses.BaseResponse
}

// CreateCreateRepoRequest creates a request to invoke CreateRepo API
func CreateCreateRepoRequest() (request *CreateRepoRequest) {
	request = &CreateRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "CreateRepo", "/repos", "", "")
	request.Method = requests.PUT
	return
}

// CreateCreateRepoResponse creates a response to parse from CreateRepo response
func CreateCreateRepoResponse() (response *CreateRepoResponse) {
	response = &CreateRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

// CreateRepoRequestBody is the request body struct for api CreateRepo
type CreateRepoRequestBody struct {
	Repo struct {
		RepoNamespace string        `json:"RepoNamespace"`
		RepoName      string        `json:"RepoName"`
		Summary       string        `json:"Summary"`
		Detail        string        `json:"Detail"`
		RepoType      repoType      `json:"RepoType"`
		Region        string        `json:"Region,omitempty"`
		RepoBuildType repoBuildType `json:"RepoBuildType,omitempty"`
	} `json:"Repo"`
	RepoSource *CreateRepoRequestBodyRepoSource `json:"RepoSource,omitempty"`
}

// CreateRepoRequestBodyRepoSource is the part of request body struct for api CreateRepo
type CreateRepoRequestBodyRepoSource struct {
	Source struct {
		SourceRepoType      sourceRepoType `json:"SourceRepoType"`
		SourceRepoNamespace string         `json:"SourceRepoNamespace"`
		SourceRepoName      string         `json:"SourceRepoName"`
	}
	BuildConfig struct {
		IsAutoBuild    bool `json:"IsAutoBuild"`
		IsOversea      bool `json:"IsOversea"`
		IsDisableCache bool `json:"IsDisableCache"`
	} `json:"BuildConfig"`
}

// NewCreateRepoRequestBody creates a request body
func NewCreateRepoRequestBody(namespace, name, summary, detail string, rtype repoType) *CreateRepoRequestBody {
	crrb := &CreateRepoRequestBody{}
	crrb.Repo.RepoNamespace = namespace
	crrb.Repo.RepoName = name
	crrb.Repo.Summary = summary
	crrb.Repo.Detail = detail
	crrb.Repo.RepoType = rtype
	return crrb
}

// SetRepoSource sets the repository source
func (crrb *CreateRepoRequestBody) SetRepoSource(stype sourceRepoType, namespace, name string) {
	if crrb.RepoSource == nil {
		crrb.RepoSource = new(CreateRepoRequestBodyRepoSource)
	}
	crrb.RepoSource.Source.SourceRepoType = stype
	crrb.RepoSource.Source.SourceRepoNamespace = namespace
	crrb.RepoSource.Source.SourceRepoName = name
}

// SetRepoBuildConfig sets the repository build config
func (crrb *CreateRepoRequestBody) SetRepoBuildConfig(isAutoBuild, isOversea, isDisableCache bool) {
	if crrb.RepoSource == nil {
		crrb.RepoSource = new(CreateRepoRequestBodyRepoSource)
	}
	crrb.RepoSource.BuildConfig.IsAutoBuild = isAutoBuild
	crrb.RepoSource.BuildConfig.IsOversea = isOversea
	crrb.RepoSource.BuildConfig.IsDisableCache = isDisableCache
}

// Marshal returns the JSON encoding of the request body
func (crrb *CreateRepoRequestBody) Marshal() (body []byte, err error) {
	return json.Marshal(crrb)
}
