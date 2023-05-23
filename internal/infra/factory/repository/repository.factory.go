package repository

import (
	"fmt"
	"github.com/dylanbatar/simple-login-system/internal/infra/adapters/mongoRepository"
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/ports"
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/repository/enum"
	"github.com/dylanbatar/simple-login-system/internal/infra/pkg/mongodb"
)

func NewRepositoryFactory(provider enum.RepositoryProvider) (ports.IRepository, error) {
	switch provider {
	case enum.MongoProvider:
		conn, err := mongodb.NewConnection()

		if err != nil {
			return nil, fmt.Errorf("error setting mongo provider %s", err.Error())
		}

		database := conn.Database("testing")
		return mongoRepository.NewMongoAdapter(database), nil
	}

	return nil, fmt.Errorf("provider not found")
}
