# Get the version.
version=`git tag`
# Write out the package.
cat << EOF > version_linux.go
package main

//go:generate bash ./get_version.sh
var version = "$version"
EOF