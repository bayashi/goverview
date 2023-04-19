package fileinfo

import (
	"bufio"
	"os"
	"regexp"
)

// GoModInfo provides go.mod file info: specified go version
func GoModInfo(filepath string, showAll bool) (*FileInfo, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	regexp := regexp.MustCompile(`^(go\s+[0-9\.]+)[\s\t]*$`)
	s := bufio.NewScanner(file)
	tag := ""
	for s.Scan() {
		if re := regexp.FindStringSubmatch(s.Text()); re != nil {
			tag = re[1]
			break
		}
	}

	return &FileInfo{
		Tag: tag,
	}, nil
}
