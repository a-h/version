# Versioning Go Services

* Tag your git repository with a version number (ideally based on the semantic versioning specification e.g. v0.0.0).
 * `git tag -a v0.0.0 -m "First tag."`
* Get your build server to run `go generate`
 * The `version_PLATFORM.go` file will be updated with data from the tag.
* Build!