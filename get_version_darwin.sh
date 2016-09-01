# Get the version.
version=`git tag -l "v*" | tail -n 1`
# Write out the package.
cat << EOF > version_darwin.go
package main

//go:generate bash ./get_version_darwin.sh
var version = "$version"
EOF