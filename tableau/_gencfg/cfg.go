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
		xsdgen.IgnoreAttributes("href", "offset"),
		xsdgen.Replace(`[._ \s-]`, ""),
		xsdgen.PackageName("ws"),
		xsdgen.HandleSOAPArrayType(),
		xsdgen.SOAPArrayAsSlice(),
		xsdgen.UseFieldNames(),

		xsdgen.ProcessTypes(func(s xsd.Schema, t xsd.Type) xsd.Type {
			switch t.(type) {
			case *xsd.SimpleType:
				// workaround for the invalid generation of LastDay type,
				// currently 'type string', which should be 'type LastDay string'" in this complexType:
				//
				//    <xs:complexType name="intervalType">
				//        <xs:attribute name="minutes">
				//            <xs:simpleType>
				//                <xs:restriction base="xs:nonNegativeInteger">
				//                    <xs:enumeration value="15" />
				//                    <xs:enumeration value="30" />
				//                </xs:restriction>
				//            </xs:simpleType>
				//        </xs:attribute>
				//        <xs:attribute name="hours">
				//            <xs:simpleType>
				//                <xs:restriction base="xs:nonNegativeInteger">
				//                    <xs:enumeration value="1" />
				//                    <xs:enumeration value="2" />
				//                    <xs:enumeration value="4" />
				//                    <xs:enumeration value="6" />
				//                    <xs:enumeration value="8" />
				//                    <xs:enumeration value="12" />
				//                </xs:restriction>
				//            </xs:simpleType>
				//        </xs:attribute>
				//        <xs:attribute name="weekDay" >
				//            <xs:simpleType>
				//                <xs:restriction base="xs:string">
				//                    <xs:enumeration value="Monday" />
				//                    <xs:enumeration value="Tuesday" />
				//                    <xs:enumeration value="Wednesday" />
				//                    <xs:enumeration value="Thursday" />
				//                    <xs:enumeration value="Friday" />
				//                    <xs:enumeration value="Saturday" />
				//                    <xs:enumeration value="Sunday" />
				//                </xs:restriction>
				//            </xs:simpleType>
				//        </xs:attribute>
				//        <xs:attribute name="monthDay" >
				//            <xs:simpleType>
				//                <xs:union>
				//                    <xs:simpleType>
				//                        <xs:restriction base="xs:integer">
				//                            <xs:minInclusive value="1" />
				//                            <xs:maxInclusive value="31" />
				//                        </xs:restriction>
				//                    </xs:simpleType>
				//                    <xs:simpleType>
				//                        <xs:restriction base="xs:string">
				//                            <xs:enumeration value="LastDay" />
				//                        </xs:restriction>
				//                    </xs:simpleType>
				//                </xs:union>
				//            </xs:simpleType>
				//        </xs:attribute>
				//    </xs:complexType>
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
