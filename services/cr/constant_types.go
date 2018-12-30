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

type namespaceVisibility string

const (
	// NamespaceVisibilityPublic sets the namespace public
	NamespaceVisibilityPublic namespaceVisibility = "PUBLIC"
	// NamespaceVisibilityPrivate sets the namespace private
	NamespaceVisibilityPrivate namespaceVisibility = "PRIVATE"
)

type repoType string

const (
	// RepoTypePublic sets the Repo public
	RepoTypePublic repoType = "PUBLIC"
	// RepoTypePrivate sets the Repo private
	RepoTypePrivate repoType = "PRIVATE"
)

type repoBuildType string

const (
	// RepoBuildTypeAutoBuild sets the build type to auto
	RepoBuildTypeAutoBuild repoBuildType = "AUTO_BUILD"
	// RepoBuildTypeManual sets the build type to manual
	RepoBuildTypeManual repoBuildType = "MANUAL"
)

type repoStatus string

const (
	// RepoStatusAll sets the Repo status to all
	RepoStatusAll repoStatus = "ALL"
	// RepoStatusNormal sets the Repo status to normal
	RepoStatusNormal repoStatus = "NORMAL"
)

type sourceRepoType string

const (
	// SourceRepoTypeCode sets the repository source type to Aliyun
	SourceRepoTypeCode sourceRepoType = "CODE"
	// SourceRepoTypeGitHub sets the repository source type to GitHub
	SourceRepoTypeGitHub sourceRepoType = "GITHUB"
	// SourceRepoTypeBitbucket sets the repository source type to Bitbucket
	SourceRepoTypeBitbucket sourceRepoType = "BITBUCKET"
	// SourceRepoTypeGitlab sets the repository source type to GitLab
	SourceRepoTypeGitlab sourceRepoType = "GITLAB"
)
