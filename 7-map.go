package main

func main() {
	person := map[string]string{
		"name":    "aziz",
		"address": "palembang",
	}

	delete(person, "")

	var book = map[string]string = make(map[string]string)

	book['test'] = "1"
	len(book)

	book['aa'] = "2"
	len(book)
}
