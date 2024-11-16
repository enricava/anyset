# anyset

A generic definition of a `Set` interface that can contain any kind of value,
and whose underlying container structure is not imposed or restricted in any
way.

Implementors of `Set` are supposed to manage their own underlying containers.

> **_Note_**: Check `set_test.go` for the intended usage of `Set` or if you are unsure where
to start.

## Examples

Two `Set`ters are provided as examples, the algorithmic performance of either
implementation has not been optimized in any way, they are just PoCs:

- MapSet: uses a `map[T]bool` as container.
- SliceSet: uses a `[]T` as container.

## Testing

Implementors of `Set` are not required to provide additional tests, since
two testing functions are provided for the `Set` interface.

Verifiers of `Set`ters are provided in `set_test.go`. Use these verifiers to
test your implementation.

## Notes

This is my first program/code in `go`.
