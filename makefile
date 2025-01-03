run:
	go run main.go -year=$(year) -day=$(day) -part=$(part)
generate:
	go run main.go -generator=true -day=$(day) -year=$(year)
test:
	go test ./aoc_$(year)/day$(day) -v