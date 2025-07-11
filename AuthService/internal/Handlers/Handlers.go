package handlers

import (
	"authService/config"
	jwtAuth "authService/pkg/Models"
	"context"
	"fmt"
	"sync"

	authService "github.com/DEEBBLUE/MailProtos/api/Auth"
	dbService "github.com/DEEBBLUE/MailProtos/api/Database"
	"github.com/DEEBBLUE/MailProtos/api/Req"
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
	userPass := req.GetUser().GetPassword()

	go func(errChan chan error){
		_,err := auth.dbClient.CreateUser(ctx,req)
		if err != nil {
				errChan <- err
		}
		claster.Done()
	}(creatersErr)


	payload := fmt.Sprintf("Email: %s",userEmail)
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

func(auth *AuthServer) Login(
	ctx context.Context,
	req *authService.LoginReq,
)(*Req.TokensRes,error) {

	user,err := auth.dbClient.RepeateUser(context.Background(),&RepeateUserReq{
		Email: req.GetEmail(),
	})
	if err != nil {
		return nil,err
	}
	if req.GetPassword() == user.GetUser().GetPassword(){
		payload := fmt.Sprintf("Email: %s",req.GetEmail())
		access_token,err := jwtAuth.CreateAccess(payload,auth.conf.Credential) 
		if err != nil {
			return nil,fmt.Errorf("Error when generate access_token,%w",err)
		}

		refresh_token,err := jwtAuth.CreateRefresh(req.GetPassword(),auth.conf.Credential)
		if err != nil {
			return nil,fmt.Errorf("Error when generate refresh_token,%w",err)
		}

		return &TokensRes{
			AccessToken: access_token,
			RefreshToken: refresh_token,
		},nil
	}
	return nil,fmt.Errorf("User data is incorrect")
}

func(auth *AuthServer) Access(
	ctx context.Context,
	req *authService.AccessReq,
)( *Req.DefaultRes,error) {
	token,err := jwtAuth.ValidateToken(req.GetAccessToken(),auth.conf.Credential)		

	if err != nil || !token.Valid{
		return nil,fmt.Errorf("Access token err")
	}

	return &DefaultRes{
		Status: "Ok",
	},nil
}
func(auth *AuthServer) Refresh(
	ctx context.Context,
	req *authService.RefreshReq,
) ( *authService.RefreshRes,error){

	token,err := jwtAuth.ValidateToken(req.GetRefreshToken(),auth.conf.Credential)	
	if err != nil || !token.Valid{
		return nil,fmt.Errorf("Refresh token err")
	}

	access_token,err := jwtAuth.CreateAccess(req.GetEmail(),auth.conf.Credential)	
	if err != nil {
		return nil,fmt.Errorf("Access token err")
	}

	return &authService.RefreshRes{
		AccessToken: access_token,
	},nil
}
