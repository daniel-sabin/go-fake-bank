package server_test

import (
	"engineecore/demobank-server/infra/repository"
	"engineecore/demobank-server/server"
	"testing"

	"goyave.dev/goyave/v4"
)

type DumbStore struct {
	Exist bool
}

func (i *DumbStore) Save(key string) {
}

func (i *DumbStore) Exists(key string) bool {
	return i.Exist
}

func router(dumbStore *DumbStore) func(router *goyave.Router) {
	accountsStore := repository.NewInMemoryAccountsStore()
	return server.RegisterRoutes(dumbStore, accountsStore, nil)
}

type ServerTestSuite struct {
	goyave.TestSuite
}

func TestServer(t *testing.T) {
	goyave.RunTest(t, new(ServerTestSuite))
}

func (suite *ServerTestSuite) TestHealthCheck() {
	suite.RunServer(router(nil), func() {
		resp, err := suite.Get("/health", nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			suite.Equal("ok", string(suite.GetBody(resp)))
		}
	})
}

func (suite *ServerTestSuite) TestSwagger() {
	suite.RunServer(router(nil), func() {
		resp, err := suite.Get("/swaggerui/", nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			suite.Contains(string(suite.GetBody(resp)), "DOCTYPE")
		}
	})
}

func (suite *ServerTestSuite) TestForbiddenApplication() {
	suite.RunServer(router(&DumbStore{Exist: false}), func() {
		resp, err := suite.Get("/applications", map[string]string{"x-api-key": "fake-key"})
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(401, resp.StatusCode)
		}
	})
}

func (suite *ServerTestSuite) TestAllowedApplication() {
	suite.RunServer(router(&DumbStore{Exist: true}), func() {
		resp, err := suite.Get("/applications", map[string]string{"x-api-key": "fake-key"})
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
		}
	})
}

func (suite *ServerTestSuite) TestAccounts() {
	suite.RunServer(router(nil), func() {
		resp, err := suite.Get("/accounts", nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			suite.Contains(string(suite.GetBody(resp)), "{\"accounts\":[{\"acc_number\":\"0000001\",\"amount\":50,\"currency\":\"EUR\"}")
		}
	})
}

func (suite *ServerTestSuite) TestTransactions() {
	suite.RunServer(router(nil), func() {
		resp, err := suite.Get("/accounts/1/transactions", nil)
		suite.Nil(err)
		suite.NotNil(resp)
		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(200, resp.StatusCode)
			suite.Equal("1", string(suite.GetBody(resp)))
		}
	})
}
