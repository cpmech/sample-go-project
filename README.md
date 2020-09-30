# Very simple example of a Go project structure

NOTE: this project structure does not work with _gopls_; set `"go.useLanguageServer": false` in VS Code.

This repository shows a simple structure for a Go project with _commands_ and _libraries_ using the **go modules** approach. Thus, we can save this project anywhere in the filesystem.

We keep the tools in the **cmd** subdirectory and the library files in the **lib** subdirectory. We need two go.mod files, one for each subdirectory.

**lib/go.mod** is:

```
module _/lib

go 1.15
```

And, as cmd imports lib, we need **cmd/go.mod**:

```
module _/cmd

go 1.15

require _/lib v0.0.0

replace _/lib => ../lib
```

The directory structure is:

```
.
├── cmd
│   ├── go.mod
│   ├── logger
│   │   └── main.go
│   └── printer
│       └── main.go
└── lib
    ├── count
    │   ├── words.go
    │   └── z_words_test.go
    ├── go.mod
    └── list
        ├── list.go
        └── z_list_test.go

6 directories, 8 files
```
