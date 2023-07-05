package c2

// ftp client defines here

type FTPClient struct {
}

func (ftp FTPClient) Get(filename string, output string) error {
	return nil
}

var defaultFTPClient *FTPClient

func init() {
	defaultFTPClient = &FTPClient{}
}
