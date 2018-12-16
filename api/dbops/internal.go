package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

//添加 session
func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmt, e := dbConn.Prepare(`INSERT INTo sessions (session_id,TTL,login_name) values(?,?,?)`)
	defer stmt.Close()
	if e != nil {
		return e
	}
	_, e = stmt.Exec(sid, ttlstr, uname)
	if e != nil {
		return e
	}
	return nil
}

//获取 session
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmt, e := dbConn.Prepare(`SELECT TTL,login_name FROM sessions where session_id=?`)
	defer stmt.Close()
	if e != nil {
		return nil, e
	}
	e = stmt.QueryRow(sid).Scan(&ss.TTL, &ss.UserName)
	if e != nil && e != sql.ErrNoRows {
		return nil, e
	}

	return ss, nil
}

//获取所有的 session
func RetrieveAllSessions() (*sync.Map, error) {
	stmt, e := dbConn.Prepare(`SELECT * FROM sessions`)
	defer stmt.Close()
	if e != nil {
		return nil, e
	}
	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}
	m := &sync.Map{}
	for rows.Next() {
		ss := &defs.SimpleSession{}
		id := ""
		if e := rows.Scan(&id, &ss.TTL, &ss.UserName); e != nil {
			continue
			log.Printf("session id : = %v ss: %v", id, ss)
		}
		m.Store(id, ss)
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmt, e := dbConn.Prepare(`DELETE FROM sessions WHERE session_id=?`)
	defer stmt.Close()
	if e != nil {
		return e
	}
	if _, e = stmt.Exec(sid); e != nil {
		return e
	}

	return nil
}
