/*
JSON <-> Go
true/false <-> true/false (bool)
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int64, int, uint8, ...
array <-> []int, []string, ..., []any
object <-> struct, map[string]any

MIA: time.Time, []byte (binary), ... (also comments)

encoding/json API
File-ish
Go -> JSON via io.Writer: json.Encoder
JSON -> Go via io.Reader: json.Decoder

In-memory
Go -> JSON via []byte: Marshal
JSON -> Go via []byte: Unmarshal
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	name, repos, err := userInfo(ctx, "tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(name, repos)
}

// userInfo return user name and number of public repos
func userInfo(ctx context.Context, login string) (string, int, error) {
	u := "https://api.github.com/users/" + url.PathEscape(login)

	// resp, err := http.Get(u)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf(resp.Status)
	}

	// We can't use map[string]string for header since:
	// - HTTP headers as case insensitive
	// - HTTP header can repeat (resp.Header.Values)
	// fmt.Println("content-type:", resp.Header.Get("Content-Type"))

	// io.Copy(os.Stdout, resp.Body)
	// var r userReply
	var r struct { // anonymous struct, avoid name pollution
		// Field names must be exported to work with encoding/json
		Name string `json:"name"`
		// Public_Repos int
		NumRepos int `json:"public_repos"` // field tag
	}
	// var i int // i == 0
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, fmt.Errorf("can't decode - %s", err)
	}
	return r.Name, r.NumRepos, nil
	// fmt.Println(r)
	// fmt.Printf("%#v\n", r)
	// See %v, %+v, %#v

	// json.NewEncoder(os.Stdout).Encode(r)
}

/*
type userReply struct {
	// Field names must be exported to work with encoding/json
	Name string `json:"name"`
	// Public_Repos int
	NumRepos int `json:"public_repos"` // field tag
}
*/
