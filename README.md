# grid

[![Build Status](https://travis-ci.org/eleniums/grid.svg?branch=master)](https://travis-ci.org/eleniums/grid) [![Go Report Card](https://goreportcard.com/badge/github.com/eleniums/grid)](https://goreportcard.com/report/github.com/eleniums/grid) [![GoDoc](https://godoc.org/github.com/eleniums/grid?status.svg)](https://godoc.org/github.com/eleniums/grid) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/eleniums/grid/blob/master/LICENSE)

2D grid of near infinite size containing values of any type.

## Installation

```
go get -u github.com/eleniums/grid
```

## Example

```go
// create a new grid
grid := NewGrid()

// add a value to the grid at (x: 1, y: 2)
grid.Add(1, 2, value)

// retrieve a value from the grid at (x: 1, y: 2)
value, ok := grid.Retrieve(1, 2)
if ok {
    // do something with value
}

// delete a value from the grid at (x: 1, y: 2)
grid.Delete(1, 2)
```
