package c2

import (
	"encoding/xml"
	"fmt"
)

const (
	SOAPENV       = "http://schemas.xmlsoap.org/soap/envelope/"
	SOAPENC       = "http://schemas.xmlsoap.org/soap/encoding/"
	XSD           = "http://www.w3.org/2001/XMLSchema"
	XSI           = "http://www.w3.org/2001/XMLSchema-instance"
	IPTV          = "iptv"
	NS1           = "iptv"
	NS2           = "iptv"
	EncodingStyle = "http://schemas.xmlsoap.ong/soap/encoding/"
	XSD_INT       = "xsd:int"
	XSD_STR       = "xsd:string"
	SOAPENC_STR   = "soapenc:string"
)

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
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ iptv:ExecCmd"`
	EncodingStyle string   `xml:"soapenv:encodingStyle,attr,omitempty"`
	CSPID         struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"CSPID"`
	LSPID struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"LSPID"`
	CorrelateID struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"CorrelateID"`
	CmdFileURL struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"CmdFileURL"`
}

func NewExecCmdReq(CSPID, LSPID, CorrelateID, CmdFileURL string) SOAPEnvelope[ExecCmdReq] {
	var envelop = SOAPEnvelope[ExecCmdReq]{}
	envelop.Soapenv = SOAPENV
	envelop.Xsd = XSD
	envelop.Xsi = XSI
	envelop.NS = IPTV
	envelop.Body.Content = &ExecCmdReq{}
	envelop.Body.Content.EncodingStyle = EncodingStyle
	envelop.Body.Content.CSPID.Type = SOAPENC_STR
	envelop.Body.Content.CSPID.Soapenc = SOAPENC
	envelop.Body.Content.CSPID.Value = CSPID

	envelop.Body.Content.LSPID.Type = SOAPENC_STR
	envelop.Body.Content.LSPID.Soapenc = SOAPENC
	envelop.Body.Content.LSPID.Value = LSPID

	envelop.Body.Content.CorrelateID.Type = SOAPENC_STR
	envelop.Body.Content.CorrelateID.Soapenc = SOAPENC
	envelop.Body.Content.CorrelateID.Value = CorrelateID

	envelop.Body.Content.CmdFileURL.Type = SOAPENC_STR
	envelop.Body.Content.CmdFileURL.Soapenc = SOAPENC
	envelop.Body.Content.CmdFileURL.Value = CmdFileURL

	return envelop
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
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ ns1:ExecCmdResponse"`
	EncodingStyle string   `xml:"soapenv:encodingStyle,attr,omitempty"`
	NS1           string   `xml:"ns1,attr,omitempty"`
	ExecCmdReturn struct {
		Type          string `xml:"xsi:type,attr,omitempty"`
		EncodingStyle string `xml:"soapenv:encodingStyle,attr,omitempty"`
		Soapenc       string `xml:"xmlns:soapenc,attr,omitempty"`
		NS2           string `xml:"ns2,attr,omitempty"`
		Result        struct {
			Type  string `xml:"xsi:type,attr,omitempty"`
			Value string `xml:",chardata"`
		}
		ErrorDescription struct {
			Type  string `xml:"xsi:type,attr,omitempty"`
			Value string `xml:",chardata"`
		}
	}
}

func NewExecCmdResponse(result int, desc string) SOAPEnvelope[ExecCmdResponse] {
	var envelop = SOAPEnvelope[ExecCmdResponse]{}
	envelop.Soapenv = SOAPENV
	envelop.Xsd = XSD
	envelop.Xsi = XSI
	envelop.NS = IPTV
	envelop.Body.Content = &ExecCmdResponse{}
	envelop.Body.Content.EncodingStyle = EncodingStyle
	envelop.Body.Content.NS1 = NS1
	envelop.Body.Content.ExecCmdReturn.NS2 = NS2
	envelop.Body.Content.ExecCmdReturn.Type = "ns2:CSPResult"
	envelop.Body.Content.ExecCmdReturn.EncodingStyle = EncodingStyle
	envelop.Body.Content.ExecCmdReturn.Soapenc = SOAPENC
	envelop.Body.Content.ExecCmdReturn.Result.Type = XSD_INT
	envelop.Body.Content.ExecCmdReturn.Result.Value = fmt.Sprint(result)
	envelop.Body.Content.ExecCmdReturn.ErrorDescription.Type = SOAPENC_STR
	envelop.Body.Content.ExecCmdReturn.ErrorDescription.Value = desc
	return envelop
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
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ iptv:ResultNotify"`

	EncodingStyle string `xml:"soapenv:encodingStyle,attr,omitempty"`
	CSPID         struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"CSPID"`
	LSPID struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"LSPID"`
	CorrelateID struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"CorrelateID"`
	CmdResult struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   int    `xml:",chardata"`
	} `xml:"CmdResult"`
	ResultFileURL struct {
		Type    string `xml:"xsi:type,attr,omitempty"`
		Soapenc string `xml:"xmlns:soapenc,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"ResultFileURL"`
}

