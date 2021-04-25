package main

import (
    "database/sql"
    "fmt"
    "runtime"
)

func main() {
    db, err := sql.Open("mysql", "root:@....../test?charset=utf8")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()
    
    // 一个结构来接收查询结果集
    type User struct {
        Id int32
        Name string
        Age int8
    }
    
    // 用户信息列表
    var user User
    
    // 查询一条记录时, 不能使用类似if err := db.QueryRow().Scan(&...); err != nil {}的处理方式
    // 因为查询单条数据时, 可能返回var ErrNoRows = errors.New("sql: no rows in result set")该种错误信息
    // 而这属于正常错误
    // 应该不需要 Wrap 这个 error，抛给上层
    err = db.QueryRow(`
        SELECT id,name,age from user WHERE id = ?
    `, 2).Scan(
        &user.Id, &user.Name, &user.Age,
    )
    switch {
    case err == sql.ErrNoRows:
    case err != nil:
        fmt.Println(err, file, line)
    }
    fmt.Println(user.Id, user.Name, user.Age)
    
}