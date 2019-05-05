package main

import "fmt"

func main() {
	m := map[string]string{
		"name":  "thomas",
		"skill": "goland"}

	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m, m2, m3)
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("getting values")
	skill := m["skill"]
	fmt.Println(skill)
	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
