// Code generated by Prisma CLI (https://github.com/prisma/prisma). DO NOT EDIT.

package prisma

import (
	"context"
	"errors"

	"github.com/prisma/prisma-client-lib-go"

	"github.com/machinebox/graphql"
)

var ErrNoResult = errors.New("query returned no result")

func Str(v string) *string { return &v }
func Int32(v int32) *int32 { return &v }
func Bool(v bool) *bool    { return &v }

type BatchPayloadExec struct {
	exec *prisma.BatchPayloadExec
}

func (exec *BatchPayloadExec) Exec(ctx context.Context) (BatchPayload, error) {
	bp, err := exec.exec.Exec(ctx)
	return BatchPayload(bp), err
}

type BatchPayload struct {
	Count int64 `json:"count"`
}

type Aggregate struct {
	Count int64 `json:"count"`
}

type Client struct {
	Client *prisma.Client
}

type Options struct {
	Endpoint string
	Secret   string
}

func New(options *Options, opts ...graphql.ClientOption) *Client {
	endpoint := DefaultEndpoint
	secret := Secret
	if options != nil {
		endpoint = options.Endpoint
		secret = options.Secret
	}
	return &Client{
		Client: prisma.New(endpoint, secret, opts...),
	}
}

func (client *Client) GraphQL(ctx context.Context, query string, variables map[string]interface{}) (map[string]interface{}, error) {
	return client.Client.GraphQL(ctx, query, variables)
}

var DefaultEndpoint = "http://localhost:4466"
var Secret = ""

func (client *Client) User(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"user",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserExec{ret}
}

type UsersParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Users(params *UsersParams) *UserExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"users",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserExecArray{ret}
}

type UsersConnectionParams struct {
	Where   *UserWhereInput   `json:"where,omitempty"`
	OrderBy *UserOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) UsersConnection(params *UsersConnectionParams) *UserConnectionExec {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"UserWhereInput", "UserOrderByInput", "User"},
		"usersConnection",
		[]string{"edges", "pageInfo"})

	return &UserConnectionExec{ret}
}

func (client *Client) CreateUser(params UserCreateInput) *UserExec {
	ret := client.Client.Create(
		params,
		[2]string{"UserCreateInput!", "User"},
		"createUser",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserExec{ret}
}

type UserUpdateParams struct {
	Data  UserUpdateInput      `json:"data"`
	Where UserWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateUser(params UserUpdateParams) *UserExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"UserUpdateInput!", "UserWhereUniqueInput!", "User"},
		"updateUser",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserExec{ret}
}

type UserUpdateManyParams struct {
	Data  UserUpdateManyMutationInput `json:"data"`
	Where *UserWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyUsers(params UserUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"UserUpdateManyMutationInput!", "UserWhereInput"},
		"updateManyUsers")
	return &BatchPayloadExec{exec}
}

type UserUpsertParams struct {
	Where  UserWhereUniqueInput `json:"where"`
	Create UserCreateInput      `json:"create"`
	Update UserUpdateInput      `json:"update"`
}

func (client *Client) UpsertUser(params UserUpsertParams) *UserExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"UserWhereUniqueInput!", "UserCreateInput!", "UserUpdateInput!", "User"},
		"upsertUser",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserExec{ret}
}

func (client *Client) DeleteUser(params UserWhereUniqueInput) *UserExec {
	ret := client.Client.Delete(
		params,
		[2]string{"UserWhereUniqueInput!", "User"},
		"deleteUser",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserExec{ret}
}

func (client *Client) DeleteManyUsers(params *UserWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "UserWhereInput", "deleteManyUsers")
	return &BatchPayloadExec{exec}
}

type UserOrderByInput string

const (
	UserOrderByInputIDAsc            UserOrderByInput = "id_ASC"
	UserOrderByInputIDDesc           UserOrderByInput = "id_DESC"
	UserOrderByInputNameAsc          UserOrderByInput = "name_ASC"
	UserOrderByInputNameDesc         UserOrderByInput = "name_DESC"
	UserOrderByInputUsernameAsc      UserOrderByInput = "username_ASC"
	UserOrderByInputUsernameDesc     UserOrderByInput = "username_DESC"
	UserOrderByInputEmailAsc         UserOrderByInput = "email_ASC"
	UserOrderByInputEmailDesc        UserOrderByInput = "email_DESC"
	UserOrderByInputPasswordAsc      UserOrderByInput = "password_ASC"
	UserOrderByInputPasswordDesc     UserOrderByInput = "password_DESC"
	UserOrderByInputRegisteredAtAsc  UserOrderByInput = "registeredAt_ASC"
	UserOrderByInputRegisteredAtDesc UserOrderByInput = "registeredAt_DESC"
	UserOrderByInputUpdatedAtAsc     UserOrderByInput = "updatedAt_ASC"
	UserOrderByInputUpdatedAtDesc    UserOrderByInput = "updatedAt_DESC"
)

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

type UserWhereUniqueInput struct {
	ID       *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
}

