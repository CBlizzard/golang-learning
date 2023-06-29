package main

import "fmt"

func main() {
	pokemons := make(map[string]string)

	pokemons["fire"] = "charmander"
	pokemons["water"] = "squirtle"
	pokemons["grass"] = "bulbasaur"
	pokemons["electric"] = "pikachu"
	pokemons["ice"] = "glalie"

	fmt.Println(pokemons)

	delete(pokemons, "ice")
	fmt.Println(pokemons)

}
