package tableau

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// TabGo is the base implementation of the TableauApi command line in GoLang
type TabGo struct {
	ServerURL     string
	ApiVersion    string
	CurrentToken  string
	CurrentSiteID string
}

type CredentialHolder struct {
	Credentials Credentials `json:"credentials,omitempty" xml:"credentials,omitempty"`
}

type ProjectsHolder struct {
	Projects Projects `json:"projects,omitempty" xml:"projects,omitempty"`
}

type Credentials struct {
	Name        string `json:"name,omitempty" xml:"name,attr,omitempty"`
	Password    string `json:"password,omitempty" xml:"password,attr,omitempty"`
	Token       string `json:"token,omitempty" xml:"token,attr,omitempty"`
	Site        *Site  `json:"site,omitempty" xml:"site,omitempty"`
	Impersonate *User  `json:"user,omitempty" xml:"user,omitempty"`
}

type Site struct {
	ID           string     `json:"id,omitempty" xml:"id,attr,omitempty"`
	Name         string     `json:"name,omitempty" xml:"name,attr,omitempty"`
	ContentUrl   string     `json:"contentUrl,omitempty" xml:"contentUrl,attr,omitempty"`
	AdminMode    string     `json:"adminMode,omitempty" xml:"adminMode,attr,omitempty"`
	UserQuota    string     `json:"userQuota,omitempty" xml:"userQuota,attr,omitempty"`
	StorageQuota int        `json:"storageQuota,omitempty" xml:"storageQuota,attr,omitempty"`
	State        string     `json:"state,omitempty" xml:"state,attr,omitempty"`
	StatusReason string     `json:"statusReason,omitempty" xml:"statusReason,attr,omitempty"`
	Usage        *SiteUsage `json:"usage,omitempty" xml:"usage,omitempty"`
}

type SiteUsage struct {
	NumberOfUsers int `json:"number-of-users" xml:"number-of-users,attr"`
	Storage       int `json:"storage" xml:"storage,attr"`
}

type User struct {
	ID       string `json:"id,omitempty" xml:"id,attr,omitempty"`
	Name     string `json:"name,omitempty" xml:"name,attr,omitempty"`
	SiteRole string `json:"siteRole,omitempty" xml:"siteRole,attr,omitempty"`
	FullName string `json:"fullName,omitempty" xml:"fullName,attr,omitempty"`
}

type Project struct {
	ID          string `json:"id,omitempty" xml:"id,attr,omitempty"`
	Name        string `json:"name,omitempty" xml:"name,attr,omitempty"`
	Description string `json:"description,omitempty" xml:"description,attr,omitempty"`
}

type Projects struct {
	Projects []Project `json:"project,omitempty" xml:"project,omitempty"`
}

