package main

type Person struct {
	user           User
	selectedPerson *Person
	isDrunk        bool
	isDead         bool
}

func (Person *Person) getUser() User {
	return Person.user
}

func (Person *Person) getRole() Role {
	return GodFather
}

func (Person *Person) setSelectedPerson(person *Person) {
	Person.selectedPerson = person
}

func (Person *Person) getSelectedPerson() *Person {
	return Person.selectedPerson
}

func (Person *Person) setDrunk() {
	Person.isDrunk = true
}

func (Person *Person) canAct() bool {
	return !Person.isDead && !Person.isDrunk
}

func (Person *Person) kill() {
	Person.isDead = true
}

func (Person *Person) heal() {
	Person.isDead = false
}
