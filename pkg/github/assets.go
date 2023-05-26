package github

import (
	gh "github.com/google/go-github/v52/github"
	"strings"
)

type PlatformAsset struct {
	Id        int64
	Name      string
	Platform  string
	Integrity string
}

func MapReleasesToPlatformAssetsByVersion(releases []*gh.RepositoryRelease, platforms []string) map[string][]PlatformAsset {
	results := make(map[string][]PlatformAsset)
	for i := 0; i < len(releases); i++ {
		release := releases[i]
		assets := release.Assets
		var foundAssets []PlatformAsset
	Assets:
		for j := 0; j < len(assets); j++ {
			asset := assets[j]
			for _, p := range platforms {
				if strings.Contains(*asset.Name, p) {
					foundAssets = append(foundAssets, PlatformAsset{*asset.ID, *asset.Name, p, ""})
					continue Assets
				}
			}
		}

		if len(foundAssets) == 0 {
			continue
		}

		version := strings.TrimPrefix(*release.TagName, "v")
		results[version] = foundAssets
	}

	return results
}
