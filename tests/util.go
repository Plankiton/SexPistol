package main

import Sex "github.com/Plankiton/SexPistol"

func main() {
	j := map[string]string{
		"joao": "maria",
	}
	m := struct {
		Joao  interface{} `json:"joao"`
		Maria interface{} `json:"maria"`
	}{
		Joao:  false,
		Maria: true,
	}
	Sex.Log(m)

	Sex.Merge(j, &m)
	Sex.Log(m)

	Sex.Merge(j, &m, true)
	Sex.Log(m.Maria)

	Sex.Merge(m, &j, true)
	Sex.Log(j["maria"])
}
