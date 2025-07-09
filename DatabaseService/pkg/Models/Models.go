package models

import (
	types "github.com/DEEBBLUE/MailProtos/api/Types"
)

type User struct{
	Email string
	Password string
	Name string
}

func CreateUser(email,password,name string) *User{
	return &User{
		Email: email,
		Password: password,
		Name: name,
	}
}

func ToGRPC(user *User) *types.User {
	return &types.User{
		Email: user.Email,
		Password: user.Password,
		Name: user.Name,
	}
}

func ToModel(grpc *types.User) *User {
	return CreateUser(
		grpc.GetEmail(),
		grpc.GetName(),
		grpc.GetPassword(),
	)
}


