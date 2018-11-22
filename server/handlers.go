package server

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

var primes []int = initPrimes()

func initPrimes() []int {
	out := make([]int, 2)
	out[0] = 2
	cod := 0
	for index := 3; index < 5000; index++ {
		if index%2 == 0 {
			continue
		}
		cod = 0
		bound := int(math.Ceil(math.Sqrt(float64(index))))
		for j := 2; j < bound; j++ {
			if index%j == 0 {
				cod++
				break
			}
		}
		if cod == 0 {
			out = append(out, index)
		}
	}
	return out
}

func calculateNSD(w http.ResponseWriter, r *http.Request) {
	primesA := make([]int, 0)
	primesB := make([]int, 0)

	input := NSDN{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(404)
	}

	largest := 0

	for _, element := range primes {
		if input.NumA%element == 0 && element < input.NumA {
			primesA = append(primesA, element)
		}
		if input.NumB%element == 0 && element < input.NumB {
			primesB = append(primesB, element)
		}
		if input.NumA < element && input.NumB < element {
			break
		}
		if input.NumB%element == 0 && input.NumA%element == 0 {
			largest = element
		}
	}

	output := NSDN{Response: largest}
	fin, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(500)
	}
	fmt.Fprintf(w, string(fin))

}

type NSDN struct {
	NumA     int `json:"first`
	NumB     int `json:"second"`
	Response int `json:"gcd"`
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hello this is server")
}
