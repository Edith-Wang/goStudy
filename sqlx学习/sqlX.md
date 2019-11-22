## sqlx简介
sqlx是基于标准database/sql库的一个扩展库，重新封装了sql.DB、sql.TX、sql.Stmt等方法，所有底层接口都保持不变，sqlx相当于是标准库的超集。所以将sqlX整合进已使用标准库的项目完全无压力

## 最新的变化（2019/11/20）
* 引入sql.ColumnType，将Go的最低版本设置为1.8
* sqlx/types.JsonText已被重新命名为JSONText，以遵循Go的命名规则
  这个修改无法向后兼容，但是它影响很小，是可以修复的（s/JsonText/JSONText/g）。types包是实验阶段，没有处于活跃研发阶段
* 在基于Go 1.6及以下版本开发时，使用types.JSONText和types.GzippedText可能会有潜在的不安全行，特别是与常见 auto-scan sqlx习惯用法(如Select和Get)一起使用时。见golang bug#13905

## 安装
``go get github.com/jmoiron/sqlx``

## 使用方法
```
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//定义结构体
type Person struct {
	ID 	 int 		`db:"id""`  //对应数据库字段名 id
	Name string		`db:"name"` //对应数据库字段名 name
	Age  int		`db:"age"`  //对应数据库字段名 age
}
```
* 连接数据库
  * 方法：
        ``sqlx.Open(driverName, dataSourceName string) (*DB, error)``
        底层调用的是``sql.Open(driverName, dataSourceName string)``
  * 返回值：
       sqlX.DB
       ```
        type DB struct {
       	*sql.DB
       	driverName string
       	unsafe     bool
       	Mapper     *reflectx.Mapper //利用反射，根据结构体的tag声明与数据库字段名进行映射
       }
       ```
  * 示例：
        ``db,err := sqlx.Open("mysql","username:password@protocol(ip:port)/db?charset=utf8")``
* 插入数据
* 更新数据
* 删除数据
* 