# rmvcsdir - remove version control directories

This utility takes in one or more paths and recursively walks their complete
directory trees deleting Git, Bzr, SVN, and Mercurial directories along the way.

[![Build Status](https://travis-ci.org/Masterminds/rmvcsdir.svg?branch=master)](https://travis-ci.org/Masterminds/rmvcsdir) [![Build status](https://ci.appveyor.com/api/projects/status/0s5dw5yyrvboy2u9?svg=true&passingText=windows%20build%20passing&failingText=windows%20build%20failing)](https://ci.appveyor.com/project/mattfarina/rmvcsdir)

_Warning: This is a destructive operation where you likely cannot recover the
files afterwards. Do this at your own risk._

## Options/Flags

* `-dryrun`: Display locations to delete without deleting them.
* `-quiet`: Do not display output unless an error occurs.
* `-skip-check`: Skip checking if locations exist.

## Examples:

Deleting files from a single directory:
```sh
$ rmvcsdir ./vendor
Deleting: /Users/mfarina/Code/go/src/github.com/Masterminds/glide/vendor/github.com/Masterminds/semver/.git
Deleting: /Users/mfarina/Code/go/src/github.com/Masterminds/glide/vendor/github.com/Masterminds/vcs/.git
Deleting: /Users/mfarina/Code/go/src/github.com/Masterminds/glide/vendor/github.com/codegangsta/cli/.git
Deleting: /Users/mfarina/Code/go/src/github.com/Masterminds/glide/vendor/gopkg.in/yaml.v2/.git
```

Deleting files from multiple directories:
```sh
$ rmvcsdir ~/Code/myproj ~/Code/myproj2/vendor
...
```

Perform a dry run where files are not deleted but you can see what would be deleted:
```sh
$ rmvcsdir -dryrun ./vendor
Deleting: /Users/mfarina/Code/go/src/github.com/Masterminds/glide/vendor/github.com/Masterminds/semver/.git
Deleting: /Users/mfarina/Code/go/src/github.com/Masterminds/glide/vendor/github.com/Masterminds/vcs/.git
Deleting: /Users/mfarina/Code/go/src/github.com/Masterminds/glide/vendor/github.com/codegangsta/cli/.git
Deleting: /Users/mfarina/Code/go/src/github.com/Masterminds/glide/vendor/gopkg.in/yaml.v2/.git
```
