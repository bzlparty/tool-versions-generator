package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mgred/tool-versions-generator/pkg/github"
	"github.com/mgred/tool-versions-generator/pkg/output"
)

var (
	repo      string
	algo      int
	count     int
	token     string
	platforms []string
	write     string
)

func RunApp() error {
	client := github.NewGithubAssetsClient(repo, token)
	releases, err := client.GetReleases(count)
	platformAssets := github.MapReleasesToPlatformAssetsByVersion(releases, platforms)

	if err != nil {
		return err
	}

	for _, assets := range platformAssets {
		for i := 0; i < len(assets); i++ {
			asset := &assets[i]
			content, err := client.DownloadAsset(asset.Id)

			if err != nil {
				return err
			}

			shasum, err := output.GenerateShaSum(&content, algo)

			if err != nil {
				return err
			}

			asset.Integrity = fmt.Sprintf("sha%v-%v", algo, shasum)
		}
	}

	output.NewOutput(output.OutputData{ResultMap: platformAssets, Repo: repo}).Write(write)

	return nil
}

func init() {
	algos := joinSupportedAlgos()
	flag.IntVar(&algo, "algo", 384, "shasum algorithm")
	flag.IntVar(&algo, "a", 384, "shasum algorithm (shorthand)")
	flag.IntVar(&count, "count", 30, "number of releases")
	flag.IntVar(&count, "c", 30, "number of releases (shorthand)")
	flag.Func("platform", "Platform to match assets with", appendPlatformFlag(&platforms))
	flag.Func("p", "Platform to match assets with", appendPlatformFlag(&platforms))
	flag.StringVar(&repo, "repo", "", "repo path, like owner/repo")
	flag.StringVar(&repo, "r", "", "repo path, like owner/repo (shorthand)")
	flag.StringVar(&token, "token", "", "Github token (optional)")
	flag.StringVar(&token, "t", "", "Github token (shorthand)")
	flag.StringVar(&write, "w", "", "Ouput file path (optional)")
	flag.StringVar(&write, "write", "", "Ouput file path (shorthand)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, `
Usage: %s -h | -r -p [-p=p0,...,pn] [-a] [-c] [-t] [-w]

Options:
      -r, --repo        repository to fetch assets from
      -p, --platform    platform part to match
      -a, --algo        algorithm for integrity, supported: %s (optional, default: 384)
      -c, --count       count of releases (optional)
      -t, --token       github token to bypass rate limits (optional)
      -w, --write       file to write (optional, default: stdout)
      -h, --help
`, os.Args[0], algos)
	}
	flag.Parse()
	err := validateFlags()

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
