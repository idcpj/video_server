package dbops

import (
	"database/sql"
	"log"

	"video_server/api/defs"
	"video_server/api/utils"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName, pwd string) error {
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

func GetUserCredential(loginName string) (*defs.UserCredential, error) {
	stmt, e := dbConn.Prepare("select  id,pwd from  users WHERE login_name=?")
	if e != nil {
		log.Printf("%s", e)
		return nil, e
	}
	user := &defs.UserCredential{}
	err := stmt.QueryRow(loginName).Scan(&user.Id, &user.Pwd)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	defer stmt.Close()
	user.UserName = loginName
	return user, nil

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

	stmt, e := dbConn.Prepare(`insert into video_info (id,author_id,name,display_ctime) VALUES (?,?,?,?)`)
	if e != nil {
		return nil, e
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
	er := stmt.QueryRow(id).Scan(&video.Id, &video.AuthorId, &video.Name, &video.DisplayCtime)
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

func AddNewComment(vid string, aid int, content string) error {
	uuid, e := utils.NewUuid()
	if e != nil {
		return e
	}
	stmt, er := dbConn.Prepare(`INSERT into comments (id,video_id,author_Id,content,time)VALUES (?,?,?,?,?)`)
	defer stmt.Close()
	if er != nil {
		return er
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	_, err := stmt.Exec(uuid, vid, aid, content, t)
	if err != nil {
		return err
	}
	return nil

}

func ListComments(vid string, from, to string) ([]*defs.Comment, error) {
	stmt, e := dbConn.Prepare(`SELECT comments.id,users.login_name,comments.content 
								from comments inner JOIN users on comments.author_id = users.id 
								WHERE  comments.video_id=? AND comments.time > ? AND comments.time <=  ?`)
	defer stmt.Close()
	if e != nil {
		return nil, e
	}
	rows, er := stmt.Query(vid, from, to)
	if er != nil {
		return nil, er
	}
	var res []*defs.Comment
	for rows.Next() {
		var id, name, content string
		if er := rows.Scan(&id, &name, &content); er != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, AuthorName: name, Content: content}
		res = append(res, c)
	}

	return res, nil
}
