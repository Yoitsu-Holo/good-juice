// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"juice/public/ent/userfollowinfo"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFollowInfoCreate is the builder for creating a UserFollowInfo entity.
type UserFollowInfoCreate struct {
	config
	mutation *UserFollowInfoMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ufic *UserFollowInfoCreate) SetUserID(u uint64) *UserFollowInfoCreate {
	ufic.mutation.SetUserID(u)
	return ufic
}

// SetFollowID sets the "follow_id" field.
func (ufic *UserFollowInfoCreate) SetFollowID(u uint64) *UserFollowInfoCreate {
	ufic.mutation.SetFollowID(u)
	return ufic
}

// SetStatus sets the "status" field.
func (ufic *UserFollowInfoCreate) SetStatus(i int8) *UserFollowInfoCreate {
	ufic.mutation.SetStatus(i)
	return ufic
}

// SetCreateTime sets the "create_time" field.
func (ufic *UserFollowInfoCreate) SetCreateTime(t time.Time) *UserFollowInfoCreate {
	ufic.mutation.SetCreateTime(t)
	return ufic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ufic *UserFollowInfoCreate) SetNillableCreateTime(t *time.Time) *UserFollowInfoCreate {
	if t != nil {
		ufic.SetCreateTime(*t)
	}
	return ufic
}

// SetUpdateTime sets the "update_time" field.
func (ufic *UserFollowInfoCreate) SetUpdateTime(t time.Time) *UserFollowInfoCreate {
	ufic.mutation.SetUpdateTime(t)
	return ufic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ufic *UserFollowInfoCreate) SetNillableUpdateTime(t *time.Time) *UserFollowInfoCreate {
	if t != nil {
		ufic.SetUpdateTime(*t)
	}
	return ufic
}

// SetID sets the "id" field.
func (ufic *UserFollowInfoCreate) SetID(i int) *UserFollowInfoCreate {
	ufic.mutation.SetID(i)
	return ufic
}

// Mutation returns the UserFollowInfoMutation object of the builder.
func (ufic *UserFollowInfoCreate) Mutation() *UserFollowInfoMutation {
	return ufic.mutation
}

// Save creates the UserFollowInfo in the database.
func (ufic *UserFollowInfoCreate) Save(ctx context.Context) (*UserFollowInfo, error) {
	return withHooks(ctx, ufic.sqlSave, ufic.mutation, ufic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ufic *UserFollowInfoCreate) SaveX(ctx context.Context) *UserFollowInfo {
	v, err := ufic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufic *UserFollowInfoCreate) Exec(ctx context.Context) error {
	_, err := ufic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufic *UserFollowInfoCreate) ExecX(ctx context.Context) {
	if err := ufic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ufic *UserFollowInfoCreate) check() error {
	if _, ok := ufic.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserFollowInfo.user_id"`)}
	}
	if _, ok := ufic.mutation.FollowID(); !ok {
		return &ValidationError{Name: "follow_id", err: errors.New(`ent: missing required field "UserFollowInfo.follow_id"`)}
	}
	if _, ok := ufic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "UserFollowInfo.status"`)}
	}
	return nil
}

func (ufic *UserFollowInfoCreate) sqlSave(ctx context.Context) (*UserFollowInfo, error) {
	if err := ufic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ufic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ufic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ufic.mutation.id = &_node.ID
	ufic.mutation.done = true
	return _node, nil
}

func (ufic *UserFollowInfoCreate) createSpec() (*UserFollowInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &UserFollowInfo{config: ufic.config}
		_spec = sqlgraph.NewCreateSpec(userfollowinfo.Table, sqlgraph.NewFieldSpec(userfollowinfo.FieldID, field.TypeInt))
	)
	if id, ok := ufic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ufic.mutation.UserID(); ok {
		_spec.SetField(userfollowinfo.FieldUserID, field.TypeUint64, value)
		_node.UserID = value
	}
	if value, ok := ufic.mutation.FollowID(); ok {
		_spec.SetField(userfollowinfo.FieldFollowID, field.TypeUint64, value)
		_node.FollowID = value
	}
	if value, ok := ufic.mutation.Status(); ok {
		_spec.SetField(userfollowinfo.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := ufic.mutation.CreateTime(); ok {
		_spec.SetField(userfollowinfo.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := ufic.mutation.UpdateTime(); ok {
		_spec.SetField(userfollowinfo.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// UserFollowInfoCreateBulk is the builder for creating many UserFollowInfo entities in bulk.
type UserFollowInfoCreateBulk struct {
	config
	err      error
	builders []*UserFollowInfoCreate
}

// Save creates the UserFollowInfo entities in the database.
func (uficb *UserFollowInfoCreateBulk) Save(ctx context.Context) ([]*UserFollowInfo, error) {
	if uficb.err != nil {
		return nil, uficb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uficb.builders))
	nodes := make([]*UserFollowInfo, len(uficb.builders))
	mutators := make([]Mutator, len(uficb.builders))
	for i := range uficb.builders {
		func(i int, root context.Context) {
			builder := uficb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserFollowInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uficb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uficb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uficb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uficb *UserFollowInfoCreateBulk) SaveX(ctx context.Context) []*UserFollowInfo {
	v, err := uficb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uficb *UserFollowInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := uficb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uficb *UserFollowInfoCreateBulk) ExecX(ctx context.Context) {
	if err := uficb.Exec(ctx); err != nil {
		panic(err)
	}
}
