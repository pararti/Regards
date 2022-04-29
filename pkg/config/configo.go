// configo
package configo

type PSQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

var CreateUsersTable = `CREATE TABLE users(
				id serial PRIMARY KEY,
				name varchar(128) NOT NULL,
				login varchar(256) UNIQUE NOT NULL,
				cookie text
				)`
var CreateMediaTable = `CREATE TABLE media (
				id serial PRIMARY KEY,
				lasting interval NULL,
				link text UNIQUE,
				type varchar(256) NOT NULL,
				name varchar(256) NOT NULL
				)`
var CreateSessionTable = `CREATE TABLE session (
				id serial PRIMARY KEY,
				meta text NOT NULL,
				users int[]
				)`
