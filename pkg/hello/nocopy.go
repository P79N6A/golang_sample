package hello

// TypeNotPermitCopy 不允许拷贝
// 将一个实现了sync.Locker的结构体内嵌后, go vet会自动检测拷贝
type TypeNotPermitCopy struct {
	noCopy
	Value string
}

// noCopy may be embedded into structs which must not be copied
// after the first use.
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// Copy ...
// func Copy(t TypeNotPermitCopy) {

// }
