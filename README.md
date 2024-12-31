# goverview

<a href="https://github.com/bayashi/goverview/blob/main/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png"></a>
<a href="https://github.com/bayashi/goverview/actions"><img src="https://github.com/bayashi/goverview/workflows/main/badge.svg?_t=1681289447"/></a>
<a href="https://pkg.go.dev/github.com/bayashi/goverview"><img src="https://pkg.go.dev/badge/github.com/bayashi/goverview.svg" alt="Go Reference"></a>

`goverview` provides an overview as ASCII tree for a Golang project.

## Usage

    $ goverview ~/go/src/github.com/bayashi/goverview
    
    ┌ goverview/
    ├─┬ .github/
    │ └─┬ workflows/
    │   └── run-tests.yaml
    ├── .gitignore
    ├── LICENSE: License MIT
    ├── README.md
    ├──* arg.go: main
    ├──* builder.go: main
    ├──* cmd.go: main
    ├─┬ fileinfo/
    │ ├──* fileinfo.go: fileinfo
    │ │     Struct: FileInfo
    │ ├──* go.go: fileinfo
    │ │     Func: GoInfo
    │ ├──* gomod.go: fileinfo
    │ │     Func: GoModInfo
    │ └──* license.go: fileinfo
    │       Func: LicenseInfo
    ├── go.mod: go 1.19
    ├── go.sum
    └──* main.go: main

You can see also private stuff with `-a` option.

    $ goverview ~/go/src/github.com/bayashi/goverview -a
    
    ┌ goverview/
    ├─┬ .github/
    │ └─┬ workflows/
    │   └── run-tests.yaml
    ├── .gitignore
    ├── LICENSE: License MIT
    ├── README.md
    ├──* arg.go: main
    │     struct: options
    │     func: parseArgs
    ├──* builder.go: main
    │     struct: walkerArgs
    │     func: fromLocal, validateDirPath, buildTree, isSkipPath, walkProcess, getFileInfo
    ├──* cmd.go: main
    │     func: putErr, putUsage, putHelp
    ├─┬ fileinfo/
    │ ├──* fileinfo.go: fileinfo
    │ │     Struct: FileInfo
    │ ├──* go.go: fileinfo
    │ │     struct: organizer
    │ │     Func: GoInfo
    │ │     func: goInfoProcessIdent, buildDescriptions
    │ ├──* gomod.go: fileinfo
    │ │     Func: GoModInfo
    │ └──* license.go: fileinfo
    │       Func: LicenseInfo
    ├── go.mod: go 1.19
    ├── go.sum
    └──* main.go: main
          func: main, run

Full options:

    Usage: goverview [OPTIONS] DIR
    Options:
      -h, --help                 Show help (This message) and exit
      -t, --hide-test            Hide contents of test files
          --ignore stringArray   Ignore path to show if a given string would match
      -a, --show-all             Show all stuff
      -v, --version              Show version and build info and exit

## Installation

    go install github.com/bayashi/goverview@latest

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
