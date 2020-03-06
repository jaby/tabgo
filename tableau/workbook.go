package tableau

import "encoding/xml"

type Datasource struct {
	Text          string `xml:",chardata"`
	Hasconnection string `xml:"hasconnection,attr"`
	Inline        string `xml:"inline,attr"`
	Name          string `xml:"name,attr"`
	Version       string `xml:"version,attr"`
	Aliases       struct {
		Text    string `xml:",chardata"`
		Enabled string `xml:"enabled,attr"`
	} `xml:"aliases"`
	Column []struct {
		Text            string `xml:",chardata"`
		Caption         string `xml:"caption,attr"`
		Datatype        string `xml:"datatype,attr"`
		Name            string `xml:"name,attr"`
		ParamDomainType string `xml:"param-domain-type,attr"`
		Role            string `xml:"role,attr"`
		Type            string `xml:"type,attr"`
		Value           string `xml:"value,attr"`
		DefaultFormat   string `xml:"default-format,attr"`
		SemanticRole    string `xml:"semantic-role,attr"`
		Hidden          string `xml:"hidden,attr"`
		Aggregation     string `xml:"aggregation,attr"`
		AutoColumn      string `xml:"auto-column,attr"`
		Calculation     struct {
			Text           string `xml:",chardata"`
			Class          string `xml:"class,attr"`
			Formula        string `xml:"formula,attr"`
			ScopeIsolation string `xml:"scope-isolation,attr"`
			Column         string `xml:"column,attr"`
			Decimals       string `xml:"decimals,attr"`
			Peg            string `xml:"peg,attr"`
			SizeParameter  string `xml:"size-parameter,attr"`
			Bin            []struct {
				Text      string   `xml:",chardata"`
				AttrValue string   `xml:"value,attr"`
				Value     []string `xml:"value"`
			} `xml:"bin"`
		} `xml:"calculation"`
		Range struct {
			Text        string `xml:",chardata"`
			Granularity string `xml:"granularity,attr"`
			Max         string `xml:"max,attr"`
			Min         string `xml:"min,attr"`
		} `xml:"range"`
		Aliases struct {
			Text  string `xml:",chardata"`
			Alias []struct {
				Text  string `xml:",chardata"`
				Key   string `xml:"key,attr"`
				Value string `xml:"value,attr"`
			} `xml:"alias"`
		} `xml:"aliases"`
	} `xml:"column"`
	Connection struct {
		Text             string `xml:",chardata"`
		Class            string `xml:"class,attr"`
		Cleaning         string `xml:"cleaning,attr"`
		Compat           string `xml:"compat,attr"`
		DataRefreshTime  string `xml:"dataRefreshTime,attr"`
		Filename         string `xml:"filename,attr"`
		Validate         string `xml:"validate,attr"`
		NamedConnections struct {
			Text            string `xml:",chardata"`
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
					Service              string `xml:"service,attr"`
					Sslmode              string `xml:"sslmode,attr"`
					Username             string `xml:"username,attr"`
					Dbname               string `xml:"dbname,attr"`
					MinimumDriverVersion string `xml:"minimum-driver-version,attr"`
					OdbcNativeProtocol   string `xml:"odbc-native-protocol,attr"`
				} `xml:"connection"`
			} `xml:"named-connection"`
		} `xml:"named-connections"`
		Relation struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Table   string `xml:"table,attr"`
			Type    string `xml:"type,attr"`
			Columns struct {
				Text    string `xml:",chardata"`
				Header  string `xml:"header,attr"`
				Outcome string `xml:"outcome,attr"`
				Column  []struct {
					Text     string `xml:",chardata"`
					Datatype string `xml:"datatype,attr"`
					Name     string `xml:"name,attr"`
					Ordinal  string `xml:"ordinal,attr"`
				} `xml:"column"`
			} `xml:"columns"`
		} `xml:"relation"`
		MetadataRecords struct {
			Text           string `xml:",chardata"`
			MetadataRecord []struct {
				Text         string `xml:",chardata"`
				Class        string `xml:"class,attr"`
				RemoteName   string `xml:"remote-name"`
				RemoteType   string `xml:"remote-type"`
				LocalName    string `xml:"local-name"`
				ParentName   string `xml:"parent-name"`
				RemoteAlias  string `xml:"remote-alias"`
				Ordinal      string `xml:"ordinal"`
				LocalType    string `xml:"local-type"`
				Aggregation  string `xml:"aggregation"`
				ContainsNull string `xml:"contains-null"`
				Attributes   struct {
					Text      string `xml:",chardata"`
					Attribute []struct {
						Text     string `xml:",chardata"`
						Datatype string `xml:"datatype,attr"`
						Name     string `xml:"name,attr"`
					} `xml:"attribute"`
				} `xml:"attributes"`
				Collation struct {
					Text string `xml:",chardata"`
					Flag string `xml:"flag,attr"`
					Name string `xml:"name,attr"`
				} `xml:"collation"`
				Precision string `xml:"precision"`
			} `xml:"metadata-record"`
		} `xml:"metadata-records"`
	} `xml:"connection"`
	ColumnInstance struct {
		Text               string `xml:",chardata"`
		Column             string `xml:"column,attr"`
		Derivation         string `xml:"derivation,attr"`
		ForecastColumnBase string `xml:"forecast-column-base,attr"`
		ForecastColumnType string `xml:"forecast-column-type,attr"`
		Name               string `xml:"name,attr"`
		Pivot              string `xml:"pivot,attr"`
		Type               string `xml:"type,attr"`
	} `xml:"column-instance"`
	Group struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		NameStyle   string `xml:"name-style,attr"`
		UiBuilder   string `xml:"ui-builder,attr"`
		Groupfilter struct {
			Text         string `xml:",chardata"`
			Count        string `xml:"count,attr"`
			End          string `xml:"end,attr"`
			Function     string `xml:"function,attr"`
			Units        string `xml:"units,attr"`
			UiMarker     string `xml:"ui-marker,attr"`
			UiTopByField string `xml:"ui-top-by-field,attr"`
			Groupfilter  struct {
				Text        string `xml:",chardata"`
				Direction   string `xml:"direction,attr"`
				Expression  string `xml:"expression,attr"`
				Function    string `xml:"function,attr"`
				UiMarker    string `xml:"ui-marker,attr"`
				Groupfilter struct {
					Text          string `xml:",chardata"`
					Function      string `xml:"function,attr"`
					Level         string `xml:"level,attr"`
					UiEnumeration string `xml:"ui-enumeration,attr"`
					UiMarker      string `xml:"ui-marker,attr"`
				} `xml:"groupfilter"`
			} `xml:"groupfilter"`
		} `xml:"groupfilter"`
	} `xml:"group"`
	DrillPaths struct {
		Text      string `xml:",chardata"`
		DrillPath []struct {
			Text  string   `xml:",chardata"`
			Name  string   `xml:"name,attr"`
			Field []string `xml:"field"`
		} `xml:"drill-path"`
	} `xml:"drill-paths"`
	Folder []struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name,attr"`
		Role       string `xml:"role,attr"`
		FolderItem []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
			Type string `xml:"type,attr"`
		} `xml:"folder-item"`
	} `xml:"folder"`
	Layout struct {
		Text              string `xml:",chardata"`
		DimOrdering       string `xml:"dim-ordering,attr"`
		DimPercentage     string `xml:"dim-percentage,attr"`
		GroupPercentage   string `xml:"group-percentage,attr"`
		MeasureOrdering   string `xml:"measure-ordering,attr"`
		MeasurePercentage string `xml:"measure-percentage,attr"`
		ShowStructure     string `xml:"show-structure,attr"`
	} `xml:"layout"`
	SemanticValues struct {
		Text          string `xml:",chardata"`
		SemanticValue struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key,attr"`
			Value string `xml:"value,attr"`
		} `xml:"semantic-value"`
	} `xml:"semantic-values"`
	DefaultSorts struct {
		Text string `xml:",chardata"`
		Sort struct {
			Text       string `xml:",chardata"`
			Class      string `xml:"class,attr"`
			Column     string `xml:"column,attr"`
			Direction  string `xml:"direction,attr"`
			Dictionary struct {
				Text   string   `xml:",chardata"`
				Bucket []string `xml:"bucket"`
			} `xml:"dictionary"`
		} `xml:"sort"`
	} `xml:"default-sorts"`
	DatasourceDependencies struct {
		Text       string `xml:",chardata"`
		Datasource string `xml:"datasource,attr"`
		Column     []struct {
			Text            string `xml:",chardata"`
			Caption         string `xml:"caption,attr"`
			Datatype        string `xml:"datatype,attr"`
			Name            string `xml:"name,attr"`
			ParamDomainType string `xml:"param-domain-type,attr"`
			Role            string `xml:"role,attr"`
			Type            string `xml:"type,attr"`
			Value           string `xml:"value,attr"`
			Calculation     struct {
				Text    string `xml:",chardata"`
				Class   string `xml:"class,attr"`
				Formula string `xml:"formula,attr"`
			} `xml:"calculation"`
			Range struct {
				Text        string `xml:",chardata"`
				Granularity string `xml:"granularity,attr"`
				Max         string `xml:"max,attr"`
				Min         string `xml:"min,attr"`
			} `xml:"range"`
		} `xml:"column"`
	} `xml:"datasource-dependencies"`
}

