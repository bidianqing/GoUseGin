package queue

var (
	_queue chan func()
)

func NewQueue(capacity int) {
	_queue = make(chan func(), capacity)
}

func Queue(f func()) {
	_queue <- f
}

func Dequeue() func() {
	return <-_queue
}
