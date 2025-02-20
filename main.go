package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/mtardy/pathgen/pkg/randpath"
)

const (
	defaultPrefix = "/tmp/pathgen"
	defaultSuffix = "bin"
)

func main() {
	// flags
	write := flag.Bool("w", false, "write the filepath on the filesystem")
	cleanup := flag.Bool("c", false, fmt.Sprintf("cleanup everything under %q", defaultPrefix))
	symlink := flag.Bool("l", true, fmt.Sprintf("create a symlink at \"%s/exe\" to the target file, used when write is enabled", defaultPrefix))
	random := flag.Bool("r", false, "generate a random string for the path")
	prefix := flag.String("p", defaultPrefix, "prefix for the random path, you will need to cleanup manually")
	binary := flag.String("b", "", "binary to put at the end of the random path, used when write is enabled")
	suffix := flag.String("s", defaultSuffix, "suffix of the random path, name of the copied binary")
	dirMaxLen := flag.Int("d", randpath.NAME_MAX, "maximum length of a directory name")
	flag.Parse()
	args := flag.Args()

	// cleanup feature
	if cleanup != nil && *cleanup {
		err := os.RemoveAll(defaultPrefix)
		if err != nil {
			panic(err)
		}
		return
	}

	// generation feature
	if len(args) < 1 {
		fmt.Fprint(os.Stderr, "please provide a length\n\n")

		fmt.Fprintf(os.Stderr, `Usage:
  %s [flags] length

Flags:
`, os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	targetLength, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	if prefix == nil || suffix == nil {
		panic("prefix or prefix is nil")
	}
	targetPath, err := randpath.Generate(*prefix, *suffix, targetLength, *random, *dirMaxLen)
	if err != nil {
		panic(err)
	}

	fmt.Println(targetPath)

	// write feature
	if write != nil && *write {
		dir, _ := path.Split(targetPath)
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			panic(err)
		}

		if binary != nil && *binary != "" {
			cpCmd := exec.Command("cp", "-f", *binary, targetPath)
			err = cpCmd.Run()
			if err != nil {
				panic(err)
			}
		}

		if symlink != nil && *symlink {
			symlink := path.Join(defaultPrefix, "exe")
			os.Remove(symlink) // try to remove the file if it exists, we ignore err
			err = os.Symlink(targetPath, symlink)
			if err != nil {
				panic(err)
			}
		}
	}
}
