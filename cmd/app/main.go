package main

func main() {
	app, clean, err := build()
	if err != nil {
		panic(err.Error())
	}
	defer clean()
	err = app.Run()
	if err != nil {
		panic(err.Error())
	}

}
