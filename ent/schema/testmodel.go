package schema

import (
	"context"
	"test/ent/hook"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"

	gen "test/ent"
)

// TestModel holds the schema definition for the TestModel entity.
type TestModel struct {
	ent.Schema
}

// Mixin defines the mixins used for this entity
func (TestModel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}

// Fields of the TestModel.
func (TestModel) Fields() []ent.Field {
	return []ent.Field{
		field.String("test"),
	}
}

// Edges of the TestModel.
func (TestModel) Edges() []ent.Edge {
	return nil
}

// Hooks of the TestModel
func (TestModel) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.TestModelFunc(func(ctx context.Context, m *gen.TestModelMutation) (ent.Value, error) {

				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
	}
}
