// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"mai.today/core/entity/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"mai.today/core/entity/base"
	"mai.today/core/entity/baseinfo"
	"mai.today/core/entity/basemember"
	"mai.today/core/entity/basenavstate"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Base is the client for interacting with the Base builders.
	Base *BaseClient
	// BaseInfo is the client for interacting with the BaseInfo builders.
	BaseInfo *BaseInfoClient
	// BaseMember is the client for interacting with the BaseMember builders.
	BaseMember *BaseMemberClient
	// BaseNavState is the client for interacting with the BaseNavState builders.
	BaseNavState *BaseNavStateClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Base = NewBaseClient(c.config)
	c.BaseInfo = NewBaseInfoClient(c.config)
	c.BaseMember = NewBaseMemberClient(c.config)
	c.BaseNavState = NewBaseNavStateClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("entity: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("entity: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Base:         NewBaseClient(cfg),
		BaseInfo:     NewBaseInfoClient(cfg),
		BaseMember:   NewBaseMemberClient(cfg),
		BaseNavState: NewBaseNavStateClient(cfg),
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
		ctx:          ctx,
		config:       cfg,
		Base:         NewBaseClient(cfg),
		BaseInfo:     NewBaseInfoClient(cfg),
		BaseMember:   NewBaseMemberClient(cfg),
		BaseNavState: NewBaseNavStateClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Base.
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
	c.Base.Use(hooks...)
	c.BaseInfo.Use(hooks...)
	c.BaseMember.Use(hooks...)
	c.BaseNavState.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Base.Intercept(interceptors...)
	c.BaseInfo.Intercept(interceptors...)
	c.BaseMember.Intercept(interceptors...)
	c.BaseNavState.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BaseMutation:
		return c.Base.mutate(ctx, m)
	case *BaseInfoMutation:
		return c.BaseInfo.mutate(ctx, m)
	case *BaseMemberMutation:
		return c.BaseMember.mutate(ctx, m)
	case *BaseNavStateMutation:
		return c.BaseNavState.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("entity: unknown mutation type %T", m)
	}
}

// BaseClient is a client for the Base schema.
type BaseClient struct {
	config
}