type TsResponse struct {
	XMLName        xml.Name `xml:"tsResponse"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Datasource     struct {
		Text                string `xml:",chardata"`
		CertificationNote   string `xml:"certificationNote,attr"`
		ContentUrl          string `xml:"contentUrl,attr"`
		CreatedAt           string `xml:"createdAt,attr"`
		EncryptExtracts     string `xml:"encryptExtracts,attr"`
		ID                  string `xml:"id,attr"`
		IsCertified         string `xml:"isCertified,attr"`
		Name                string `xml:"name,attr"`
		Type                string `xml:"type,attr"`
		UpdatedAt           string `xml:"updatedAt,attr"`
		UseRemoteQueryAgent string `xml:"useRemoteQueryAgent,attr"`
		WebpageUrl          string `xml:"webpageUrl,attr"`
		Project             struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"project"`
		Owner struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"owner"`
		Tags string `xml:"tags"`
	} `xml:"datasource"`
	Workbook struct {
		Text            string `xml:",chardata"`
		ID              string `xml:"id,attr"`
		Name            string `xml:"name,attr"`
		ContentUrl      string `xml:"contentUrl,attr"`
		WebpageUrl      string `xml:"webpageUrl,attr"`
		ShowTabs        string `xml:"showTabs,attr"`
		Size            string `xml:"size,attr"`
		CreatedAt       string `xml:"createdAt,attr"`
		UpdatedAt       string `xml:"updatedAt,attr"`
		EncryptExtracts string `xml:"encryptExtracts,attr"`
		DefaultViewId   string `xml:"defaultViewId,attr"`
		Project         struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"project"`
		Owner struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"owner"`
		Tags  string `xml:"tags"`
		Views struct {
			Text string `xml:",chardata"`
			View []struct {
				Text        string `xml:",chardata"`
				ID          string `xml:"id,attr"`
				Name        string `xml:"name,attr"`
				ContentUrl  string `xml:"contentUrl,attr"`
				CreatedAt   string `xml:"createdAt,attr"`
				UpdatedAt   string `xml:"updatedAt,attr"`
				ViewUrlName string `xml:"viewUrlName,attr"`
				Tags        string `xml:"tags"`
			} `xml:"view"`
		} `xml:"views"`
	} `xml:"workbook"`
}

type Connection struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	EmbedPassword bool   `json:"embedPassword"`
	ServerAddress string `json:"serverAddress"`
	ServerPort    string `json:"serverPort"`
	UserName      string `json:"userName"`
	DbName        string `json:"dbName"`
	PassWord      string `json:"password"`
}

type ConnectionFinder interface {
	FindConnection(caption string) (Connection, error)
}

func (tabl *TabGo) ApiURL() string {
	return fmt.Sprintf("%s/api/%s", tabl.ServerURL, tabl.ApiVersion)
}

// Signin signs in to a tableau site on the give Tablea.URL
// It remembers the current token and site ID for subsequent calls
// cfr https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_concepts_auth.htm
func (tabl *TabGo) Signin(username, password, siteName string) error {

	credentialHolder := CredentialHolder{
		Credentials: Credentials{
			Name:     username,
			Password: password,
			Site: &Site{
				ContentUrl: siteName,
			},
		},
	}

	jsonStr, err := json.Marshal(credentialHolder)
	if err != nil {
		return errors.Wrapf(err, "can not marshall json: %+v", tabl)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/signin", tabl.ApiURL()), bytes.NewBuffer(jsonStr))
	if err != nil {
		return errors.Wrapf(err, "can not post json %s", string(jsonStr))
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "can not client.Do(request) to signin to tableau")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if strings.Contains(strings.ToLower(string(body)), "error") {
		return fmt.Errorf("signin error details: %s", string(body))
	}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&credentialHolder)
	if err != nil {
		return errors.Wrapf(err, "can not json decode")
	}

	tabl.CurrentToken = credentialHolder.Credentials.Token
	tabl.CurrentSiteID = credentialHolder.Credentials.Site.ID
	return nil
}

// Signout signs out of tableau,
// forgetting previously stored site and token id
// cfr https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_concepts_auth.htm
func (tabl *TabGo) Signout() error {

	if tabl.CurrentToken == "" {
		return fmt.Errorf("can not sign out from tableau if not signed in")
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/signout", tabl.ApiURL()), nil)
	if err != nil {
		return errors.Wrapf(err, "can not post")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-tableau-auth", tabl.CurrentToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "can not client.Do(request) to sign out from tableau")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "" {
		return fmt.Errorf("can not sign out from tableau with token '%s', error returned from tableau: '%s'", tabl.CurrentToken, string(body))
	}

	tabl.CurrentToken = ""
	tabl.CurrentSiteID = ""
	return nil
}

// cfr https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_concepts_publish.htm
func (tabl *TabGo) PublishDocument(documentPath, projectName string, targetConnectionFinder ConnectionFinder) (TsResponse, error) {

	var tsResponse TsResponse
	documentName, documentExtension := GetDocumentNameFromPath(documentPath)
	projectID, err := tabl.GetProjectID(projectName)
	if err != nil {
		return tsResponse, errors.Wrapf(err, "can not get project id")
	}

	// Response success

	captionRe, err := regexp.Compile(`named-connection caption='([^']+)'`)
	if err != nil {
		return tsResponse, errors.Wrapf(err, "can not compile regex")
	}

	namedConnectionsRe, err := regexp.Compile(`(?s)(<named-connections>.*</named-connections>)`)
	if err != nil {
		return tsResponse, errors.Wrapf(err, "can not compile regex")
	}

	switch documentExtension {
	case "twb", "twbx":
		var connections string
		if documentExtension == "twbx" {

			// twbx is a zip containing twb files ...
			tmpdir, err := ioutil.TempDir("", "twbx")
			if err != nil {
				return tsResponse, errors.Wrapf(err, "can not create tmp dir")
			}

			defer os.RemoveAll(tmpdir)

			err = Unzip(documentPath, tmpdir)
			if err != nil {
				return tsResponse, errors.Wrapf(err, "can not unzip twbx")
			}

			var connections string
			twbxFiles, err := ioutil.ReadDir(tmpdir)
			if err != nil {
				return tsResponse, errors.Wrapf(err, "can not read files from dir '%s'", tmpdir)
			}
			for _, fi := range twbxFiles {
				if "twb" == filepath.Ext(fi.Name()) {
					fileConnections, err := ConnectionLinesXml(filepath.Join(tmpdir, fi.Name()), tsResponse, captionRe, targetConnectionFinder)
					if err != nil {
						return tsResponse, errors.Wrapf(err, "can not get ConnectionLines")
					}
					connections += fileConnections
				}
			}

		} else {
			connections, err = ConnectionLinesXml(documentPath, tsResponse, captionRe, targetConnectionFinder)
			if err != nil {
				return tsResponse, errors.Wrapf(err, "can not get ConnectionLines")
			}
		}

		tsRequest := fmt.Sprintf(`<tsRequest><workbook name="%s" showTabs="true">%s<project id="%s"/></workbook></tsRequest>`, documentName, connections, projectID)

		return uploadFile("request_payload", "text/xml", tsRequest, "tableau_workbook", documentPath,
			fmt.Sprintf("%s/sites/%s/workbooks?workbookType=%s&overwrite=true", tabl.ApiURL(), tabl.CurrentSiteID, documentExtension),
			documentExtension,
			tabl.CurrentToken)

	case "tds", "tdsx":
		//// Following works, but does not embed connection password
		tsRequest := fmt.Sprintf(`<tsRequest><datasource name="%s"><project id="%s"/></datasource></tsRequest>`, documentName, projectID)

		tsResponse, err := uploadFile("request_payload", "text/xml", tsRequest, "tableau_datasource", documentPath,
			fmt.Sprintf("%s/sites/%s/datasources?datasourceType=%s&overwrite=true", tabl.ApiURL(), tabl.CurrentSiteID, documentExtension),
			documentExtension,
			tabl.CurrentToken,
		)
		if err != nil {
			return tsResponse, errors.Wrapf(err, "can not upload datasource")
		}

		datasourceId := tsResponse.Datasource.ID

		connections, err := tabl.DataSourceConnections(datasourceId)
		if err != nil {
			return tsResponse, errors.Wrapf(err, "can not get DataSourceConnections")
		}

		namedConnections, err := NamedConnections(documentPath, namedConnectionsRe)
		if err != nil {
			return tsResponse, errors.Wrapf(err, "can not get NamedConnections for '%s'", documentPath)
		}

		for _, connection := range connections {
			var caption string
			ok := false
			connectionKey := fmt.Sprintf("%s|%s|%s|%s",
				connection.Type,
				connection.ServerAddress,
				//connection.DbName,
				connection.ServerPort,
				connection.UserName)
			if caption, ok = namedConnections[connectionKey]; !ok {
				return tsResponse, errors.Wrapf(err, "no named connection '%+v' found in '%+v'", connection, namedConnections)
			}
			err := tabl.EmbedDatasourceConnection(datasourceId, connection, targetConnectionFinder, caption)
			if err != nil {
				return tsResponse, errors.Wrapf(err, "can not embed datasource connection")
			}
		}

		return tsResponse, nil
	default:
		return tsResponse, fmt.Errorf("invalid document extension '', expecting one of 'tds', 'tdsx', 'twb', 'twbx'")
	}

}

