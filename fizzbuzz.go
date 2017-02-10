package main

import "fmt"

func main() {
  fizzbuzz(1)
  fizzbuzz(2)
  fizzbuzz(3)
  fizzbuzz(4)
  fizzbuzz(5)
  fizzbuzz(6)
  fizzbuzz(7)
  fizzbuzz(8)
  fizzbuzz(9)
  fizzbuzz(10)
  fizzbuzz(11)
  fizzbuzz(12)
  fizzbuzz(13)
  fizzbuzz(14)
  fizzbuzz(15)
  fizzbuzz(27)
  fizzbuzz(30)
  fizzbuzz(50)
}

func fizzbuzz(number int) {
  fizz := "fizz"
  buzz := "buzz"
  fizzbuzz := fizz + buzz

  if 0 == number % 3 && 0 == number % 5 {
    fmt.Println(number, fizzbuzz)
  } else if 0 == number % 3  {
    fmt.Println(number, fizz)
  } else if 0 == number % 5 {
    fmt.Println(number, buzz)
  } else {
    fmt.Println(number)
  }
}
