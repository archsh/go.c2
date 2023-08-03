package c2

import (
	"bytes"
	"encoding/xml"
	"github.com/jlaffaye/ftp"
	"github.com/sirupsen/logrus"
	"io"
	"net/url"
)

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

func (c FTPClient) Get(s string, w io.Writer) (int64, error) {
	if u, e := url.Parse(s); nil != e {
		return 0, e
	} else {
		if nil != u.User {
			c.username = u.User.Username()
			if p, b := u.User.Password(); b {
				c.password = p
			}
		}
		if "" != u.Port() {
			c.address = u.Host + ":" + u.Port()
		} else {
			c.address = u.Host + ":21"
		}
		if e := c.login(); nil != e {
			return 0, e
		}
		defer func() { _ = c.logout() }()
		if resp, e := c.conn.Retr(u.Path); nil != e {
			return 0, e
		} else {
			defer func() { _ = resp.Close() }()
			return io.Copy(w, resp)
		}
	}
}

func FtpGet(s string, w io.Writer) (int64, error) {
	return defaultFTPClient.Get(s, w)
}

func FTPGetADI(s string) (ADI, error) {
	var adi ADI
	var buf = &bytes.Buffer{}
	if _, e := FtpGet(s, buf); nil != e {
		logrus.Errorln("FTPGetADI:>", e)
		return adi, e
	} else if e := xml.Unmarshal(buf.Bytes(), &adi); nil != e {
		logrus.Errorln("FTPGetADI:>", e)
		return adi, e
	} else {
		return adi, nil
	}
}

func FtpConfig(username string, password string) {
	defaultFTPClient.username = username
	defaultFTPClient.password = password
}

var defaultFTPClient *FTPClient

func init() {
	defaultFTPClient = &FTPClient{username: "anonymous", password: "anonymous"}
}
