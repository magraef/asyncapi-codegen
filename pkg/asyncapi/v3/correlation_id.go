package asyncapiv3

import "github.com/lerenn/asyncapi-codegen/pkg/utils/template"

// CorrelationID is a representation of the corresponding asyncapi object filled
// from an asyncapi specification that will be used to generate code.
// Source: https://www.asyncapi.com/docs/reference/specification/v3.0.0#correlationIdObject
type CorrelationID struct {
	// --- AsyncAPI fields -----------------------------------------------------

	Description string `json:"description"`
	Location    string `json:"location"`
	Reference   string `json:"$ref"`

	// --- Non AsyncAPI fields -------------------------------------------------

	Name        string   `json:"-"`
	ReferenceTo *Channel `json:"-"`
}

// Process processes the CorrelationID to make it ready for code generation.
func (c *CorrelationID) Process(path string, spec Specification) error {
	// Prevent modification if nil
	if c == nil {
		return nil
	}

	// Set name
	c.Name = template.Namify(path)

	// Add pointer to reference if there is one
	if c.Reference != "" {
		refTo, err := spec.ReferenceChannel(c.Reference)
		if err != nil {
			return err
		}
		c.ReferenceTo = refTo
	}

	return nil
}

// Exists checks that the correlation exists (and that the location is set).
func (c *CorrelationID) Exists() bool {
	return c != nil && c.Location != ""
}
