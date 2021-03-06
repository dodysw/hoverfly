package templating

import (
	"strings"

	"github.com/SpectoLabs/hoverfly/core/models"
	"github.com/aymerick/raymond"
)

type TemplatingData struct {
	Request Request
	State   map[string]string
}

type Request struct {
	QueryParam map[string][]string
	Path       []string
	Scheme     string
}

type Templator struct {
}

var helpersRegistered = false

func NewTemplator() *Templator {

	if !helpersRegistered {
		raymond.RegisterHelper("iso8601DateTime", iso8601DateTime)
		raymond.RegisterHelper("iso8601DateTimePlusDays", iso8601DateTimePlusDays)
		raymond.RegisterHelper("randomString", randomString)
		raymond.RegisterHelper("randomStringLength", randomStringLength)
		raymond.RegisterHelper("randomBoolean", randomBoolean)
		raymond.RegisterHelper("randomInteger", randomInteger)
		raymond.RegisterHelper("randomIntegerRange", randomIntegerRange)
		raymond.RegisterHelper("randomFloat", randomFloat)
		raymond.RegisterHelper("randomFloatRange", randomFloatRange)
		raymond.RegisterHelper("randomEmail", randomEmail)
		raymond.RegisterHelper("randomIPv4", randomIPv4)
		raymond.RegisterHelper("randomIPv6", randomIPv6)
		raymond.RegisterHelper("randomUuid", randomUuid)

		helpersRegistered = true
	}

	return &Templator{}
}

func (*Templator) ApplyTemplate(requestDetails *models.RequestDetails, state map[string]string, responseBody string) (string, error) {

	t := NewTemplatingDataFromRequest(requestDetails, state)

	if rendered, err := raymond.Render(responseBody, t); err == nil {
		responseBody = rendered
		return responseBody, nil
	} else {
		return "", err
	}
}

func NewTemplatingDataFromRequest(requestDetails *models.RequestDetails, state map[string]string) *TemplatingData {
	return &TemplatingData{
		Request: Request{
			Path:       strings.Split(requestDetails.Path, "/")[1:],
			QueryParam: requestDetails.Query,
			Scheme:     requestDetails.Scheme,
		},
		State: state,
	}

}
