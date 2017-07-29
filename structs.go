package godork

type PackageDoc struct {
	PackageName string
	ImportPath  string
	Doc         string
	Examples    []*Example
	Constants   []*Constant
	Variables   []*Variable
	Functions   []*Function
	Types       []*Type
}

type Example struct {
	Name     string
	FullName string
	Doc      string
	Code     string
	Output   string
}

type Constant struct {
	Doc  string
	Code string
}

type Variable struct {
	Doc  string
	Code string
}

type Function struct {
	Name      string
	Doc       string
	Signature string
	Examples  []*Example
}

type Type struct {
	Name     string
	Doc      string
	Code     string
	Examples []*Example
	Methods  []*Method
}

type Method struct {
	Name      string
	Doc       string
	Signature string
	Receiver  string
	Examples  []*Example
}
