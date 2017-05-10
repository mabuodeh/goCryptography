package utilities

import "strings"

// Parse takes an email, uid, and role, and stores them in a structure
func Parse(str string) map[string]interface{} {
	// foo=bar&baz=qux&zap=zazzle
	m := make(map[string]interface{})

	vals := strings.Split(str, "&")
	for _, val := range vals {
		temp := strings.Split(val, "=")
		m[temp[0]] = temp[1]
	}
	// fmt.Println(m)
	return m
}

// Unparse takes a structure of email, uid, and role, and creates a stream of tokens
func Unparse(vals map[string]interface{}) string {
	// str := ""
	// map[foo:bar baz:qux zap:zazzle]

	str := "email=" + vals["email"].(string) + "&" + "uid=" + vals["uid"].(string) + "&" + "role=" + vals["role"].(string)

	// // iterate over map
	// for key, val := range vals {
	// 	// append key to string
	// 	// append = to string
	// 	// append value to string
	// 	str += "&"
	// 	str += key + "=" + val.(string)
	// 	// append & to string
	// }
	// // remove last &
	// str = str[1:]
	// print
	// fmt.Println(str)
	return str
}
