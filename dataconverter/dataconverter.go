package dataconverter

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/piquette/finance-go"
	"github.com/shopspring/decimal"
)

type Params map[string]interface{}

func (p Params) ToCSV() ([]string, []string, error) {
	headers := make([]string, len(p))
	rows := make([]string, len(p))
	i := 0
	for k, v := range p {
		sv, err := obtainString(v)
		if err != nil {
			log.Println(err)
			continue
		}
		headers[i] = k
		rows[i] = sv
		i++
	}
	if i == 0 {
		return nil, nil, ErrEmptyParam
	}
	return headers, rows, nil
}

func obtainString(v interface{}) (string, error) {
	switch result := v.(type) {
	case string:
		return result, nil
	case int:
		return strconv.Itoa(result), nil
	case float64:
		return strconv.FormatFloat(result, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(result), nil
	case int64:
		return strconv.FormatInt(result, 64), nil
	case fmt.Stringer:
		return result.String(), nil
	case fmt.GoStringer:
		return result.GoString(), nil
	}
	return "", ErrIncompatibleType
}

func (p Params) ToJsonString() (string, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(bs), err
}

func FromGFStructToParam(lowerCaseFields map[string]struct{}, structValue interface{}) (Params, error) {
	if structValue == nil {
		return nil, ErrNilArgument
	}
	v := reflect.ValueOf(structValue)
	typeOfStruct := v.Type()
	if typeOfStruct.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}
	params := make(Params)
	for i := 0; i < v.NumField(); i++ {
		field := typeOfStruct.Field(i).Name
		value := v.Field(i).Interface()
		if _, ok := lowerCaseFields[strings.ToLower(field)]; ok {
			switch nv := value.(type) {
			case decimal.Decimal:
				params[field] = nv.String()
			case finance.MarketState:
				params[field] = string(nv)
			case finance.QuoteType:
				params[field] = string(nv)
			case int:
				if strings.EqualFold(field, "timestamp") {
					params[field] = time.Unix(int64(nv), 0)
				}
			default:
				params[field] = value
			}
		}
	}
	return params, nil
}