type UserWhereInput struct {
	ID                    *string          `json:"id,omitempty"`
	IDNot                 *string          `json:"id_not,omitempty"`
	IDIn                  []string         `json:"id_in,omitempty"`
	IDNotIn               []string         `json:"id_not_in,omitempty"`
	IDLt                  *string          `json:"id_lt,omitempty"`
	IDLte                 *string          `json:"id_lte,omitempty"`
	IDGt                  *string          `json:"id_gt,omitempty"`
	IDGte                 *string          `json:"id_gte,omitempty"`
	IDContains            *string          `json:"id_contains,omitempty"`
	IDNotContains         *string          `json:"id_not_contains,omitempty"`
	IDStartsWith          *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith       *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith            *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith         *string          `json:"id_not_ends_with,omitempty"`
	Name                  *string          `json:"name,omitempty"`
	NameNot               *string          `json:"name_not,omitempty"`
	NameIn                []string         `json:"name_in,omitempty"`
	NameNotIn             []string         `json:"name_not_in,omitempty"`
	NameLt                *string          `json:"name_lt,omitempty"`
	NameLte               *string          `json:"name_lte,omitempty"`
	NameGt                *string          `json:"name_gt,omitempty"`
	NameGte               *string          `json:"name_gte,omitempty"`
	NameContains          *string          `json:"name_contains,omitempty"`
	NameNotContains       *string          `json:"name_not_contains,omitempty"`
	NameStartsWith        *string          `json:"name_starts_with,omitempty"`
	NameNotStartsWith     *string          `json:"name_not_starts_with,omitempty"`
	NameEndsWith          *string          `json:"name_ends_with,omitempty"`
	NameNotEndsWith       *string          `json:"name_not_ends_with,omitempty"`
	Username              *string          `json:"username,omitempty"`
	UsernameNot           *string          `json:"username_not,omitempty"`
	UsernameIn            []string         `json:"username_in,omitempty"`
	UsernameNotIn         []string         `json:"username_not_in,omitempty"`
	UsernameLt            *string          `json:"username_lt,omitempty"`
	UsernameLte           *string          `json:"username_lte,omitempty"`
	UsernameGt            *string          `json:"username_gt,omitempty"`
	UsernameGte           *string          `json:"username_gte,omitempty"`
	UsernameContains      *string          `json:"username_contains,omitempty"`
	UsernameNotContains   *string          `json:"username_not_contains,omitempty"`
	UsernameStartsWith    *string          `json:"username_starts_with,omitempty"`
	UsernameNotStartsWith *string          `json:"username_not_starts_with,omitempty"`
	UsernameEndsWith      *string          `json:"username_ends_with,omitempty"`
	UsernameNotEndsWith   *string          `json:"username_not_ends_with,omitempty"`
	Email                 *string          `json:"email,omitempty"`
	EmailNot              *string          `json:"email_not,omitempty"`
	EmailIn               []string         `json:"email_in,omitempty"`
	EmailNotIn            []string         `json:"email_not_in,omitempty"`
	EmailLt               *string          `json:"email_lt,omitempty"`
	EmailLte              *string          `json:"email_lte,omitempty"`
	EmailGt               *string          `json:"email_gt,omitempty"`
	EmailGte              *string          `json:"email_gte,omitempty"`
	EmailContains         *string          `json:"email_contains,omitempty"`
	EmailNotContains      *string          `json:"email_not_contains,omitempty"`
	EmailStartsWith       *string          `json:"email_starts_with,omitempty"`
	EmailNotStartsWith    *string          `json:"email_not_starts_with,omitempty"`
	EmailEndsWith         *string          `json:"email_ends_with,omitempty"`
	EmailNotEndsWith      *string          `json:"email_not_ends_with,omitempty"`
	Password              *string          `json:"password,omitempty"`
	PasswordNot           *string          `json:"password_not,omitempty"`
	PasswordIn            []string         `json:"password_in,omitempty"`
	PasswordNotIn         []string         `json:"password_not_in,omitempty"`
	PasswordLt            *string          `json:"password_lt,omitempty"`
	PasswordLte           *string          `json:"password_lte,omitempty"`
	PasswordGt            *string          `json:"password_gt,omitempty"`
	PasswordGte           *string          `json:"password_gte,omitempty"`
	PasswordContains      *string          `json:"password_contains,omitempty"`
	PasswordNotContains   *string          `json:"password_not_contains,omitempty"`
	PasswordStartsWith    *string          `json:"password_starts_with,omitempty"`
	PasswordNotStartsWith *string          `json:"password_not_starts_with,omitempty"`
	PasswordEndsWith      *string          `json:"password_ends_with,omitempty"`
	PasswordNotEndsWith   *string          `json:"password_not_ends_with,omitempty"`
	RegisteredAt          *string          `json:"registeredAt,omitempty"`
	RegisteredAtNot       *string          `json:"registeredAt_not,omitempty"`
	RegisteredAtIn        []string         `json:"registeredAt_in,omitempty"`
	RegisteredAtNotIn     []string         `json:"registeredAt_not_in,omitempty"`
	RegisteredAtLt        *string          `json:"registeredAt_lt,omitempty"`
	RegisteredAtLte       *string          `json:"registeredAt_lte,omitempty"`
	RegisteredAtGt        *string          `json:"registeredAt_gt,omitempty"`
	RegisteredAtGte       *string          `json:"registeredAt_gte,omitempty"`
	UpdatedAt             *string          `json:"updatedAt,omitempty"`
	UpdatedAtNot          *string          `json:"updatedAt_not,omitempty"`
	UpdatedAtIn           []string         `json:"updatedAt_in,omitempty"`
	UpdatedAtNotIn        []string         `json:"updatedAt_not_in,omitempty"`
	UpdatedAtLt           *string          `json:"updatedAt_lt,omitempty"`
	UpdatedAtLte          *string          `json:"updatedAt_lte,omitempty"`
	UpdatedAtGt           *string          `json:"updatedAt_gt,omitempty"`
	UpdatedAtGte          *string          `json:"updatedAt_gte,omitempty"`
	And                   []UserWhereInput `json:"AND,omitempty"`
}

