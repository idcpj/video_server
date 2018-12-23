package dbops

import "log"

func AddVideoDeleionReocrd(vid string) error {
	stmt, e := dbConn.Prepare(`Insert Into video_del_rec(video_id) values (?)`)
	defer stmt.Close()
	if e != nil {
		return e
	}
	_, e = stmt.Exec(vid)
	if e != nil {
		log.Printf("AddVideoDeleionReocrd error %v", e)
		return e
	}

	return nil

}
