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
	return &User{ Email: email,
		Password: password,
		Name: name,
	}
}

func ToGRPCUser(user *User) *types.User {
	return &types.User{
		Email: user.Email,
		Password: user.Password,
		Name: user.Name,
	}
}

func ToModelUser(grpc *types.User) *User {
	return CreateUser(
		grpc.GetEmail(),
		grpc.GetName(),
		grpc.GetPassword(),
	)
}

type Message struct{
	Id int
	Name string 
	Publisher_email string
	Consumers_emails []string
	Message string
	TimeDispatch string
}
func CreateMessage(id int,name string,message string,time_dispatch string,
									 publisher_email string,consumers_emails []string,) *Message{
	return &Message{
		Id: id,
		Name: name,
		Publisher_email: publisher_email,
		Consumers_emails: consumers_emails,
		Message: message,
		TimeDispatch: time_dispatch,
	}
}

func ToGRPCMessage(message *Message) *types.Message{
	return &types.Message{
		Id: int32(message.Id),
		Name: message.Name,
		PublisherEmail: message.Publisher_email,
		EmailsConsumers: message.Consumers_emails,
		MessageValue: message.Message,
		TimeDispatch: message.TimeDispatch,
	}
}

func ToModelMessage(message *types.Message) *Message{
	return CreateMessage(int(message.GetId()),message.GetName(),message.MessageValue,message.GetTimeDispatch(),
											 message.GetPublisherEmail(),message.EmailsConsumers)	
}
