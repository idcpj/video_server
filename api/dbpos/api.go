package dbpos

import (
	"database/sql"
	"log"

	"video_server/api/defs"
	"video_server/api/utils"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	stmt, e := dbConn.Prepare("insert into users  (login_name,pwd) values(?,?)")
	if e != nil {
		return e
	}

	_, er := stmt.Exec(loginName, pwd)
	if er != nil {
		return er
	}
	defer stmt.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmt, e := dbConn.Prepare("select  pwd from  users WHERE login_name=?")
	if e != nil {
		log.Printf("%s", e)
		return "", e
	}
	var pwd string
	err := stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil || err == sql.ErrNoRows {
		return "", err
	}
	defer stmt.Close()
	return pwd, nil

}

func DeleteUser(loginName string, pwd string) error {
	stmt, e := dbConn.Prepare("delete from users WHERE login_name=? AND pwd=?")
	if e != nil {
		log.Printf("deleteuser error %s", e)
		return e
	}
	_, err := stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

//创建视频列表
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	id, e := utils.NewUuid()
	if e != nil {
		return nil, err
	}
	t := time.Now().Format("Jan 02 2006,15:04:05")

	stmt, err := dbConn.Prepare(`insert into video_info (id,author_id,name,display_ctime) VALUES (?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, error := stmt.Exec(id, aid, name, t)
	if error != nil {
		return nil, nil
	}
	defer stmt.Close()
	info := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: t}
	return info, nil
}
func GetVideo(id string) (*defs.VideoInfo, error) {
	stmt, e := dbConn.Prepare("select id,author_id,name,display_ctime FROM video_info WHERE id=?")
	if e != nil {
		return nil, e
	}
	video := &defs.VideoInfo{}
	er := stmt.QueryRow(id).Scan(&video.Id, &video.AuthorId, &video.Name, &video.DisplayCtime) //todo
	if er != nil {
		return nil, er
	}
	defer stmt.Close()
	return video, nil
}
func DeleteVideo(id string) error {
	stmt, e := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if e != nil {
		return e
	}
	defer stmt.Close()
	_, er := stmt.Exec(id)
	if er != nil {
		return er
	}
	return nil
}
