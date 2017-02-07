package controllers

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"mysite/helper/apicode"
)

func (c *DockerClientController) DockerSwarm() {
	cli, err := client.NewEnvClient()
	// cli, err := client.NewClient("tcp://www.bigolive.site:5555", "v1.22", nil, nil)
	if err != nil {
		c.RenderApiJson(apicode.DockerClientError, apicode.Msg(apicode.DockerClientError), err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		c.RenderApiJson(apicode.DockerContainerListError, apicode.Msg(apicode.DockerContainerListError), err)
	}

	c.RenderApiJsonSlice(apicode.Success, apicode.Msg(apicode.Success), containers)
}
