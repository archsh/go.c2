package c2

import "github.com/jlaffaye/ftp"

// ftp client defines here

type FTPClient struct {
	address  string
	username string
	password string
	conn     *ftp.ServerConn
}

func (c *FTPClient) login() error {
	if conn, e := ftp.Dial(c.address); nil != e {
		return e
	} else if e := conn.Login(c.username, c.password); nil != e {
		return e
	} else {
		c.conn = conn
		return nil
	}
}

func (c *FTPClient) logout() error {
	if c.conn == nil {
		return nil
	}
	if e := c.conn.Quit(); nil != e {
		return e
	} else {
		c.conn = nil
		return nil
	}
}

func (c FTPClient) Get(filename string, output string) error {
	return nil
}

func FtpGet(filename string, output string) error {
	return defaultFTPClient.Get(filename, output)
}

var defaultFTPClient *FTPClient

func init() {
	defaultFTPClient = &FTPClient{username: "anonymous", password: "anonymous"}
}
