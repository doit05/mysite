package controllers

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"mysite/helper/apicode"
)

type DockerClientController struct {
	BaseController
}

func (c *DockerClientController) DockerPs() {
	cli, err := client.NewEnvClient()
	if err != nil {
		c.RenderApiJson(apicode.DockerClientError, apicode.Msg(apicode.DockerClientError), err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		c.RenderApiJson(apicode.DockerContainerListError, apicode.Msg(apicode.DockerContainerListError), err)
	}

	c.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), containers)
}
