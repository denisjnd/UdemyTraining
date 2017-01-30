package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.hitzmuzik.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

/*
Use _ to define an identifier that is not used...
Here we used the http.Get() to send a request and it returns a responds and an error but since we don't need
 the error to be stored in a variable, we use an _ so that th compiler will throw the error away.
*/
