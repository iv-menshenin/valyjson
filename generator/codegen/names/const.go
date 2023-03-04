package names

import ast "github.com/iv-menshenin/go-ast"

// names
const (
	VarNameJsonValue         = "v"
	VarNameError             = "err"
	VarNameObjPath           = "objPath"
	VarNameReceiver          = "s"
	VarNameListOfArrayValues = "listA"
	VarNameListElem          = "listElem"
	VarNameBuf               = "buf"
	VarPrefixPool            = "jsonParser"
	VarNameData              = "data"

	MethodNameFill      = "FillFromJSON"
	MethodNameUnmarshal = "UnmarshalJSON"
	MethodNameMarshal   = "MarshalJSON"
	MethodNameValidate  = "validate"
	MethodNameAppend    = "MarshalAppend"
)

var (
	FastJsonValue      = ast.SimpleSelector("fastjson", "Value")
	FastJsonParserPool = ast.SimpleSelector("fastjson", "ParserPool")
)
