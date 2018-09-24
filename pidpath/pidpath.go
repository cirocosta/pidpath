package pidpath

type Resolver interface {
	PidPath(pid uint64) (path string, err error)
}
