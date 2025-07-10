package handlers

import (
	"authService/config"
	jwtAuth "authService/pkg/Models"
	"context"
	"fmt"
	"sync"

	authService "github.com/DEEBBLUE/MailProtos/api/Auth"
	dbService "github.com/DEEBBLUE/MailProtos/api/Database"
	_ "github.com/DEEBBLUE/MailProtos/api/Types"

	. "github.com/DEEBBLUE/MailProtos/api/Req"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServer struct{
	authService.UnimplementedAuthServer		
	dbClient dbService.DatabaseClient 
	conf *config.AuthConfig
}

func InitAuthServer(grpcServer *grpc.Server) (error){
	//ClientInit
	clientConn,err := grpc.Dial("localhost:4444",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}			
	dbClient := dbService.NewDatabaseClient(clientConn)

	//ConfigInit
	conf,err := config.InitConfig("/home/lifter/mail/backend/AuthService/config/config.yaml")
	if err != nil {
		panic(err)
	}
	
	//Register handlers
	authService.RegisterAuthServer(grpcServer,&AuthServer{
		dbClient: dbClient,
		conf: conf,
	})

	return nil
}

func(auth *AuthServer) Register(
	ctx context.Context,
	req *CreateUserReq,
) (*TokensRes,error) {
	var claster sync.WaitGroup
	claster.Add(1)
	creatersErr := make(chan error,1)	

	userEmail := req.GetUser().GetEmail()
	userName := req.GetUser().GetName()
	userPass := req.GetUser().GetPassword()
	
	go func(errChan chan error){
		_,err := auth.dbClient.CreateUser(ctx,req)
		if err != nil {
				errChan <- err
		}
		claster.Done()
	}(creatersErr)


	payload := fmt.Sprintf("Email: %s,Name: %s",userEmail,userName)
	access_token,err := jwtAuth.CreateAccess(payload,auth.conf.Credential) 
	if err != nil {
		return nil,fmt.Errorf("Error when generate access_token,%w",err)
	}

	refresh_token,err := jwtAuth.CreateRefresh(userPass,auth.conf.Credential)
	if err != nil {
		return nil,fmt.Errorf("Error when generate refresh_token,%w",err)
	}

	claster.Wait()
	close(creatersErr)

	for err := range creatersErr{
		if err != nil {
			return nil,err
		}
	}

	return &TokensRes{
		AccessToken: access_token,
		RefreshToken: refresh_token,
	},nil
}
