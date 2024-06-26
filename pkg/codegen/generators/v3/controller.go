package generatorv3

import (
	"bytes"

	asyncapi "github.com/lerenn/asyncapi-codegen/pkg/asyncapi/v3"
	"github.com/lerenn/asyncapi-codegen/pkg/codegen/generators"
)

// ControllerGenerator is a code generator for controllers that will turn an
// asyncapi specification into controller golang code.
type ControllerGenerator struct {
	Operations ActionOperations
	Prefix     string
	Version    string
}

// NewControllerGenerator will create a new controller code generator.
func NewControllerGenerator(side generators.Side, spec asyncapi.Specification) ControllerGenerator {
	var gen ControllerGenerator

	// Generate receive send operations
	gen.Operations = NewActionOperations(side, spec)

	// Set generation name
	if side == generators.SideIsApplication {
		gen.Prefix = "App"
	} else {
		gen.Prefix = "User"
	}

	// Set version
	gen.Version = spec.Info.Version

	return gen
}

func shouldControllerRespondToReply(side generators.Side, op *asyncapi.Operation) bool {
	if op.Reply == nil || op.Reply.Channel == nil {
		return false
	}

	switch {
	case side == generators.SideIsApplication && op.Action.IsReceive():
		return true
	case side == generators.SideIsUser && op.Action.IsSend():
		return true
	default:
		return false
	}
}

func isControllerReceiveOperation(side generators.Side, op *asyncapi.Operation) bool {
	switch {
	case side == generators.SideIsApplication && op.Action.IsReceive():
		return true
	case side == generators.SideIsUser && op.Action.IsSend():
		return true
	default:
		return false
	}
}

func isControllerSendOperation(side generators.Side, op *asyncapi.Operation) bool {
	switch {
	case side == generators.SideIsApplication && op.Action.IsSend():
		return true
	case side == generators.SideIsUser && op.Action.IsReceive():
		return true
	default:
		return false
	}
}

// Generate will generate the controller code.
func (asg ControllerGenerator) Generate() (string, error) {
	tmplt, err := loadTemplate(
		controllerTemplatePath,
		schemaDefinitionTemplatePath,
		schemaNameTemplatePath,
		messageTemplatePath,
	)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := tmplt.Execute(buf, asg); err != nil {
		return "", err
	}

	return buf.String(), nil
}
