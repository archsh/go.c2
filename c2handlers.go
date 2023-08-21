package c2

import (
	"encoding/xml"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ExecCmdReqEnvelope struct {
	XMLName      xml.Name `xml:"Envelope"`
	XmlnsXsi     string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd     string   `xml:"xmlns:xsd,attr"`
	XmlnsSoapenv string   `xml:"xmlns:soapenv,attr"`
	XmlnsIptv    string   `xml:"xmlns:iptv,attr"`
	Header       *struct {
		XMLName xml.Name      `xml:"Header"`
		Items   []interface{} `xml:",omitempty"`
	} `xml:",omitempty"`
	Body struct {
		XMLName xml.Name `"Body"`

		Fault *struct {
			XMLName xml.Name `xml:"Fault"`
			Code    string   `xml:"faultcode,omitempty"`
			String  string   `xml:"faultstring,omitempty"`
			Actor   string   `xml:"faultactor,omitempty"`
			Detail  string   `xml:"detail,omitempty"`
		} `xml:",omitempty"`
		Content *struct {
			XMLName     xml.Name `xml:"ExecCmd"`
			CSPID       string
			LSPID       string
			CorrelateID string
			CmdFileURL  string
		} `xml:",omitempty"`
	}
}

type RequestCmdHandleFunc = func(CSPID, LSPID, CorrelateID, CmdFileURL string) error

func MakeRequestCmdHandler(f RequestCmdHandleFunc) func(ctx *fiber.Ctx) error {
	var ff = func(ctx *fiber.Ctx) error {
		var envelope = ExecCmdReqEnvelope{}
		var resp SOAPEnvelope[ExecCmdResponse]
		//resp.Soapenv = SOAPENV
		//resp.Xsi = XSI
		//resp.Xsd = XSD
		//resp.NS = NS1
		//resp.Body.Content.EncodingStyle = EncodingStyle
		//resp.Body.Content.NS1 = NS1
		//resp.Body.Content.ExecCmdReturn.Type = "ns2:CSPResult"
		//resp.Body.Content.ExecCmdReturn.EncodingStyle = EncodingStyle
		//resp.Body.Content.ExecCmdReturn.NS2 = "iptv"
		//resp.Body.Content.ExecCmdReturn.Soapenc = SOAPENC
		//resp.Body.Content.ExecCmdReturn.Result.Type = XSD_INT
		//resp.Body.Content.ExecCmdReturn.ErrorDescription.Type = SOAPENC_STR
		if e := xml.Unmarshal(ctx.Body(), &envelope); nil != e {
			logrus.Errorln("HandleRequestCmd:> unmarshal failed:", e)
			logrus.Debugln("===================================================")
			logrus.Debugln("\n" + string(ctx.Body()))
			logrus.Debugln("===================================================")
			if bs, e := xml.MarshalIndent(envelope, "", " "); nil == e {
				logrus.Debugf("%s\n", string(bs))
			} else {
				logrus.Debugf("%+v\n", envelope)
			}
			logrus.Debugln("===================================================")
			_ = ctx.SendStatus(fiber.StatusBadRequest)
			resp = NewExecCmdResponse(fiber.StatusBadRequest, e.Error())
			//resp.Body.Content.ExecCmdReturn.Result.Value = fmt.Sprint(fiber.StatusBadRequest)
			//resp.Body.Content.ExecCmdReturn.ErrorDescription.Value = e.Error()

		} else if nil == envelope.Body.Content {
			_ = ctx.SendStatus(fiber.StatusBadRequest)
			resp = NewExecCmdResponse(fiber.StatusBadRequest, "Missing content body")
			//resp.Body.Content.ExecCmdReturn.Result.Value = fmt.Sprint(fiber.StatusBadRequest)
			//resp.Body.Content.ExecCmdReturn.ErrorDescription.Value = "Missing content body"
		} else if e := f(envelope.Body.Content.CSPID, envelope.Body.Content.LSPID, envelope.Body.Content.CorrelateID, envelope.Body.Content.CmdFileURL); nil != e {
			_ = ctx.SendStatus(fiber.StatusInternalServerError)
			resp = NewExecCmdResponse(fiber.StatusInternalServerError, e.Error())
			//resp.Body.Content.ExecCmdReturn.Result.Value = fmt.Sprint(fiber.StatusInternalServerError)
			//resp.Body.Content.ExecCmdReturn.ErrorDescription.Value = e.Error()
		} else {
			if bs, e := xml.MarshalIndent(envelope, "", " "); nil == e {
				logrus.Debugf("%s\n", string(bs))
			} else {
				logrus.Debugf("%+v\n", envelope)
			}
			resp = NewExecCmdResponse(0, "Success")
			//resp.Body.Content.ExecCmdReturn.Result.Value = ""
			//resp.Body.Content.ExecCmdReturn.ErrorDescription.Value = "Success"
		}
		if bs, e := xml.MarshalIndent(resp, "", " "); nil != e {
			return e
		} else {
			_, _ = ctx.WriteString(xml.Header)
			_, _ = ctx.Write(bs)
			ctx.Response().Header.SetContentType(fiber.MIMEApplicationXML)
		}
		return nil // ctx.XML(resp)
	}
	return ff
}
