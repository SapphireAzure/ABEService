package mydatabase

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type DBConnector struct{
	targetDB *sql.DB
}

/**
负责连接数据库
dbName 为要连接数据的名字
注意不要忘记关闭DB
 */
func ConnectDB(dbName string )(dbco DBConnector){
	var targetDB string = "root:111111@tcp(127.0.0.1:3306)/"+dbName+"?charset=utf8"
	db, err := sql.Open("mysql", targetDB)
	if err != nil{
		fmt.Println("connect err\n")
	}
	if err == nil{
		fmt.Println("Connect success\n")
	}
	connector :=DBConnector{
		targetDB: db,
	}
	return connector
}
/**
负责插入一条数据到目标数据库
op Operation表示更新类型 类型为 INSERT UPDATE 和DELETE 插入更新删除
这个地方暂且固定成一个固定插入形式
 */
func (co DBConnector) StateDeal(op Operation,tableName string,id int ,policy string,user string){
	var sqlstr string=""
	if op == Insert{
		sqlstr = "INSERT INTO "+tableName+" VALUES(?,?,?)"
		myStmt,err :=co.targetDB.Prepare(sqlstr)
		if err !=nil{
			log.Fatal(err)
		}
		defer myStmt.Close()
		myStmt.Exec(id,policy,user)
	}
	if op ==Update{
		sqlstr = "UPDATE "+tableName+" SET policy = ?,user = ? WHERE id = ?"
		myStmt,err :=co.targetDB.Prepare(sqlstr)
		if err !=nil{
			log.Fatal(err)
		}
		defer myStmt.Close()
		myStmt.Exec(policy,user,id)
	}
	if op == Delete{
		sqlstr = "DELETE FROM "+tableName+" WHERE id = ?"
		myStmt,err :=co.targetDB.Prepare(sqlstr)
		if err !=nil{
			log.Fatal(err)
		}
		defer myStmt.Close()
		myStmt.Exec(id)
	}
	if sqlstr == "" {
		fmt.Println("Worng Operation Type")
	}

}


/**
Query一条在db中的信息
tableName 表示要查询表的名称
会打印返回整个表的信息
 */
func (co DBConnector) Query(tableName string){
	var queryStr string = "SELECT * from "+tableName

	rows,err := co.targetDB.Query(queryStr)
	if err != nil {
		log.Fatal(err)

	}
	defer rows.Close()


	for rows.Next(){
		var(
			id int
			policy string
			user string
		)
		err := rows.Scan(&id,&policy,&user)
		if(err != nil) {
			log.Fatal(err)
		}
		fmt.Println(id,policy,user)
	}
	err=rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

/**
Connector关闭连接
 */
func (co DBConnector) Close(){
	co.targetDB.Close()
}

/**
内部方法用来得到ID最大值
 */
func (co DBConnector)GetMaxID(tableName string)(id int){
	sqlstr := "SELECT MAX(id) FROM "+tableName
	rows,err :=co.targetDB.Query(sqlstr)
	if err !=nil{
		log.Fatal(err)
	}
	var maxID int
	rows.Next()
	rows.Scan(&maxID)
	return maxID
}
