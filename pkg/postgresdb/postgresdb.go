// postgresdb
package psqldb

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	pb "github.com/pararti/Regards/api/golang"
	configo "github.com/pararti/Regards/pkg/config"
)

var tablenames = []string{"media", "session", "users"}

type DataBase struct {
	DB *sql.DB
}

func (d DataBase) GetMedia(mid *pb.MediaID) (*pb.Media, error) {
	query := `SELECT id, lasting, link, type, name FROM media
			WHERE id = $1`
	row := d.DB.QueryRow(query, mid.Id)
	m := &pb.Media{}
	err := row.Scan(&m.Id.Id, &m.Lasting, &m.Link, &m.Type, &m.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		} else {
			return nil, err
		}
	}
	return m, nil
}

func (d DataBase) SetMedia(m *pb.Media) (*pb.MediaID, error) {
	query := `SELECT EXISTS(SELECT 1 FROM media WHERE link = $1)`
	var exists bool
	var id *pb.MediaID
	err := d.DB.QueryRow(query, m.Link).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		updt := `UPDATE media SET name = $1 WHERE link = $2 RETURNING id`
		err = d.DB.QueryRow(updt, m.Name, m.Link).Scan(&id.Id)
		if err != nil {
			return nil, err
		}
	} else {
		ins := `INSERT INTO media (lasting, link, type, name)
				VALUES($1,$2,$3,$4) RETURNING id`
		err := d.DB.QueryRow(ins, m.Lasting, m.Link, m.Type, m.Name).Scan(&id.Id)
		if err != nil {
			return nil, err
		}
	}
	return id, nil
}

func (d DataBase) GetSession(sid *pb.SessionID) (*pb.Session, error) {
	query := `SELECT id, meta, users FROM session
			WHERE id = $1`
	row := d.DB.QueryRow(query, sid.Id)
	s := &pb.Session{}
	err := row.Scan(&s.Id.Id, &s.Meta, &m.L, &m.Type, &m.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (d DataBase) SetSession(s *pb.Session) (*pb.SessionID, error) {
	query := `INSERT INTO session (meta,users) VALUES($1,$2) RETURNING id`
	var id *pb.SessionID
	err := d.DB.QueryRow(query, s.Meta, s.Users).Scan(&id.Id)

	return id, nil
}

func (d DataBase) CheckTable(names ...string) ([]bool, err) {
	query := `SELECT EXISTS(SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = $1)`
	booleans := make([]bool, len(names))
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	for i, e := range names {
		err := stmt.QueryRow(e).Scan(&booleans[i])
		if err != nil {
			return nil, err
		}
	}
	return booleans, nil
}

func (d DataBase) CreateTable(booleans ...bool) error {
	var err error
	for i, e := range booleans {
		if !e {
			switch tablenames[i] {
			case "media":
				_, err = d.DB.Exec(configo.CreateMediaTable)
			case "session":
				_, err = d.DB.Exec(configo.CreateSessionTable)
			case "users":
				_, err = d.DB.Exec(configo.CreateUsersTable)

			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func NewDataBase(c *configo.PSQLConfig) (*DataBase, error) {
	psqlconn := "host=" + c.Host + " port=" + c.Port + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " sslmode=" + c.SSLMode
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	d := DataBase{DB: db}

	booleans, err := d.CheckTable(tablenames)
	if err != nil {
		return nil, err
	}
	err = d.CreateTable(booleans)
	if err != nil {
		return nil, err
	}

	return db, nil
}
