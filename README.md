




* [Basic Auth](#basic-auth)




## Basic Auth
``` go
func BasicAuth(realm string, validate func(username, password string, r *http.Request) (err error)) func(http.Handler) http.Handler
```

```go
	f := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintln(w, "DONE")
	})
	
	validate := func(username, password string, r *http.Request) (err error) {
	    if username == "u" && password == "p" {
	        return
	    }
	    err = errors.New("Wrong")
	    return
	}
	
	ts := httptest.NewServer(basicauth.BasicAuth("myrealm", validate)(f))
	defer ts.Close()
	
	resp, _ := http.Get(ts.URL)
	if resp.StatusCode != 401 {
	    fmt.Printf("should be 401, but was %d", resp.StatusCode)
	}
	
	req, _ := http.NewRequest("GET", ts.URL, strings.NewReader(""))
	req.SetBasicAuth("u", "p")
	resp, _ = http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
	    fmt.Printf("should be 200, but was %d", resp.StatusCode)
	}
	
	// Output:
	//
```




