package auth

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	userID := 10

	utoken := CreateToken(userID)

	fmt.Printf("CreateToken response: %v", utoken)
}

func BenchmarkCreateToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateToken(n)
	}
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE0NDQ0Nzg0MDAsInVzZXJpZCI6MTB9.y3yWjqvopwq32NjT84LSYTPA8v3WH9poYokqtglop7I"
	userID, err := ParseToken(token)

	if err != nil {
		t.Errorf("Failed to parse token: %v", err)
		return
	}

	if userID != 10 {
		t.Errorf("Unable to parse token! expected: %v, got: %v",
			10, userID)
		return
	}
}
