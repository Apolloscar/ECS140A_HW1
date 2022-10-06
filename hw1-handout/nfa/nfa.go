package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

// You may define helper functions here.

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO: Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise
	//panic("TODO: implement this!")

	if len(input) == 0 { //check if sring is empty
		if start == final { //if so then start and final need to be same since no transitions will occur, it is reachable if so
			return true
		} else {
			return false //if start and final are not same then it is un reachable with blank input
		}
	}

	var next_state []state = transitions(start, input[0]) //with next symbol, determine what states can be reached

	var result bool

	for i := 0; i < len(next_state); i++ { //traverse to all possible next states and see if from there are any reachable with the same string minus the first symbol
		result = Reachable(transitions, next_state[i], final, input[1:]) //go one possible path with string without the first symbol
		if result == true {                                              // if it reached final then it is reachable so return true, if not then check another path
			return true
		}
	}

	return false // all paths checked and none coul reach final state

}

//---------------------------------------------------------------------------------------------------------------------------------------------------------------

//Logic for the recursion:
//1. if symbols are blank then start and final state must be same to return true else false
//2. In order for string to make it reachable from start state to final state, the same string without the first symbol must make it reachable from any of the next states to final state
//3. Step 2 gives all possible paths that the string can take from start state so if noe are the final state then it is unreachable

//Example from the model in handout:
//Is state 2 reachable from state 0 with "aba"?

//For that to be true state 2 must be reachable from state 1 or 2 with "ba", since those are next states for 0 given 'a'

//For that to be true state 2 must be reachable from state 0 with "a", since 1 goes to 0 with 'b' but 2 goes nowhere with 'b'

//For that to be true state 2 must be reachable from state 2 or 1 with "", since 0 goes to 2 or 1 with 'a'

//with state 2 go to 2 with "", that is our base case so it is true
