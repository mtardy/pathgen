package randpath

import (
	"strings"
	"testing"
)

func TestGenerateRandomPath(t *testing.T) {
	for i := 1; i < PATH_MAX; i++ {
		for _, prefix := range []string{"", "/tmp", "a"} {
			for _, suffix := range []string{"", "exe", "b"} {
				got, err := Generate(prefix, suffix, i, true)
				if err != nil {
					t.Log(err)
					continue
				}
				if len(got) != i {
					t.Errorf("len = %v, want %v (prefix: %q, suffix: %q) %s", len(got), i, prefix, suffix, got)
					continue
				}
				sections := strings.Split(got, "/")
				for _, section := range sections {
					if len(section) > NAME_MAX {
						t.Errorf("invalid file path name, len(%s)=%d > NAME_MAX=%d", section, len(section), NAME_MAX)
						continue
					}
				}
			}
		}
	}
}
