// Generate a string with characters that have odd counts
package main

import (
  "fmt"
  "math/rand"
  "time"
  "strings"
)

func init() {
  rand.Seed(time.Now().UnixNano())
}

var lowerCaseLetters = []rune("abcdefghijklmnopqrstuvwxyz")

func randCharacter() string {
  return string(lowerCaseLetters[rand.Intn(len(lowerCaseLetters))])
}

func generateTheString1(n int) string {
  var result string
  for n > 0 {
    c := randCharacter()
    m := rand.Intn(n)
    if m == 0 {
      m += 1
    }
    if m % 2 == 0 {
      m -= 1
    }
    result = result + strings.Repeat(c, m)
    fmt.Println(result, n, c, m)
    n = n -m
  }

  return result
}

func generateTheString2(n int) string {
  var result string
  c := randCharacter()

  for i := 0; i < n - 1; i++ {
    result += c
  }

  if n % 2 == 0 {
    c2 := randCharacter()
    for c2 == c {
      c2 = randCharacter()
    }
    result += c2
  } else {
    result += c
  }

  return result
}

func main () {

  test := generateTheString2(4)

  fmt.Println(test)
}

