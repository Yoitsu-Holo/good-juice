// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"juice/internal/data/ent/videometadata"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoMetadataCreate is the builder for creating a VideoMetadata entity.
type VideoMetadataCreate struct {
	config
	mutation *VideoMetadataMutation
	hooks    []Hook
}

// SetVideoID sets the "video_id" field.
func (vmc *VideoMetadataCreate) SetVideoID(i int64) *VideoMetadataCreate {
	vmc.mutation.SetVideoID(i)
	return vmc
}

// SetUserID sets the "user_id" field.
func (vmc *VideoMetadataCreate) SetUserID(i int64) *VideoMetadataCreate {
	vmc.mutation.SetUserID(i)
	return vmc
}

// SetCoverURL sets the "cover_url" field.
func (vmc *VideoMetadataCreate) SetCoverURL(s string) *VideoMetadataCreate {
	vmc.mutation.SetCoverURL(s)
	return vmc
}

// SetVideoURL sets the "video_url" field.
func (vmc *VideoMetadataCreate) SetVideoURL(s string) *VideoMetadataCreate {
	vmc.mutation.SetVideoURL(s)
	return vmc
}

// SetVideoIntro sets the "video_intro" field.
func (vmc *VideoMetadataCreate) SetVideoIntro(s string) *VideoMetadataCreate {
	vmc.mutation.SetVideoIntro(s)
	return vmc
}

// SetNillableVideoIntro sets the "video_intro" field if the given value is not nil.
func (vmc *VideoMetadataCreate) SetNillableVideoIntro(s *string) *VideoMetadataCreate {
	if s != nil {
		vmc.SetVideoIntro(*s)
	}
	return vmc
}

// SetVideoType sets the "video_type" field.
func (vmc *VideoMetadataCreate) SetVideoType(i int64) *VideoMetadataCreate {
	vmc.mutation.SetVideoType(i)
	return vmc
}

// SetPublishAddress sets the "publish_address" field.
func (vmc *VideoMetadataCreate) SetPublishAddress(i int32) *VideoMetadataCreate {
	vmc.mutation.SetPublishAddress(i)
	return vmc
}

// SetNillablePublishAddress sets the "publish_address" field if the given value is not nil.
func (vmc *VideoMetadataCreate) SetNillablePublishAddress(i *int32) *VideoMetadataCreate {
	if i != nil {
		vmc.SetPublishAddress(*i)
	}
	return vmc
}

