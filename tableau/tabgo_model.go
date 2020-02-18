package tableau

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
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
	Workbook       struct {
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
		errors.Wrapf(err, "can not post")
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
func (tabl *TabGo) PublishDocument(documentPath, projectName string) (TsResponse, error) {

	var tsResponse TsResponse
	documentName, documentExtension := GetDocumentNameFromPath(documentPath)
	projectID, err := tabl.GetProjectID(projectName)
	if err != nil {
		return tsResponse, err
	}

	switch documentExtension {
	case "twb", "twbx":
		tsRequest := fmt.Sprintf(`<tsRequest><workbook name="%s" showTabs="true"><project id="%s"/></workbook></tsRequest>`, documentName, projectID)
		return uploadFile("request_payload", "text/xml", tsRequest, "tableau_workbook", documentPath,
			fmt.Sprintf("%s/sites/%s/workbooks?workbookType=%s&overwrite=true", tabl.ApiURL(), tabl.CurrentSiteID, documentExtension),
			tabl.CurrentToken,
			map[string]string{})
	case "tds", "tdsx":
		tsRequest := fmt.Sprintf(`<tsRequest><datasource name="%s"><project id="%s"/></datasource></tsRequest>`, documentName, projectID)
		return uploadFile("request_payload", "text/xml", tsRequest, "tableau_datasource", documentPath,
			fmt.Sprintf("%s/sites/%s/datasources?datasourceType=%s&overwrite=true", tabl.ApiURL(), tabl.CurrentSiteID),
			tabl.CurrentToken,
			map[string]string{})
	default:
		return tsResponse, fmt.Errorf("invalid document extension '', expecting one of 'tds', 'tdsx', 'twb', 'twbx'")
	}

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
		errors.Wrapf(err, "can not post")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-tableau-auth", tabl.CurrentToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return projectID, errors.Wrapf(err, "can not client.Do(request) to get ProjectID from tableau")
	}

	body, _ := ioutil.ReadAll(resp.Body)
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

func uploadFile(payloadFieldName, payloadContentType, payloadContent, fileFieldName, filePath, uri string, tablToken string, extraParams map[string]string) (TsResponse, error) {

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
