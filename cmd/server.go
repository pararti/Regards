// server
package main

import (
	"context"
	"log"
	
	pb "github.com/pararti/Regards/api/golang"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMediaAndSessionServer
}
//recive session id from client
//get session from database
//and return response to client
func (s *server) GetSession(ctx context.Context, *pb.SessionID)(*pb.Session, error){
	
	
}
//recive new session object
//add new session to database
//return new session id to client
func (s *server) SetSession(ctx context.Context, *pb.Session)(*pb.SessionID, error){
	
}
//recive media id from client
//get media object from database
//return response media object to client
func (s *server) GetMedia(ctx context.Context, *pb.MediaID)(*pb.Media, error){
	
}
//recive new media object from client
//add new media to database
//return new media id to client
func (s *server) SetMedia(ctx context.Context, *pb.Media)(*pb.MediaID, error){
	
}
