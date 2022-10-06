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

	if len(input) == 0 {
		if start == final {
			return true
		} else {
			return false
		}
	}

	var next_state []state = transitions(start, input[0])

	var result bool

	for i := 0; i < len(next_state); i++ {
		result = Reachable(transitions, next_state[i], final, input[1:])
		if result == true {
			return true
		}
	}

	return false

}
