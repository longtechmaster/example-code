# example-code : REGISTRY _CONSUL_OTHER_SERVER

```
cd example-code/registry-consul-other-server/greeter/srv/
```

### Customize func newService 

```
service := micro.NewService(
		micro.Name("greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(registry.NewRegistry(registry.Addrs("192.168.1.133:8500"))),
	)
```
