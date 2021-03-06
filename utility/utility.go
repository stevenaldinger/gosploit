package utility

import (
	"bufio"
	"os"
	"net/http"
	"bytes"
)

var lines []string
var body string

func ReadLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

//Return the string of response body from HTTP request
func HTTPResponseBodyString(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		body = ""
	}
	if resp != nil {
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
	    buf.ReadFrom(resp.Body)
	    body = buf.String()
	}
	return body, err
}
