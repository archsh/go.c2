package main

import "time"

type ExecCommandRequest struct {
	CorrelateID string     `xql:"name=correlate_id,pk,size=48,nullable=false"`
	CSPID       string     `xql:"name=csp_id,size=48"`
	LSPID       string     `xql:"name=lsp_id,size=48"`
	CmdFileURL  string     `xql:"name=cmd_file_url,size=256"`
	Status      int        `xql:"name=status,default=0"`
	Created     *time.Time `xql:"type=timestamp,default=Now()"`
	Updated     *time.Time `xql:"type=timestamp,default=Now()"`
}

func (r ExecCommandRequest) TableName() string {
	return "exec_command_requests"
}