// NewBaseClient returns a client for the Base from the given config.
func NewBaseClient(c config) *BaseClient {
	return &BaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `base.Hooks(f(g(h())))`.
func (c *BaseClient) Use(hooks ...Hook) {
	c.hooks.Base = append(c.hooks.Base, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `base.Intercept(f(g(h())))`.
func (c *BaseClient) Intercept(interceptors ...Interceptor) {
	c.inters.Base = append(c.inters.Base, interceptors...)
}

// Create returns a builder for creating a Base entity.
func (c *BaseClient) Create() *BaseCreate {
	mutation := newBaseMutation(c.config, OpCreate)
	return &BaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Base entities.
func (c *BaseClient) CreateBulk(builders ...*BaseCreate) *BaseCreateBulk {
	return &BaseCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BaseClient) MapCreateBulk(slice any, setFunc func(*BaseCreate, int)) *BaseCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BaseCreateBulk{err: fmt.Errorf("calling to BaseClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BaseCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BaseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Base.
func (c *BaseClient) Update() *BaseUpdate {
	mutation := newBaseMutation(c.config, OpUpdate)
	return &BaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaseClient) UpdateOne(b *Base) *BaseUpdateOne {
	mutation := newBaseMutation(c.config, OpUpdateOne, withBase(b))
	return &BaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaseClient) UpdateOneID(id string) *BaseUpdateOne {
	mutation := newBaseMutation(c.config, OpUpdateOne, withBaseID(id))
	return &BaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Base.
func (c *BaseClient) Delete() *BaseDelete {
	mutation := newBaseMutation(c.config, OpDelete)
	return &BaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BaseClient) DeleteOne(b *Base) *BaseDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BaseClient) DeleteOneID(id string) *BaseDeleteOne {
	builder := c.Delete().Where(base.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaseDeleteOne{builder}
}

// Query returns a query builder for Base.
func (c *BaseClient) Query() *BaseQuery {
	return &BaseQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBase},
		inters: c.Interceptors(),
	}
}

// Get returns a Base entity by its id.
func (c *BaseClient) Get(ctx context.Context, id string) (*Base, error) {
	return c.Query().Where(base.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaseClient) GetX(ctx context.Context, id string) *Base {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BaseClient) Hooks() []Hook {
	return c.hooks.Base
}

// Interceptors returns the client interceptors.
func (c *BaseClient) Interceptors() []Interceptor {
	return c.inters.Base
}

func (c *BaseClient) mutate(ctx context.Context, m *BaseMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BaseCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BaseUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BaseDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown Base mutation op: %q", m.Op())
	}
}

// BaseInfoClient is a client for the BaseInfo schema.
type BaseInfoClient struct {
	config
}

// NewBaseInfoClient returns a client for the BaseInfo from the given config.
func NewBaseInfoClient(c config) *BaseInfoClient {
	return &BaseInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `baseinfo.Hooks(f(g(h())))`.
func (c *BaseInfoClient) Use(hooks ...Hook) {
	c.hooks.BaseInfo = append(c.hooks.BaseInfo, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `baseinfo.Intercept(f(g(h())))`.
func (c *BaseInfoClient) Intercept(interceptors ...Interceptor) {
	c.inters.BaseInfo = append(c.inters.BaseInfo, interceptors...)
}

// Create returns a builder for creating a BaseInfo entity.
func (c *BaseInfoClient) Create() *BaseInfoCreate {
	mutation := newBaseInfoMutation(c.config, OpCreate)
	return &BaseInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaseInfo entities.
func (c *BaseInfoClient) CreateBulk(builders ...*BaseInfoCreate) *BaseInfoCreateBulk {
	return &BaseInfoCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BaseInfoClient) MapCreateBulk(slice any, setFunc func(*BaseInfoCreate, int)) *BaseInfoCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BaseInfoCreateBulk{err: fmt.Errorf("calling to BaseInfoClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BaseInfoCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BaseInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaseInfo.
func (c *BaseInfoClient) Update() *BaseInfoUpdate {
	mutation := newBaseInfoMutation(c.config, OpUpdate)
	return &BaseInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaseInfoClient) UpdateOne(bi *BaseInfo) *BaseInfoUpdateOne {
	mutation := newBaseInfoMutation(c.config, OpUpdateOne, withBaseInfo(bi))
	return &BaseInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaseInfoClient) UpdateOneID(id string) *BaseInfoUpdateOne {
	mutation := newBaseInfoMutation(c.config, OpUpdateOne, withBaseInfoID(id))
	return &BaseInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaseInfo.
func (c *BaseInfoClient) Delete() *BaseInfoDelete {
	mutation := newBaseInfoMutation(c.config, OpDelete)
	return &BaseInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BaseInfoClient) DeleteOne(bi *BaseInfo) *BaseInfoDeleteOne {
	return c.DeleteOneID(bi.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BaseInfoClient) DeleteOneID(id string) *BaseInfoDeleteOne {
	builder := c.Delete().Where(baseinfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaseInfoDeleteOne{builder}
}

// Query returns a query builder for BaseInfo.
func (c *BaseInfoClient) Query() *BaseInfoQuery {
	return &BaseInfoQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBaseInfo},
		inters: c.Interceptors(),
	}
}

// Get returns a BaseInfo entity by its id.
func (c *BaseInfoClient) Get(ctx context.Context, id string) (*BaseInfo, error) {
	return c.Query().Where(baseinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaseInfoClient) GetX(ctx context.Context, id string) *BaseInfo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BaseInfoClient) Hooks() []Hook {
	return c.hooks.BaseInfo
}

// Interceptors returns the client interceptors.
func (c *BaseInfoClient) Interceptors() []Interceptor {
	return c.inters.BaseInfo
}

func (c *BaseInfoClient) mutate(ctx context.Context, m *BaseInfoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BaseInfoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BaseInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BaseInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BaseInfoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown BaseInfo mutation op: %q", m.Op())
	}
}

// BaseMemberClient is a client for the BaseMember schema.
type BaseMemberClient struct {
	config
}

// NewBaseMemberClient returns a client for the BaseMember from the given config.
func NewBaseMemberClient(c config) *BaseMemberClient {
	return &BaseMemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `basemember.Hooks(f(g(h())))`.
func (c *BaseMemberClient) Use(hooks ...Hook) {
	c.hooks.BaseMember = append(c.hooks.BaseMember, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `basemember.Intercept(f(g(h())))`.
func (c *BaseMemberClient) Intercept(interceptors ...Interceptor) {
	c.inters.BaseMember = append(c.inters.BaseMember, interceptors...)
}

// Create returns a builder for creating a BaseMember entity.
func (c *BaseMemberClient) Create() *BaseMemberCreate {
	mutation := newBaseMemberMutation(c.config, OpCreate)
	return &BaseMemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaseMember entities.
func (c *BaseMemberClient) CreateBulk(builders ...*BaseMemberCreate) *BaseMemberCreateBulk {
	return &BaseMemberCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BaseMemberClient) MapCreateBulk(slice any, setFunc func(*BaseMemberCreate, int)) *BaseMemberCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BaseMemberCreateBulk{err: fmt.Errorf("calling to BaseMemberClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BaseMemberCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BaseMemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaseMember.
func (c *BaseMemberClient) Update() *BaseMemberUpdate {
	mutation := newBaseMemberMutation(c.config, OpUpdate)
	return &BaseMemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaseMemberClient) UpdateOne(bm *BaseMember) *BaseMemberUpdateOne {
	mutation := newBaseMemberMutation(c.config, OpUpdateOne, withBaseMember(bm))
	return &BaseMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaseMemberClient) UpdateOneID(id string) *BaseMemberUpdateOne {
	mutation := newBaseMemberMutation(c.config, OpUpdateOne, withBaseMemberID(id))
	return &BaseMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaseMember.
func (c *BaseMemberClient) Delete() *BaseMemberDelete {
	mutation := newBaseMemberMutation(c.config, OpDelete)
	return &BaseMemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BaseMemberClient) DeleteOne(bm *BaseMember) *BaseMemberDeleteOne {
	return c.DeleteOneID(bm.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BaseMemberClient) DeleteOneID(id string) *BaseMemberDeleteOne {
	builder := c.Delete().Where(basemember.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaseMemberDeleteOne{builder}
}

// Query returns a query builder for BaseMember.
func (c *BaseMemberClient) Query() *BaseMemberQuery {
	return &BaseMemberQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBaseMember},
		inters: c.Interceptors(),
	}
}

// Get returns a BaseMember entity by its id.
func (c *BaseMemberClient) Get(ctx context.Context, id string) (*BaseMember, error) {
	return c.Query().Where(basemember.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaseMemberClient) GetX(ctx context.Context, id string) *BaseMember {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BaseMemberClient) Hooks() []Hook {
	return c.hooks.BaseMember
}

// Interceptors returns the client interceptors.
func (c *BaseMemberClient) Interceptors() []Interceptor {
	return c.inters.BaseMember
}

func (c *BaseMemberClient) mutate(ctx context.Context, m *BaseMemberMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BaseMemberCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BaseMemberUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BaseMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BaseMemberDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown BaseMember mutation op: %q", m.Op())
	}
}

// BaseNavStateClient is a client for the BaseNavState schema.
type BaseNavStateClient struct {
	config
}

// NewBaseNavStateClient returns a client for the BaseNavState from the given config.
func NewBaseNavStateClient(c config) *BaseNavStateClient {
	return &BaseNavStateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `basenavstate.Hooks(f(g(h())))`.
func (c *BaseNavStateClient) Use(hooks ...Hook) {
	c.hooks.BaseNavState = append(c.hooks.BaseNavState, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `basenavstate.Intercept(f(g(h())))`.
func (c *BaseNavStateClient) Intercept(interceptors ...Interceptor) {
	c.inters.BaseNavState = append(c.inters.BaseNavState, interceptors...)
}

// Create returns a builder for creating a BaseNavState entity.
func (c *BaseNavStateClient) Create() *BaseNavStateCreate {
	mutation := newBaseNavStateMutation(c.config, OpCreate)
	return &BaseNavStateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BaseNavState entities.
func (c *BaseNavStateClient) CreateBulk(builders ...*BaseNavStateCreate) *BaseNavStateCreateBulk {
	return &BaseNavStateCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BaseNavStateClient) MapCreateBulk(slice any, setFunc func(*BaseNavStateCreate, int)) *BaseNavStateCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BaseNavStateCreateBulk{err: fmt.Errorf("calling to BaseNavStateClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BaseNavStateCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BaseNavStateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BaseNavState.
func (c *BaseNavStateClient) Update() *BaseNavStateUpdate {
	mutation := newBaseNavStateMutation(c.config, OpUpdate)
	return &BaseNavStateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BaseNavStateClient) UpdateOne(bns *BaseNavState) *BaseNavStateUpdateOne {
	mutation := newBaseNavStateMutation(c.config, OpUpdateOne, withBaseNavState(bns))
	return &BaseNavStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BaseNavStateClient) UpdateOneID(id string) *BaseNavStateUpdateOne {
	mutation := newBaseNavStateMutation(c.config, OpUpdateOne, withBaseNavStateID(id))
	return &BaseNavStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BaseNavState.
func (c *BaseNavStateClient) Delete() *BaseNavStateDelete {
	mutation := newBaseNavStateMutation(c.config, OpDelete)
	return &BaseNavStateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BaseNavStateClient) DeleteOne(bns *BaseNavState) *BaseNavStateDeleteOne {
	return c.DeleteOneID(bns.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BaseNavStateClient) DeleteOneID(id string) *BaseNavStateDeleteOne {
	builder := c.Delete().Where(basenavstate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BaseNavStateDeleteOne{builder}
}

// Query returns a query builder for BaseNavState.
func (c *BaseNavStateClient) Query() *BaseNavStateQuery {
	return &BaseNavStateQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBaseNavState},
		inters: c.Interceptors(),
	}
}

// Get returns a BaseNavState entity by its id.
func (c *BaseNavStateClient) Get(ctx context.Context, id string) (*BaseNavState, error) {
	return c.Query().Where(basenavstate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BaseNavStateClient) GetX(ctx context.Context, id string) *BaseNavState {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BaseNavStateClient) Hooks() []Hook {
	return c.hooks.BaseNavState
}

// Interceptors returns the client interceptors.
func (c *BaseNavStateClient) Interceptors() []Interceptor {
	return c.inters.BaseNavState
}

func (c *BaseNavStateClient) mutate(ctx context.Context, m *BaseNavStateMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BaseNavStateCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BaseNavStateUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BaseNavStateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BaseNavStateDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entity: unknown BaseNavState mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Base, BaseInfo, BaseMember, BaseNavState []ent.Hook
	}
	inters struct {
		Base, BaseInfo, BaseMember, BaseNavState []ent.Interceptor
	}
)