func NewResultNotifyReq(CSPID, LSPID, CorrelateID string, result int, resultFileUrl string) SOAPEnvelope[ResultNotifyReq] {
	var envelop = SOAPEnvelope[ResultNotifyReq]{}
	envelop.Soapenv = SOAPENV
	envelop.Xsd = XSD
	envelop.Xsi = XSI
	envelop.NS = IPTV
	envelop.Body.Content = &ResultNotifyReq{}
	envelop.Body.Content.EncodingStyle = EncodingStyle
	envelop.Body.Content.CSPID.Type = SOAPENC_STR
	envelop.Body.Content.CSPID.Soapenc = SOAPENC
	envelop.Body.Content.CSPID.Value = CSPID

	envelop.Body.Content.LSPID.Type = SOAPENC_STR
	envelop.Body.Content.LSPID.Soapenc = SOAPENC
	envelop.Body.Content.LSPID.Value = LSPID

	envelop.Body.Content.CorrelateID.Type = SOAPENC_STR
	envelop.Body.Content.CorrelateID.Soapenc = SOAPENC
	envelop.Body.Content.CorrelateID.Value = CorrelateID

	envelop.Body.Content.CmdResult.Type = XSD_INT
	envelop.Body.Content.CmdResult.Soapenc = SOAPENC
	envelop.Body.Content.CmdResult.Value = result

	envelop.Body.Content.ResultFileURL.Type = SOAPENC_STR
	envelop.Body.Content.ResultFileURL.Soapenc = SOAPENC
	envelop.Body.Content.ResultFileURL.Value = resultFileUrl
	return envelop
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
	XMLName            xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ iptv:ResultNotifyResponse"`
	EncodingStyle      string   `xml:"soapenv:encodingStyle,attr,omitempty"`
	NS1                string   `xml:"ns1,attr,omitempty"`
	ResultNotifyReturn struct {
		Type          string `xml:"xsi:type,attr,omitempty"`
		EncodingStyle string `xml:"soapenv:encodingStyle,attr,omitempty"`
		Soapenc       string `xml:"xmlns:soapenc,attr,omitempty"`
		NS2           string `xml:"ns2,attr,omitempty"`
		Result        struct {
			Type  string `xml:"xsi:type,attr,omitempty"`
			Value string `xml:",chardata"`
		}
		ErrorDescription struct {
			Type  string `xml:"xsi:type,attr,omitempty"`
			Value string `xml:",chardata"`
		}
	}
}

func NewResultNotifyResponse(result int, desc string) SOAPEnvelope[ResultNotifyResponse] {
	var envelop = SOAPEnvelope[ResultNotifyResponse]{}
	envelop.Soapenv = SOAPENV
	envelop.Xsd = XSD
	envelop.Xsi = XSI
	envelop.NS = IPTV
	envelop.Body.Content = &ResultNotifyResponse{}
	envelop.Body.Content.EncodingStyle = EncodingStyle
	envelop.Body.Content.NS1 = NS1
	envelop.Body.Content.ResultNotifyReturn.NS2 = NS2
	envelop.Body.Content.ResultNotifyReturn.Type = "ns2:CSPResult"
	envelop.Body.Content.ResultNotifyReturn.EncodingStyle = EncodingStyle
	envelop.Body.Content.ResultNotifyReturn.Soapenc = SOAPENC
	envelop.Body.Content.ResultNotifyReturn.Result.Type = XSD_INT
	envelop.Body.Content.ResultNotifyReturn.Result.Value = fmt.Sprint(result)
	envelop.Body.Content.ResultNotifyReturn.ErrorDescription.Type = SOAPENC_STR
	envelop.Body.Content.ResultNotifyReturn.ErrorDescription.Value = desc
	return envelop
}
