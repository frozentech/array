package array

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// DeleteEmpty ...
func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// Sort sort map string by string
func Sort(source map[string]string) (output map[string]string) {
	keys := make([]string, 0, len(source))
	output = make(map[string]string)
	for k := range source {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		output[k] = source[k]
	}
	return output
}

// InArray checks if the given value exists in an array
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

// KeyExist checks if key exist in map string
func KeyExist(find string, source map[string]string) (exists bool) {
	exists = false
	for key, value := range source {
		if find == key && value != "" {
			exists = true
			return
		}
	}
	return
}

// Flip exchanges all keys with corresponding value
func Flip(source map[string]string) (output map[string]string) {
	output = make(map[string]string)
	for key, value := range source {
		output[value] = key
	}
	return
}

// ToString Array String to string
func ToString(source []string) (output string) {
	output = fmt.Sprintf("(%s)", strings.Join(source, ","))
	return
}

// Unique remove duplicates
func Unique(source []string) (output []string) {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	for v := range source {
		if encountered[source[v]] != true {
			encountered[source[v]] = true
			output = append(output, source[v])
		}
	}
	return
}
