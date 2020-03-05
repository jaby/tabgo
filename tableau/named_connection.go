package tableau

import "encoding/xml"

type NamedConnection struct {
	XMLName    xml.Name `xml:"named-connection"`
	Text       string   `xml:",chardata"`
	Caption    string   `xml:"caption,attr"`
	Name       string   `xml:"name,attr"`
	Connection struct {
		Text                 string `xml:",chardata"`
		Authentication       string `xml:"authentication,attr"`
		Class                string `xml:"class,attr"`
		OneTimeSql           string `xml:"one-time-sql,attr"`
		Port                 string `xml:"port,attr"`
		Schema               string `xml:"schema,attr"`
		Server               string `xml:"server,attr"`
		ServerOauth          string `xml:"server-oauth,attr"`
		Service              string `xml:"service,attr"`
		Sslmode              string `xml:"sslmode,attr"`
		Username             string `xml:"username,attr"`
		WorkgroupAuthMode    string `xml:"workgroup-auth-mode,attr"`
		Dbname               string `xml:"dbname,attr"`
		MinimumDriverVersion string `xml:"minimum-driver-version,attr"`
		OdbcNativeProtocol   string `xml:"odbc-native-protocol,attr"`
	} `xml:"connection"`
}

type NamedConnections struct {
	XMLName         xml.Name          `xml:"named-connections"`
	Text            string            `xml:",chardata"`
	NamedConnection []NamedConnection `xml:"named-connection"`
}
