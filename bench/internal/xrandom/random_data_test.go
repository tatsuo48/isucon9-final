package xrandom

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/chibiegg/isucon9-final/bench/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestRandomUseAt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	assert.NoError(t, config.SetAvailReserveDays(30))
	for i := 0; i < 10; i++ {
		log.Println(GetRandomUseAt().String())
	}
}

func TestRandomSection(t *testing.T) {
	for i := 0; i < 10; i++ {
		s1, s2 := GetRandomSection()
		log.Printf("[*] s1=%s, s2=%s\n", s1, s2)
	}
}
