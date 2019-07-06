package designpattern

func Example() {
	var player Player

	player = &Boy{Person{"will"}}
	player.Play()

	player = &Girl{Person{"jean"}}
	player.Play()

	player = &Woman{Girl{Person{Name: "jean"}}}
	player.Play()

	woman := Woman{}
	woman.Name = "jean"
	player = &woman
	player.Play()
}
