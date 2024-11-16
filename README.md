# anyset

A generic definition of a `Set` interface that can contain any kind of value,
and whose underlying container structure is not imposed or restricted in any
way.

Implementors of `Set` are supposed to manage their own underlying containers.

> **_Note_**: Check `set_test.go` for the intended usage of `Set` or if you are unsure where
to start.

## Why

Suppose you are designing a library, and your library uses a `Set` for some operations. Your users may want to provide a different implementation of the `Set` interface to accommodate their performance needs.

Or.. I just wanted to try out generics in `go`.

## Examples

Two `Set`ters are provided as examples, the algorithmic performance of either
implementation has not been optimized in any way, they are just PoCs:

- MapSet: uses a `map[T]bool` as container.
- SliceSet: uses a `[]T` as container.

## Testing

Verifiers of `Set`ters are provided in `set_test.go`. Use these verifiers to
test your implementation.

## Notes

This is my first program/code in `go`.
