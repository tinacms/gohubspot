package gohubspot

import (
	"encoding/json"
)

type ItemDataType string

type ItemFieldType string

const (
	TextAreaField        ItemFieldType = "textarea"
	TextField            ItemFieldType = "text"
	DateField            ItemFieldType = "date"
	FileField            ItemFieldType = "file"
	NumberField          ItemFieldType = "number"
	SelectField          ItemFieldType = "select"
	RadioField           ItemFieldType = "radio"
	CheckboxField        ItemFieldType = "checkbox"
	BooleanCheckboxField ItemFieldType = "booleancheckbox"
)

const (
	String      ItemDataType = "string"
	Number      ItemDataType = "number"
	Date        ItemDataType = "date"
	DateTime    ItemDataType = "datetime"
	Enumeration ItemDataType = "enumeration"
)

type Property struct {
	Property string      `json:"property"`
	Value    interface{} `json:"value"`
	Versions Versions    `json:"versions,omitempty"`
}

type Properties struct {
	Properties []Property `json:"properties"`
}

func (p *Properties) UnmarshalJSON(data []byte) error {
	v := map[string]Property{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	*p = Properties{}

	for k, v := range v {
		p.Properties = append(p.Properties, Property{Property: k, Value: v.Value, Versions: v.Versions})
	}

	return nil
}

// AddProperty addes a new property to list
func (p *Properties) AddProperty(prop string, value interface{}) {
	p.Properties = append(p.Properties, Property{Property: prop, Value: value})
}

func (p *Properties) UpdateProperty(prop string, value interface{}) {
	p.DeleteProperty(prop)
	p.AddProperty(prop, value)
}

func (p *Properties) DeleteProperty(prop string) {
	props := &Properties{}
	for i := range p.Properties {
		if p.Properties[i].Property != prop {
			props.AddProperty(p.Properties[i].Property, p.Properties[i].Value)
		}
	}

	p.Properties = props.Properties
}

func (p *Properties) GetProperty(prop string) (Property, bool) {
	for i := range p.Properties {
		if p.Properties[i].Property == prop {
			return p.Properties[i], true
		}
	}

	return Property{}, false
}

// ItemPropertyOption definition
type ItemPropertyOption struct {
	Description  string      `json:"description"`
	Label        string      `json:"label"`
	DisplayOrder int         `json:"displayOrder"`
	Hidden       bool        `json:"hidden"`
	Value        interface{} `json:"value"`
}

// ItemProperty definition
type ItemProperty struct {
	Name                          string               `json:"name"`
	Label                         string               `json:"label,omitempty"`
	Description                   string               `json:"description,omitempty"`
	DataType                      ItemDataType         `json:"type,omitempty"`
	FieldType                     ItemFieldType        `json:"fieldType,omitempty"`
	GroupName                     string               `json:"groupName,omitempty"`
	Options                       []ItemPropertyOption `json:"options"`
	Deleted                       bool                 `json:"deleted"`
	FormField                     bool                 `json:"formField"`
	DisplayOrder                  int                  `json:"displayOrder"`
	ReadOnlyValue                 bool                 `json:"readOnlyValue,omitempty"`
	ReadOnlyDefinition            bool                 `json:"readOnlyDefinition,omitempty"`
	Hidden                        bool                 `json:"hidden,omitempty"`
	MutableDefinitionNotDeletable bool                 `json:"mutableDefinitionNotDeletable,omitempty"`
	SearchableInGlobalSearch      bool                 `json:"searchableInGlobalSearch,omitempty"`
	Calculated                    bool                 `json:"calculated"`
	ExternalOptions               bool                 `json:"externalOptions"`
}

// ItemProperties list of ItemProperty type
type ItemProperties []ItemProperty

// ItemPropertyGroup item property group definition
type ItemPropertyGroup struct {
	Name           string `json:"name"`
	DisplayName    string `json:"displayName"`
	DisplayOrder   int    `json:"displayOrder"`
	HubspotDefined bool   `json:"hubspotDefined"`
}

// ItemPropertyGroups ItemPropertyGroup list type
type ItemPropertyGroups []ItemPropertyGroup
