// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go-test/ent/migrate"

	"go-test/ent/middleware"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Middleware is the client for interacting with the Middleware builders.
	Middleware *MiddlewareClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Middleware = NewMiddlewareClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Middleware: NewMiddlewareClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Middleware: NewMiddlewareClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Middleware.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Middleware.Use(hooks...)
}

// MiddlewareClient is a client for the Middleware schema.
type MiddlewareClient struct {
	config
}

// NewMiddlewareClient returns a client for the Middleware from the given config.
func NewMiddlewareClient(c config) *MiddlewareClient {
	return &MiddlewareClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `middleware.Hooks(f(g(h())))`.
func (c *MiddlewareClient) Use(hooks ...Hook) {
	c.hooks.Middleware = append(c.hooks.Middleware, hooks...)
}

// Create returns a builder for creating a Middleware entity.
func (c *MiddlewareClient) Create() *MiddlewareCreate {
	mutation := newMiddlewareMutation(c.config, OpCreate)
	return &MiddlewareCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Middleware entities.
func (c *MiddlewareClient) CreateBulk(builders ...*MiddlewareCreate) *MiddlewareCreateBulk {
	return &MiddlewareCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Middleware.
func (c *MiddlewareClient) Update() *MiddlewareUpdate {
	mutation := newMiddlewareMutation(c.config, OpUpdate)
	return &MiddlewareUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MiddlewareClient) UpdateOne(m *Middleware) *MiddlewareUpdateOne {
	mutation := newMiddlewareMutation(c.config, OpUpdateOne, withMiddleware(m))
	return &MiddlewareUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MiddlewareClient) UpdateOneID(id int) *MiddlewareUpdateOne {
	mutation := newMiddlewareMutation(c.config, OpUpdateOne, withMiddlewareID(id))
	return &MiddlewareUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Middleware.
func (c *MiddlewareClient) Delete() *MiddlewareDelete {
	mutation := newMiddlewareMutation(c.config, OpDelete)
	return &MiddlewareDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MiddlewareClient) DeleteOne(m *Middleware) *MiddlewareDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MiddlewareClient) DeleteOneID(id int) *MiddlewareDeleteOne {
	builder := c.Delete().Where(middleware.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MiddlewareDeleteOne{builder}
}

// Query returns a query builder for Middleware.
func (c *MiddlewareClient) Query() *MiddlewareQuery {
	return &MiddlewareQuery{
		config: c.config,
	}
}

// Get returns a Middleware entity by its id.
func (c *MiddlewareClient) Get(ctx context.Context, id int) (*Middleware, error) {
	return c.Query().Where(middleware.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MiddlewareClient) GetX(ctx context.Context, id int) *Middleware {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MiddlewareClient) Hooks() []Hook {
	return c.hooks.Middleware
}
