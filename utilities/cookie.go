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
	// done manually to maintain order
	return "email=" + vals["email"].(string) + "&" + "uid=" + vals["uid"].(string) + "&" + "role=" + vals["role"].(string)
}
