run:
	go run main.go -year=$(year) -day=$(day) -part=$(part) -test=$(test)
generate:
	go run main.go -generator=true -day=$(day) -year=$(year)
