package value_objects

import "golang.org/x/crypto/bcrypt"

type HashPassword struct {
	value string
}

func NewHashPassword(value string) (*HashPassword, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value),
		bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &HashPassword{value: string(hashedPassword)}, nil
}

func (h *HashPassword) Compare(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(h.value), []byte(password))
}

func (h *HashPassword) GetValue() string {
	return h.value
}
