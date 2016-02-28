package main

import (
    //"github.com/garyburd/redigo/redis"
    "log"
    "github.com/docker/engine-api/types/container"
    "github.com/docker/engine-api/types/network"
    "github.com/docker/engine-api/client"
)

type Container struct {
    name string
    id string
    started bool
    h *hub
}

func create(cli *client.Client, name string, h *hub) *Container {
    containerOptions := container.Config{
        Hostname:"",
        User:"",
        AttachStdin:true,
        AttachStdout:true,
        AttachStderr:true,
        Tty:true,
        OpenStdin:true,
        StdinOnce:true,
        Image:"ros",
        WorkingDir:"",
    }
    hostOptions := container.HostConfig{
        NetworkMode: "bridge",
    }
    id, err := cli.ContainerCreate(&containerOptions, &hostOptions, &network.NetworkingConfig{}, name)
    if err != nil {
        log.Println(err)
        return nil
    }
    log.Println("Created container:", name)
    return &Container{name: name, id: id.ID, started: false, h: h}
}

func (c *Container) start() {
    err := c.h.cli.ContainerStart(c.id)
    if err != nil {
        return
    }
    c.started = true
    log.Println("Started container:", c.name)
}

func (c *Container) stop() {
    c.h.cli.ContainerStop(c.id, 10)
    c.started = false
    log.Println("Stopped container:", c.name)
}