package db

import (
	"os"
	"sync"

	"github.com/msanvarov/gin-rest-prisma-boilerplate/prisma-client"
)

var (
	client *prisma.Client
	once   sync.Once
)

// DB method is responsible for establishing a prisma connection and returning a pointer to it.
func DB() *prisma.Client {
	once.Do(func() {
		if prismaEndpoint := os.Getenv("PRISMA_ENDPOINT"); prismaEndpoint != "" {
			client = prisma.New(&prisma.Options{
				Endpoint: prismaEndpoint,
			})
		} else {
			client = prisma.New(nil)
		}
	})
	return client
}
