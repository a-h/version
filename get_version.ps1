# Get the version.
$version = &("C:\Program Files\Git\cmd\git.exe") tag | Select-Object -Last 1
$version = $version.Trim()
# Write out the package.
@"
package main

//go:generate powershell .\get_version.ps1
var version = "$version"
"@ | Out-File -Encoding ASCII -FilePath version_windows.go