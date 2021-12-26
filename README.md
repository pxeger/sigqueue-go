# sigqueue-go

This is a small module which provides an interface to C's `sigqueue` (via the `rt_sigqueueinfo` system call) in Go,
which allows passing values when signalling processes on Linux. This may work on other POSIX operating systems, but I
won't explicitly support them.

Requires `cgo` and `libc`.

[**Documentation**](https://pkg.go.dev/github.com/pxeger/sigqueue-go)
