package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	name := "World"
	user := User{Name: "Krystyna", Age: 67}
	items := []string{"go", "rust", "python"}
	price := 49.99
	qty := 3

	st := "standard string declaration: %{user.Name}."
	rs := `raw string declaration: %{user.Name}.`

	fmt.Println("--- go-fmts string interpolation ---")
	fmt.Println(st)
	fmt.Println(rs)
	fmt.Println("Hello %{name}!")
	fmt.Println("User: %{user.Name}, age: %{user.Age:d}")
	fmt.Println("First language: %{items[0]}")
	fmt.Println("Total: %{len(items)} languages")
	fmt.Println("Price: $%{price:.2f} x %{qty} = $%{price * float64(qty):.2f}")
	fmt.Println("Braces: { single are literal }")

	_, _, _, _, _ = name, user, items, price, qty
}
