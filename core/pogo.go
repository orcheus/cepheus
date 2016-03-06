package core

// RegisterContextRequest : a requet to register a new context
type RegisterContextRequest struct {
	ContextRegistrations []ContextRegistration
	Duration             string
}

type ContextRegistration struct {
	Entities             []Entity
	Attributes           []Attribute
	ProvidingApplication string
}

// Attribute : a context attribute
type Attribute struct {
	Name  string
	Type  string
	Value string
}

type Entity struct {
	Type string
	Id   string
}
