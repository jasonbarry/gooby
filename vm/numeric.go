package vm

// Numeric currently represents a class that support some numeric conversions.
// At this stage, it's not meant to be a Gooby class in a strict sense, but only
// a convenient interface.
type Numeric interface {
	floatValue() float64
	lessThan(object Object) bool
}
