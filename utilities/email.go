package utilities

import (
	"strconv"
	"strings"
)

// ProfileFor creates a profile for the user
func ProfileFor(email string) string {

	if strings.Contains(email, "&") || strings.Contains(email, "=") {
		panic("Invalid email!")
	}

	uid := 10 // anything for learning purposes
	role := "user"

	strVals := "email=" + email + "&uid=" + strconv.Itoa(uid) + "&role=" + role

	vals := Unparse(Parse(strVals))

	return vals
}
