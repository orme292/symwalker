# SymWalker

[![Go Reference](https://pkg.go.dev/badge/github.com/orme292/symwalker.svg)](https://pkg.go.dev/github.com/orme292/symwalker)

SymWalker is a directory tree walker with symlink loop protection. It works by building a
separate history for each sub-directory branch (internally called lineHistory) underneath
a given starting path. Each branch history is checked when evaluating new symlinks. The
return results include three separate objects: directories, regular files, and all
other directory entries.

## Import this Module

```shell
go get github.com/orme292/symwalker@latest
```

```go
import (
    sw "github.com/orme292/symwalker"
)
```

## Usage

SymWalker uses a functional options pattern to apply an optional configuration.

### Build Optional Configuration

```go
func main() {

opts := sw.NewSymConf("/home/andrew/",
        sw.WithFollowedSymLinks(),
        sw.WithLogging(), 
    )

}
```

### Options

`WithDepth(n)`: Used to tell SymWalker how deep to walk from the starting path. `0` is infinite while `1` would walk
only the provided starting path.

`WithFileData()`: SymWalker will populate the FileObj field in each DirEntry with file information.

`WithFollowedSymLinks()`: Used to tell SymWalker that it SHOULD evaluate and/or walk symlinks.

`WithLogging()`: Tells SymWalker to emit logs. SymWalker uses `log.Printf` to output messages.

`WithoutFiles()`: Tells SymWalker to skip processing non-directory entries; Results.Files will be empty.

### Call SymWalk

The `SymWalk` function is the starting point for SymWalker. Call the function by passing the starting path and an
options configuration object to begin the directory walk.

```go
results, err := sw.SymWalker("root_path", opts)
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
Path    string
FileObj objectify.FileObj
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

## `FileObj` Field

For the `FileObj` field, see the [Objectify](https://github.com/orme292/objectify) library
[documentation](https://pkg.go.dev/github.com/orme292/objectify) for information.

## Example

Here's a full example on how to use SymWalker. The easiest way to work with Results is to loop over
each DirEntries type with a range loop (see below).

```go
package main

import (
    "fmt"
    "os"
    
    sw "github.com/orme292/symwalker"
)

func main() {

    conf := sw.NewSymConf("/home/andrew/",
        sw.WithFollowedSymLinks(),
        sw.WithLogging(),
    )

    results, err := sw.SymWalker(conf)
    if err != nil {
        fmt.Printf("Error occurred: %s", err.Error())
        os.Exit(1)
    }
    
    for _, dir := range results.Dirs {
        fmt.Printf("Dir: %s\n", dir.Path)
    }
 

    for _, file := range results.Files {
        fmt.Printf("File: %s\n", file.Path)
    }
    
    os.Exit(0)
}

```
