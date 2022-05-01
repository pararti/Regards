// server
package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path"

	pb "github.com/pararti/Regards/api/golang"
	"github.com/pararti/Regards/pkg/config"
	db "github.com/pararti/Regards/pkg/postgresdb"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

//this var contains config for connect to postgresql database
var psqlConf = configo.PSQLConfig{Host: "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "",
	dbname:   "testdb",
	sslmode:  "disable",
}

var connConf configo.ConnConf

//this var contains database object like postgresql
var psqldb *db.DataBase

var log3 = Log3{Info: log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile),
	Warn:  log.New(os.Stdout, "WARN: ", log.LstdFlags|log.Lshortfile),
	Error: log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile),
}

func loadConfig(filepath string) error {
	data, err := ioutil.ReadFile(path.Join(filepath, "psqlconf.yml"))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &psqlConf)
	if err != nil {
		return err
	}
	data, err = ioutil.ReadFile(path.Join(filepath, "connconf.yml"))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &connConf)
	if err != nil {
		return err
	}
}

type server struct {
	pb.UnimplementedMediaAndSessionServer
}

//recive session id from client
//get session from database
//and return response to client
func (s *server) GetSession(ctx context.Context, id *pb.SessionID) (*pb.Session, error) {
	session, err := db.GetSession(id)
	if err != nil {
		return nil, err
	}
	return session, nil
}

//recive new session object
//add new session to database
//return new session id to client
func (s *server) SetSession(ctx context.Context, session *pb.Session) (*pb.SessionID, error) {
	id, err := db.SetSession(session)
	if err != nil {
		return nil, err
	}
	return id, nil
}

//recive media id from client
//get media object from database
//return response media object to client
func (s *server) GetMedia(ctx context.Context, id *pb.MediaID) (*pb.Media, error) {
	media, err := db.GetMedia(id)
	if err != nil {
		return nil, err
	}
	return media, nil
}

//recive new media object from client
//add new media to database
//return new media id to client
func (s *server) SetMedia(ctx context.Context, media *pb.Media) (*pb.MediaID, error) {
	id, err := db.SetMedia(media)
	if err != nil {
		return nil, err
	}
	return id, nil

}
