package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	Elem     []int
	Capacity int
	Front    int
	Rear     int
	Lock     sync.Locker
	Cond     *sync.Cond
}

func New() *Queue {
	theQueue := &Queue{}
	theQueue.Capacity = 10
	theQueue.Elem = make([]int, 10)
	theQueue.Front, theQueue.Rear = 0, 0
	theQueue.Lock = &sync.Mutex{}
	theQueue.Cond = sync.NewCond(theQueue.Lock)
	return theQueue
}

func (self *Queue) Put(e int) {
	self.Cond.L.Lock()
	// the Queue is full, Producer waits here
	// note that we use for not if to test the condition
	for self.Full() {
		self.Cond.Wait()
	}

	self.Elem[self.Rear] = e
	self.Rear = (self.Rear + 1) % self.Capacity
	self.Cond.Signal()
	defer self.Cond.L.Unlock()
}

func (self *Queue) Get() int {
	self.Cond.L.Lock()
	// the Queue is empty, Consumer waits here
	// note that we use for not if to test the condition
	for self.Empty() {
		self.Cond.Wait()
	}

	p := self.Elem[self.Front]
	self.Front = (self.Front + 1) % self.Capacity
	self.Cond.Signal()
	defer self.Cond.L.Unlock()
	return p
}

func (self *Queue) Empty() bool {
	if self.Front == self.Rear {
		return true
	}
	return false
}

func (self *Queue) Full() bool {
	if ((self.Rear + 1) % self.Capacity) == self.Front {
		return true
	}
	return false
}

func main() {
	theQueue := New()

	// producer puts
	go func() {
		for i := 1; i <= 100; i++ {
			//time.Sleep(100 * time.Millisecond)
			theQueue.Put(i)
			fmt.Println("Bob puts ", i)
		}
	}()

	// consumer gets
	for i := 1; i <= 100; i++ {
		//time.Sleep(100 * time.Millisecond)
		p := theQueue.Get()
		fmt.Println("Alice gets : ", p)
	}
}
