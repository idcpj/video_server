package taskrunner

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"video_server/shceduler/dbops"

	"github.com/kataras/iris/core/errors"
)

func deleteVideoFile(vid string) error {
	path := VIDEO_DIR_2 + vid
	f, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	joinpath := filepath.Join(f, path)

	e := os.Remove(joinpath)
	if e != nil || !os.IsNotExist(e) {
		log.Printf(" remove file is error:%v ", e)
		return e
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, e := dbops.ReadVideoDeleteRecord(3)
	if e != nil {
		log.Printf("video clear dispathcer error :%v", e)
		return e
	}
	if len(res) == 0 {
		return errors.New("all tasks finised")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errorMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if e := deleteVideoFile(id.(string)); e != nil {
					errorMap.Store(id, e)
					return
				}
				if e := dbops.DelVideoDeleteionRecord(id.(string)); e != nil {
					errorMap.Store(id, e)
					return
				}
			}(vid)
		default:
			break forloop
		}

	}
	//判断是否 有 错误
	errorMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
