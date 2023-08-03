package main

import (
	"database/sql"
	"fmt"
	xql "github.com/archsh/go.xql"
	_ "github.com/archsh/go.xql/dialects/postgres"
	"time"
)

type ExecCommandRequest struct {
	CorrelateID string     `xql:"name=correlate_id,pk,size=48,nullable=false"`
	CSPID       string     `xql:"name=csp_id,size=48"`
	LSPID       string     `xql:"name=lsp_id,size=48"`
	CmdFileURL  string     `xql:"name=cmd_file_url,size=256"`
	Status      int        `xql:"name=status,default=0"`
	Result      string     `xql:"name=result,size=1024,default=''"`
	Created     *time.Time `xql:"type=timestamp,default=Now()"`
	Updated     *time.Time `xql:"type=timestamp,default=Now()"`
}

func (r ExecCommandRequest) TableName() string {
	return "exec_command_requests"
}

type PostgresqlConfig struct {
	Host     string
	Port     uint16
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func DefaultPgConfig() PostgresqlConfig {
	return PostgresqlConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		DBName:   "c2db",
		SSLMode:  "disable",
	}
}

func ConnectSQL(host string, port uint16, username string, password string, dbname string, sslmode string) (*sql.DB, error) {
	ds := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, username, password, dbname, sslmode)
	db, e := sql.Open("postgres", ds)
	if nil != e {
		return nil, e
	}
	return db, nil
}

var (
	CmdRequestTable = xql.DeclareTable(&ExecCommandRequest{})
)

var tables = []*xql.Table{
	CmdRequestTable,
}
