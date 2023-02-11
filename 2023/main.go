package main

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fetch()
}
