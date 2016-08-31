# Get the version.
version=`git tag | tail --lines 1`
# Write out the package.
cat << EOF > version_linux.go
package main

//go:generate bash ./get_version.sh
var version = "$version"
EOF