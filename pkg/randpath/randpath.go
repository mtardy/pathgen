package randpath

import (
	"fmt"
	"math/rand"
	"path"
)

const (
	NAME_MAX = 255
	PATH_MAX = 4096
)

func selectChar(random bool) byte {
	if random {
		return byte('a' + rand.Intn('z'-'a'+1))
	}
	return 'a'
}

func Generate(prefix, suffix string, length int, random bool) (string, error) {
	if length == 0 {
		return "", nil
	}

	starterPath := path.Join(prefix, suffix)
	if len(starterPath) == length {
		return starterPath, nil
	} else if len(starterPath) > length {
		return "", fmt.Errorf("len(path.Join(prefix, suffix)) = len(%s) = %d > %d", starterPath, len(starterPath), length)
	}

	if length == len(starterPath)+1 {
		return "", fmt.Errorf("can't fit enough \"/\" for desired length, len(path.Join(prefix,suffix)) = len(%s) = %d, and want length of %d", starterPath, len(starterPath), length)
	}

	slashBeforeSuffix := 1
	if suffix == "" && prefix == "" {
		slashBeforeSuffix = 0
	}

	randomPart := make([]byte, 0, PATH_MAX)
	for currentLength, randomSectionLength := 0, 0; currentLength < length-slashBeforeSuffix; currentLength = len(randomPart) + len(starterPath) {
		var char byte
		if randomSectionLength != 0 && randomSectionLength%NAME_MAX == 0 {
			char = '/'
			randomSectionLength = 0
		} else {
			char = selectChar(random)
			randomSectionLength++
		}
		randomPart = append(randomPart, char)
	}

	if len(randomPart) > 2 {
		if randomPart[len(randomPart)-1] == '/' {
			randomPart[len(randomPart)-2] = '/'
			randomPart[len(randomPart)-1] = selectChar(random)
		}
	}

	return path.Join(prefix, string(randomPart), suffix), nil
}
