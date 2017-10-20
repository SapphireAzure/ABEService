package ABEservice

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
	"./mydatabase"
)

/**
这个里面就是我们服务层操作的东西
暂时提供 新明文加密 更新一条明文 更新一个user所有明文
待补充 回滚
根据算法要求能还要在调整
 */

/**
这两个是测试的模拟函数
 */
func E(plaintext string,properties []string)(policy string){
	r :=rand.New(rand.NewSource(time.Now().UnixNano()))
	var p string= "p:"
	var i =0
	for ;i<3;i++{
		tempStr :=strconv.Itoa(r.Intn(10))
		p+=tempStr
	}
	fmt.Println(p)
	return p
}
func d(policy string,properties []string)(plaintext string){
	r :=rand.New(rand.NewSource(time.Now().UnixNano()))
	var p string= "m:"
	var i =0
	for ;i<3;i++{
		tempStr :=strconv.Itoa(r.Intn(10))
		p+=tempStr
	}
	fmt.Println(p)
	return p
}

/**
假定服务层为一个类 提供各种接口
 */
type ABEService struct{
	db mydatabase.DBConnector
}

/**
负责将服务层初始化 为包下方法
DBName需要一个数据库的名称 test为PolicyDb
 */
func ServiceInit(DBName string)(ABES ABEService){
	tar := ABEService{}
	tar.db = mydatabase.ConnectDB(DBName)
	return tar
}
/**
服务层调用的一些基础操作
 */
func (ABES ABEService) Query(tableName string) {
	ABES.db.Query(tableName)
}
/**
关闭数据库连接
 */
func (ABES ABEService) Close(){
	ABES.db.Close()
}

/**
面向客户
负责调用方法来创建一个密文并存储进数据库中
user 为用户姓名
plaintext 为明文
properties 为属性集合
 */
func (ABES ABEService) Encrypt (user string,plaintext string,properties []string)(){

	//调用算法产生密文
	policy := E(plaintext,properties)
	id := ABES.db.GetMaxID("policy")+1
	//将他插入到数据库 注ID自动生成加1
	ABES.db.StateDeal(mydatabase.Insert,"policy",id,policy,user)
}

/**
面向客户
负责调用方法来解密一个密文并更新密文存储进数据库中
user 为用户姓名
plaintext 为明文
properties 为属性集合
 */
func Update(user string,policy string,properties []string){

}

/**
面向众安
全面更新 传入用户 和一个属性集合将其密文全部更新
 */
func PropertyUpdate(user string,properties []string){

}

