package minimax

type gameState interface {
  terminal() bool
  value() float64
  nextState() []gameState
}

func maximize(g gameState) (gameState, float64) {
 if g.terminal() {
   return g, g.value()
 }
 states := g.nextState()
 if len(states) == 0 {
   return g, 0
 }
}

func maximizeLimit(g gameState, lower float64, upper float64) float64 {

}

func minimize(gameState g) (gameState, float64) { 

}

func minimizeLimit(gameState g, float64 value) (float64) {

}
