package main

import (
	"aqwari.net/xml/xsd"
	"log"
	"os"

	"aqwari.net/xml/xsdgen"
)

func main() {
	var cfg xsdgen.Config

	cfg.Option(
		xsdgen.IgnoreAttributes("href", "offset", "monthDay"),
		xsdgen.Replace(`[._ \s-]`, ""),
		xsdgen.PackageName("ws"),
		xsdgen.HandleSOAPArrayType(),
		xsdgen.SOAPArrayAsSlice(),
		xsdgen.UseFieldNames(),

		xsdgen.ProcessTypes(func(s xsd.Schema, t xsd.Type) xsd.Type {
			switch t.(type) {
			case *xsd.SimpleType:
				st := t.(*xsd.SimpleType)
				if st.Name.Local == "" {
					st.Name.Local = "LastDay"
				}
			}
			return t
		}),
	)
	if err := cfg.GenCLI(os.Args[1:]...); err != nil {
		log.Fatal(err)
	}
}
