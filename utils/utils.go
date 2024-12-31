package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"time"
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

func fecthDataSet(year, day int) ([]byte, error) {

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

func RunAoc(test bool, part, day, year int, aocs map[int]Code) error {

	var data []byte
	var err error

	if !test {

		data, err = fecthDataSet(year, day)

		if err != nil {
			return fmt.Errorf("Can't fetch data for day %d\n", day)
		}
	}

	aoc, ok := aocs[year]

	if !ok {
		log.Fatalf("year %d doesn't exist\n", year)
	}

	start := time.Now()

	res, err := aoc.Execute(data, part, day)

	if err != nil {
		log.Fatalf("Error during advent code %d for day %d : %s\n", year, day, err)
	}

	fmt.Printf("Time : %v\n", time.Since(start))

	fmt.Printf("Resultat for day %d : %v\n", day, res)

	return nil
}

func NumberOfDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func GenerateFiles(day, year int) error {

	const runnerTemplate = `
package day{{.Day}}

import (
	"github.com/michelprogram/adventofcode/registry"
   	"github.com/michelprogram/adventofcode/utils"
)

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) Part1(data []byte) (any, error) {
    // TODO: Implement Part 1 logic here
    return nil, nil
}

func (d Runner) Part2(data []byte) (any, error) {
    // TODO: Implement Part 2 logic here
    return nil, nil
}

func init() {
    registry.RegisterChallenge({{.Day}}, Runner{})
}
    `
	const testRunnerTemplate = `
package day{{.Day}}_test

import (
    "testing"

    "github.com/michelprogram/adventofcode/aoc_{{.Year}}/day{{.Day}}"
)

func Test1(t *testing.T) {
    //TODO implement test
}
    `

	var output bytes.Buffer

	dayFolder := fmt.Sprintf("aoc_%d/day%02d", year, day)
	if err := os.MkdirAll(dayFolder, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create folder %s: %v", dayFolder, err)
	}

	runnerFilePath := filepath.Join(dayFolder, "runner.go")
	if _, err := os.Stat(runnerFilePath); err == nil {
		return fmt.Errorf("File %s already exists", runnerFilePath)
	}

	runnerTestFilePath := filepath.Join(dayFolder, "runner_test.go")
	if _, err := os.Stat(runnerTestFilePath); err == nil {
		return fmt.Errorf("File %s already exists", runnerTestFilePath)
	}

	tmpl, err := template.New("runnerTest").Parse(testRunnerTemplate)
	if err != nil {
		return fmt.Errorf("Failed to parse template: %v", err)
	}

	if err := tmpl.Execute(&output, struct {
		Day  int
		Year int
	}{Day: day, Year: year}); err != nil {
		return fmt.Errorf("Failed to execute template: %v", err)
	}

	if err := os.WriteFile(runnerTestFilePath, output.Bytes(), 0644); err != nil {
		return fmt.Errorf("Failed to write file %s: %v", runnerTestFilePath, err)
	}

	output.Reset()

	tmpl, err = template.New("runner").Parse(runnerTemplate)
	if err != nil {
		return fmt.Errorf("Failed to parse template: %v", err)
	}

	if err := tmpl.Execute(&output, struct{ Day int }{Day: day}); err != nil {
		return fmt.Errorf("Failed to execute template: %v", err)
	}

	if err := os.WriteFile(runnerFilePath, output.Bytes(), 0644); err != nil {
		return fmt.Errorf("Failed to write file %s: %v", runnerFilePath, err)
	}

	fmt.Printf("Generated day %02d runner: %s\n", day, runnerFilePath)

	return nil

}
