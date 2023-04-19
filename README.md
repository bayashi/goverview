# goverview

`goverview` provides an overview of Golang project

Don't try big project yet :D

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

    Version v0.0.1
    Usage: goverview [OPTIONS] FILE
    Options:
      -h, --help                 Show help (This message) and exit
          --ignore stringArray   Ignore path to show if a given string would match
      -a, --show-all             Show all stuff

## Installation

    go install github.com/bayashi/goverview

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
