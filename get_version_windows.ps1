# Get the version.
$version = &("C:\Program Files\Git\cmd\git.exe") describe --tags --long
$version = $version.Trim()
# Write out the package.
@"
package main

//go:generate powershell .\get_version_windows.ps1
var version = "$version"
"@ | Out-File -Encoding ASCII -FilePath version_windows.go