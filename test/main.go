package main
import (
  "minimax"
  "fmt"
  "net/http"
)

func main() {
  game := minimax.TicTacToe("ooxxx o  ")
  a, b := minimax.Maximize(game)
  yuh := a.(minimax.TicTacToe)
  fmt.Println(yuh[:3])
  fmt.Println(yuh[3:6])
  fmt.Println(yuh[6:])
  fmt.Println(b)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", "")
  })
}
