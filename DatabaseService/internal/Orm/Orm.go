package orm

import (
	. "database/pkg/Models"
	"database/sql"
	"fmt"
)

type Orm struct {
	db *sql.DB
}

func Init(db *sql.DB) Orm{
	return Orm{
		db: db,
	}
}

func(orm *Orm) CreateUser(user User) (error) {
	_,err := orm.db.Exec("insert into users (email,password,name) values ($1,$2,$3)",user.Email,user.Password,user.Name)	
	if err != nil{
		return fmt.Errorf("Database error %w",err)	
	}
	return nil
}

func(orm *Orm) RepeateUser(email string) (*User,error){
	var user User
	err := orm.db.QueryRow("select * from users where email=$1",email).Scan(&user.Email,&user.Password,&user.Name)
	if err != nil {
		return nil,fmt.Errorf("Database error %w",err) 
	}	
	return &user,nil
}

func(orm *Orm) DeleteUser(email string) (error){
	_,err := orm.db.Exec("delete from users where email=$1",email)
	if err != nil{
		return fmt.Errorf("Database error %w",err)
	}
	return nil
}

func(orm *Orm) UpdateUser(email string,field_name string,field_value string) (error){
	_,err := orm.db.Exec("update users set $1=$2 where email=$3",field_name,field_value,email)
	if err != nil{
		return fmt.Errorf("Database error %w",err)
	}
	return nil
}

func(orm *Orm) CreateMessage(message *Message) (error){
	var msg_id int
	err := orm.db.QueryRow("insert into messages (name,message,time_dispatch,publisher) values ($1,$2,$3,$4) returning id",
													message.Name,message.Message,message.TimeDispatch,message.Publisher_email).Scan(&msg_id)
	if err != nil{
		return err
	}

	for _,mail := range message.Consumers_emails{
		_,err := orm.db.Exec("insert into message_consumers (message_id,consumer_email) values ($1,$2)",msg_id,mail)	
		if err != nil {
			return err
		}
	}

	return nil
}

func(orm *Orm) RepeateMessage(msg_id int) (*Message,error) {
	var msg Message
	var consumer_emails []string
	err := orm.db.QueryRow("select * from messages where messages.id=$1",msg_id,
												).Scan(&msg.Id,&msg.Message,&msg.TimeDispatch,&msg.Publisher_email,&msg.Name)
	if err != nil{
		return nil,fmt.Errorf("database error %w",err)
	}
	res,err := orm.db.Query("select (consumer_email) from message_consumers where message_consumers.message_id=$1",msg.Id)
	if err != nil{
		return nil,fmt.Errorf("database error %w",err)
	}	
	defer res.Close()

	for res.Next(){
		var email string
		err := res.Scan(&email)
		if err != nil {
			return nil,fmt.Errorf("database error %w",err)
		}
		consumer_emails = append(consumer_emails, email)
	}
	msg.Consumers_emails = consumer_emails

	return &msg,nil
}

func(orm *Orm) RepeateUserMessages(email string) ([]LightMessage,error){
	var listMsg []LightMessage
	var listId []int
		
	list,err := orm.db.Query("select (id) from messages where publisher=$1",email)
	if err != nil{
		return nil,fmt.Errorf("database error %w",err)
	}
	defer list.Close()

	for list.Next(){
		var id int
		if err:=list.Scan(&id);err != nil{
			return nil,fmt.Errorf("database error %w",err)
		}
		listId = append(listId, id)

	}

	listConsume,err := orm.db.Query("select (message_id) from message_consumers where consumer_email=$1",email)	

	if err != nil{
		return nil,fmt.Errorf("database error %w",err)
	}
	defer listConsume.Close()

	for listConsume.Next(){
		var id int
		if err:=listConsume.Scan(&id);err != nil{
			return nil,fmt.Errorf("database error %w",err)
		}
		listId = append(listId, id)
	}
	
	for _,id := range listId{
		var msg LightMessage 
	
		res := orm.db.QueryRow("select id,name,time_dispatch from messages where id=$1",id)

		if err := res.Scan(&msg.Id,&msg.Name,&msg.TimeDispatch); err != nil{
			return nil,fmt.Errorf("database error this is QueryRow %w",err)
		}
		listMsg = append(listMsg, msg)
	}
	return listMsg,nil
}
