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

Pushes a tag to git.

The tag prefix is inferred from existing tags: if the latest tag is prefixed (e.g. `v0.0.2`), the new tag uses the same prefix. When the repository has no tags, the prefix defaults to `v` to match Go's module versioning convention.

This default changed: previous versions pushed an unprefixed tag (e.g. `0.0.1`) when no `--prefix` was given. To push unprefixed tags, pass `--prefix ""`.

The `--prefix` flag overrides the inferred prefix. If the chosen prefix differs from the prefix used by existing tags, the push fails to prevent an inconsistent tag history. Use `--force` to push anyway.

```
version push
```

```
version push --prefix "" --force
```

## Tasks

### version

```bash
go run . set
```

### push

Push a new release tag.

```bash
go run . push
```
