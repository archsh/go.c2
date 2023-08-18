package c2

import "encoding/xml"

// the c2 command definitions

/*
<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:iptv="iptv">

	<soapenv:Header/>
	<soapenv:Body>
	   <iptv:ExecCmd soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	      <CSPID xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</CSPID>
	      <LSPID xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</LSPID>
	      <CorrelateID xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</CorrelateID>
	      <CmdFileURL xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</CmdFileURL>
	   </iptv:ExecCmd>
	</soapenv:Body>

</soapenv:Envelope>
*/
type ExecCmdReq struct {
	XMLName     xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ iptv:ExecCmd"`
	CSPID       string
	LSPID       string
	CorrelateID string
	CmdFileURL  string
}

/*
<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:iptv="iptv">

	<soapenv:Header/>
	<soapenv:Body>
	   <iptv:ExecCmdResponse soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	      <ExecCmdReturn xsi:type="iptv:CSPResult">
	         <Result xsi:type="xsd:int">?</Result>
	         <ErrorDescription xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</ErrorDescription>
	      </ExecCmdReturn>
	   </iptv:ExecCmdResponse>
	</soapenv:Body>

</soapenv:Envelope>
*/
type ExecCmdResponse struct {
	XMLName          xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ iptv:ExecCmdResponse"`
	Result           string
	ErrorDescription string
}

/*
<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:iptv="iptv">

	<soapenv:Header/>
	<soapenv:Body>
	   <iptv:ResultNotify soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	      <CSPID xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</CSPID>
	      <LSPID xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</LSPID>
	      <CorrelateID xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</CorrelateID>
	      <CmdResult xsi:type="xsd:int">?</CmdResult>
	      <ResultFileURL xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</ResultFileURL>
	   </iptv:ResultNotify>
	</soapenv:Body>

</soapenv:Envelope>
*/
type ResultNotifyReq struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ iptv:ResultNotify"`
	CSPID         string
	LSPID         string
	CorrelateID   string
	CmdResult     int
	ResultFileURL string
}

/*
<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:iptv="iptv">

	<soapenv:Header/>
	<soapenv:Body>
	   <iptv:ResultNotifyResponse soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	      <ResultNotifyReturn xsi:type="iptv:CSPResult">
	         <Result xsi:type="xsd:int">?</Result>
	         <ErrorDescription xsi:type="soapenc:string" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">?</ErrorDescription>
	      </ResultNotifyReturn>
	   </iptv:ResultNotifyResponse>
	</soapenv:Body>

</soapenv:Envelope>
*/
type ResultNotifyResponse struct {
	XMLName          xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ iptv:ResultNotifyResponse"`
	Result           int
	ErrorDescription string
}
