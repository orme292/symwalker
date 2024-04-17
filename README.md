# SymWalker

---
[![Go Reference](https://pkg.go.dev/badge/github.com/orme292/symwalker@v0.1.6.svg)](https://pkg.go.dev/github.com/orme292/symwalker@v0.1.6)

SymWalker is a directory tree walker with symlink loop protection. It works by building a
separate history for each sub-directory branch (internally called lineHistory) underneath
a given starting path. Each branch history is checked when evaluating new symlinks. The
return results include three separate objects: directories, regular files, and all
other directory entries.

## Import this Module

---

```shell
go get github.com/orme292/symwalker@v0.1.6
```

```go
import (
"github.com/orme292/symwalker"
)
```

## Usage

---
Symwalker uses a functional options pattern to create a configuration object.

### Build Configuration

```go
func main() {

conf := NewSymConf(
WithStartPath("/home/andrew/"),
WithFollowedSymLinks(),
WithLogging(),
)

}
```

### Options

`WithStartPath("path")`: required in order to specify the path to be walked. This has to be
a path to a directory, not a file or a symlink.

`WithFollowedSymLinks()`: Used to tell SymWalker that it SHOULD evaluate and/or walk symlinks.

`WithLogging()`: Tells SymWalker to emit logs. SymWalker uses `log.Printf` to output messages.

### Call SymWalk

The `SymWalk` function is the starting point for SymWalker. Call the function and pass a SymConf
configuration object to it to begin the directory walk.

```go
results, err := SymWalker(conf)
```

### The *Results* Type Struct

`SymWalker` returns `Results` and an error. The Results struct has three exported fields:

```go
type Results struct {
Dirs   DirEntries
Files  DirEntries
Others DirEntries
}
```

A `DirEntries` is a slice of `DirEntry` structs, and that has one exported field:

```go
type DirEntries []DirEntry
```

```go
type DirEntry struct {
Path   string
}
```

Each `Results.Dirs` DirEntry will have a `Path` value set to a walked directory. If following symlinks, some
paths could be for a symlink with a directory target.

Each `Results.Files` DirEntry will have a `Path` value set to a regular file. If following symlinks, some paths
could be for a symlink with a regular file target.

Each `Results.Others` DirEntry could have a `Path` value set to a directory entry that is unknown.

See documentation for the [os.FileMode](https://pkg.go.dev/os#FileMode) type. `Results.Dirs` are paths which
match `ModeDir` (os.FileMode *or* fs.FileMode). `Results.Others` files are paths which match any type other
than `ModeDir`. `Results.Files` are paths which match no `os.FileMode` type.

---

## Working with Results

You can work with Results using a ranged loop.

```go
for _, dir := range Results.Dirs {
fmt.Printf("Dir: %s\n", dir.Path)
}
// Dir: /home/andrew/documents
// Dir: /home/andrew/pictures
// Dir: /home/andrew/work
// Dir: /home/andrew/work/january
// Dir: /home/andrew/work/january/meetings
// ....

for ), file := range Results.Files {
fmt.Printf("File: %s\n", file.Path)
}
// File: /home/andrew/important.doc
// File: /home/andrew/documents/taxes.pdf
// File: /home/andrew/documents/passwords.txt 😉
```