package dbcassandra

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"sort"
	"strconv"
	"strings"
)

func BuildWhereQuery(inValues map[string]interface{}) string {
	if len(inValues) == 0 {
		return " "
	}
	query := ""
	x := 0
	for key, val := range inValues {
		and := ""
		if x > 0 {
			and = "and"
		}
		if IsStringValue(val) {
			query = fmt.Sprintf("%v %v %v='%v' ", query, and, key, val)
		} else {
			query = fmt.Sprintf("%v %v %v=%v ", query, and, key, val)
		}
		x++
	}
	return fmt.Sprintf("where %v", query)
}
func IsStringValue(valueIn interface{}) bool {
	var myType = fmt.Sprintf("%T", valueIn)
	if myType == "string" {
		return true
	}
	return false
}
func IsNumberValue(valueIn interface{}) bool {
	var myType = fmt.Sprintf("%T", valueIn)
	if myType == "int" || myType == "int64" || myType == "int32" || myType == "float64" ||
		myType == "float34" {
		return true
	}
	return false
}
func IsInArrayString(data []string, checkValue string) bool {
	for _, item := range data {
		if item == checkValue {
			return true
		}
	}
	return false
}
func FetchEntityDataFromDb(session *gocql.Session, dbName, table, appName, org string) string {
	qry := fmt.Sprintf("select * from %v.%v where appName='%v' ",
		dbName, table, appName)
	if org != "" {
		qry = fmt.Sprintf("select * from %v.%v where appName='%v' and org='%v' ",
			dbName, table, appName, org)
	}

	res, _ := RunQueryCass2(session, qry)
	return res
}
func CompareEqualBool(val1, val2 interface{}) bool {
	v1 := ToBool(val1)
	v2 := ToBool(val2)
	if v1 != v2 {
		return false
	}
	return true
}
func CompareEqualNumbers(val1, val2 interface{}) bool {
	v1 := ToFloat64(val1)
	v2 := ToFloat64(val2)
	if v1 != v2 {
		return false
	}
	return true
}
func CompareEqualFloat64(val1, val2 float64) bool {
	if val1 != val2 {
		return false
	}
	return true
}
func CompareEqualFloat32(val1, val2 float32) bool {
	if val1 != val2 {
		return false
	}
	return true
}
func CompareEqualInt32(val1, val2 int32) bool {
	if val1 != val2 {
		return false
	}
	return true
}
func CompareEqualInt(val1, val2 int) bool {
	if val1 != val2 {
		return false
	}
	return true
}
func CompareEqualInt64(val1, val2 int64) bool {
	if val1 != val2 {
		return false
	}
	return true
}
func toInt(valueIn interface{}) int {
	val := ToString(valueIn)
	out, _ := strconv.Atoi(val)
	return out
}
func ToInt32(valueIn interface{}) int32 {
	val := toInt(valueIn)
	return int32(val)
}
func ToInt64(valueIn interface{}) int64 {
	val := toInt(valueIn)
	return int64(val)
}
func ToFloat32(valueIn interface{}) float32 {
	val := ToString(valueIn)
	out, _ := strconv.ParseFloat(val, 32)
	return float32(out)
}
func ToFloat64(valueIn interface{}) float64 {
	val := ToString(valueIn)
	out, _ := strconv.ParseFloat(val, 64)
	return out
}
func ToBool(valueIn interface{}) bool {
	val := ToString(valueIn)
	out, _ := strconv.ParseBool(val)
	return out
}
func ToString(valueIn interface{}) string {
	return fmt.Sprintf("%v", valueIn)
}
func GenerateAutoNumber(session *gocql.Session, dbName, appName, entity, field, prefix string, startWith int) string {
	var data []interface{}
	info := FetchEntityFromDbAll(session, entity, dbName, appName)
	_ = json.Unmarshal(info, &data)
	var arr []int
	for _, row := range data {
		var record = make(map[string]interface{})
		str, _ := json.Marshal(row)
		_ = json.Unmarshal(str, &record)
		field = strings.ToLower(field)
		f, _ := record[field]
		c := fmt.Sprintf("%v", f)
		c = strings.Replace(c, prefix, "", 10)
		c1, _ := strconv.Atoi(c)
		arr = append(arr, c1)
	}
	sort.Ints(arr)
	if len(arr) == 0 {
		arr = append(arr, startWith)
	}
	var index = len(arr) - 1
	var currentCode = arr[index]
	var nextCode = currentCode + 1
	return fmt.Sprintf("%v%v", prefix, nextCode)
}
