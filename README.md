# HOW TO USE

```go
const salt = "???ExX4Mpl33rR%($(($$("

func main() {
	plaintext := "test_11_@s"

	h := new(hasher.Hasher)
	h.Plaintext = plaintext
	h.Salt = salt
	e := h.Hash()

	if e == nil {
		fmt.Println("PLAIN =", plaintext)
		fmt.Println("ENCRYPT = ", h.Hashed)
	}
}
```