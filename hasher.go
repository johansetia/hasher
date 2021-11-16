package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type (
	// Hasher is used to hash a plaintext
	Hasher struct {
		Plaintext string
		Hashed    string
		Salt      string
		process   []rune
		code      rune
	}
)

func (h *Hasher) reverse() {
	for i, j := 0, len(h.process)-1; i < j; i, j = i+1, j-1 {
		h.process[i], h.process[j] = h.process[j], h.process[i]
	}
}

func (h *Hasher) rearrange() {
	var nr []rune
	for _, x := range h.process {
		nr = append(nr, x, h.code)
	}
	h.process = nr
}

func (h *Hasher) md5Salt() error {
	hash := md5.New()
	_, err := hash.Write([]byte(h.Salt + string(h.process) + h.Salt))
	if err != nil {
		return err
	}
	h.Hashed = hex.EncodeToString(hash.Sum(nil))
	return nil
}

// Hash is used to hash a plaintext
func (h *Hasher) Hash() error {

	h.process = []rune(h.Plaintext)
	l := len(h.process)
	if l == 0 {
		return errors.New("cannot pass an empty string in a hasher")
	}
	x := l * 7
	h.code = rune(l + x)

	h.reverse()
	h.rearrange()
	return h.md5Salt()

}
