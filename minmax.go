package minimax

import (
  "fmt"
  "math"
)

type TicTacToe string

func (t TicTacToe) Terminal() bool {
  count := 0
  for i := 0; i < 9; i++ {
    if t[i] != ' ' {
      count ++
    }
  }
  return t.Value() != 0 || count == 9
}

func (t TicTacToe) Value() float64 {
  count := 0
  for i := 0; i < 9; i++ {
    if t[i] != ' ' {
      count ++
    }
  }
  w := ' ';
  for i:=0; i < 2; i++ {
    if t[i] == t[i+3] && t[i] == t[i+6] && t[i] != ' ' {
      w = rune(t[i])
    }
  }
  for i:=0; i < 9; i+=3 {
    if t[i] == t[i+1] && t[i] == t[i+2] && t[i] != ' ' {
      w = rune(t[i])
    }
  }
  if t[0] == t[4] && t[0] == t[8] && t[0] != ' ' {
    w = rune(t[0])
  }
  if t[2] == t[4] && t[2] == t[6] && t[2] != ' ' {
    w = rune(t[2])
  }
  if w == ' ' {
    return 0
  }
  if w == 'o' {
    return 1.0/float64(count)
  }
  return -1.0/float64(count)
}

func (t TicTacToe) NextState() []GameState {
  count := 0
  for i := 0; i < 9; i++ {
    if t[i] != ' ' {
      count ++
    }
  }
  var states []GameState
  next := "x"
  if count % 2 == 0 {
    next = "o"
  }
  for i := 0; i < 9; i++ {
    if t[i] == ' ' {
      states = append(states, TicTacToe(string(t[:i]) + next + string(t[i+1:])))
    }
  }
  return states
}

type GameState interface {
  Terminal() bool
  Value() float64
  NextState() []GameState
}


func Maximize(g GameState) (GameState, float64) {
  return MaximizeLimit(g, math.Inf(-1), math.Inf(1))
}

func Minimize(g GameState) (GameState, float64) {
  return MinimizeLimit(g, math.Inf(-1), math.Inf(1))
}

func MaximizeLimit(g GameState, lower float64, upper float64) (GameState, float64) {
 if g.Terminal() {
   return g, g.Value()
 }
 states := g.NextState()
 if len(states) == 0 {
   fmt.Println("ayoooo no child states on terminal node")
   return g, lower 
 }
 mState := states[0]

 for _, state := range states {
   _, score := MinimizeLimit(state, lower, upper)
   if score >= upper {
     return state, score
   }
   if score > lower {
     lower = score
     mState = state 
   }
 }
  
 return mState, lower
}

func MinimizeLimit(g GameState, lower, upper float64) (GameState, float64) {
 if g.Terminal() {
   return g, g.Value()
 }
 states := g.NextState()
 if len(states) == 0 {
   fmt.Println("ayoooo no child states on terminal node")
   return g, lower 
 }
 mState := states[0]

 for _, state := range states {
   _, score := MaximizeLimit(state, lower, upper)
   if score <= lower {
     return state, score
   }
   if score < upper {
     upper = score
     mState = state 
   }
 }
  
 return mState, upper
}
