package github

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	gh "github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

const (
	defaultCount   = 30
	defaultPerPage = 30
	maxPerPage     = 100
)

type GithubAssetsClient struct {
	Owner  string
	Repo   string
	http   *http.Client
	github *gh.Client
}

func NewGithubAssetsClient(repoPath string, token string) *GithubAssetsClient {
	r := strings.Split(repoPath, "/")
	owner := r[0]
	repo := r[1]

	return &GithubAssetsClient{
		Owner:  owner,
		Repo:   repo,
		github: createGithubClent(token),
		http:   &http.Client{},
	}
}

func (gac *GithubAssetsClient) DownloadAsset(id int64) (content io.ReadCloser, err error) {
	content, url, err := gac.github.Repositories.DownloadReleaseAsset(context.Background(), gac.Owner, gac.Repo, id, gac.http)

	if err != nil || url == "" {
		return
	}

	if url != "" {
		var resp *http.Response
		resp, err = gac.http.Get(url)

		if err != nil {
			return
		}
		defer resp.Body.Close()

		content = resp.Body
	}

	return
}

func (gac *GithubAssetsClient) GetReleases(count int) (results []*gh.RepositoryRelease) {
	perPage := defaultPerPage
	if count > maxPerPage {
		perPage = maxPerPage
	}
	options := &gh.ListOptions{Page: 1, PerPage: perPage}
	for {
		ctx := context.Background()
		releases, response, err := gac.github.Repositories.ListReleases(ctx, gac.Owner, gac.Repo, options)
		options.Page = response.NextPage

		if err != nil {
			fmt.Println(err)
			break
		}

		if len(releases) == 0 {
			break
		}

		results = append(results, releases[:]...)

		if len(results) >= count {
			return results[:count]
		}

		// if the size of the current list of releases is lesser than the count per
		// page, we can assume, that this was the last page and save another
		// request that will response an empty list.
		if len(releases) < options.PerPage {
			break
		}
	}

	return
}

func createGithubClent(token string) *gh.Client {
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		client := oauth2.NewClient(context.Background(), ts)

		return gh.NewClient(client)
	}

	return gh.NewClient(nil)

}
