package main

import (
	"fmt"
	"log"
)

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func (c ConnLimiter) String() string {
	return fmt.Sprintf("c.concurrentConn = %v bucket = %v", c.concurrentConn, c.bucket)
}

// POST http://127.0.0.1:8000/user/cpj HTTP/1.1
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Println("Readched thr rate limitation", len(cl.bucket), " ", cl.concurrentConn)
		return false
	}
	cl.bucket <- 1

	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Println("New Connection coming ", c)
}
