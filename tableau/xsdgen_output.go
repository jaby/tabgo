package tableau

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ContentOnly, ContentAndUsers
type AdminMode string

type Anon4 struct {
	TotalViewCount int `xml:"totalViewCount,attr"`
}

// May be one of -1
type Anon7 int

type BackgroundJobListType struct {
	BackgroundJob []BackgroundJobType `xml:"http://tableau.com/api backgroundJob,omitempty"`
}

type BackgroundJobType struct {
	Id        ResourceIdType `xml:"id,attr,omitempty"`
	Status    Status         `xml:"status,attr,omitempty"`
	CreatedAt time.Time      `xml:"createdAt,attr,omitempty"`
	StartedAt time.Time      `xml:"startedAt,attr,omitempty"`
	EndedAt   time.Time      `xml:"endedAt,attr,omitempty"`
	Priority  int            `xml:"priority,attr,omitempty"`
	JobType   string         `xml:"jobType,attr,omitempty"`
	Title     string         `xml:"title,attr,omitempty"`
	Subtitle  string         `xml:"subtitle,attr,omitempty"`
}

func (t *BackgroundJobType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BackgroundJobType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		StartedAt *xsdDateTime `xml:"startedAt,attr,omitempty"`
		EndedAt   *xsdDateTime `xml:"endedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.StartedAt = (*xsdDateTime)(&layout.T.StartedAt)
	layout.EndedAt = (*xsdDateTime)(&layout.T.EndedAt)
	return e.EncodeElement(layout, start)
}
func (t *BackgroundJobType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BackgroundJobType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		StartedAt *xsdDateTime `xml:"startedAt,attr,omitempty"`
		EndedAt   *xsdDateTime `xml:"endedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.StartedAt = (*xsdDateTime)(&overlay.T.StartedAt)
	overlay.EndedAt = (*xsdDateTime)(&overlay.T.EndedAt)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern [0-9]{5}.[0-9]{2}.[0-9]{4}.[0-9]{4}
type Build string

type Capabilities struct {
	Capability []CapabilityType `xml:"http://tableau.com/api capability"`
}

type CapabilityType struct {
	Name Name `xml:"name,attr"`
	Mode Mode `xml:"mode,attr"`
}

type ColumnListType struct {
	Column []ColumnType `xml:"http://tableau.com/api column,omitempty"`
}

type ColumnType struct {
	Site          SiteType       `xml:"http://tableau.com/api site,omitempty"`
	Id            ResourceIdType `xml:"id,attr,omitempty"`
	Name          string         `xml:"name,attr,omitempty"`
	Description   string         `xml:"description,attr,omitempty"`
	RemoteType    string         `xml:"remoteType,attr,omitempty"`
	ParentTableId ResourceIdType `xml:"parentTableId,attr,omitempty"`
}

type ConnectionCredentialsType struct {
	Name     string `xml:"name,attr,omitempty"`
	Password string `xml:"password,attr,omitempty"`
	Embed    bool   `xml:"embed,attr,omitempty"`
	OAuth    bool   `xml:"oAuth,attr,omitempty"`
}

type ConnectionListType struct {
	Connection []ConnectionType `xml:"http://tableau.com/api connection,omitempty"`
}

type ConnectionType struct {
	Datasource            DataSourceType            `xml:"http://tableau.com/api datasource,omitempty"`
	ConnectionCredentials ConnectionCredentialsType `xml:"http://tableau.com/api connectionCredentials,omitempty"`
	Id                    ResourceIdType            `xml:"id,attr,omitempty"`
	Type                  string                    `xml:"type,attr,omitempty"`
	DbClass               string                    `xml:"dbClass,attr,omitempty"`
	Scope                 string                    `xml:"scope,attr,omitempty"`
	Filename              string                    `xml:"filename,attr,omitempty"`
	GoogleSheetId         string                    `xml:"googleSheetId,attr,omitempty"`
	EmbedPassword         bool                      `xml:"embedPassword,attr,omitempty"`
	ServerAddress         string                    `xml:"serverAddress,attr,omitempty"`
	ServerPort            int                       `xml:"serverPort,attr,omitempty"`
	UserName              string                    `xml:"userName,attr,omitempty"`
	Password              string                    `xml:"password,attr,omitempty"`
}

// May be one of LockedToDatabase, ManagedByOwner
type ContentPermissions string

type DataAlertListType struct {
	DataAlert []DataAlertType `xml:"http://tableau.com/api dataAlert,omitempty"`
}

type DataAlertType struct {
	Owner          UserType                    `xml:"http://tableau.com/api owner"`
	View           ViewType                    `xml:"http://tableau.com/api view"`
	Recipients     DataAlertsRecipientListType `xml:"http://tableau.com/api recipients,omitempty"`
	Id             ResourceIdType              `xml:"id,attr,omitempty"`
	Subject        string                      `xml:"subject,attr,omitempty"`
	CreatorId      ResourceIdType              `xml:"creatorId,attr,omitempty"`
	CreatedAt      time.Time                   `xml:"createdAt,attr,omitempty"`
	UpdatedAt      time.Time                   `xml:"updatedAt,attr,omitempty"`
	Frequency      Frequency                   `xml:"frequency,attr,omitempty"`
	AlertCondition string                      `xml:"alertCondition,attr,omitempty"`
	AlertThreshold string                      `xml:"alertThreshold,attr,omitempty"`
}

func (t *DataAlertType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DataAlertType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *DataAlertType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DataAlertType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

type DataAlertsRecipientListType struct {
	Recipient []DataAlertsRecipientType `xml:"http://tableau.com/api recipient,omitempty"`
}

type DataAlertsRecipientType struct {
	Id       ResourceIdType `xml:"id,attr,omitempty"`
	Name     string         `xml:"name,attr,omitempty"`
	LastSent time.Time      `xml:"lastSent,attr,omitempty"`
}

func (t *DataAlertsRecipientType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DataAlertsRecipientType
	var layout struct {
		*T
		LastSent *xsdDateTime `xml:"lastSent,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastSent = (*xsdDateTime)(&layout.T.LastSent)
	return e.EncodeElement(layout, start)
}
func (t *DataAlertsRecipientType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DataAlertsRecipientType
	var overlay struct {
		*T
		LastSent *xsdDateTime `xml:"lastSent,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastSent = (*xsdDateTime)(&overlay.T.LastSent)
	return d.DecodeElement(&overlay, &start)
}

type DataQualityWarningListType struct {
	DataQualityWarning []DataQualityWarningType `xml:"http://tableau.com/api dataQualityWarning,omitempty"`
}

type DataQualityWarningType struct {
	Site            SiteType       `xml:"http://tableau.com/api site,omitempty"`
	Owner           UserType       `xml:"http://tableau.com/api owner,omitempty"`
	Id              ResourceIdType `xml:"id,attr,omitempty"`
	UserDisplayName string         `xml:"userDisplayName,attr,omitempty"`
	ContentId       ResourceIdType `xml:"contentId,attr,omitempty"`
	ContentType     string         `xml:"contentType,attr,omitempty"`
	Message         string         `xml:"message,attr,omitempty"`
	Type            string         `xml:"type,attr,omitempty"`
	IsActive        bool           `xml:"isActive,attr,omitempty"`
	CreatedAt       time.Time      `xml:"createdAt,attr,omitempty"`
	UpdatedAt       time.Time      `xml:"updatedAt,attr,omitempty"`
}

func (t *DataQualityWarningType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DataQualityWarningType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *DataQualityWarningType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DataQualityWarningType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

type DataRoleType struct {
	Site         SiteType         `xml:"http://tableau.com/api site,omitempty"`
	Project      ProjectType      `xml:"http://tableau.com/api project,omitempty"`
	Owner        UserType         `xml:"http://tableau.com/api owner,omitempty"`
	Name         string           `xml:"http://tableau.com/api name"`
	Description  string           `xml:"http://tableau.com/api description,omitempty"`
	CreatedAt    time.Time        `xml:"http://tableau.com/api createdAt,omitempty"`
	UpdatedAt    time.Time        `xml:"http://tableau.com/api updatedAt,omitempty"`
	FieldConcept FieldConceptType `xml:"http://tableau.com/api fieldConcept"`
	Url          string           `xml:"http://tableau.com/api url"`
}

func (t *DataRoleType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DataRoleType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"http://tableau.com/api createdAt,omitempty"`
		UpdatedAt *xsdDateTime `xml:"http://tableau.com/api updatedAt,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *DataRoleType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DataRoleType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"http://tableau.com/api createdAt,omitempty"`
		UpdatedAt *xsdDateTime `xml:"http://tableau.com/api updatedAt,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

type DataSourceListType struct {
	Datasource []DataSourceType `xml:"http://tableau.com/api datasource,omitempty"`
}

type DataSourceType struct {
	ConnectionCredentials   ConnectionCredentialsType `xml:"http://tableau.com/api connectionCredentials,omitempty"`
	Site                    SiteType                  `xml:"http://tableau.com/api site,omitempty"`
	Project                 ProjectType               `xml:"http://tableau.com/api project,omitempty"`
	Owner                   UserType                  `xml:"http://tableau.com/api owner,omitempty"`
	Tags                    TagListType               `xml:"http://tableau.com/api tags,omitempty"`
	Id                      ResourceIdType            `xml:"id,attr,omitempty"`
	Name                    string                    `xml:"name,attr,omitempty"`
	ContentUrl              string                    `xml:"contentUrl,attr,omitempty"`
	WebpageUrl              string                    `xml:"webpageUrl,attr,omitempty"`
	Description             string                    `xml:"description,attr,omitempty"`
	Type                    string                    `xml:"type,attr,omitempty"`
	CreatedAt               time.Time                 `xml:"createdAt,attr,omitempty"`
	UpdatedAt               time.Time                 `xml:"updatedAt,attr,omitempty"`
	IsCertified             bool                      `xml:"isCertified,attr,omitempty"`
	CertificationNote       string                    `xml:"certificationNote,attr,omitempty"`
	ServerName              string                    `xml:"serverName,attr,omitempty"`
	DatabaseName            string                    `xml:"databaseName,attr,omitempty"`
	HasExtracts             bool                      `xml:"hasExtracts,attr,omitempty"`
	HasAlert                bool                      `xml:"hasAlert,attr,omitempty"`
	Size                    int                       `xml:"size,attr,omitempty"`
	IsPublished             bool                      `xml:"isPublished,attr,omitempty"`
	ConnectedWorkbooksCount int                       `xml:"connectedWorkbooksCount,attr,omitempty"`
	FavoritesTotal          int                       `xml:"favoritesTotal,attr,omitempty"`
	EncryptExtracts         string                    `xml:"encryptExtracts,attr,omitempty"`
	UseRemoteQueryAgent     bool                      `xml:"useRemoteQueryAgent,attr,omitempty"`
}

func (t *DataSourceType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DataSourceType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *DataSourceType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DataSourceType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

type DataSourceValueStoreType struct {
	DatasourceID string `xml:"http://tableau.com/api datasourceID,omitempty"`
	FieldName    string `xml:"http://tableau.com/api fieldName,omitempty"`
}

// May be one of DATA_TYPE_UNSPECIFIED, DATE, DATETIME, STRING, INT, FLOAT, BOOL
type DataTypeType string

type DatabaseListType struct {
	Database []DatabaseType `xml:"http://tableau.com/api database,omitempty"`
}

type DatabaseType struct {
	Site               SiteType           `xml:"http://tableau.com/api site,omitempty"`
	Contact            UserType           `xml:"http://tableau.com/api contact"`
	Certifier          UserType           `xml:"http://tableau.com/api certifier"`
	Id                 ResourceIdType     `xml:"id,attr,omitempty"`
	Name               string             `xml:"name,attr,omitempty"`
	ConnectionType     string             `xml:"connectionType,attr,omitempty"`
	IsEmbedded         bool               `xml:"isEmbedded,attr,omitempty"`
	Description        string             `xml:"description,attr,omitempty"`
	IsCertified        bool               `xml:"isCertified,attr,omitempty"`
	CertificationNote  string             `xml:"certificationNote,attr,omitempty"`
	Type               DatabaseTypeType   `xml:"type,attr,omitempty"`
	HostName           string             `xml:"hostName,attr,omitempty"`
	Port               int                `xml:"port,attr,omitempty"`
	FilePath           string             `xml:"filePath,attr,omitempty"`
	Provider           string             `xml:"provider,attr,omitempty"`
	MimeType           string             `xml:"mimeType,attr,omitempty"`
	FileId             string             `xml:"fileId,attr,omitempty"`
	ConnectorUrl       string             `xml:"connectorUrl,attr,omitempty"`
	RequestUrl         string             `xml:"requestUrl,attr,omitempty"`
	FileExtension      string             `xml:"fileExtension,attr,omitempty"`
	ContentPermissions ContentPermissions `xml:"contentPermissions,attr,omitempty"`
}

// May be one of DatabaseServer, File, WebDataConnector, CloudFile
type DatabaseTypeType string

type DegradationListType struct {
	Degradation []DegradationType `xml:"http://tableau.com/api degradation,omitempty"`
}

type DegradationType struct {
	Name     string `xml:"name,attr,omitempty"`
	Severity string `xml:"severity,attr,omitempty"`
}

type DistinctValueListType struct {
	DistinctValue []DistinctValueType `xml:"http://tableau.com/api distinctValue,omitempty"`
}

type DistinctValueType struct {
	Value     SemanticsValueType `xml:"http://tableau.com/api value"`
	Frequency int64              `xml:"http://tableau.com/api frequency"`
}

type DomainDirectiveType struct {
	Name string `xml:"name,attr"`
}

type ErrorType struct {
	Summary string `xml:"http://tableau.com/api summary"`
	Detail  string `xml:"http://tableau.com/api detail"`
	Code    int    `xml:"code,attr"`
}

// May be one of Parallel, Serial
type ExecutionOrder string

type ExtractListType struct {
	Extract []ExtractType `xml:"http://tableau.com/api extract,omitempty"`
}

type ExtractRefreshJobType struct {
	Notes      string         `xml:"http://tableau.com/api notes"`
	View       ViewType       `xml:"http://tableau.com/api view"`
	Workbook   WorkbookType   `xml:"http://tableau.com/api workbook"`
	Datasource DataSourceType `xml:"http://tableau.com/api datasource"`
}

type ExtractType struct {
	Datasource DataSourceType `xml:"http://tableau.com/api datasource"`
	Workbook   WorkbookType   `xml:"http://tableau.com/api workbook"`
	Id         ResourceIdType `xml:"id,attr,omitempty"`
	Priority   int            `xml:"priority,attr,omitempty"`
	Type       Type           `xml:"type,attr,omitempty"`
}

type FavoriteListType struct {
	Favorite []FavoriteType `xml:"http://tableau.com/api favorite,omitempty"`
}

type FavoriteType struct {
	View              ViewType       `xml:"http://tableau.com/api view"`
	Workbook          WorkbookType   `xml:"http://tableau.com/api workbook"`
	Datasource        DataSourceType `xml:"http://tableau.com/api datasource"`
	Project           ProjectType    `xml:"http://tableau.com/api project"`
	Flow              FlowType       `xml:"http://tableau.com/api flow"`
	Label             string         `xml:"label,attr"`
	ParentProjectName string         `xml:"parentProjectName,attr,omitempty"`
	TargetOwnerName   string         `xml:"targetOwnerName,attr,omitempty"`
}

type FieldConceptType struct {
	Uri                   string                   `xml:"http://tableau.com/api uri"`
	ObjectConceptURI      string                   `xml:"http://tableau.com/api objectConceptURI"`
	Names                 []NameType               `xml:"http://tableau.com/api names,omitempty"`
	NameCharacteristics   NameCharacteristicsType  `xml:"http://tableau.com/api nameCharacteristics,omitempty"`
	Description           string                   `xml:"http://tableau.com/api description,omitempty"`
	ParentFieldConceptURI string                   `xml:"http://tableau.com/api parentFieldConceptURI,omitempty"`
	DataTypes             []DataTypeType           `xml:"http://tableau.com/api dataTypes,omitempty"`
	FieldRoles            []FieldRoleType          `xml:"http://tableau.com/api fieldRoles,omitempty"`
	DefaultFormats        []string                 `xml:"http://tableau.com/api defaultFormats,omitempty"`
	ValueCharacteristics  ValueCharacteristicsType `xml:"http://tableau.com/api valueCharacteristics,omitempty"`
	OwnerID               string                   `xml:"http://tableau.com/api ownerID,omitempty"`
	CreatedAt             time.Time                `xml:"http://tableau.com/api createdAt,omitempty"`
	UpdatedAt             time.Time                `xml:"http://tableau.com/api updatedAt,omitempty"`
	ValueSource           ValueSourceType          `xml:"http://tableau.com/api valueSource,omitempty"`
}

func (t *FieldConceptType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FieldConceptType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"http://tableau.com/api createdAt,omitempty"`
		UpdatedAt *xsdDateTime `xml:"http://tableau.com/api updatedAt,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *FieldConceptType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FieldConceptType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"http://tableau.com/api createdAt,omitempty"`
		UpdatedAt *xsdDateTime `xml:"http://tableau.com/api updatedAt,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

type FieldMatchListType struct {
	FieldMatch []FieldMatchType `xml:"http://tableau.com/api fieldMatch,omitempty"`
}

type FieldMatchType struct {
	FieldConceptURI string           `xml:"http://tableau.com/api fieldConceptURI"`
	Weight          float64          `xml:"http://tableau.com/api weight"`
	FieldConcept    FieldConceptType `xml:"http://tableau.com/api fieldConcept,omitempty"`
	ValueMatches    []ValueMatchType `xml:"http://tableau.com/api valueMatches,omitempty"`
}

// May be one of FIELD_ROLE_UNSPECIFIED, DIMENSION, MEASURE
type FieldRoleType string

type FieldType struct {
	SampleValues []DistinctValueType `xml:"http://tableau.com/api sampleValues,omitempty"`
	DataType     DataTypeType        `xml:"http://tableau.com/api dataType,omitempty"`
	FieldRole    FieldRoleType       `xml:"http://tableau.com/api fieldRole,omitempty"`
	Name         string              `xml:"http://tableau.com/api name,omitempty"`
}

// Must match the pattern ([0-9]+:[0-9a-fA-F]+)-([0-9]+:[0-9]+)
type FileUploadSessionIdType string

type FileUploadType struct {
	UploadSessionId FileUploadSessionIdType `xml:"uploadSessionId,attr"`
	FileSize        int                     `xml:"fileSize,attr,omitempty"`
}

type FlowListType struct {
	Flow []FlowType `xml:"http://tableau.com/api flow,omitempty"`
}

type FlowOutputStepListType struct {
	FlowOutputStep []FlowOutputStepType `xml:"http://tableau.com/api flowOutputStep,omitempty"`
}

type FlowOutputStepType struct {
	Id   ResourceIdType `xml:"id,attr,omitempty"`
	Name string         `xml:"name,attr,omitempty"`
}

type FlowType struct {
	Site        SiteType       `xml:"http://tableau.com/api site,omitempty"`
	Project     ProjectType    `xml:"http://tableau.com/api project,omitempty"`
	Owner       UserType       `xml:"http://tableau.com/api owner,omitempty"`
	Tags        TagListType    `xml:"http://tableau.com/api tags,omitempty"`
	Id          ResourceIdType `xml:"id,attr,omitempty"`
	Name        string         `xml:"name,attr,omitempty"`
	Description string         `xml:"description,attr,omitempty"`
	WebpageUrl  string         `xml:"webpageUrl,attr,omitempty"`
	FileType    string         `xml:"fileType,attr,omitempty"`
	CreatedAt   time.Time      `xml:"createdAt,attr,omitempty"`
	UpdatedAt   time.Time      `xml:"updatedAt,attr,omitempty"`
}

func (t *FlowType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FlowType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *FlowType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FlowType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

type FlowWarningsListContainerType struct {
	ConnectionWarnings WarningListType `xml:"http://tableau.com/api connectionWarnings,omitempty"`
	NodeWarnings       WarningListType `xml:"http://tableau.com/api nodeWarnings,omitempty"`
}

// May be one of Hourly, Daily, Weekly, Monthly
type Frequency string

type FrequencyDetailsType struct {
	Intervals Intervals `xml:"http://tableau.com/api intervals"`
	Start     time.Time `xml:"start,attr"`
	End       time.Time `xml:"end,attr,omitempty"`
}

func (t *FrequencyDetailsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T FrequencyDetailsType
	var layout struct {
		*T
		Start *xsdTime `xml:"start,attr"`
		End   *xsdTime `xml:"end,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Start = (*xsdTime)(&layout.T.Start)
	layout.End = (*xsdTime)(&layout.T.End)
	return e.EncodeElement(layout, start)
}
func (t *FrequencyDetailsType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FrequencyDetailsType
	var overlay struct {
		*T
		Start *xsdTime `xml:"start,attr"`
		End   *xsdTime `xml:"end,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Start = (*xsdTime)(&overlay.T.Start)
	overlay.End = (*xsdTime)(&overlay.T.End)
	return d.DecodeElement(&overlay, &start)
}

type GranteeCapabilitiesType struct {
	Group        GroupType    `xml:"http://tableau.com/api group"`
	User         UserType     `xml:"http://tableau.com/api user"`
	Capabilities Capabilities `xml:"http://tableau.com/api capabilities"`
}

type GroupListType struct {
	Group []GroupType `xml:"http://tableau.com/api group,omitempty"`
}

type GroupType struct {
	Domain          DomainDirectiveType `xml:"http://tableau.com/api domain,omitempty"`
	Import          ImportDirectiveType `xml:"http://tableau.com/api import,omitempty"`
	Id              ResourceIdType      `xml:"id,attr,omitempty"`
	Name            string              `xml:"name,attr,omitempty"`
	UserCount       int                 `xml:"userCount,attr,omitempty"`
	MinimumSiteRole SiteRoleType        `xml:"minimumSiteRole,attr,omitempty"`
}

// May be one of 1, 2, 4, 6, 8, 12
type Hours int

type ImportDirectiveType struct {
	Source     ImportSourceType `xml:"source,attr"`
	DomainName string           `xml:"domainName,attr"`
	SiteRole   SiteRoleType     `xml:"siteRole,attr"`
}

// May be one of ActiveDirectory
type ImportSourceType string

type IndexingStatusType struct {
	IndexingStatusCode         string `xml:"http://tableau.com/api indexingStatusCode"`
	IndexingErrorCode          string `xml:"http://tableau.com/api indexingErrorCode"`
	IndexedValueConceptVersion int64  `xml:"http://tableau.com/api indexedValueConceptVersion,omitempty"`
}

type IntervalType struct {
	Minutes Minutes `xml:"minutes,attr,omitempty"`
	Hours   Hours   `xml:"hours,attr,omitempty"`
	WeekDay WeekDay `xml:"weekDay,attr,omitempty"`
}

type Intervals struct {
	Interval []IntervalType `xml:"http://tableau.com/api interval"`
}

type JobType struct {
	StatusNotes       StatusNoteListType    `xml:"http://tableau.com/api statusNotes,omitempty"`
	ExtractRefreshJob ExtractRefreshJobType `xml:"http://tableau.com/api extractRefreshJob,omitempty"`
	RunFlowJobType    RunFlowJobType        `xml:"http://tableau.com/api runFlowJobType,omitempty"`
	Id                ResourceIdType        `xml:"id,attr,omitempty"`
	Mode              Mode                  `xml:"mode,attr,omitempty"`
	Type              Type                  `xml:"type,attr,omitempty"`
	Progress          int                   `xml:"progress,attr,omitempty"`
	CreatedAt         time.Time             `xml:"createdAt,attr,omitempty"`
	StartedAt         time.Time             `xml:"startedAt,attr,omitempty"`
	UpdatedAt         time.Time             `xml:"updatedAt,attr,omitempty"`
	CompletedAt       time.Time             `xml:"completedAt,attr,omitempty"`
	FinishCode        int                   `xml:"finishCode,attr,omitempty"`
}

func (t *JobType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T JobType
	var layout struct {
		*T
		CreatedAt   *xsdDateTime `xml:"createdAt,attr,omitempty"`
		StartedAt   *xsdDateTime `xml:"startedAt,attr,omitempty"`
		UpdatedAt   *xsdDateTime `xml:"updatedAt,attr,omitempty"`
		CompletedAt *xsdDateTime `xml:"completedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.StartedAt = (*xsdDateTime)(&layout.T.StartedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	layout.CompletedAt = (*xsdDateTime)(&layout.T.CompletedAt)
	return e.EncodeElement(layout, start)
}
func (t *JobType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T JobType
	var overlay struct {
		*T
		CreatedAt   *xsdDateTime `xml:"createdAt,attr,omitempty"`
		StartedAt   *xsdDateTime `xml:"startedAt,attr,omitempty"`
		UpdatedAt   *xsdDateTime `xml:"updatedAt,attr,omitempty"`
		CompletedAt *xsdDateTime `xml:"completedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.StartedAt = (*xsdDateTime)(&overlay.T.StartedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	overlay.CompletedAt = (*xsdDateTime)(&overlay.T.CompletedAt)
	return d.DecodeElement(&overlay, &start)
}

// May be one of LastDay
type LastDay string

type ListFieldConceptsType struct {
	FieldConcepts []FieldConceptType `xml:"http://tableau.com/api fieldConcepts,omitempty"`
	NextPageToken string             `xml:"http://tableau.com/api nextPageToken"`
}

type ListValueConceptsType struct {
	ValueConcepts []ValueConceptType `xml:"http://tableau.com/api valueConcepts,omitempty"`
	NextPageToken string             `xml:"http://tableau.com/api nextPageToken,omitempty"`
}

type MatchValuesResultType struct {
	AverageMatchWeight float64          `xml:"http://tableau.com/api averageMatchWeight,omitempty"`
	ValueMatches       []ValueMatchType `xml:"http://tableau.com/api valueMatches,omitempty"`
	MatchResponseSet   string           `xml:"http://tableau.com/api matchResponseSet,omitempty"`
}

type MaterializedViewsEnablementConfigType struct {
	MaterializedViewsEnabled bool `xml:"materializedViewsEnabled,attr,omitempty"`
	MaterializeNow           bool `xml:"materializeNow,attr,omitempty"`
}

func (t *MaterializedViewsEnablementConfigType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MaterializedViewsEnablementConfigType
	var overlay struct {
		*T
		MaterializeNow *bool `xml:"materializeNow,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.MaterializeNow = (*bool)(&overlay.T.MaterializeNow)
	return d.DecodeElement(&overlay, &start)
}

// May be one of POST
type Method string

type MetricListType struct {
	Metric []MetricType `xml:"http://tableau.com/api metric,omitempty"`
}

type MetricType struct {
	Site        SiteType       `xml:"http://tableau.com/api site,omitempty"`
	Project     ProjectType    `xml:"http://tableau.com/api project,omitempty"`
	Owner       UserType       `xml:"http://tableau.com/api owner,omitempty"`
	Id          ResourceIdType `xml:"id,attr,omitempty"`
	Name        string         `xml:"name,attr,omitempty"`
	Description string         `xml:"description,attr,omitempty"`
	WebpageUrl  string         `xml:"webpageUrl,attr,omitempty"`
	CreatedAt   time.Time      `xml:"createdAt,attr,omitempty"`
	UpdatedAt   time.Time      `xml:"updatedAt,attr,omitempty"`
	Definition  string         `xml:"definition,attr,omitempty"`
}

func (t *MetricType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MetricType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *MetricType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MetricType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

// May be one of 15, 30
type Minutes int

// May be one of Allow, Deny
type Mode string

type MonthDay string

// May be one of AddComment, ChangeHierarchy, ChangePermissions, Connect, Delete, ExportData, ExportImage, ExportXml, Filter, ProjectLeader, Read, ShareView, ViewComments, ViewUnderlyingData, WebAuthoring, Write
type Name string

type NameCharacteristicsType struct {
	TextPattern TextPatternType `xml:"http://tableau.com/api textPattern,omitempty"`
}

type NameType struct {
	Locale    string  `xml:"http://tableau.com/api locale,omitempty"`
	Name      string  `xml:"http://tableau.com/api name,omitempty"`
	Weight    float64 `xml:"http://tableau.com/api weight,omitempty"`
	IsPrimary bool    `xml:"http://tableau.com/api isPrimary,omitempty"`
}

type PaginationType struct {
	PageNumber     int `xml:"pageNumber,attr"`
	PageSize       int `xml:"pageSize,attr"`
	TotalAvailable int `xml:"totalAvailable,attr"`
}

type ParentType struct {
	Id   ResourceIdType `xml:"id,attr"`
	Type Type           `xml:"type,attr"`
	Name LastDay        `xml:"name,attr,omitempty"`
}

type PermissionsType struct {
	Parent              ParentType                `xml:"http://tableau.com/api parent,omitempty"`
	Flow                FlowType                  `xml:"http://tableau.com/api flow"`
	Database            DatabaseType              `xml:"http://tableau.com/api database"`
	Datasource          DataSourceType            `xml:"http://tableau.com/api datasource"`
	Project             ProjectType               `xml:"http://tableau.com/api project"`
	Table               TableType                 `xml:"http://tableau.com/api table"`
	View                ViewType                  `xml:"http://tableau.com/api view"`
	Workbook            WorkbookType              `xml:"http://tableau.com/api workbook"`
	GranteeCapabilities []GranteeCapabilitiesType `xml:"http://tableau.com/api granteeCapabilities,omitempty"`
}

type ProductVersion struct {
	Value string `xml:",chardata"`
	Build Build  `xml:"build,attr,omitempty"`
}

type ProjectListType struct {
	Project []ProjectType `xml:"http://tableau.com/api project,omitempty"`
}

type ProjectType struct {
	Owner              UserType           `xml:"http://tableau.com/api owner,omitempty"`
	Id                 ResourceIdType     `xml:"id,attr,omitempty"`
	Name               string             `xml:"name,attr,omitempty"`
	Description        string             `xml:"description,attr,omitempty"`
	TopLevelProject    bool               `xml:"topLevelProject,attr,omitempty"`
	ParentProjectId    ResourceIdType     `xml:"parentProjectId,attr,omitempty"`
	CreatedAt          time.Time          `xml:"createdAt,attr,omitempty"`
	UpdatedAt          time.Time          `xml:"updatedAt,attr,omitempty"`
	FavoritesTotal     int                `xml:"favoritesTotal,attr,omitempty"`
	ContentPermissions ContentPermissions `xml:"contentPermissions,attr,omitempty"`
}

func (t *ProjectType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ProjectType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *ProjectType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ProjectType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

// Must match the pattern [0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}
type ResourceIdType string

// Must match the pattern [0-9]{5}.[0-9]{2}.[0-9]{4}.[0-9]{4}
type RestApiVersion string

type RevisionLimitType string

type RevisionListType struct {
	Revision []RevisionType `xml:"http://tableau.com/api revision,omitempty"`
}

type RevisionType struct {
	Publisher      UserType  `xml:"http://tableau.com/api publisher,omitempty"`
	RevisionNumber int       `xml:"revisionNumber,attr,omitempty"`
	PublishedAt    time.Time `xml:"publishedAt,attr,omitempty"`
	Deleted        bool      `xml:"deleted,attr,omitempty"`
	Current        bool      `xml:"current,attr,omitempty"`
	SizeInBytes    int64     `xml:"sizeInBytes,attr,omitempty"`
}

func (t *RevisionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T RevisionType
	var layout struct {
		*T
		PublishedAt *xsdDateTime `xml:"publishedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PublishedAt = (*xsdDateTime)(&layout.T.PublishedAt)
	return e.EncodeElement(layout, start)
}
func (t *RevisionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T RevisionType
	var overlay struct {
		*T
		PublishedAt *xsdDateTime `xml:"publishedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PublishedAt = (*xsdDateTime)(&overlay.T.PublishedAt)
	return d.DecodeElement(&overlay, &start)
}

type RunFlowJobType struct {
	Notes string   `xml:"http://tableau.com/api notes"`
	Flow  FlowType `xml:"http://tableau.com/api flow"`
}

type ScheduleListType struct {
	Schedule []ScheduleType `xml:"http://tableau.com/api schedule,omitempty"`
}

type ScheduleType struct {
	FrequencyDetails FrequencyDetailsType `xml:"http://tableau.com/api frequencyDetails,omitempty"`
	Id               ResourceIdType       `xml:"id,attr,omitempty"`
	Name             string               `xml:"name,attr,omitempty"`
	State            StateType            `xml:"state,attr,omitempty"`
	Priority         int                  `xml:"priority,attr,omitempty"`
	CreatedAt        time.Time            `xml:"createdAt,attr,omitempty"`
	UpdatedAt        time.Time            `xml:"updatedAt,attr,omitempty"`
	Type             Type                 `xml:"type,attr,omitempty"`
	Frequency        Frequency            `xml:"frequency,attr,omitempty"`
	NextRunAt        time.Time            `xml:"nextRunAt,attr,omitempty"`
	EndScheduleAt    time.Time            `xml:"endScheduleAt,attr,omitempty"`
	ExecutionOrder   ExecutionOrder       `xml:"executionOrder,attr,omitempty"`
}

func (t *ScheduleType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ScheduleType
	var layout struct {
		*T
		CreatedAt     *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt     *xsdDateTime `xml:"updatedAt,attr,omitempty"`
		NextRunAt     *xsdDateTime `xml:"nextRunAt,attr,omitempty"`
		EndScheduleAt *xsdDateTime `xml:"endScheduleAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	layout.NextRunAt = (*xsdDateTime)(&layout.T.NextRunAt)
	layout.EndScheduleAt = (*xsdDateTime)(&layout.T.EndScheduleAt)
	return e.EncodeElement(layout, start)
}
func (t *ScheduleType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ScheduleType
	var overlay struct {
		*T
		CreatedAt     *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt     *xsdDateTime `xml:"updatedAt,attr,omitempty"`
		NextRunAt     *xsdDateTime `xml:"nextRunAt,attr,omitempty"`
		EndScheduleAt *xsdDateTime `xml:"endScheduleAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	overlay.NextRunAt = (*xsdDateTime)(&overlay.T.NextRunAt)
	overlay.EndScheduleAt = (*xsdDateTime)(&overlay.T.EndScheduleAt)
	return d.DecodeElement(&overlay, &start)
}

type SemanticsValueType struct {
	NumberValue float64   `xml:"http://tableau.com/api numberValue,omitempty"`
	StringValue string    `xml:"http://tableau.com/api stringValue,omitempty"`
	TimeValue   time.Time `xml:"http://tableau.com/api timeValue,omitempty"`
	BoolValue   bool      `xml:"http://tableau.com/api boolValue,omitempty"`
}

func (t *SemanticsValueType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T SemanticsValueType
	var layout struct {
		*T
		TimeValue *xsdDateTime `xml:"http://tableau.com/api timeValue,omitempty"`
	}
	layout.T = (*T)(t)
	layout.TimeValue = (*xsdDateTime)(&layout.T.TimeValue)
	return e.EncodeElement(layout, start)
}
func (t *SemanticsValueType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SemanticsValueType
	var overlay struct {
		*T
		TimeValue *xsdDateTime `xml:"http://tableau.com/api timeValue,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.TimeValue = (*xsdDateTime)(&overlay.T.TimeValue)
	return d.DecodeElement(&overlay, &start)
}

type ServerInfo struct {
	ProductVersion       ProductVersion `xml:"http://tableau.com/api productVersion"`
	PrepConductorVersion string         `xml:"http://tableau.com/api prepConductorVersion"`
	RestApiVersion       RestApiVersion `xml:"http://tableau.com/api restApiVersion"`
	Platform             string         `xml:"http://tableau.com/api platform"`
	ServerSettings       ServerSettings `xml:"http://tableau.com/api serverSettings"`
}

type ServerSettings struct {
	OAuthEnabled                     bool `xml:"http://tableau.com/api oAuthEnabled"`
	SheetImageMaxAgeFloor            int  `xml:"http://tableau.com/api sheetImageMaxAgeFloor"`
	SheetImageMaxAgeCeiling          int  `xml:"http://tableau.com/api sheetImageMaxAgeCeiling"`
	OfflineInteractionSupportedPhase int  `xml:"http://tableau.com/api offlineInteractionSupportedPhase"`
}

type SiteListType struct {
	Site []SiteType `xml:"http://tableau.com/api site,omitempty"`
}

// May be one of Creator, Explorer, ExplorerCanPublish, Guest, Interactor, Publisher, ReadOnly, ServerAdministrator, SiteAdministrator, SiteAdministratorCreator, SiteAdministratorExplorer, Unlicensed, UnlicensedWithPublish, Viewer, ViewerWithPublish
type SiteRoleType string

type SiteType struct {
	Usage                        Usage             `xml:"http://tableau.com/api usage,omitempty"`
	Id                           ResourceIdType    `xml:"id,attr,omitempty"`
	Name                         string            `xml:"name,attr,omitempty"`
	ContentUrl                   string            `xml:"contentUrl,attr,omitempty"`
	AdminMode                    AdminMode         `xml:"adminMode,attr,omitempty"`
	UserQuota                    int               `xml:"userQuota,attr,omitempty"`
	StorageQuota                 int               `xml:"storageQuota,attr,omitempty"`
	TierCreatorCapacity          int               `xml:"tierCreatorCapacity,attr,omitempty"`
	TierExplorerCapacity         int               `xml:"tierExplorerCapacity,attr,omitempty"`
	TierViewerCapacity           int               `xml:"tierViewerCapacity,attr,omitempty"`
	DisableSubscriptions         bool              `xml:"disableSubscriptions,attr,omitempty"`
	State                        StateType         `xml:"state,attr,omitempty"`
	RevisionHistoryEnabled       bool              `xml:"revisionHistoryEnabled,attr,omitempty"`
	RevisionLimit                RevisionLimitType `xml:"revisionLimit,attr,omitempty"`
	SubscribeOthersEnabled       bool              `xml:"subscribeOthersEnabled,attr,omitempty"`
	AllowSubscriptionAttachments bool              `xml:"allowSubscriptionAttachments,attr,omitempty"`
	GuestAccessEnabled           bool              `xml:"guestAccessEnabled,attr,omitempty"`
	CacheWarmupEnabled           bool              `xml:"cacheWarmupEnabled,attr,omitempty"`
	CommentingEnabled            bool              `xml:"commentingEnabled,attr,omitempty"`
	CacheeWarmupThreshold        int               `xml:"cacheeWarmupThreshold,attr,omitempty"`
	FlowsEnabled                 bool              `xml:"flowsEnabled,attr,omitempty"`
	ExtractEncryptionMode        string            `xml:"extractEncryptionMode,attr,omitempty"`
	MaterializedViewsMode        string            `xml:"materializedViewsMode,attr,omitempty"`
	MobileBiometricsEnabled      bool              `xml:"mobileBiometricsEnabled,attr,omitempty"`
	SheetImageEnabled            bool              `xml:"sheetImageEnabled,attr,omitempty"`
	CatalogingEnabled            bool              `xml:"catalogingEnabled,attr,omitempty"`
	DerivedPermissionsEnabled    bool              `xml:"derivedPermissionsEnabled,attr,omitempty"`
}

// May be one of ServerDefault, SAML
type SiteUserAuthSettingType string

// May be one of Active, Suspended
type StateType string

// May be one of Pending, InProgress, Success, Failed, Cancelled
type Status string

type StatusNoteListType struct {
	StatusNote []StatusNoteType `xml:"http://tableau.com/api statusNote,omitempty"`
}

type StatusNoteType struct {
	Type  Type   `xml:"type,attr"`
	Value string `xml:"value,attr,omitempty"`
	Text  string `xml:"text,attr,omitempty"`
}

type SubscriptionContentType struct {
	Id   ResourceIdType `xml:"id,attr"`
	Type Type           `xml:"type,attr"`
	Name LastDay        `xml:"name,attr,omitempty"`
}

type SubscriptionListType struct {
	Subscription []SubscriptionType `xml:"http://tableau.com/api subscription,omitempty"`
}

type SubscriptionType struct {
	Content     SubscriptionContentType `xml:"http://tableau.com/api content"`
	Schedule    ScheduleType            `xml:"http://tableau.com/api schedule"`
	User        UserType                `xml:"http://tableau.com/api user"`
	Id          ResourceIdType          `xml:"id,attr,omitempty"`
	Subject     string                  `xml:"subject,attr"`
	AttachImage bool                    `xml:"attachImage,attr,omitempty"`
	AttachPdf   bool                    `xml:"attachPdf,attr,omitempty"`
}

type TableListType struct {
	Table []TableType `xml:"http://tableau.com/api table,omitempty"`
}

type TableType struct {
	Site              SiteType       `xml:"http://tableau.com/api site,omitempty"`
	Contact           UserType       `xml:"http://tableau.com/api contact"`
	Certifier         UserType       `xml:"http://tableau.com/api certifier"`
	Id                ResourceIdType `xml:"id,attr,omitempty"`
	Name              string         `xml:"name,attr,omitempty"`
	Description       string         `xml:"description,attr,omitempty"`
	IsEmbedded        bool           `xml:"isEmbedded,attr,omitempty"`
	IsCertified       bool           `xml:"isCertified,attr,omitempty"`
	CertificationNote string         `xml:"certificationNote,attr,omitempty"`
	Schema            string         `xml:"schema,attr,omitempty"`
}

type TableauCredentialsType struct {
	Site     SiteType `xml:"http://tableau.com/api site"`
	User     UserType `xml:"http://tableau.com/api user,omitempty"`
	Name     string   `xml:"name,attr,omitempty"`
	Password string   `xml:"password,attr,omitempty"`
	Token    string   `xml:"token,attr,omitempty"`
}

type TagListType struct {
	Tag []TagType `xml:"http://tableau.com/api tag,omitempty"`
}

type TagType struct {
	Label string `xml:"label,attr"`
}

type TaskExtractRefreshType struct {
	Schedule               ScheduleType   `xml:"http://tableau.com/api schedule,omitempty"`
	View                   ViewType       `xml:"http://tableau.com/api view"`
	Workbook               WorkbookType   `xml:"http://tableau.com/api workbook"`
	Datasource             DataSourceType `xml:"http://tableau.com/api datasource"`
	Id                     string         `xml:"id,attr,omitempty"`
	Priority               int            `xml:"priority,attr,omitempty"`
	ConsecutiveFailedCount int            `xml:"consecutiveFailedCount,attr,omitempty"`
	Type                   string         `xml:"type,attr,omitempty"`
	Incremental            bool           `xml:"incremental,attr,omitempty"`
}

type TaskListType struct {
	Task []TaskType `xml:"http://tableau.com/api task,omitempty"`
}

type TaskRunFlowType struct {
	Schedule               ScheduleType `xml:"http://tableau.com/api schedule,omitempty"`
	Flow                   FlowType     `xml:"http://tableau.com/api flow"`
	Id                     string       `xml:"id,attr,omitempty"`
	Priority               int          `xml:"priority,attr,omitempty"`
	ConsecutiveFailedCount int          `xml:"consecutiveFailedCount,attr,omitempty"`
	Type                   string       `xml:"type,attr,omitempty"`
}

type TaskType struct {
	ExtractRefresh         TaskExtractRefreshType `xml:"http://tableau.com/api extractRefresh"`
	FlowRun                TaskRunFlowType        `xml:"http://tableau.com/api flowRun"`
	Schedule               ScheduleType           `xml:"http://tableau.com/api schedule,omitempty"`
	RunNow                 bool                   `xml:"runNow,attr,omitempty"`
	State                  string                 `xml:"state,attr,omitempty"`
	Id                     string                 `xml:"id,attr,omitempty"`
	Priority               int                    `xml:"priority,attr,omitempty"`
	ConsecutiveFailedCount int                    `xml:"consecutiveFailedCount,attr,omitempty"`
	Type                   string                 `xml:"type,attr,omitempty"`
}

type TextPatternType struct {
	Regex []string `xml:"http://tableau.com/api regex,omitempty"`
}

type TsRequest struct {
	Column             ColumnType                    `xml:"http://tableau.com/api column"`
	Connection         ConnectionType                `xml:"http://tableau.com/api connection"`
	Connections        ConnectionListType            `xml:"http://tableau.com/api connections"`
	Credentials        TableauCredentialsType        `xml:"http://tableau.com/api credentials"`
	DataAlert          DataAlertType                 `xml:"http://tableau.com/api dataAlert"`
	DataQualityWarning DataQualityWarningType        `xml:"http://tableau.com/api dataQualityWarning"`
	DataRole           DataRoleType                  `xml:"http://tableau.com/api dataRole"`
	Database           DatabaseType                  `xml:"http://tableau.com/api database"`
	Datasource         DataSourceType                `xml:"http://tableau.com/api datasource"`
	DistinctValues     DistinctValueListType         `xml:"http://tableau.com/api distinctValues"`
	Favorite           FavoriteType                  `xml:"http://tableau.com/api favorite"`
	Field              FieldType                     `xml:"http://tableau.com/api field"`
	FieldConcept       FieldConceptType              `xml:"http://tableau.com/api fieldConcept"`
	Flow               FlowType                      `xml:"http://tableau.com/api flow"`
	FlowWarnings       FlowWarningsListContainerType `xml:"http://tableau.com/api flowWarnings"`
	Group              GroupType                     `xml:"http://tableau.com/api group"`
	Metric             MetricType                    `xml:"http://tableau.com/api metric"`
	Permissions        PermissionsType               `xml:"http://tableau.com/api permissions"`
	Project            ProjectType                   `xml:"http://tableau.com/api project"`
	Schedule           ScheduleType                  `xml:"http://tableau.com/api schedule"`
	Site               SiteType                      `xml:"http://tableau.com/api site"`
	Table              TableType                     `xml:"http://tableau.com/api table"`
	Tags               TagListType                   `xml:"http://tableau.com/api tags"`
	User               UserType                      `xml:"http://tableau.com/api user"`
	Webhook            WebhookType                   `xml:"http://tableau.com/api webhook"`
	Workbook           WorkbookType                  `xml:"http://tableau.com/api workbook"`
	Subscription       SubscriptionType              `xml:"http://tableau.com/api subscription"`
	Task               TaskType                      `xml:"http://tableau.com/api task"`
}

type TsResponse struct {
	Pagination              PaginationType              `xml:"http://tableau.com/api pagination"`
	Columns                 ColumnListType              `xml:"http://tableau.com/api columns"`
	Databases               DatabaseListType            `xml:"http://tableau.com/api databases"`
	Datasources             DataSourceListType          `xml:"http://tableau.com/api datasources"`
	Extracts                ExtractListType             `xml:"http://tableau.com/api extracts"`
	Flows                   FlowListType                `xml:"http://tableau.com/api flows"`
	FlowOutputSteps         FlowOutputStepListType      `xml:"http://tableau.com/api flowOutputSteps"`
	Groups                  GroupListType               `xml:"http://tableau.com/api groups"`
	Metrics                 MetricListType              `xml:"http://tableau.com/api metrics"`
	Projects                ProjectListType             `xml:"http://tableau.com/api projects"`
	Revisions               RevisionListType            `xml:"http://tableau.com/api revisions"`
	Schedules               ScheduleListType            `xml:"http://tableau.com/api schedules"`
	Sites                   SiteListType                `xml:"http://tableau.com/api sites"`
	Tables                  TableListType               `xml:"http://tableau.com/api tables"`
	Users                   UserListType                `xml:"http://tableau.com/api users"`
	Workbooks               WorkbookListType            `xml:"http://tableau.com/api workbooks"`
	Subscriptions           SubscriptionListType        `xml:"http://tableau.com/api subscriptions"`
	BackgroundJob           BackgroundJobType           `xml:"http://tableau.com/api backgroundJob"`
	BackgroundJobs          BackgroundJobListType       `xml:"http://tableau.com/api backgroundJobs"`
	Column                  ColumnType                  `xml:"http://tableau.com/api column"`
	Connection              ConnectionType              `xml:"http://tableau.com/api connection"`
	Connections             ConnectionListType          `xml:"http://tableau.com/api connections"`
	Credentials             TableauCredentialsType      `xml:"http://tableau.com/api credentials"`
	DataAlert               DataAlertType               `xml:"http://tableau.com/api dataAlert"`
	DataAlerts              DataAlertListType           `xml:"http://tableau.com/api dataAlerts"`
	DataQualityWarning      DataQualityWarningType      `xml:"http://tableau.com/api dataQualityWarning"`
	DataQualityWarningList  DataQualityWarningListType  `xml:"http://tableau.com/api dataQualityWarningList"`
	DataRole                DataRoleType                `xml:"http://tableau.com/api dataRole"`
	Database                DatabaseType                `xml:"http://tableau.com/api database"`
	Datasource              DataSourceType              `xml:"http://tableau.com/api datasource"`
	Error                   ErrorType                   `xml:"http://tableau.com/api error"`
	Favorites               FavoriteListType            `xml:"http://tableau.com/api favorites"`
	FileUpload              FileUploadType              `xml:"http://tableau.com/api fileUpload"`
	Flow                    FlowType                    `xml:"http://tableau.com/api flow"`
	Group                   GroupType                   `xml:"http://tableau.com/api group"`
	Job                     JobType                     `xml:"http://tableau.com/api job"`
	Metric                  MetricType                  `xml:"http://tableau.com/api metric"`
	Permissions             PermissionsType             `xml:"http://tableau.com/api permissions"`
	Project                 ProjectType                 `xml:"http://tableau.com/api project"`
	DataAlertsRecipient     DataAlertsRecipientType     `xml:"http://tableau.com/api dataAlertsRecipient"`
	DataAlertsRecipientList DataAlertsRecipientListType `xml:"http://tableau.com/api dataAlertsRecipientList"`
	Schedule                ScheduleType                `xml:"http://tableau.com/api schedule"`
	ServerInfo              ServerInfo                  `xml:"http://tableau.com/api serverInfo"`
	ServerSettings          ServerSettings              `xml:"http://tableau.com/api serverSettings"`
	Site                    SiteType                    `xml:"http://tableau.com/api site"`
	Table                   TableType                   `xml:"http://tableau.com/api table"`
	Tags                    TagListType                 `xml:"http://tableau.com/api tags"`
	User                    UserType                    `xml:"http://tableau.com/api user"`
	View                    ViewType                    `xml:"http://tableau.com/api view"`
	Views                   ViewListType                `xml:"http://tableau.com/api views"`
	Webhook                 WebhookType                 `xml:"http://tableau.com/api webhook"`
	Webhooks                WebhookListType             `xml:"http://tableau.com/api webhooks"`
	WebhookTestResult       WebhookTestResultType       `xml:"http://tableau.com/api webhookTestResult"`
	Workbook                WorkbookType                `xml:"http://tableau.com/api workbook"`
	Subscription            SubscriptionType            `xml:"http://tableau.com/api subscription"`
	Task                    TaskType                    `xml:"http://tableau.com/api task"`
	Tasks                   TaskListType                `xml:"http://tableau.com/api tasks"`
	Warnings                WarningListType             `xml:"http://tableau.com/api warnings"`
	Degradations            DegradationListType         `xml:"http://tableau.com/api degradations"`
	ListFieldConcepts       ListFieldConceptsType       `xml:"http://tableau.com/api listFieldConcepts"`
	FieldMatches            FieldMatchListType          `xml:"http://tableau.com/api fieldMatches"`
	ValueMatches            MatchValuesResultType       `xml:"http://tableau.com/api valueMatches"`
	ValueConceptCount       ValueConceptCountType       `xml:"http://tableau.com/api valueConceptCount"`
	ListValueConcepts       ListValueConceptsType       `xml:"http://tableau.com/api listValueConcepts"`
	FieldConcept            FieldConceptType            `xml:"http://tableau.com/api fieldConcept"`
	GetIndexingStatus       IndexingStatusType          `xml:"http://tableau.com/api getIndexingStatus"`
}

// May be one of CountOfUsersAddedToGroup, CountOfUsersAddedToSite, CountOfUsersWithInsufficientLicenses, CountOfUsersInActiveDirectoryGroup, CountOfUsersProcessed, CountOfUsersSkipped, CountOfUsersInformationUpdated, CountOfUsersSiteRoleUpdated, CountOfUsersRemovedFromGroup, CountOfUsersUnlicensed
type Type string

// Must match the pattern https?://.+
type Url string

type Usage struct {
	NumUsers int `xml:"numUsers,attr"`
	Storage  int `xml:"storage,attr"`
}

type UserListType struct {
	User []UserType `xml:"http://tableau.com/api user,omitempty"`
}

type UserType struct {
	Domain             DomainDirectiveType     `xml:"http://tableau.com/api domain,omitempty"`
	Id                 ResourceIdType          `xml:"id,attr,omitempty"`
	Name               string                  `xml:"name,attr,omitempty"`
	FullName           string                  `xml:"fullName,attr,omitempty"`
	Email              string                  `xml:"email,attr,omitempty"`
	Password           string                  `xml:"password,attr,omitempty"`
	SiteRole           SiteRoleType            `xml:"siteRole,attr,omitempty"`
	AuthSetting        SiteUserAuthSettingType `xml:"authSetting,attr,omitempty"`
	LastLogin          time.Time               `xml:"lastLogin,attr,omitempty"`
	ExternalAuthUserId string                  `xml:"externalAuthUserId,attr,omitempty"`
}

func (t *UserType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T UserType
	var layout struct {
		*T
		LastLogin *xsdDateTime `xml:"lastLogin,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LastLogin = (*xsdDateTime)(&layout.T.LastLogin)
	return e.EncodeElement(layout, start)
}
func (t *UserType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T UserType
	var overlay struct {
		*T
		LastLogin *xsdDateTime `xml:"lastLogin,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LastLogin = (*xsdDateTime)(&overlay.T.LastLogin)
	return d.DecodeElement(&overlay, &start)
}

type ValueCharacteristicsType struct {
	TextPattern TextPatternType `xml:"http://tableau.com/api textPattern,omitempty"`
}

type ValueConceptCountType struct {
	Count int64 `xml:"http://tableau.com/api count"`
}

type ValueConceptSignature struct {
	FieldConceptURI string `xml:"http://tableau.com/api fieldConceptURI"`
	ValueConceptURI string `xml:"http://tableau.com/api valueConceptURI"`
}

type ValueConceptType struct {
	Uri                   string                  `xml:"http://tableau.com/api uri"`
	FieldConceptURI       string                  `xml:"http://tableau.com/api fieldConceptURI"`
	Names                 []NameType              `xml:"http://tableau.com/api names,omitempty"`
	NameCharacteristics   NameCharacteristicsType `xml:"http://tableau.com/api nameCharacteristics,omitempty"`
	Description           string                  `xml:"http://tableau.com/api description,omitempty"`
	ParentValueConceptURI string                  `xml:"http://tableau.com/api parentValueConceptURI,omitempty"`
	Value                 SemanticsValueType      `xml:"http://tableau.com/api value,omitempty"`
}

type ValueMatchType struct {
	Value                 SemanticsValueType    `xml:"http://tableau.com/api value,omitempty"`
	Weight                float64               `xml:"http://tableau.com/api weight,omitempty"`
	ValueConcept          ValueConceptType      `xml:"http://tableau.com/api valueConcept,omitempty"`
	ValueConceptSignature ValueConceptSignature `xml:"http://tableau.com/api valueConceptSignature,omitempty"`
}

type ValueSourceType struct {
	DatasourceValueStore DataSourceValueStoreType `xml:"http://tableau.com/api datasourceValueStore,omitempty"`
}

type ViewListType struct {
	View []ViewType `xml:"http://tableau.com/api view,omitempty"`
}

type ViewType struct {
	Workbook       WorkbookType   `xml:"http://tableau.com/api workbook,omitempty"`
	Owner          UserType       `xml:"http://tableau.com/api owner,omitempty"`
	Project        ProjectType    `xml:"http://tableau.com/api project,omitempty"`
	Tags           TagListType    `xml:"http://tableau.com/api tags,omitempty"`
	Usage          Anon4          `xml:"http://tableau.com/api usage,omitempty"`
	Id             ResourceIdType `xml:"id,attr,omitempty"`
	Name           string         `xml:"name,attr,omitempty"`
	ContentUrl     string         `xml:"contentUrl,attr,omitempty"`
	CreatedAt      time.Time      `xml:"createdAt,attr,omitempty"`
	UpdatedAt      time.Time      `xml:"updatedAt,attr,omitempty"`
	SheetType      string         `xml:"sheetType,attr,omitempty"`
	FavoritesTotal int            `xml:"favoritesTotal,attr,omitempty"`
	Hidden         bool           `xml:"hidden,attr,omitempty"`
	RecentlyViewed bool           `xml:"recentlyViewed,attr,omitempty"`
	ViewUrlName    string         `xml:"viewUrlName,attr,omitempty"`
}

func (t *ViewType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ViewType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *ViewType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ViewType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

type WarningListType struct {
	Warning []WarningType `xml:"http://tableau.com/api warning,omitempty"`
}

type WarningType struct {
	Message   string `xml:"message,attr,omitempty"`
	Id        string `xml:"id,attr,omitempty"`
	ErrorCode string `xml:"errorCode,attr,omitempty"`
}

type WebhookDestinationHttpType struct {
	Method Method `xml:"method,attr,omitempty"`
	Url    Url    `xml:"url,attr,omitempty"`
}

type WebhookDestinationType struct {
	Webhookdestinationhttp WebhookDestinationHttpType `xml:"http://tableau.com/api webhook-destination-http"`
}

type WebhookListType struct {
	Webhook []WebhookType `xml:"http://tableau.com/api webhook,omitempty"`
}

type WebhookSourceEventDatasourceCreatedType struct {
}

type WebhookSourceEventDatasourceDeletedType struct {
}

type WebhookSourceEventDatasourceRefreshFailedType struct {
}

type WebhookSourceEventDatasourceRefreshStartedType struct {
}

type WebhookSourceEventDatasourceRefreshSucceededType struct {
}

type WebhookSourceEventDatasourceUpdatedType struct {
}

type WebhookSourceEventViewDeletedType struct {
}

type WebhookSourceEventWorkbookCreatedType struct {
}

type WebhookSourceEventWorkbookDeletedType struct {
}

type WebhookSourceEventWorkbookRefreshFailedType struct {
}

type WebhookSourceEventWorkbookRefreshStartedType struct {
}

type WebhookSourceEventWorkbookRefreshSucceededType struct {
}

type WebhookSourceEventWorkbookUpdatedType struct {
}

type WebhookSourceType struct {
	Webhooksourceeventdatasourcerefreshstarted   WebhookSourceEventDatasourceRefreshStartedType   `xml:"http://tableau.com/api webhook-source-event-datasource-refresh-started"`
	Webhooksourceeventdatasourcerefreshsucceeded WebhookSourceEventDatasourceRefreshSucceededType `xml:"http://tableau.com/api webhook-source-event-datasource-refresh-succeeded"`
	Webhooksourceeventdatasourcerefreshfailed    WebhookSourceEventDatasourceRefreshFailedType    `xml:"http://tableau.com/api webhook-source-event-datasource-refresh-failed"`
	Webhooksourceeventdatasourceupdated          WebhookSourceEventDatasourceUpdatedType          `xml:"http://tableau.com/api webhook-source-event-datasource-updated"`
	Webhooksourceeventdatasourcecreated          WebhookSourceEventDatasourceCreatedType          `xml:"http://tableau.com/api webhook-source-event-datasource-created"`
	Webhooksourceeventdatasourcedeleted          WebhookSourceEventDatasourceDeletedType          `xml:"http://tableau.com/api webhook-source-event-datasource-deleted"`
	Webhooksourceeventworkbookupdated            WebhookSourceEventWorkbookUpdatedType            `xml:"http://tableau.com/api webhook-source-event-workbook-updated"`
	Webhooksourceeventworkbookcreated            WebhookSourceEventWorkbookCreatedType            `xml:"http://tableau.com/api webhook-source-event-workbook-created"`
	Webhooksourceeventworkbookdeleted            WebhookSourceEventWorkbookDeletedType            `xml:"http://tableau.com/api webhook-source-event-workbook-deleted"`
	Webhooksourceeventviewdeleted                WebhookSourceEventViewDeletedType                `xml:"http://tableau.com/api webhook-source-event-view-deleted"`
	Webhooksourceeventworkbookrefreshstarted     WebhookSourceEventWorkbookRefreshStartedType     `xml:"http://tableau.com/api webhook-source-event-workbook-refresh-started"`
	Webhooksourceeventworkbookrefreshsucceeded   WebhookSourceEventWorkbookRefreshSucceededType   `xml:"http://tableau.com/api webhook-source-event-workbook-refresh-succeeded"`
	Webhooksourceeventworkbookrefreshfailed      WebhookSourceEventWorkbookRefreshFailedType      `xml:"http://tableau.com/api webhook-source-event-workbook-refresh-failed"`
}

type WebhookTestResultType struct {
	Body   string         `xml:"http://tableau.com/api body"`
	Id     ResourceIdType `xml:"id,attr,omitempty"`
	Status int            `xml:"status,attr,omitempty"`
}

type WebhookType struct {
	Webhooksource      WebhookSourceType      `xml:"http://tableau.com/api webhook-source"`
	Webhookdestination WebhookDestinationType `xml:"http://tableau.com/api webhook-destination"`
	Owner              UserType               `xml:"http://tableau.com/api owner,omitempty"`
	Id                 ResourceIdType         `xml:"id,attr,omitempty"`
	Name               string                 `xml:"name,attr,omitempty"`
	Enabled            bool                   `xml:"enabled,attr,omitempty"`
	CreatedAt          time.Time              `xml:"createdAt,attr,omitempty"`
	UpdatedAt          time.Time              `xml:"updatedAt,attr,omitempty"`
}

func (t *WebhookType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WebhookType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *WebhookType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WebhookType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

// May be one of Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
type WeekDay string

type WorkbookListType struct {
	Workbook []WorkbookType `xml:"http://tableau.com/api workbook,omitempty"`
}

type WorkbookType struct {
	Connections                       ConnectionListType                    `xml:"http://tableau.com/api connections,omitempty"`
	ConnectionCredentials             ConnectionCredentialsType             `xml:"http://tableau.com/api connectionCredentials,omitempty"`
	Site                              SiteType                              `xml:"http://tableau.com/api site,omitempty"`
	Project                           ProjectType                           `xml:"http://tableau.com/api project,omitempty"`
	Owner                             UserType                              `xml:"http://tableau.com/api owner,omitempty"`
	Tags                              TagListType                           `xml:"http://tableau.com/api tags,omitempty"`
	Views                             ViewListType                          `xml:"http://tableau.com/api views,omitempty"`
	MaterializedViewsEnablementConfig MaterializedViewsEnablementConfigType `xml:"http://tableau.com/api materializedViewsEnablementConfig,omitempty"`
	Id                                ResourceIdType                        `xml:"id,attr,omitempty"`
	Name                              string                                `xml:"name,attr,omitempty"`
	Description                       string                                `xml:"description,attr,omitempty"`
	ContentUrl                        string                                `xml:"contentUrl,attr,omitempty"`
	WebpageUrl                        string                                `xml:"webpageUrl,attr,omitempty"`
	ShowTabs                          bool                                  `xml:"showTabs,attr,omitempty"`
	ThumbnailsUserId                  ResourceIdType                        `xml:"thumbnailsUserId,attr,omitempty"`
	Size                              int                                   `xml:"size,attr,omitempty"`
	CreatedAt                         time.Time                             `xml:"createdAt,attr,omitempty"`
	UpdatedAt                         time.Time                             `xml:"updatedAt,attr,omitempty"`
	SheetCount                        int                                   `xml:"sheetCount,attr,omitempty"`
	HasExtracts                       bool                                  `xml:"hasExtracts,attr,omitempty"`
	EncryptExtracts                   string                                `xml:"encryptExtracts,attr,omitempty"`
	RecentlyViewed                    bool                                  `xml:"recentlyViewed,attr,omitempty"`
	DefaultViewId                     ResourceIdType                        `xml:"defaultViewId,attr,omitempty"`
}

func (t *WorkbookType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T WorkbookType
	var layout struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.CreatedAt = (*xsdDateTime)(&layout.T.CreatedAt)
	layout.UpdatedAt = (*xsdDateTime)(&layout.T.UpdatedAt)
	return e.EncodeElement(layout, start)
}
func (t *WorkbookType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T WorkbookType
	var overlay struct {
		*T
		CreatedAt *xsdDateTime `xml:"createdAt,attr,omitempty"`
		UpdatedAt *xsdDateTime `xml:"updatedAt,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.CreatedAt = (*xsdDateTime)(&overlay.T.CreatedAt)
	overlay.UpdatedAt = (*xsdDateTime)(&overlay.T.UpdatedAt)
	return d.DecodeElement(&overlay, &start)
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("15:04:05.999999999")), nil
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
