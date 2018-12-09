package dbpos

import(
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func AddUserCredential(loginName string,pwd string)error{
	stmt, e := dbConn.Prepare("insert into users  (login_name,pwd) values(?,?)")
	if e!=nil {
		return e
	}
	_, er := stmt.Exec(loginName, pwd)
	if er != nil {
		return er
	}
	stmt.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error){
	stmt, e := dbConn.Prepare("select  pwd from  users WHERE login_name=?")
	if e != nil {
		log.Printf("%s",e)
		return "",err
	}
	var pwd string
	stmt.QueryRow(loginName).Scan(&pwd)
	stmt.Close()
	return pwd,nil

}

func DeleteUser(loginName string,pwd string)error{
	stmt, e := dbConn.Prepare("delete from users WHERE login_name=? AND pwd=?")
	if e!=nil {
		log.Printf("deleteuser error %s",e)
		return e
	}
	_, err := stmt.Exec(loginName, pwd)
	if err!=nil{
		log.Printf("deleteuser error %s",err)
		return err
	}
	return nil
}
