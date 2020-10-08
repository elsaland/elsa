// Package core lists all the OPS in elsa
// Will be updated as an when new OPS are added
package core

// FileSystem ops
const (
	FSWrite     = 1
	FSRead      = 2
	FSExists    = 3
	FSDirExists = 4
	FSCwd       = 5
	Serve       = 25
	FSStat      = 6
	FSRemove    = 7
	FSMkdir     = 9
)

// console binding ops
const (
	Log = 10
)

// plugin ops
const (
	Plugin = 15
)

// fetch ops
const (
	Fetch = 20
)

// env ops
const (
	Env = 11
)
