package main

import (
	"fmt"

	"github.com/mndrix/golog"
)

func main() {
	m := golog.NewMachine().Consult(`
      father(john).
      father(jacob).

      mother(sue).

      parent(X) :-
          father(X).
      parent(X) :-
          mother(X).
  `)
	if m.CanProve(`father(john).`) {
		fmt.Printf("john is a father\n")
	}

	solutions := m.ProveAll(`parent(X).`)
	for _, solution := range solutions {
		fmt.Printf("%s is a parent\n", solution.ByName_("X"))
	}

}
