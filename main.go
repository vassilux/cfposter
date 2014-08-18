package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	apiUrl := "http://www.123contactform.com"
	resource := "/form-1075094/ContactMe"
	data := url.Values{}
	data.Set("action", "verify")
	data.Add("special_autoresponder", "")
	data.Add("language", "en")
	data.Add("languageChanged", "no")
	data.Add("control8792015", "totoname")
	data.Add("control8792016", "toto@tttt.com")
	data.Add("control8792017", "text")
	data.Add("go_back_and_edit", "0")
	data.Add("PHPSESSID", "seoecorp0m2n4021mi9j8rrrt2")
	data.Add("hiddenfields", "")
	data.Add("hiddenfields_pages", "")
	data.Add("activepage", "1")
	data.Add("totalpages", "1")
	data.Add("nextpagenr", "1")
	data.Add("prevpagenr", "0")
	data.Add("", "")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u) //

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Autorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}
