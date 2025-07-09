package orm

import (
	models "database/pkg/Models"
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

func(orm *Orm) CreateUser(user models.User) (error) {
	_,err := orm.db.Exec("insert into users (email,password,name) values ($1,$2,$3)",user.Email,user.Password,user.Name)	
	if err != nil{
		return fmt.Errorf("Database error %w",err)	
	}
	return nil
}

func(orm *Orm) RepeateUser(email string) (*models.User,error){
	var user models.User
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
