package fileinfo

import (
	"os"

	"github.com/google/licenseclassifier/v2/assets"
)

// LicenseInfo provides LICENSE file info
func LicenseInfo(args *Args) (*FileInfo, error) {
	file, err := os.Open(args.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	c, err := assets.DefaultClassifier()
	if err != nil {
		return nil, err
	}
	res, err := c.MatchFrom(file)
	if err != nil {
		return nil, err
	}
	if len(res.Matches) == 0 {
		return &FileInfo{}, nil
	}

	license := ""
	for _, m := range res.Matches {
		license = m.MatchType + " " + m.Name
	}

	return &FileInfo{
		Tag: license,
	}, nil
}
