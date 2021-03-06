package server

import (
	"fmt"
	"log"
	"net"
	"os"
	rt "runtime"
	"strings"
	"sync"
	"time"

	"github.com/appcelerator/amp/api/rpc/function"
	"github.com/appcelerator/amp/api/rpc/logs"
	"github.com/appcelerator/amp/api/rpc/oauth"
	"github.com/appcelerator/amp/api/rpc/service"
	"github.com/appcelerator/amp/api/rpc/stack"
	"github.com/appcelerator/amp/api/rpc/stats"
	"github.com/appcelerator/amp/api/rpc/storage"
	"github.com/appcelerator/amp/api/rpc/topic"
	"github.com/appcelerator/amp/api/rpc/version"
	"github.com/appcelerator/amp/api/runtime"
	"github.com/appcelerator/amp/data/influx"
	"github.com/appcelerator/amp/data/storage/etcd"
	"github.com/appcelerator/amp/pkg/config"
	"github.com/docker/docker/client"
	"google.golang.org/grpc"
)

const (
	defaultTimeOut = 30 * time.Second
)

type (
	clientInitializer  func(Config) error
	serviceInitializer func(Config, *grpc.Server)
)

// Client initializers open connections to required backend services
// Clients are stored as members of runtime
var clientInitializers = []clientInitializer{
//initEtcd,
//initElasticsearch,
//initNats,
//initInfluxDB,
//initDocker,
}

// Service initializers register the services with the grpc server
var serviceInitializers = []serviceInitializer{
	registerVersionServer,
	//registerStorageServer,
	//registerLogsServer,
	//registerStatsServer,
	//registerServiceServer,
	//registerStackServiceServer,
	//registerTopicServer,
	//registerFunctionServer,
	//registerGithubServer,
}

// Start starts the amplifier server
func Start(c Config) {
	// initialize clients
	initClients(c)

	// register services
	s := grpc.NewServer()
	registerServices(c, s)

	// start listening
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Unable to listen on %s: %v\n", c.Port[1:], err)
	}
	log.Println("Listening on port:", c.Port[1:])
	log.Fatalln(s.Serve(lis))
}

func initClients(config Config) {
	// ensure all initialization code fails fast on errors; there is no point in
	// attempting to continue in a degraded state if there are problems at start up

	var wg sync.WaitGroup
	for _, f := range clientInitializers {
		wg.Add(1)
		go func(f clientInitializer) {
			defer wg.Done()
			if err := f(config); err != nil {
				log.Fatalln(err)
			}
		}(f)
	}

	// Wait for all inits to complete.
	wg.Wait()
}

func initEtcd(config Config) error {
	log.Println("Connecting to etcd at", strings.Join(config.EtcdEndpoints, ","))
	runtime.Store = etcd.New(config.EtcdEndpoints, "amp")
	if err := runtime.Store.Connect(defaultTimeOut); err != nil {
		return fmt.Errorf("unable to connect to etcd at %s: %v", config.EtcdEndpoints, err)
	}
	log.Println("Connected to etcd at", strings.Join(runtime.Store.Endpoints(), ","))
	return nil
}

func initElasticsearch(config Config) error {
	log.Println("Connecting to elasticsearch at", config.ElasticsearchURL)
	if err := runtime.Elasticsearch.Connect(config.ElasticsearchURL, defaultTimeOut); err != nil {
		return fmt.Errorf("unable to connect to elasticsearch at %s: %v", config.ElasticsearchURL, err)
	}
	log.Println("Connected to elasticsearch at", config.ElasticsearchURL)
	return nil
}

func initInfluxDB(config Config) error {
	log.Println("Connecting to InfluxDB at", config.InfluxURL)
	runtime.Influx = influx.New(config.InfluxURL, "telegraf", "", "")
	if err := runtime.Influx.Connect(defaultTimeOut); err != nil {
		return fmt.Errorf("unable to connect to influxDB at %s: %v", config.InfluxURL, err)
	}
	log.Println("Connected to influxDB at", config.InfluxURL)
	return nil
}

func initNats(config Config) error {
	// NATS
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("unable to get hostname: %v", err)
	}
	if runtime.NatsStreaming.Connect(config.NatsURL, amp.NatsClusterID, os.Args[0]+"-"+hostname, amp.DefaultTimeout) != nil {
		return err
	}
	return nil
}

func initDocker(config Config) error {
	log.Printf("Connecting to Docker API at %s version API: %s\n", config.DockerURL, config.DockerVersion)
	defaultHeaders := map[string]string{"User-Agent": "amplifier-1.0"}
	var err error
	runtime.Docker, err = client.NewClient(config.DockerURL, config.DockerVersion, nil, defaultHeaders)
	if err != nil {
		return fmt.Errorf("unable to connect to Docker at %s: %v", config.DockerURL, err)
	}
	log.Println("Connected to Docker API at", config.DockerURL)
	return nil
}

func registerServices(c Config, s *grpc.Server) {
	var wg sync.WaitGroup
	for _, f := range serviceInitializers {
		wg.Add(1)
		go func(f serviceInitializer) {
			defer wg.Done()
			f(c, s)
		}(f)
	}

	// Wait for all service registrations to complete.
	wg.Wait()
}

func registerVersionServer(c Config, s *grpc.Server) {
	version.RegisterVersionServer(s, &version.Server{
		Version:   c.Version,
		Port:      c.Port,
		GoVersion: rt.Version(),
		Os:        rt.GOOS,
		Arch:      rt.GOARCH,
	})
}

func registerLogsServer(c Config, s *grpc.Server) {
	logs.RegisterLogsServer(s, &logs.Server{
		Es:            &runtime.Elasticsearch,
		Store:         runtime.Store,
		NatsStreaming: runtime.NatsStreaming,
	})
}

func registerStorageServer(c Config, s *grpc.Server) {
	storage.RegisterStorageServer(s, &storage.Server{
		Store: runtime.Store,
	})
}

func registerStatsServer(c Config, s *grpc.Server) {
	stats.RegisterStatsServer(s, &stats.Stats{
		Influx: runtime.Influx,
	})
}

func registerServiceServer(c Config, s *grpc.Server) {
	service.RegisterServiceServer(s, &service.Service{
		Docker: runtime.Docker,
	})
}

func registerStackServiceServer(c Config, s *grpc.Server) {
	stack.RegisterStackServiceServer(s, stack.NewServer(
		runtime.Store,
		runtime.Docker,
	))
}

func registerTopicServer(c Config, s *grpc.Server) {
	topic.RegisterTopicServer(s, &topic.Server{
		Store:         runtime.Store,
		NatsStreaming: runtime.NatsStreaming,
	})
}

func registerFunctionServer(c Config, s *grpc.Server) {
	function.RegisterFunctionServer(s, &function.Server{
		Store:         runtime.Store,
		NatsStreaming: runtime.NatsStreaming,
	})
}

func registerGithubServer(c Config, s *grpc.Server) {
	oauth.RegisterGithubServer(s, &oauth.Oauth{
		Store:        runtime.Store,
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
	})
}
