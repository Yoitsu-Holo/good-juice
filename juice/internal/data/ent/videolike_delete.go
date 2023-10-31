// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"juice/internal/data/ent/predicate"
	"juice/internal/data/ent/videolike"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoLikeDelete is the builder for deleting a VideoLike entity.
type VideoLikeDelete struct {
	config
	hooks    []Hook
	mutation *VideoLikeMutation
}

// Where appends a list predicates to the VideoLikeDelete builder.
func (vld *VideoLikeDelete) Where(ps ...predicate.VideoLike) *VideoLikeDelete {
	vld.mutation.Where(ps...)
	return vld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vld *VideoLikeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vld.sqlExec, vld.mutation, vld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vld *VideoLikeDelete) ExecX(ctx context.Context) int {
	n, err := vld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vld *VideoLikeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(videolike.Table, sqlgraph.NewFieldSpec(videolike.FieldID, field.TypeInt))
	if ps := vld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vld.mutation.done = true
	return affected, err
}

// VideoLikeDeleteOne is the builder for deleting a single VideoLike entity.
type VideoLikeDeleteOne struct {
	vld *VideoLikeDelete
}

// Where appends a list predicates to the VideoLikeDelete builder.
func (vldo *VideoLikeDeleteOne) Where(ps ...predicate.VideoLike) *VideoLikeDeleteOne {
	vldo.vld.mutation.Where(ps...)
	return vldo
}

// Exec executes the deletion query.
func (vldo *VideoLikeDeleteOne) Exec(ctx context.Context) error {
	n, err := vldo.vld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{videolike.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vldo *VideoLikeDeleteOne) ExecX(ctx context.Context) {
	if err := vldo.Exec(ctx); err != nil {
		panic(err)
	}
}