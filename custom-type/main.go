package customtype

import "fmt"

type Celcius float64
type Benang int

func (c Celcius) ToFahrenheit() float64 {
	return float64(Celcius(c*9/5 + 32))
}

func (c *Celcius) add(value float64) Benang {
	*c += Celcius(value)
	return Benang(*c)
}

func main() {
	fmt.Println("")
}
