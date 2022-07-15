package Helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (error, string) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err, string(hashed)
	}
	return err, string(hashed)
}
func CheckPassword(password string, hashedpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}
