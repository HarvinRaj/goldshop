//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package neon

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __skip_one_fast_entry__() uintptr

var (
	_subr__skip_one_fast uintptr = __skip_one_fast_entry__() + 32
)

const (
	_stack__skip_one_fast = 192
)

var (
	_ = _subr__skip_one_fast
)

const (
	_ = _stack__skip_one_fast
)
