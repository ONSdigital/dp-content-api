package steps

import (
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
)

func (c *Component) RegisterSteps(ctx *godog.ScenarioContext) {
	c.apiFeature.RegisterSteps(ctx)

	ctx.Step(`^I should receive a hello-world response$`, c.iShouldReceiveAHelloworldResponse)
}

func (c *Component) iShouldReceiveAHelloworldResponse() error {
	responseBody := c.apiFeature.HttpResponse.Body
	body, _ := ioutil.ReadAll(responseBody)

	assert.Equal(c, `{"message":"Hello, World!"}`, strings.TrimSpace(string(body)))

	return c.StepError()
}
