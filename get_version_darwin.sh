# Get the version.
version=`git describe --tags --long`
# Write out the package.
cat << EOF > version_darwin.go
package main

//go:generate bash ./get_version_darwin.sh
var version = "$version"
EOF