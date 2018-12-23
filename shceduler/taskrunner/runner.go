package taskrunner

import "log"

type Runner struct {
	Controller controllerChan
	Error      controllerChan
	Data       dataChan
	DataSize   int
	longLived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longLived bool, d, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		DataSize:   size,
		longLived:  longLived,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			r.Close()
		}
	}()
	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				e := r.Dispatcher(r.Data)
				if e != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			} else if c == READY_TO_EXECUTE {
				e := r.Executor(r.Data)
				if e != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:
			log.Println("default ... ")
		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH //要先给个ready_to_dispatch ,不然，select 会卡住
	r.startDispatch()
}

func (r *Runner) Close() {
	close(r.Controller)
	close(r.Error)
	close(r.Data)
}
