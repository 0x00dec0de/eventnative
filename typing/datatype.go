package typing

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type DataType int

const (
	//IMPORTANT: order of iota values. Int values according to Typecast tree (see typing.typecastTree)
	UNKNOWN DataType = iota
	INT64
	FLOAT64
	STRING
	TIMESTAMP
)

var (
	inputStringToType = map[string]DataType{
		"string":    STRING,
		"integer":   INT64,
		"double":    FLOAT64,
		"timestamp": TIMESTAMP,
	}
	typeToInputString = map[DataType]string{
		STRING:    "string",
		INT64:     "integer",
		FLOAT64:   "double",
		TIMESTAMP: "timestamp",
	}
)

func (dt DataType) String() string {
	switch dt {
	default:
		return ""
	case STRING:
		return "STRING"
	case INT64:
		return "INT64"
	case FLOAT64:
		return "FLOAT64"
	case TIMESTAMP:
		return "TIMESTAMP"
	case UNKNOWN:
		return "UNKNOWN"
	}
}

func TypeFromString(t string) (DataType, error) {
	trimmed := strings.TrimSpace(t)
	lowerTrimmed := strings.ToLower(trimmed)
	dataType, ok := inputStringToType[lowerTrimmed]
	if !ok {
		return UNKNOWN, fmt.Errorf("Unknown casting type: %s", t)
	}
	return dataType, nil
}

func StringFromType(dataType DataType) (string, error) {
	str, ok := typeToInputString[dataType]
	if !ok {
		return "", fmt.Errorf("Unable to get string from DataType for: %s", dataType.String())
	}
	return str, nil
}

//TypeFromValue return DataType from v type
//note: json.Unmarshal doesn't return int type at all, but returns float64 instead
//      we have to check math.Trunc(floatV) == floatV and than it will be INT64
func TypeFromValue(v interface{}) (DataType, error) {
	switch v.(type) {
	case string:
		return STRING, nil
	case float32:
		floatV := float64(v.(float32))
		if math.Trunc(floatV) == floatV {
			return INT64, nil
		}
		return FLOAT64, nil
	case float64:
		floatV := v.(float64)
		if math.Trunc(floatV) == floatV {
			return INT64, nil
		}
		return FLOAT64, nil
	case int, int8, int16, int32, int64:
		return INT64, nil
	case time.Time:
		return TIMESTAMP, nil
	default:
		return UNKNOWN, fmt.Errorf("Unknown DataType for value: %v type: %t", v, v)
	}
}
