package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "abc"
	s2 := strings.Clone(s1)
	fmt.Println(s1)
	fmt.Println(s2)
	b := strings.Builder{}
	// b.Write([]byte("Here is an example"))
	b.WriteString("Here is an example")

	fmt.Println(b.String())

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s2))
	s2 = " Hello Welcome Blahb lah  "
	fmt.Println("s2", len(s2))
	s3 := strings.TrimSpace(s2)
	fmt.Println("s2", len(s3))
	fmt.Println(strings.HasSuffix("test@gmail.com", "@gmail.com"))
	fmt.Println(strings.HasPrefix("test@gmail.com", "test"))
	fmt.Println(strings.Replace("test@gmail.com", "test", "john", 1))

	parts := strings.Split("test@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)

	parts = strings.Fields("jame gmail.com") //fields is using to split using space
	username, domain = parts[0], parts[1]
	fmt.Println(username, domain)

}
