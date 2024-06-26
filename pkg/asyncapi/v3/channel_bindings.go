package asyncapiv3

import (
	"github.com/lerenn/asyncapi-codegen/pkg/utils/template"
)

// ChannelBindings is a representation of the corresponding asyncapi object filled
// from an asyncapi specification that will be used to generate code.
// Source: https://www.asyncapi.com/docs/reference/specification/v3.0.0#channelBindingsObject
type ChannelBindings struct {
	// --- AsyncAPI fields -----------------------------------------------------

	HTTP         HTTPBinding         `json:"http"`
	WS           WsBinding           `json:"ws"`
	Kafka        KafkaBinding        `json:"kafka"`
	AnyPointMQ   AnyPointMqBinding   `json:"anypointmq"`
	AMQP         AMQPBinding         `json:"amqp"`
	AMQP1        AMQP1Binding        `json:"amqp1"`
	MQTT         MQTTBinding         `json:"mqtt"`
	MQTT5        MQTT5Binding        `json:"mqtt5"`
	NATS         NATSBinding         `json:"nats"`
	JMS          JMSBinding          `json:"jms"`
	SNS          SNSBinding          `json:"sns"`
	Solace       SolaceBinding       `json:"solace"`
	SQS          SQSBinding          `json:"sqs"`
	Stomp        StompBinding        `json:"stomp"`
	Redis        RedisBinding        `json:"redis"`
	Mercure      MercureBinding      `json:"mercure"`
	IBMMQ        IBMMQBinding        `json:"ibmmq"`
	GooglePubSub GooglePubSubBinding `json:"googlepubsub"`
	Pulsar       PulsarBinding       `json:"pulsar"`
	Reference    string              `json:"$ref"`

	// --- Non AsyncAPI fields -------------------------------------------------

	Name        string           `json:"-"`
	ReferenceTo *ChannelBindings `json:"-"`
}

// Process processes the ChannelBindings to make it ready for code generation.
func (chb *ChannelBindings) Process(name string, spec Specification) error { // Return if empty
	// Prevent modification if nil
	if chb == nil {
		return nil
	}

	// Set name
	chb.Name = template.Namify(name)

	// Add pointer to reference if there is one
	if chb.Reference != "" {
		refTo, err := spec.ReferenceChannelBindings(chb.Reference)
		if err != nil {
			return err
		}
		chb.ReferenceTo = refTo
	}

	return nil
}
