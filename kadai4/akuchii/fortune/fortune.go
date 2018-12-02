package fortune

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var fortuneList = []string{"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶"}

// Clock is interface to get current time
type Clock interface {
	GetCurrentTime() time.Time
}

// DefaultClock equals current time
type DefaultClock struct{}

// Fortune lots fortune depends on time
type Fortune struct {
	clock Clock
}

// LotResult contains result of fortune
type LotResult struct {
	Result string `json:"result"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetCurrentTime returns current time
func (d DefaultClock) GetCurrentTime() time.Time {
	return time.Now()
}

// NewFortune returns fortune instance
func NewFortune(c Clock) *Fortune {
	return &Fortune{clock: c}
}

func (f Fortune) lotForNewYearDay() string {
	return "大吉"
}

func (f Fortune) defaultLot() string {
	return fortuneList[rand.Intn(len(fortuneList))]
}

func (f Fortune) isNewYearDay() bool {
	c := f.clock.GetCurrentTime()

	return c.Month() == 1 && 1 <= c.Day() && c.Day() <= 3
}

// Lot lots fortune
func (f Fortune) Lot() string {
	if f.isNewYearDay() {
		return f.lotForNewYearDay()
	}
	return f.defaultLot()
}

// Handler returns fortune result
func (f Fortune) Handler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	fr := LotResult{Result: f.Lot()}
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(fr); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, buf.String())
}
