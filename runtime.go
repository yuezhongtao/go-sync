package go_sync

import _ "unsafe"

const (
	active_spin     = 4
	active_spin_cnt = 30
	passive_spin    = 1
)

//go:linktime ncpu runtime.ncpu
var ncpu int32

//go:linkname gomaxprocs runtime.gomaxprocs
var gomaxprocs int32

//var sched schedt
//
////go:linkname getg runtime.getg
//func getg() *g
//
//func sync_runtime_canSpin(i int) bool {
//	// sync.Mutex is cooperative, so we are conservative with spinning.
//	// Spin only few times and only if running on a multicore machine and
//	// GOMAXPROCS>1 and there is at least one other running P and local runq is empty.
//	// As opposed to runtime mutex we don't do passive spinning here,
//	// because there can be work on global runq or on other Ps.
//	if i >= active_spin || ncpu <= 1 || gomaxprocs <= int32(sched.npidle+sched.nmspinning)+1 {
//		return false
//	}
//	if p := getg().m.p.ptr(); !runqempty(p) {
//		return false
//	}
//	return true
//}

//go:linkname runtime_canSpin sync.runtime_canSpin
func runtime_canSpin(i int) bool

//go:linkname runtime_doSpin sync.runtime_doSpin
func runtime_doSpin()

//go:linkname runtime_nanotime runtime.nanotime
func runtime_nanotime() int64

//go:linkname throw runtime.throw
func throw(string)

//go:linkname fatal runtime.fatal
func fatal(string)

//go:linkname runtime_SemacquireMutex sync.runtime_SemacquireMutex
func runtime_SemacquireMutex(s *uint32, lifo bool, skipframes int)

//go:linkname runtime_Semrelease runtime.semrelease1
func runtime_Semrelease(addr *uint32, handoff bool, skipframes int)
