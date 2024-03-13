package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"net/http"
)

func main() {
	app := gin.Default()
	app.Static("/static", "./consul-gin")
	app.GET("/hello", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"hello": "world"})
	})
	app.GET("/health", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"health": "ok"})
	})
	app.GET("/detail", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": "SUCCESS",
			"msg":  "成功",
			"data": map[string]any{
				"ip":         "127.0.0.1",
				"country":    "中国",
				"country_id": "CN",
				"region":     "广东省",
			},
		})
	})
	RegisterGateway(&NacosCfg{
		NacosIP:     "192.168.3.237",
		NacosPort:   31006,
		ServiceName: "demo" + ".http",
		ClientIP:    "192.168.3.4",
		ClientPort:  8888,
		Weight:      100,
	})
	app.Run("0.0.0.0:8888")
}

type NacosCfg struct {
	NacosIP   string
	NacosPort uint64

	//
	ServiceName string
	ClientIP    string
	ClientPort  uint64
	Weight      float64
}

// RegisterGateway 注册网关
func RegisterGateway(nacosCfg *NacosCfg) {
	var (
		err error
	)
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(nacosCfg.NacosIP, nacosCfg.NacosPort),
	}
	var (
		nameingClient naming_client.INamingClient
		success       bool
	)
	if nameingClient, err = clients.NewNamingClient(vo.NacosClientParam{
		ServerConfigs: sc},
	); err != nil {
		log.Fatalln(err)
	}

	if success, err = nameingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          nacosCfg.ClientIP,
		Port:        nacosCfg.ClientPort,
		ServiceName: nacosCfg.ServiceName,
		Metadata: map[string]string{
			"kind": "http",
		},
		Weight:      nacosCfg.Weight,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		ClusterName: "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP

	}); err != nil {
		log.Fatalln(err)
	}
	if !success {
		log.Fatalln("gateway register failed")
	}
	log.Println("gateway register success")
}