type UserCreateInput struct {
	ID       *string `json:"id,omitempty"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type UserUpdateInput struct {
	Name     *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserUpdateManyMutationInput struct {
	Name     *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *UserWhereInput              `json:"node,omitempty"`
	And                        []UserSubscriptionWhereInput `json:"AND,omitempty"`
}

type UserExec struct {
	exec *prisma.Exec
}

func (instance UserExec) Exec(ctx context.Context) (*User, error) {
	var v User
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserExecArray struct {
	exec *prisma.Exec
}

func (instance UserExecArray) Exec(ctx context.Context) ([]User, error) {
	var v []User
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RegisteredAt string `json:"registeredAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type UserConnectionExec struct {
	exec *prisma.Exec
}

func (instance *UserConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *UserConnectionExec) Edges() *UserEdgeExecArray {
	edges := instance.exec.Client.GetMany(
		instance.exec,
		nil,
		[3]string{"UserWhereInput", "UserOrderByInput", "UserEdge"},
		"edges",
		[]string{"cursor"})

	nodes := edges.Client.GetMany(
		edges,
		nil,
		[3]string{"", "", "User"},
		"node",
		[]string{"id", "createdAt", "updatedAt", "name", "desc"})

	return &UserEdgeExecArray{nodes}
}

func (instance *UserConnectionExec) Aggregate(ctx context.Context) (*Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateUser"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return &v, err
}

func (instance UserConnectionExec) Exec(ctx context.Context) (*UserConnection, error) {
	var v UserConnection
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance UserConnectionExecArray) Exec(ctx context.Context) ([]UserConnection, error) {
	var v []UserConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserConnection struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Edges    []UserEdge `json:"edges"`
}

type PageInfoExec struct {
	exec *prisma.Exec
}

func (instance PageInfoExec) Exec(ctx context.Context) (*PageInfo, error) {
	var v PageInfo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PageInfoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PageInfoExecArray struct {
	exec *prisma.Exec
}

func (instance PageInfoExecArray) Exec(ctx context.Context) ([]PageInfo, error) {
	var v []PageInfo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type UserEdgeExec struct {
	exec *prisma.Exec
}

func (instance *UserEdgeExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserExec{ret}
}

func (instance UserEdgeExec) Exec(ctx context.Context) (*UserEdge, error) {
	var v UserEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance UserEdgeExecArray) Exec(ctx context.Context) ([]UserEdge, error) {
	var v []UserEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserEdge struct {
	Node   User   `json:"node"`
	Cursor string `json:"cursor"`
}

type UserSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *UserSubscriptionPayloadExec) Node() *UserExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "User"},
		"node",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserExec{ret}
}

func (instance *UserSubscriptionPayloadExec) PreviousValues() *UserPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "UserPreviousValues"},
		"previousValues",
		[]string{"id", "name", "username", "email", "password", "registeredAt", "updatedAt"})

	return &UserPreviousValuesExec{ret}
}

func (instance UserSubscriptionPayloadExec) Exec(ctx context.Context) (*UserSubscriptionPayload, error) {
	var v UserSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance UserSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]UserSubscriptionPayload, error) {
	var v []UserSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserSubscriptionPayload struct {
	Mutation      MutationType `json:"mutation"`
	Node          *User        `json:"node,omitempty"`
	UpdatedFields []string     `json:"updatedFields,omitempty"`
}

type UserPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExec) Exec(ctx context.Context) (*UserPreviousValues, error) {
	var v UserPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance UserPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type UserPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance UserPreviousValuesExecArray) Exec(ctx context.Context) ([]UserPreviousValues, error) {
	var v []UserPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type UserPreviousValues struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RegisteredAt string `json:"registeredAt"`
	UpdatedAt    string `json:"updatedAt"`
}
