package main

func main() {
	screen := Screen{}

	screen.Init(&screenConfig{
		width:  80,
		height: 20,
		backgroundChar: '_',
	})
	screen.render()
}