func ConnectionLinesXml(documentPath string, tsResponse TsResponse, captionRe *regexp.Regexp, targetConnectionFinder ConnectionFinder) (string, error) {
	documentContent, err := ioutil.ReadFile(documentPath)
	if err != nil {
		return "", errors.Wrapf(err, "can not read content from '%s'", documentPath)
	}

	matches := captionRe.FindAllStringSubmatch(string(documentContent), -1)
	connections := ""

	for _, match := range matches {
		caption := match[1]
		targetConnection, err := targetConnectionFinder.FindConnection(caption)
		if err != nil {
			return "", errors.Wrapf(err, "can not find namedConnection with caption '%s'", caption)
		}
		connections += fmt.Sprintf(`<connection serverAddress="%s"><connectionCredentials name="%s" password="%s" embed="true" /></connection>`,
			targetConnection.ServerAddress, targetConnection.UserName, targetConnection.PassWord)

	}
	if connections != "" {
		return fmt.Sprintf("<connections>%s</connections>", connections), nil
	}
	return "", nil
}

// NamedConnections returns a map of Connections (as key) with the value being the caption of the named connection
func NamedConnections(documentPath string, namedConnectionsRe *regexp.Regexp) (map[string]string, error) {
	namedConnections := make(map[string]string)
	documentContent, err := ioutil.ReadFile(documentPath)
	if err != nil {
		return namedConnections, errors.Wrapf(err, "can not read content from '%s'", documentPath)
	}

	matches := namedConnectionsRe.FindStringSubmatch(string(documentContent))

	type NamedConnections struct {
		XMLName         xml.Name `xml:"named-connections"`
		Text            string   `xml:",chardata"`
		NamedConnection []struct {
			Text       string `xml:",chardata"`
			Caption    string `xml:"caption,attr"`
			Name       string `xml:"name,attr"`
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
		} `xml:"named-connection"`
	}

	parsedNamedConnections := NamedConnections{}
	err = xml.Unmarshal([]byte(matches[0]), &parsedNamedConnections)
	if err != nil {
		fmt.Printf("error: %v", err)
		return namedConnections, errors.Wrapf(err, "can not unmarshall named connections from '%s'", string(documentContent))
	}

	for _, parsedNamedConnection := range parsedNamedConnections.NamedConnection {
		namedConnections[fmt.Sprintf("%s|%s|%s|%s",
			parsedNamedConnection.Connection.Class,
			parsedNamedConnection.Connection.Server,
			//parsedNamedConnection.Connection.Dbname,
			parsedNamedConnection.Connection.Port,
			parsedNamedConnection.Connection.Username)] = parsedNamedConnection.Caption
	}

	return namedConnections, nil
}

