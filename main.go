package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var skipFoundCheck bool
var quiet bool
var dryrun bool

var help = `rmvcsdir - remove version control directories

The version control directories for Git, SVN, Bzr, and Mercurial are recursively
removed from the filesystem. For example:

    $ rmvcsdir ./vendor
    Deleting: ./vendor/github.com/Masterminds/semver/.git
    Deleting: ./vendor/github.com/Masterminds/vcs/.git
    Deleting: ./vendor/github.com/codegangsta/cli/.git
    Deleting: ./vendor/gopkg.in/yaml.v2/.git

One or more directories can be passed in. For example:

    $ rmvcsdir ~/Code/myproj ~/Code/myproj2/vendor

Options:`

func init() {
	flag.BoolVar(&skipFoundCheck, "skip-check", false, "Skip checking if locations exist.")
	flag.BoolVar(&quiet, "quiet", false, "Do not display output unless an error occurs.")
	flag.BoolVar(&dryrun, "dryrun", false, "Display locations to delete without deleting them.")
}

func main() {
	flag.Parse()

	args := flag.Args()

	// If nothing was passed in display help
	if len(args) == 0 {
		fmt.Println(help)
		flag.PrintDefaults()
		return
	}

	if !skipFoundCheck {
		// Make sure the passed in locations exist before proceeding.
		found := true
		for _, dir := range args {
			_, err := os.Stat(dir)
			if err != nil {
				fmt.Println("Unable to find directory:", dir)
				found = false
			}
		}
		if !found {
			os.Exit(1)
		}
	}

	for _, dir := range args {
		err := filepath.Walk(dir, handler)
		if err != nil {
			fmt.Printf("Error walking %s: %s", dir, err)
			os.Exit(2)
		}
	}
}

func handler(path string, info os.FileInfo, err error) error {

	name := info.Name()
	if name == ".git" || name == ".bzr" || name == ".svn" || name == ".hg" {
		if _, err := os.Stat(path); err == nil {
			if info.IsDir() {
				if !quiet || dryrun {
					fmt.Println("Deleting:", path)
				}
				if !dryrun {
					err := os.RemoveAll(path)
					return err
				}
				return nil
			}
		}
	}
	return nil
}
