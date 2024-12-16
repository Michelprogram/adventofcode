package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func getEnvVariables() (map[string]string, error) {

	vars := make(map[string]string)

	data, err := os.ReadFile(".env")

	if err != nil {
		return nil, err
	}

	for _, line := range bytes.Split(data, []byte("\n")) {
		equalIdx := bytes.IndexByte(line, '=')

		if equalIdx == -1 {
			continue
		}

		vars[string(line[:equalIdx])] = string(line[equalIdx+1:])
	}

	return vars, nil

}

func FecthDataSet(year, day int) ([]byte, error) {

	vars, err := getEnvVariables()

	if err != nil {
		return nil, err
	}

	u, err := url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))

	log.Println(u.String())

	if err != nil {
		return nil, err
	}

	jar, _ := cookiejar.New(nil)

	jar.SetCookies(u, []*http.Cookie{
		{
			Name:   "session",
			Value:  vars["SESSION_ID"],
			Path:   "/",
			Domain: ".adventofcode.com",
		},
	})

	client := http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func IsOutOfBound(x, y, maxX, maxY int) bool {
	return x < 0 || x >= maxX || y >= maxY || y < 0
}
