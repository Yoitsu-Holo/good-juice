// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"juice/public/ent/videocollection"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoCollectionCreate is the builder for creating a VideoCollection entity.
type VideoCollectionCreate struct {
	config
	mutation *VideoCollectionMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (vcc *VideoCollectionCreate) SetUserID(i int) *VideoCollectionCreate {
	vcc.mutation.SetUserID(i)
	return vcc
}

// SetVideoID sets the "video_id" field.
func (vcc *VideoCollectionCreate) SetVideoID(i int) *VideoCollectionCreate {
	vcc.mutation.SetVideoID(i)
	return vcc
}

// SetStatus sets the "status" field.
func (vcc *VideoCollectionCreate) SetStatus(i int8) *VideoCollectionCreate {
	vcc.mutation.SetStatus(i)
	return vcc
}

// SetCreateTime sets the "create_time" field.
func (vcc *VideoCollectionCreate) SetCreateTime(t time.Time) *VideoCollectionCreate {
	vcc.mutation.SetCreateTime(t)
	return vcc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (vcc *VideoCollectionCreate) SetNillableCreateTime(t *time.Time) *VideoCollectionCreate {
	if t != nil {
		vcc.SetCreateTime(*t)
	}
	return vcc
}

// SetUpdateTime sets the "update_time" field.
func (vcc *VideoCollectionCreate) SetUpdateTime(t time.Time) *VideoCollectionCreate {
	vcc.mutation.SetUpdateTime(t)
	return vcc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (vcc *VideoCollectionCreate) SetNillableUpdateTime(t *time.Time) *VideoCollectionCreate {
	if t != nil {
		vcc.SetUpdateTime(*t)
	}
	return vcc
}

// SetID sets the "id" field.
func (vcc *VideoCollectionCreate) SetID(i int) *VideoCollectionCreate {
	vcc.mutation.SetID(i)
	return vcc
}

// Mutation returns the VideoCollectionMutation object of the builder.
func (vcc *VideoCollectionCreate) Mutation() *VideoCollectionMutation {
	return vcc.mutation
}

// Save creates the VideoCollection in the database.
func (vcc *VideoCollectionCreate) Save(ctx context.Context) (*VideoCollection, error) {
	return withHooks(ctx, vcc.sqlSave, vcc.mutation, vcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vcc *VideoCollectionCreate) SaveX(ctx context.Context) *VideoCollection {
	v, err := vcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcc *VideoCollectionCreate) Exec(ctx context.Context) error {
	_, err := vcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcc *VideoCollectionCreate) ExecX(ctx context.Context) {
	if err := vcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vcc *VideoCollectionCreate) check() error {
	if _, ok := vcc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "VideoCollection.user_id"`)}
	}
	if _, ok := vcc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video_id", err: errors.New(`ent: missing required field "VideoCollection.video_id"`)}
	}
	if _, ok := vcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "VideoCollection.status"`)}
	}
	return nil
}

func (vcc *VideoCollectionCreate) sqlSave(ctx context.Context) (*VideoCollection, error) {
	if err := vcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	vcc.mutation.id = &_node.ID
	vcc.mutation.done = true
	return _node, nil
}

func (vcc *VideoCollectionCreate) createSpec() (*VideoCollection, *sqlgraph.CreateSpec) {
	var (
		_node = &VideoCollection{config: vcc.config}
		_spec = sqlgraph.NewCreateSpec(videocollection.Table, sqlgraph.NewFieldSpec(videocollection.FieldID, field.TypeInt))
	)
	if id, ok := vcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vcc.mutation.UserID(); ok {
		_spec.SetField(videocollection.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := vcc.mutation.VideoID(); ok {
		_spec.SetField(videocollection.FieldVideoID, field.TypeInt, value)
		_node.VideoID = value
	}
	if value, ok := vcc.mutation.Status(); ok {
		_spec.SetField(videocollection.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := vcc.mutation.CreateTime(); ok {
		_spec.SetField(videocollection.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := vcc.mutation.UpdateTime(); ok {
		_spec.SetField(videocollection.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// VideoCollectionCreateBulk is the builder for creating many VideoCollection entities in bulk.
type VideoCollectionCreateBulk struct {
	config
	err      error
	builders []*VideoCollectionCreate
}

// Save creates the VideoCollection entities in the database.
func (vccb *VideoCollectionCreateBulk) Save(ctx context.Context) ([]*VideoCollection, error) {
	if vccb.err != nil {
		return nil, vccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vccb.builders))
	nodes := make([]*VideoCollection, len(vccb.builders))
	mutators := make([]Mutator, len(vccb.builders))
	for i := range vccb.builders {
		func(i int, root context.Context) {
			builder := vccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoCollectionMutation)
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
					_, err = mutators[i+1].Mutate(root, vccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vccb *VideoCollectionCreateBulk) SaveX(ctx context.Context) []*VideoCollection {
	v, err := vccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vccb *VideoCollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := vccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vccb *VideoCollectionCreateBulk) ExecX(ctx context.Context) {
	if err := vccb.Exec(ctx); err != nil {
		panic(err)
	}
}
