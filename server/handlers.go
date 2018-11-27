package server

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

var primes = initPrimes()

func initPrimes() []int {
	out := make([]int, 2)
	out[0] = 2
	cod := 0
	for index := 3; index < 500; index++ {
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

type source struct {
	Things []int `json:"list,omitempty"`
}

func calculateNSD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("aaaa")
	primesA := make([]int, 0)
	primesB := make([]int, 0)

	input := nSDN{}
	//some, err := r.Body.Read()
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("shit happens")
		w.WriteHeader(404)
		fmt.Fprintf(w, "error:\"unexpected\"")
		return
	}

	largest := 0

	for _, element := range primes {
		if input.First%element == 0 && element < input.First {
			primesA = append(primesA, element)
		}
		if input.Second%element == 0 && element < input.Second {
			primesB = append(primesB, element)
		}
		if input.First < element && input.Second < element {
			break
		}
		if input.Second%element == 0 && input.First%element == 0 {
			largest = element
		}
	}

	input.Response = largest
	fin, err := json.Marshal(input)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error:\"internal\"")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, string(fin))
}

//
type nSDN struct {
	First    int `json:"first,omitempty"`
	Second   int `json:"second,omitempty"`
	Response int `json:"gcd,ommitempty,omitempty"`
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hello this is server")
}
