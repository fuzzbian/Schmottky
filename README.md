# Schmottky
Playground for fractal fun loosely based on the ideas presented in the book [Indra's Pearls](https://en.wikipedia.org/wiki/Indra%27s_Pearls_(book)).

# Getting started
## Dependencies
* [Golang](https://go.dev/doc/install)
* Python3
     * sympy
     * numpy
     * matplotlib
## Build from src

```bash
# build
go mod init Schmottky
go mod tidy
go build
# run
./Schmottky
```

# TODO
* ~~add sender/receiver for raw data output~~
* ~~add receiver for matplotlib~~
* add Grandma's recipes
    * ~~Special Parabolic~~
* add UI
    * CLI
    * GUI
* ~~add P/Q calculation~~
* ~~add color scripting + optional lev channel in dfs~~
* add Lua-Api to costomize coloring via external scripts
    * ~~fractal coloring~~
        * very slow! :<
        * make optional
        * better add presets for internal color calculation
    * background
* add venv for python scripts
* add installation script
* add lines operation for Img
* add calculation for optical eps
* add algo for cropped Image
* investigate special words algo
* maybe add Dockerfile
* maybe cache for pq/mu calculation
