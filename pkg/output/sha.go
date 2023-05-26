package output

import (
	"crypto"
	"encoding/base64"
	"hash"
	"io"
)

func GenerateShaSum(content *io.ReadCloser, algo int) (string, error) {
	var sha hash.Hash

	switch algo {
	case 1:
		sha = crypto.SHA1.New()
	case 512:
		sha = crypto.SHA512.New()
	case 256:
		sha = crypto.SHA256.New()
	case 384:
		sha = crypto.SHA384.New()
	}

	if _, err := io.Copy(sha, *content); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sha.Sum(nil)), nil
}
