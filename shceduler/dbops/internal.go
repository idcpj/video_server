package dbops

import (
	"log"
)

//读
func ReadVideoDeleteRecord(count int) ([]string, error) {
	stmt, e := dbConn.Prepare("select video_id from video_del_rec limit ?")
	defer stmt.Close()
	var ids []string
	if e != nil {
		log.Println("ReadVideoDeleteRecord", e)
		return ids, e
	}
	rows, e := stmt.Query(count)
	if e != nil {
		log.Printf("query ReadVideoDeleteRecord error %v", e)
		return ids, e
	}
	for rows.Next() {
		var id string
		if e := rows.Scan(&id); e != nil {
			log.Println("ReadVideoDeleteRecord next", e)
			return ids, e
		}
		ids = append(ids, id)
	}
	return ids, nil

}

//删
func DelVideoDeleteionRecord(vid string) error {
	stmt, e := dbConn.Prepare("Delete from video_del_rec where video_id=?")
	defer stmt.Close()
	if e != nil {
		return e
	}
	_, e = stmt.Exec(vid)
	if e != nil {
		return e
	}
	return nil

}
