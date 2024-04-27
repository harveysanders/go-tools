package pkg

import "errors"

var Err = errors.New("error")

func fn(err, target error) {
	errors.Is(err, target)

	errors.Is(target, err) //@ diag("flipped err and target arguments")
	errors.Is(Err, err)    //@ diag("flipped err and target arguments")
	errors.Is(err, Err)
}
