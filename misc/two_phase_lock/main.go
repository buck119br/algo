package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	N = 100000
	M = 1
	L = 10000
)

func main() {
	d := new(data)
	d.init()

	for i := 0; i < M; i++ {
		w := &worker{id: i, data: d}
		go w.run()
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
}

type data struct {
	d []int
	l []sync.RWMutex
}

func (d *data) init() {
	d.d = make([]int, N)
	d.l = make([]sync.RWMutex, N)
	for i := 0; i < N; i++ {
		d.d[i] = rand.Intn(N)
	}
}

func (d *data) transaction2PL(i, j int) {
	ii := (i + 1) % N
	iii := (i + 2) % N

	// lock
	d.l[i].RLock()
	defer d.l[i].RUnlock()
	d.l[ii].RLock()
	defer d.l[ii].RUnlock()
	d.l[iii].RLock()
	defer d.l[iii].RUnlock()
	// read
	a := d.d[i]
	b := d.d[ii]
	c := d.d[iii]
	sum := a + b + c
	// write
	d.l[j].Lock()
	defer d.l[j].Unlock()
	d.d[j] = sum
}

type worker struct {
	id   int
	data *data
}

func (w *worker) run() {
	for i := 0; i < L; i++ {
		w.randomUpdate()
		fmt.Printf("w: %d, r: %d\n", w.id, i)
	}
}

func (w *worker) randomUpdate() {
	i := rand.Intn(N)
	j := rand.Intn(N)
	w.data.transaction2PL(i, j)
}
