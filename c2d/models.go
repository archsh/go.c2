package main

import (
	"database/sql"
	"fmt"
	xql "github.com/archsh/go.xql"
	_ "github.com/archsh/go.xql/dialects/postgres"
	"time"
)

type ExecRequest struct {
	CorrelateID string     `xql:"name=correlate_id,pk,size=48,nullable=false"`
	CSPID       string     `xql:"name=csp_id,size=48"`
	LSPID       string     `xql:"name=lsp_id,size=48"`
	CmdFileURL  string     `xql:"name=cmd_file_url,size=256"`
	Status      int        `xql:"name=status,default=0"`
	Result      string     `xql:"name=result,size=1024,default=''"`
	Created     *time.Time `xql:"type=timestamp,default=Now()"`
	Updated     *time.Time `xql:"type=timestamp,default=Now()"`
}

func (ExecRequest) TableName() string {
	return "exec_requests"
}

type C2Object struct {
	ID          string         `xql:"name=id,pk,size=48,nullable=false"`
	Code        string         `xql:"name=code,size=48,nullable=false,unique=true"`
	ElementType string         `xql:"name=type,size=48,nullable=false"`
	ParentID    string         `xql:"name=parent_id,size=48,nullable=false,default=''"`
	ParentCode  string         `xql:"name=parent_code,size=48,nullable=false,default=''"`
	ParentType  string         `xql:"name=parent_type,size=48,nullable=false,default=''"`
	Properties  JSONDictionary `xql:"name=props,type=JSONB,nullable=true"`
	MappingType int            `xql:"name=map_type,default=-1"`
	Sequence    int64          `xql:"name=sequence,default=-1"`
	ValidStart  *time.Time     `xql:"type=timestamp,nullable=True"`
	ValidEnd    *time.Time     `xql:"type=timestamp,nullable=True"`
	Status      int            `xql:"name=status,default=0"`
	Sync        int            `xql:"name=sync,default=0"`
	Created     *time.Time     `xql:"type=timestamp,default=Now()"`
	Updated     *time.Time     `xql:"type=timestamp,default=Now()"`
}

func (C2Object) TableName() string {
	return "c2_objects"
}

type LoggingConfig struct {
	Level    string `yaml:"level"`
	Filename string `yaml:"filename"`
}

func DefaultLoggingConfig() LoggingConfig {
	return LoggingConfig{
		Level:    "debug",
		Filename: "/var/log/goc2d.log",
	}
}

type C2Config struct {
	CSPID  string `yaml:"cspid"`
	LSPID  string `yaml:"lspid"`
	Notify string `yaml:"notify"`
}

type PostgresqlConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db-name"`
	SSLMode  string `yaml:"ssl-mode"`
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
	CmdRequestTable = xql.DeclareTable(&ExecRequest{})
	ObjectTable     = xql.DeclareTable(&C2Object{})
)

var tables = []*xql.Table{
	CmdRequestTable,
	ObjectTable,
}
