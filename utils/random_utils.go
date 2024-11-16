package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateJobID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("job-%d", rand.Intn(100000))
}
