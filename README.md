# version

Versions a git repository based on the number of commits.

## Usage

### check

Checks that the .version file is up-to-date with the current version, and returns a non-zero exit code if it isn't.

```bash
version check
```

```
Version file .version is up-to-date with version 0.0.1
```

### set

Sets the .version file to the latest version number. If the git repo is dirty, then the version number will be the number of existing commits plus one, so that when the changes are committed, the version number will be correct.

```bash
version set
```

```
0.0.2
```

### push

Pushes a tag to git. If the `--prefix` flag is used, the version will be prefixed, e.g. with a v.

```
version push
```
