package main

type RingBuff struct {
	data    []string
	size    int
	start   int
	counter int
}

func (q *RingBuff) PushToStart(data string) {
	if q.size != q.counter {
		q.counter++
	}
	q.start = (q.start - 1 + q.size) % q.size
	q.data[q.start] = data
}

func (q *RingBuff) PushToEnd(data string) {
	if q.size == q.counter {
		q.start = (q.start + 1) % q.size
	} else {
		q.counter++
	}
	end := (q.start + q.counter - 1) % q.size
	q.data[end] = data
}

func (q *RingBuff) GetAsQueue() []string {
	r := make([]string, q.counter)
	for i := range q.size {
		r = append(r, q.data[(i+q.start)%q.size])
	}

	return r
}
