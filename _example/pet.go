// Code generated by entc, DO NOT EDIT.

package _example

import (
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/masseelch/elk/_example/owner"
	"github.com/masseelch/elk/_example/pet"
	"github.com/masseelch/elk/_example/schema"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `groups:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" groups:"pet:list"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty" groups:"pet:list"`
	// Color holds the value of the "color" field.
	Color schema.Color `json:"color,omitempty" groups:"pet:list"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges      PetEdges `json:"edges" groups:"pet:list"`
	owner_pets *int
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Owner `json:"owner" groups:"pet:list"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) OwnerOrErr() (*Owner, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: owner.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullInt64{},  // age
		&sql.NullInt64{},  // color
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Pet) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // owner_pets
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(values ...interface{}) error {
	if m, n := len(values), len(pet.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pe.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		pe.Name = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[1])
	} else if value.Valid {
		pe.Age = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field color", values[2])
	} else if value.Valid {
		pe.Color = schema.Color(value.Int64)
	}
	values = values[3:]
	if len(values) == len(pet.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field owner_pets", value)
		} else if value.Valid {
			pe.owner_pets = new(int)
			*pe.owner_pets = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Pet.
func (pe *Pet) QueryOwner() *OwnerQuery {
	return (&PetClient{config: pe.config}).QueryOwner(pe)
}

// Update returns a builder for updating this Pet.
// Note that, you need to call Pet.Unwrap() before calling this method, if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("_example: Pet is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// Get rid of the fmt.Stringer implementation since it breaks liip/sheriff.
// This lines have to be here since template/text does skip empty templates.

// Pets is a parsable slice of Pet.
type Pets []*Pet

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
