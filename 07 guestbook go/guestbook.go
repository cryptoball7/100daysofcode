package main

import (
	"fmt"
	"net/http"
	"os"
)

func main () {
	http.HandleFunc("/", guestbookHandler)
	http.ListenAndServe(":8080", nil)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func guestbookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<form method='post' action='/'>
Sign my guestbook!<br>
<input type=text name=text><input type=submit>
</form>
</html>`)

	data, err := os.ReadFile("data.txt")
	check(err)

	dataString := string(data)
	err = os.WriteFile("data.txt", []byte(r.PostFormValue("text")+"<br>\n"+dataString), 0600)
	check(err)

	fmt.Fprintf(w, dataString)
}