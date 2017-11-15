# Versioning Go Services

* Tag your git repository with a version number (ideally based on the semantic versioning specification e.g. v0.0.0).
  * `git tag -a v0.0.0 -m "First tag."`
* Push the version number to the remote repository.
  * `git push --tags`
* Generate the version.go file
  * `go generate`
    * The `version_PLATFORM.go` file will be updated with data from the tag.
* Build!
  * `go build`
  
# A better way?

There's no need to use Go Generate, you can use `ldflags` instead.

https://www.atatus.com/blog/golang-auto-build-versioning/
