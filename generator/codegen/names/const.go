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

	MarshalTime   = "marshalTime"
	MarshalString = "marshalString"
)

var (
	FastJsonValue      = ast.SimpleSelector("fastjson", "Value")
	FastJsonParserPool = ast.SimpleSelector("fastjson", "ParserPool")

	MarshalTimeFunc = ast.CallFunctionDescriber{
		FunctionName:             ast.NewIdent(MarshalTime),
		MinimumNumberOfArguments: 3,
	}
	MarshalStringFunc = ast.CallFunctionDescriber{
		FunctionName:             ast.NewIdent(MarshalString),
		MinimumNumberOfArguments: 2,
	}

	TimeDefaultLayout = ast.SimpleSelector("time", "RFC3339Nano")
)
