package dbops

import (
	"database/sql"
	"testing"
	"time"
	"video_server/api/utils"
)

var (
	video_id   string
	user_id    int
	session_id string
	user_name  = "cpj"
	pwd        = "123"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
	dbConn.Exec("truncate video_info")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	//clearTables()
}

//按顺序执行
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("delete", testDeleteUser)
	t.Run("reget", testReget)
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("add", testAddNewVideo)
	t.Run("get", testGetVideo)
	t.Run("delete", testDeleteVideo)
	t.Run("reget", reGetVideo)
}
func TestCommentWorkFlow(t *testing.T) {
	clearTables()
	video_id = ""
	t.Run("create_user", testAddUserCredential)
	t.Run("get_user", testGetUserCredential)
	t.Run("create_video", testAddNewVideo)
	t.Run("create_comment", testAddNewComment)
	t.Run("create_comment", testAddNewComment)
	t.Run("list_comment", testListComments)
}
func TestSessionWorkFlow(t *testing.T) {
	clearTables()
	session_id = ""
	t.Run("create_session", testInsertSession)
	t.Run("get_session", testRetrieveSession)
	t.Run("get_all_session", testRetrieveAllSessions)
	//t.Run("delete_session", TestDeleteSession)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential(user_name, pwd)
	if err != nil {
		t.Errorf("error of Adduser %v", err)
	}
}
func testGetUserCredential(t *testing.T) {
	user, err := GetUserCredential(user_name)
	if user == nil || err != nil {
		t.Errorf("error of GetUSer %v", err)
	}
	user_id = user.Id

}
func testDeleteUser(t *testing.T) {
	err := DeleteUser(user_name, pwd)
	if err != nil {
		t.Errorf("error deleteUser %v", err)
	}
}

func testReget(t *testing.T) {
	user, err := GetUserCredential(user_name)
	if user != nil || err != sql.ErrNoRows {
		t.Errorf("error of GetUSer %v", err)
	}
}

func testAddNewVideo(t *testing.T) {
	info, e := AddNewVideo(1, "test")
	if info == nil || e != nil {
		t.Errorf("info : %v, error:%v", info, e)
	}
	video_id = info.Id
}

func testGetVideo(t *testing.T) {
	info, e := GetVideo(video_id)
	if info == nil || e != nil {
		t.Errorf("info : %v , error :%v", info, e)
	}
}

func testDeleteVideo(t *testing.T) {
	e := DeleteVideo(video_id)
	if e != nil {
		t.Errorf("error : %v", e)
	}
}
func reGetVideo(t *testing.T) {
	info, e := GetVideo(video_id)
	if info != nil || e != sql.ErrNoRows {
		t.Errorf("info : %v , error :%v", info, e)
	}
}

func testAddNewComment(t *testing.T) {
	err := AddNewComment(video_id, user_id, "test content")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func testListComments(t *testing.T) {
	comments, e := ListComments(video_id, time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"), time.Now().AddDate(0, 0, +1).Format("2006-01-02 15:04:05"))
	if comments == nil || e != nil {
		t.Errorf("%v", e)
	}
}

//创建两个 session
func testInsertSession(t *testing.T) {
	session_id, _ = utils.NewUuid()
	ct := time.Now().Unix()
	ttl := ct + 30*60
	err := InsertSession(session_id, ttl, user_name)
	if err != nil {
		t.Errorf("%v", err)
	}
	session_id, _ = utils.NewUuid()
	err = InsertSession(session_id, ttl, user_name)
	if err != nil {
		t.Errorf("%v", err)
	}
}
func testRetrieveSession(t *testing.T) {
	session, e := RetrieveSession(session_id)
	if session == nil || e != nil {
		t.Errorf("%s", e)
	}
}
func testRetrieveAllSessions(t *testing.T) {
	syncMap, e := RetrieveAllSessions()

	if e != nil {
		t.Errorf("%v", e)
	}
	value, ok := syncMap.Load(session_id)
	if !ok {
		t.Errorf("获取 session 失败 session_id=%v  value=%v", session_id, value)
	}

}

func TestDeleteSession(t *testing.T) {
	e := DeleteSession(session_id)
	if e != nil {
		t.Errorf("%v", e)
	}
}
