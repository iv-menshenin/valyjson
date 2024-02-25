package names

import ast "github.com/iv-menshenin/go-ast"

// names
const (
	VarNameJsonValue         = "v"
	VarNameError             = "err"
	VarNameReceiver          = "s"
	VarNameListOfArrayValues = "listA"
	VarNameListElem          = "listElem"
	VarPrefixPool            = "jsonParser"
	VarNameData              = "data"

	VarNameWriter = "result"

	MethodNameFill      = "FillFromJSON"
	MethodNameUnmarshal = "UnmarshalJSON"
	MethodNameMarshal   = "MarshalJSON"
	MethodNameValidate  = "validate"
	MethodNameMarshalTo = "MarshalTo"
	MethodNameZero      = "IsZero"
	UnpackObjFunc       = "unpackObject"
	MethodNameReset     = "Reset"

	WriteTime = "writeTime"

	ParsingError = "newParsingError"
)

var (
	FastJsonValue      = ast.SimpleSelector("fastjson", "Value")
	FastJsonParserPool = ast.SimpleSelector("fastjson", "ParserPool")

	WriteTimeFunc = ast.CallFunctionDescriber{
		FunctionName:             ast.NewIdent(WriteTime),
		MinimumNumberOfArguments: 3,
	}
	// WriteInt64Func equal result.Int64(?)
	WriteInt64Func = ast.CallFunctionDescriber{
		FunctionName:             ast.SimpleSelector(VarNameWriter, "Int64"),
		MinimumNumberOfArguments: 1,
	}
	// WriteUint64Func equal result.Uint64(?)
	WriteUint64Func = ast.CallFunctionDescriber{
		FunctionName:             ast.SimpleSelector(VarNameWriter, "Uint64"),
		MinimumNumberOfArguments: 1,
	}
	// WriteFloat64Func equal result.Float64(?)
	WriteFloat64Func = ast.CallFunctionDescriber{
		FunctionName:             ast.SimpleSelector(VarNameWriter, "Float64"),
		MinimumNumberOfArguments: 1,
	}
	// WriteStringFunc equal result.String(?)
	WriteStringFunc = ast.CallFunctionDescriber{
		FunctionName:             ast.SimpleSelector(VarNameWriter, "String"),
		MinimumNumberOfArguments: 1,
	}

	TimeDefaultLayout = ast.SimpleSelector("time", "RFC3339Nano")
)
