package sigqueue

// #include <errno.h>
// #include <string.h>
// #include <signal.h>
// char *get_errno() { return strerror(errno); }
// int sigqueue_int(int pid, int signal, int value) {
// union sigval sigval;
// sigval.sival_int = value;
// return sigqueue(pid, signal, sigval);
// }
import "C"

type Error struct {
	error
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

func Sigqueue(pid, signal, value int) error {
	ret := C.sigqueue_int(C.int(pid), C.int(signal), C.int(value))
	switch ret {
	case 0:
		// success
		return nil
	case -1:
		return &Error{msg: C.GoString(C.get_errno())}
	default:
		panic(ret) // unexpected return value from sigqueue
	}
}