func (tabl *TabGo) EmbedDatasourceConnection(datasourceId string, connection Connection, pwFinder ConnectionFinder, caption string) error {
	connectionURL := fmt.Sprintf("%s/sites/%s/datasources/%s/connections/%s", tabl.ApiURL(), tabl.CurrentSiteID, datasourceId, connection.ID)

	targetConnection, err := pwFinder.FindConnection(caption)
	if err != nil {
		return errors.Wrapf(err, "can not find the connection for caption '%s' ", caption)
	}

	payload := fmt.Sprintf(`<tsRequest><connection serverAddress="%s" userName="%s" password="%s" embedPassword="true" /></tsRequest>`,
		targetConnection.ServerAddress, targetConnection.UserName, targetConnection.PassWord)

	req, err := http.NewRequest("PUT", connectionURL, strings.NewReader(payload))
	if err != nil {
		return errors.Wrapf(err, "can not get")
	}
	//req.Header.Set("Accept", "text/xml")
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("X-tableau-auth", tabl.CurrentToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "can not client.Do(request)")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "can not read response body")
	}

	if strings.Contains(strings.ToLower(string(body)), "error") {
		return fmt.Errorf("get datasource connections failed: %s", string(body))
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected success response %s, but got %s: %s", http.StatusOK, resp.StatusCode, body)
	}
	return nil
}