type Workbook struct {
	XMLName        xml.Name `xml:"workbook"`
	Text           string   `xml:",chardata"`
	SourcePlatform string   `xml:"source-platform,attr"`
	Version        string   `xml:"version,attr"`
	User           string   `xml:"user,attr"`
	Preferences    struct {
		Text       string `xml:",chardata"`
		Preference []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"preference"`
	} `xml:"preferences"`
	Datasources struct {
		Text       string       `xml:",chardata"`
		Datasource []Datasource `xml:"datasource"`
	} `xml:"datasources"`
	Mapsources struct {
		Text      string `xml:",chardata"`
		Mapsource struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"mapsource"`
	} `xml:"mapsources"`
	Worksheets struct {
		Text      string `xml:",chardata"`
		Worksheet []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Table struct {
				Text string `xml:",chardata"`
				View struct {
					Text        string `xml:",chardata"`
					Datasources struct {
						Text       string `xml:",chardata"`
						Datasource []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"datasource"`
					} `xml:"datasources"`
					Mapsources struct {
						Text      string `xml:",chardata"`
						Mapsource struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"mapsource"`
					} `xml:"mapsources"`
					DatasourceDependencies []struct {
						Text       string `xml:",chardata"`
						Datasource string `xml:"datasource,attr"`
						Column     []struct {
							Text            string `xml:",chardata"`
							Datatype        string `xml:"datatype,attr"`
							Name            string `xml:"name,attr"`
							Role            string `xml:"role,attr"`
							SemanticRole    string `xml:"semantic-role,attr"`
							Type            string `xml:"type,attr"`
							DefaultFormat   string `xml:"default-format,attr"`
							Caption         string `xml:"caption,attr"`
							ParamDomainType string `xml:"param-domain-type,attr"`
							Value           string `xml:"value,attr"`
							Calculation     struct {
								Text    string `xml:",chardata"`
								Class   string `xml:"class,attr"`
								Formula string `xml:"formula,attr"`
							} `xml:"calculation"`
							Range struct {
								Text        string `xml:",chardata"`
								Granularity string `xml:"granularity,attr"`
								Max         string `xml:"max,attr"`
								Min         string `xml:"min,attr"`
							} `xml:"range"`
						} `xml:"column"`
						ColumnInstance []struct {
							Text       string `xml:",chardata"`
							Column     string `xml:"column,attr"`
							Derivation string `xml:"derivation,attr"`
							Name       string `xml:"name,attr"`
							Pivot      string `xml:"pivot,attr"`
							Type       string `xml:"type,attr"`
						} `xml:"column-instance"`
					} `xml:"datasource-dependencies"`
					Aggregation struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value,attr"`
					} `xml:"aggregation"`
					Filter []struct {
						Text        string `xml:",chardata"`
						Class       string `xml:"class,attr"`
						Column      string `xml:"column,attr"`
						Groupfilter struct {
							Text     string `xml:",chardata"`
							Function string `xml:"function,attr"`
							Level    string `xml:"level,attr"`
						} `xml:"groupfilter"`
					} `xml:"filter"`
					Slices struct {
						Text   string   `xml:",chardata"`
						Column []string `xml:"column"`
					} `xml:"slices"`
				} `xml:"view"`
				Style struct {
					Text      string `xml:",chardata"`
					StyleRule []struct {
						Text    string `xml:",chardata"`
						Element string `xml:"element,attr"`
						Format  []struct {
							Text  string `xml:",chardata"`
							Attr  string `xml:"attr,attr"`
							ID    string `xml:"id,attr"`
							Value string `xml:"value,attr"`
						} `xml:"format"`
					} `xml:"style-rule"`
				} `xml:"style"`
				Panes struct {
					Text string `xml:",chardata"`
					Pane struct {
						Text string `xml:",chardata"`
						View struct {
							Text      string `xml:",chardata"`
							Breakdown struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value,attr"`
							} `xml:"breakdown"`
						} `xml:"view"`
						Mark struct {
							Text  string `xml:",chardata"`
							Class string `xml:"class,attr"`
						} `xml:"mark"`
						Encodings struct {
							Text string `xml:",chardata"`
							Size struct {
								Text   string `xml:",chardata"`
								Column string `xml:"column,attr"`
							} `xml:"size"`
							Lod []struct {
								Text   string `xml:",chardata"`
								Column string `xml:"column,attr"`
							} `xml:"lod"`
						} `xml:"encodings"`
					} `xml:"pane"`
				} `xml:"panes"`
				Rows string `xml:"rows"`
				Cols string `xml:"cols"`
			} `xml:"table"`
		} `xml:"worksheet"`
	} `xml:"worksheets"`
	Windows struct {
		Text   string `xml:",chardata"`
		Window []struct {
			Text         string `xml:",chardata"`
			Class        string `xml:"class,attr"`
			SourceHeight string `xml:"source-height,attr"`
			AutoHidden   string `xml:"auto-hidden,attr"`
			Maximized    string `xml:"maximized,attr"`
			Name         string `xml:"name,attr"`
			Cards        struct {
				Text string `xml:",chardata"`
				Edge []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Strip []struct {
						Text string `xml:",chardata"`
						Size string `xml:"size,attr"`
						Card []struct {
							Text                string `xml:",chardata"`
							Type                string `xml:"type,attr"`
							Param               string `xml:"param,attr"`
							Mode                string `xml:"mode,attr"`
							PaneSpecificationID string `xml:"pane-specification-id,attr"`
						} `xml:"card"`
					} `xml:"strip"`
				} `xml:"edge"`
			} `xml:"cards"`
			Highlight struct {
				Text        string `xml:",chardata"`
				ColorOneWay struct {
					Text  string   `xml:",chardata"`
					Field []string `xml:"field"`
				} `xml:"color-one-way"`
			} `xml:"highlight"`
			Viewpoint string `xml:"viewpoint"`
		} `xml:"window"`
	} `xml:"windows"`
}
