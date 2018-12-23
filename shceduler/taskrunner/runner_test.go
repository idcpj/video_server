package taskrunner

import (
	"log"
	"testing"
	"time"
)

func TestRuner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("dispathcher sent: %d", i)
		}
		return nil
	}
	e := func(dc dataChan) error {
	forLoop:
		for {
			select {
			case d := <-dc:
				log.Printf("Execture received %v", d)
			default:
				break forLoop
			}
		}
		//return errors.New("exectuer error")
		return nil
	}
	runner := NewRunner(30, false, d, e)
	go runner.StartAll()
	time.Sleep(1 * time.Second)
}
