package util

import (
	"math/rand"
	"time"
)


func init()  {
	rand.Seed(time.Now().UnixNano())
}

// Random integer generates a random integer between min and max

func RandomInit(min, max int64) int64 {
	return rand.Int63n(max - min) + min
}


//Random string generator

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}




func RandomRole() string  {
	roles := []string{"admin", "user"}
	return roles[rand.Intn(len(roles))]
}