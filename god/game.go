package main

type Game struct {
	Citizens map[Role]*Person
	Users map[int]*Person
	IsNight  bool
}

func (game Game) process() {
	if game.Citizens[Cupbearer].canAct() {
		game.Users[game.Citizens[Cupbearer].getSelectedPerson().getUser().Id].setDrunk()
	}

	if game.Citizens[GodFather].canAct() {
		game.Users[game.Citizens[GodFather].getSelectedPerson().getUser().Id].kill()
	}

	if game.Citizens[Angle].canAct() {
		game.Users[game.Citizens[Angle].getSelectedPerson().getUser().Id].heal()
	}

	if game.Citizens[Detective].canAct() {
		game.Users[game.Citizens[Detective].getSelectedPerson().getUser().Id].heal()
	}
}
