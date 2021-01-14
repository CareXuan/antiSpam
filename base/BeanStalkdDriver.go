package base

import (
	"fmt"
	"github.com/beanstalkd/go-beanstalk"
	"time"
)

func GetBeanConn() *beanstalk.Conn {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		fmt.Print(err)
	}
	return c
}

func Put(c *beanstalk.Conn, tubeName string, js string, pri uint32) (uint64, error) {
	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true
	id, err := c.Put([]byte(js), pri, 1, 120*time.Second)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Get(c *beanstalk.Conn, tubeName string) (uint64, []byte, error) {
	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true
	id, body, err := c.Reserve(1000 * time.Second)
	if err != nil {
		return 0, nil, err
	}
	err = c.Delete(id)
	if err != nil {
		return 0, nil, err
	}
	return id, body, err
}
