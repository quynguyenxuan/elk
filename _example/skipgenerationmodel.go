// Code generated by entc, DO NOT EDIT.

package _example

import (
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/masseelch/elk/_example/skipgenerationmodel"
)

// SkipGenerationModel is the model entity for the SkipGenerationModel schema.
type SkipGenerationModel struct {
	config `groups:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" groups:"SkipGenerationModel:list"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SkipGenerationModel) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SkipGenerationModel fields.
func (sgm *SkipGenerationModel) assignValues(values ...interface{}) error {
	if m, n := len(values), len(skipgenerationmodel.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	sgm.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		sgm.Name = value.String
	}
	return nil
}

// Update returns a builder for updating this SkipGenerationModel.
// Note that, you need to call SkipGenerationModel.Unwrap() before calling this method, if this SkipGenerationModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (sgm *SkipGenerationModel) Update() *SkipGenerationModelUpdateOne {
	return (&SkipGenerationModelClient{config: sgm.config}).UpdateOne(sgm)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (sgm *SkipGenerationModel) Unwrap() *SkipGenerationModel {
	tx, ok := sgm.config.driver.(*txDriver)
	if !ok {
		panic("_example: SkipGenerationModel is not a transactional entity")
	}
	sgm.config.driver = tx.drv
	return sgm
}

// Get rid of the fmt.Stringer implementation since it breaks liip/sheriff.
// This lines have to be here since template/text does skip empty templates.

// SkipGenerationModels is a parsable slice of SkipGenerationModel.
type SkipGenerationModels []*SkipGenerationModel

func (sgm SkipGenerationModels) config(cfg config) {
	for _i := range sgm {
		sgm[_i].config = cfg
	}
}
