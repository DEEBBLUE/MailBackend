package main

import (
	"database/config"
	handlers "database/internal/Handlers"
	orm "database/internal/Orm"
	"database/sql"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {	
	conf,err := config.LoadConfig("/home/lifter/mail/backend/DatabaseService/config/config.yaml")
	if err != nil {
		panic(err)
	}

	db,err := sql.Open("postgres",conf.GetStringConf())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ormDB := orm.Init(db)

	listener,err := net.Listen("tcp",":4444")
	if err != nil {
		panic(err)
	}
	
	server := grpc.NewServer()
	handlers.InitDatabaseService(server,ormDB)

	if err := server.Serve(listener); err != nil{
		panic(err)
	}
}
