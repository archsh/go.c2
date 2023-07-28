package c2

import "encoding/xml"

// the c2 adi definitions
/*
<?xml version="1.0" encoding="UTF-8"?>
<ADI xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<Objects>
		<Object ElementType="object_type" ID="object_id" Action="REGIST">
			<Property Name="property_name1">property_value1</Property>
			<Property Name="property_name2">property_value2</Property>
			<Property Name="property_name3">property_value3</Property>
		</Object>
    </Objects>
	<Mappings>
		<Mapping ID=”mapping_id” ParentType="parent_type" ParentID="parent_id" ElementType=”element_type” ElementID=”element_id” Action="REGIST">
			<Property name="property_name1">property_value1</Property>
			<Property name="property_name2">property_value2</Property>
		</Mapping>
	</Mappings>
</ADI>
*/

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Object struct {
	ID          string     `xml:",attr"`
	ElementType string     `xml:",attr"`
	Action      string     `xml:",attr"`
	Properties  []Property `xml:"Property"`
}

type Mapping struct {
	ID          string     `xml:",attr"`
	ParentType  string     `xml:",attr"`
	ParentID    string     `xml:",attr"`
	ElementType string     `xml:",attr"`
	ElementID   string     `xml:",attr"`
	Action      string     `xml:",attr"`
	Properties  []Property `xml:"Property"`
}

type ADI struct {
	XMLName  xml.Name  `xml:"http://www.w3.org/2001/XMLSchema-instance ADI"`
	Objects  []Object  `xml:"Objects>Object"`
	Mappings []Mapping `xml:"Mappings>Mapping"`
}
