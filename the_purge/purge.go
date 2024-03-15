package purge

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^Folder: (\d+)`)

func Delete(filePath string) (int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	purgeFolders := make(map[string]struct{})
	purgeAll := false
	var delSize int64

	for scanner.Scan() {
		text := scanner.Text()

		if re.MatchString(text) {
			numStr := re.FindStringSubmatch(text)[1]
			_, purgeAll = purgeFolders[numStr]
			continue
		}

		if !(strings.Contains(text, "delete") || strings.Contains(text, "temporary")) && !purgeAll {
			continue
		}

		strSlice := strings.Fields(text)
		lastItem := strSlice[len(strSlice)-1]

		if strings.HasSuffix(lastItem, "]") {
			purgeFolders[strings.TrimSuffix(lastItem, "]")] = struct{}{}
			continue
		}

		num, err := strconv.Atoi(lastItem)
		if err != nil {
			return 0, err
		}
		delSize += int64(num)
	}

	return delSize, nil
}
