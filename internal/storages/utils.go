package storages

type isCheck interface {
	Check() bool
}

func deleteErrData[T isCheck](storage []T) {
	for i, el := range storage {
		if !el.Check() {
			drop(storage, i)
		}
	}
}

func drop[T any](ss []T, i int) {
	ss[i] = ss[len(ss)-1]
	ss = ss[:len(ss)-1]
}