// SetCreateTime sets the "create_time" field.
func (vmc *VideoMetadataCreate) SetCreateTime(t time.Time) *VideoMetadataCreate {
	vmc.mutation.SetCreateTime(t)
	return vmc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (vmc *VideoMetadataCreate) SetNillableCreateTime(t *time.Time) *VideoMetadataCreate {
	if t != nil {
		vmc.SetCreateTime(*t)
	}
	return vmc
}

// SetUpdateTime sets the "update_time" field.
func (vmc *VideoMetadataCreate) SetUpdateTime(t time.Time) *VideoMetadataCreate {
	vmc.mutation.SetUpdateTime(t)
	return vmc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (vmc *VideoMetadataCreate) SetNillableUpdateTime(t *time.Time) *VideoMetadataCreate {
	if t != nil {
		vmc.SetUpdateTime(*t)
	}
	return vmc
}

// Mutation returns the VideoMetadataMutation object of the builder.
func (vmc *VideoMetadataCreate) Mutation() *VideoMetadataMutation {
	return vmc.mutation
}

// Save creates the VideoMetadata in the database.
func (vmc *VideoMetadataCreate) Save(ctx context.Context) (*VideoMetadata, error) {
	vmc.defaults()
	return withHooks(ctx, vmc.sqlSave, vmc.mutation, vmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vmc *VideoMetadataCreate) SaveX(ctx context.Context) *VideoMetadata {
	v, err := vmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vmc *VideoMetadataCreate) Exec(ctx context.Context) error {
	_, err := vmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmc *VideoMetadataCreate) ExecX(ctx context.Context) {
	if err := vmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vmc *VideoMetadataCreate) defaults() {
	if _, ok := vmc.mutation.PublishAddress(); !ok {
		v := videometadata.DefaultPublishAddress
		vmc.mutation.SetPublishAddress(v)
	}
	if _, ok := vmc.mutation.CreateTime(); !ok {
		v := videometadata.DefaultCreateTime()
		vmc.mutation.SetCreateTime(v)
	}
	if _, ok := vmc.mutation.UpdateTime(); !ok {
		v := videometadata.DefaultUpdateTime()
		vmc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vmc *VideoMetadataCreate) check() error {
	if _, ok := vmc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video_id", err: errors.New(`ent: missing required field "VideoMetadata.video_id"`)}
	}
	if _, ok := vmc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "VideoMetadata.user_id"`)}
	}
	if _, ok := vmc.mutation.CoverURL(); !ok {
		return &ValidationError{Name: "cover_url", err: errors.New(`ent: missing required field "VideoMetadata.cover_url"`)}
	}
	if _, ok := vmc.mutation.VideoURL(); !ok {
		return &ValidationError{Name: "video_url", err: errors.New(`ent: missing required field "VideoMetadata.video_url"`)}
	}
	if _, ok := vmc.mutation.VideoType(); !ok {
		return &ValidationError{Name: "video_type", err: errors.New(`ent: missing required field "VideoMetadata.video_type"`)}
	}
	if _, ok := vmc.mutation.PublishAddress(); !ok {
		return &ValidationError{Name: "publish_address", err: errors.New(`ent: missing required field "VideoMetadata.publish_address"`)}
	}
	if _, ok := vmc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "VideoMetadata.create_time"`)}
	}
	if _, ok := vmc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "VideoMetadata.update_time"`)}
	}
	return nil
}

func (vmc *VideoMetadataCreate) sqlSave(ctx context.Context) (*VideoMetadata, error) {
	if err := vmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vmc.mutation.id = &_node.ID
	vmc.mutation.done = true
	return _node, nil
}

func (vmc *VideoMetadataCreate) createSpec() (*VideoMetadata, *sqlgraph.CreateSpec) {
	var (
		_node = &VideoMetadata{config: vmc.config}
		_spec = sqlgraph.NewCreateSpec(videometadata.Table, sqlgraph.NewFieldSpec(videometadata.FieldID, field.TypeInt))
	)
	if value, ok := vmc.mutation.VideoID(); ok {
		_spec.SetField(videometadata.FieldVideoID, field.TypeInt64, value)
		_node.VideoID = value
	}
	if value, ok := vmc.mutation.UserID(); ok {
		_spec.SetField(videometadata.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := vmc.mutation.CoverURL(); ok {
		_spec.SetField(videometadata.FieldCoverURL, field.TypeString, value)
		_node.CoverURL = value
	}
	if value, ok := vmc.mutation.VideoURL(); ok {
		_spec.SetField(videometadata.FieldVideoURL, field.TypeString, value)
		_node.VideoURL = value
	}
	if value, ok := vmc.mutation.VideoIntro(); ok {
		_spec.SetField(videometadata.FieldVideoIntro, field.TypeString, value)
		_node.VideoIntro = value
	}
	if value, ok := vmc.mutation.VideoType(); ok {
		_spec.SetField(videometadata.FieldVideoType, field.TypeInt64, value)
		_node.VideoType = value
	}
	if value, ok := vmc.mutation.PublishAddress(); ok {
		_spec.SetField(videometadata.FieldPublishAddress, field.TypeInt32, value)
		_node.PublishAddress = value
	}
	if value, ok := vmc.mutation.CreateTime(); ok {
		_spec.SetField(videometadata.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := vmc.mutation.UpdateTime(); ok {
		_spec.SetField(videometadata.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// VideoMetadataCreateBulk is the builder for creating many VideoMetadata entities in bulk.
type VideoMetadataCreateBulk struct {
	config
	err      error
	builders []*VideoMetadataCreate
}

// Save creates the VideoMetadata entities in the database.
func (vmcb *VideoMetadataCreateBulk) Save(ctx context.Context) ([]*VideoMetadata, error) {
	if vmcb.err != nil {
		return nil, vmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vmcb.builders))
	nodes := make([]*VideoMetadata, len(vmcb.builders))
	mutators := make([]Mutator, len(vmcb.builders))
	for i := range vmcb.builders {
		func(i int, root context.Context) {
			builder := vmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoMetadataMutation)
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
					_, err = mutators[i+1].Mutate(root, vmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, vmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vmcb *VideoMetadataCreateBulk) SaveX(ctx context.Context) []*VideoMetadata {
	v, err := vmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vmcb *VideoMetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := vmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmcb *VideoMetadataCreateBulk) ExecX(ctx context.Context) {
	if err := vmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