func (tabl *TabGo) DataSourceConnections(datasourceId string) ([]Connection, error) {
	dsConnections := []Connection{}
	connectionURL := fmt.Sprintf("%s/sites/%s/datasources/%s/connections", tabl.ApiURL(), tabl.CurrentSiteID, datasourceId)

	req, err := http.NewRequest("GET", connectionURL, nil)
	if err != nil {
		return dsConnections, errors.Wrapf(err, "can not get connections for datasource")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-tableau-auth", tabl.CurrentToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dsConnections, errors.Wrapf(err, "can not client.Do")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dsConnections, errors.Wrapf(err, "can not read response body")
	}

	if strings.Contains(strings.ToLower(string(body)), "error") {
		return dsConnections, fmt.Errorf("get datasource connections failed: %s", string(body))
	}

	type DatasourceConnectionsHolder struct {
		Connections struct {
			Connection []Connection `json:"connection"`
		} `json:"connections"`
	}
	datasourceConnectionsHolder := DatasourceConnectionsHolder{}

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&datasourceConnectionsHolder)
	if err != nil {
		return dsConnections, errors.Wrapf(err, "can not json decode response body: %s", string(body))
	}
	dsConnections = datasourceConnectionsHolder.Connections.Connection
	return dsConnections, nil
}

func GetDocumentNameFromPath(fullpath string) (string, string) {
	baseName := filepath.Base(fullpath)
	extension := filepath.Ext(fullpath)
	return baseName[0 : len(baseName)-len(extension)], extension[1:]
}

func (tabl *TabGo) GetProjectID(projectName string) (string, error) {
	projectID := ""

	pageNum := 1    // default
	pageSize := 100 // default

	uri := fmt.Sprintf("%s/sites/%s/projects?pageSize=%d&pageNumber=%d", tabl.ApiURL(), tabl.CurrentSiteID, pageSize, pageNum)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return projectID, errors.Wrapf(err, "can not post")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-tableau-auth", tabl.CurrentToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return projectID, errors.Wrapf(err, "can not client.Do(request) to get ProjectID from tableau")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return projectID, errors.Wrapf(err, "can not read response body")
	}
	projectsHolder := ProjectsHolder{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&projectsHolder)
	if err != nil {
		return projectID, errors.Wrapf(err, "can not json decode")
	}

	for _, project := range projectsHolder.Projects.Projects {
		if project.Name == projectName {
			return project.ID, nil
		}
	}

	return projectID, fmt.Errorf("project '%s' can not be found on this site", projectName)
}

func uploadFile(payloadFieldName, payloadContentType, payloadContent, fileFieldName, filePath, uri string, documentExtension string, tablToken string) (TsResponse, error) {

	var tsResponse TsResponse
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	var g errgroup.Group

	if !fileExists(filePath) {
		return tsResponse, fmt.Errorf("document does not exist '%s'", filePath)
	}

	// write the request asynchronously
	g.Go(func() error {
		defer w.Close()
		defer m.Close()

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`name="%s"`, payloadFieldName))
		h.Set("Content-Type", payloadContentType)
		part1, err := m.CreatePart(h)
		if err != nil {
			return err
		}
		_, err = part1.Write([]byte(payloadContent))
		if err != nil {
			return err
		}

		part2, err := m.CreateFormFile(fileFieldName, filepath.Base(filePath))
		if err != nil {
			return err
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err = io.Copy(part2, file); err != nil {
			return err
		}

		return nil
	})

	// post the request
	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		return tsResponse, errors.Wrapf(err, "can not http.NewRequest")
	}
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/mixed; boundary=%s", m.Boundary()))
	req.Header.Set("x-tableau-auth", tablToken)

	//dump, err := httputil.DumpRequestOut(req, true)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%q", dump)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return tsResponse, errors.Wrapf(err, "can not client.Do(request) to upload file '%s'", filePath)
	}

	// response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tsResponse, errors.Wrapf(err, "can not read response body")
	}

	// Is the response an error ?
	bodyS := string(body)
	if strings.Contains(bodyS, "error") {
		return tsResponse, fmt.Errorf("upload failed with error:\n%s", bodyS)
	}

	// Response success
	err = xml.Unmarshal(body, &tsResponse)
	if err != nil {
		return tsResponse, err
	}

	return tsResponse, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Unzip(src, dest string) error {
	dest = filepath.Clean(dest) + string(os.PathSeparator)

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		path := filepath.Join(dest, f.Name)
		// Check for ZipSlip: https://snyk.io/research/zip-slip-vulnerability
		if !strings.HasPrefix(path, dest) {
			return fmt.Errorf("%s: illegal file path", path)
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
