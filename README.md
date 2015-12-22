#RepoFiles

Simple library to recursively list files in a GitHub repo and collect the file contents for each. Makes use of the [`trees`](https://developer.github.com/v3/git/trees/) endpoint.

## Usage

```go
package main

import (
  "fmt"

  "github.com/charlieegan3/repofiles"
)

func main() {
    repo := repofiles.NewRepo("username", "repoName", "master or SHA")
    repo.List() //collect the list files
    file := repo.Files() //request the contents of each file

    fmt.Println(file)
}

```

Output:

```
[{file1.ext file1content} {file2.ext file1content} ...]
```
