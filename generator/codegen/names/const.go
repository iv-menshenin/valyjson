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

	WriteString  = "writeString"
	WriteTime    = "writeTime"
	WriteInt64   = "writeInt64"
	WriteUint64  = "writeUint64"
	WriteFloat64 = "writeFloat64"

	ParsingError = "newParsingError"
)

var (
	FastJsonValue      = ast.SimpleSelector("fastjson", "Value")
	FastJsonParserPool = ast.SimpleSelector("fastjson", "ParserPool")

	WriteTimeFunc = ast.CallFunctionDescriber{
		FunctionName:             ast.NewIdent(WriteTime),
		MinimumNumberOfArguments: 3,
	}
	WriteInt64Func = ast.CallFunctionDescriber{
		FunctionName:             ast.NewIdent(WriteInt64),
		MinimumNumberOfArguments: 2,
	}
	WriteUint64Func = ast.CallFunctionDescriber{
		FunctionName:             ast.NewIdent(WriteUint64),
		MinimumNumberOfArguments: 2,
	}
	WriteFloat64Func = ast.CallFunctionDescriber{
		FunctionName:             ast.NewIdent(WriteFloat64),
		MinimumNumberOfArguments: 2,
	}
	WriteStringFunc = ast.CallFunctionDescriber{
		FunctionName:             ast.NewIdent(WriteString),
		MinimumNumberOfArguments: 2,
	}

	TimeDefaultLayout = ast.SimpleSelector("time", "RFC3339Nano")
)
