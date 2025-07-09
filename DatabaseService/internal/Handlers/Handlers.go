package handlers

import (
	"context"
	orm "database/internal/Orm"
	. "database/pkg/Models"
	"fmt"

	dbService "github.com/DEEBBLUE/MailProtos/api/Database"
	. "github.com/DEEBBLUE/MailProtos/api/Req"
	"google.golang.org/grpc"
)

type DatabaseServer struct{
	dbService.UnimplementedDatabaseServer
	orm orm.Orm
}

func InitDatabaseService(server *grpc.Server,orm orm.Orm){
	dbService.RegisterDatabaseServer(server,&DatabaseServer{orm: orm})
}

func(serv *DatabaseServer) CreateUser(
	ctx context.Context,
	req *CreateUserReq,
) (*DefaultRes,error) {

	if err := serv.orm.CreateUser(*ToModel(req.GetUser())); err != nil {
		return &DefaultRes{
			Status: "Error",
		},fmt.Errorf("Error in db,%w",err)
	}

	return &DefaultRes{
			Status: "Ok",
	},nil
}

func(serv *DatabaseServer) RepeateUser(
	ctx context.Context,
	req *RepeateUserReq,
) (*RepeateUserRes,error){

	user,err := serv.orm.RepeateUser(req.GetEmail())
	if err != nil {
		return nil,fmt.Errorf("Error in db,%w",err)
	}

	return &RepeateUserRes{
		User: ToGRPC(user),
	},nil
}

func(serv *DatabaseServer) DeleteUser(
	ctx context.Context,
	req *DeleteUserReq,
)(*DefaultRes,error){
	if err := serv.orm.DeleteUser(req.GetEmail()); err != nil{
		return &DefaultRes{
			Status: "Error",
		},fmt.Errorf("%w",err)
	}

	return &DefaultRes{
		Status: "Ok",
	},nil
}

func(serv *DatabaseServer) UpdateUserPassword(
	ctx context.Context,
	req *UpdateUserPasswordReq,
) (*DefaultRes,error){
	if err := serv.orm.UpdateUser(req.GetEmail(),"password",req.GetNewPassword()); err != nil{
		return &DefaultRes{
			Status: "Error",
		},nil
	}
	return &DefaultRes{
		Status: "Ok",
	},nil
}

func(serv *DatabaseServer) UpdateUserName(
	ctx context.Context,
	req *UpdateUserNameReq,
) (*DefaultRes,error){
	if err := serv.orm.UpdateUser(req.GetEmail(),"name",req.GetNewUserName()); err != nil{
		return &DefaultRes{
			Status: "Error",
		},nil
	}
	return &DefaultRes{
		Status: "Ok",
	},nil
}
