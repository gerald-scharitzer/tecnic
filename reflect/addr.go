package reflect

// Get the address of a non-addressable value as expression
//
//	a := Addr(value)
//
// instead of multiple statements:
//
//	v := value
//	a := &v
//
// This is for constant and literal values,
// because the [address operator](https://go.dev/ref/spec#Address_operators)
// is defined for addressable operands only.
//
//	func f(number *int) {
//	  fmt.Println(*number)
//	}
//
// There is no automatic boxing in Go,
// so function arguments are not referenced automatically.
//
//	const c = 42
//	// f(c) // error: cannot use c (untyped int constant 42) as *int value in argument to f compiler(IncompatibleAssign)
//
//	v := c
//	// f(v) // error: cannot use v (variable of type int) as *int value in argument to f compiler(IncompatibleAssign)
//
//	// f(42) // error: cannot use 42 (untyped int constant) as *int value in argument to f compiler(IncompatibleAssign)
//
// Instead the address of a function argument must be obtained explicitly,
// but the address operator is defined for addressable operands only.
//
//	// f(&c) // invalid operation: cannot take address of c (untyped int constant 42) compiler(UnaddressableOperand)
//	f(&v) // requires declaration of variable v
//	// f(&42) // invalid operation: cannot take address of 42 (untyped int constant) compiler(UnaddressableOperand)
//
// The function `Addr` allocates a variable implicitly and returns its address.
//
//	f(Addr(c))
//	f(Addr(42))
//	f(Addr(v)) // is less efficient than f(&v) and points to the same value, but a different variable
//
// The argument is passed by value, so the address of the copy is returned instead of the original argument.
// The memory is allocated implicitly through pass by value.
func Addr[T any](x T) *T {
	return &x
}
