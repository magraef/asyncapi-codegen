package asyncapiv3

import "github.com/lerenn/asyncapi-codegen/pkg/utils/template"

const (
	// ExternalDocsNameSuffix is the suffix that is added to the name of external docs.
	ExternalDocsNameSuffix = "ExternalDocs"
)

// ExternalDocumentation is a representation of the corresponding asyncapi object filled
// from an asyncapi specification that will be used to generate code.
// Source: https://www.asyncapi.com/docs/reference/specification/v3.0.0#externalDocumentationObject
type ExternalDocumentation struct {
	// --- AsyncAPI fields -----------------------------------------------------

	Description string `json:"description"`
	URL         string `json:"url"`
	Reference   string `json:"$ref"`

	// --- Non AsyncAPI fields -------------------------------------------------

	Name        string                 `json:"-"`
	ReferenceTo *ExternalDocumentation `json:"-"`
}

// Process processes the ExternalDocumentation to make it ready for code generation.
func (doc *ExternalDocumentation) Process(name string, spec Specification) error {
	// Return if empty
	if doc == nil {
		return nil
	}

	// Set name
	doc.Name = template.Namify(name)

	// Add pointer to reference if there is one
	if doc.Reference != "" {
		refTo, err := spec.ReferenceExternalDocumentation(doc.Reference)
		if err != nil {
			return err
		}
		doc.ReferenceTo = refTo
	}

	return nil
}
