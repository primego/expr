package expr

import "time"

type (
	// Bool is an alias for Expr[bool].
	Bool = Expr[bool]
	// Float32 is an alias for Expr[float32].
	Float32 = Expr[float32]
	// Float64 is an alias for Expr[float64].
	Float64 = Expr[float64]
	// Int is an alias for Expr[int].
	Int = Expr[int]
	// Int8 is an alias for Expr[int8].
	Int8 = Expr[int8]
	// Int16 is an alias for Expr[int16].
	Int16 = Expr[int16]
	// Int32 is an alias for Expr[int32].
	Int32 = Expr[int32]
	// Int64 is an alias for Expr[int64].
	Int64 = Expr[int64]
	// String is an alias for Expr[string].
	String = Expr[string]
	// Time is an alias for Expr[time.Time].
	Time = Expr[time.Time]
	// Uint is an alias for Expr[uint].
	Uint = Expr[uint]
	// Uint8 is an alias for Expr[uint8].
	Uint8 = Expr[uint8]
	// Uint16 is an alias for Expr[uint16].
	Uint16 = Expr[uint16]
	// Uint32 is an alias for Expr[uint32].
	Uint32 = Expr[uint32]
	// Uint64 is an alias for Expr[uint64].
	Uint64 = Expr[uint64]
)
