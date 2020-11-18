// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/masseelch/elk/_example"
)

// The OwnerFunc type is an adapter to allow the use of ordinary
// function as Owner mutator.
type OwnerFunc func(context.Context, *_example.OwnerMutation) (_example.Value, error)

// Mutate calls f(ctx, m).
func (f OwnerFunc) Mutate(ctx context.Context, m _example.Mutation) (_example.Value, error) {
	mv, ok := m.(*_example.OwnerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *_example.OwnerMutation", m)
	}
	return f(ctx, mv)
}

// The PetFunc type is an adapter to allow the use of ordinary
// function as Pet mutator.
type PetFunc func(context.Context, *_example.PetMutation) (_example.Value, error)

// Mutate calls f(ctx, m).
func (f PetFunc) Mutate(ctx context.Context, m _example.Mutation) (_example.Value, error) {
	mv, ok := m.(*_example.PetMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *_example.PetMutation", m)
	}
	return f(ctx, mv)
}

// The SkipGenerationModelFunc type is an adapter to allow the use of ordinary
// function as SkipGenerationModel mutator.
type SkipGenerationModelFunc func(context.Context, *_example.SkipGenerationModelMutation) (_example.Value, error)

// Mutate calls f(ctx, m).
func (f SkipGenerationModelFunc) Mutate(ctx context.Context, m _example.Mutation) (_example.Value, error) {
	mv, ok := m.(*_example.SkipGenerationModelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *_example.SkipGenerationModelMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, _example.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m _example.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m _example.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m _example.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op _example.Op) Condition {
	return func(_ context.Context, m _example.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m _example.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m _example.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m _example.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk _example.Hook, cond Condition) _example.Hook {
	return func(next _example.Mutator) _example.Mutator {
		return _example.MutateFunc(func(ctx context.Context, m _example.Mutation) (_example.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, _example.Delete|_example.Create)
//
func On(hk _example.Hook, op _example.Op) _example.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, _example.Update|_example.UpdateOne)
//
func Unless(hk _example.Hook, op _example.Op) _example.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) _example.Hook {
	return func(_example.Mutator) _example.Mutator {
		return _example.MutateFunc(func(context.Context, _example.Mutation) (_example.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []_example.Hook {
//		return []_example.Hook{
//			Reject(_example.Delete|_example.Update),
//		}
//	}
//
func Reject(op _example.Op) _example.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []_example.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ..._example.Hook) Chain {
	return Chain{append([]_example.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() _example.Hook {
	return func(mutator _example.Mutator) _example.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ..._example.Hook) Chain {
	newHooks := make([]_example.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
