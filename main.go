package main

import (
	"bytes"
	"fmt"
	//"io/ioutil"
	"net/http"
	"net/url"
	//"os"
	"strconv"
	"time"
)

func do_send() int {
	apiUrl := "http://www.123contactform.com"
	//http://www.123contactform.com/js-form--828647.html
	//resource := "/form-1075094/ContactMe" //test of mine
	resource := "/js-form--828647.html"
	data := url.Values{}
	data.Set("action", "verify")
	data.Add("special_autoresponder", "")
	data.Add("language", "en")
	data.Add("languageChanged", "no")
	data.Add("control6521519", "odin.troll")
	data.Add("control6521520", "odin.troll@mail.ru")
	data.Add("control6521521", "What are all patches ? Stop selling !!! ")
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
	defer resp.Body.Close()
	if resp == nil {
		return -1
	}

	/*contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}*/
	//fmt.Printf("%s\n", string(contents))
	return resp.StatusCode

}

func main() {
	max := 50
	for i := 0; i < max; i++ {
		resp := do_send()
		fmt.Printf("Send %d with result %d.", i, resp)
		fmt.Println("")
		fmt.Println("Slepping.")
		time.Sleep(3 * 1000)

	}
}
