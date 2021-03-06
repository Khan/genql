package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
)

// QueryWithStructsResponse is returned by QueryWithStructs on success.
type QueryWithStructsResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithStructsUser `json:"user"`
}

// QueryWithStructsUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithStructsUser struct {
	AuthMethods []QueryWithStructsUserAuthMethodsAuthMethod `json:"authMethods"`
}

// QueryWithStructsUserAuthMethodsAuthMethod includes the requested fields of the GraphQL type AuthMethod.
type QueryWithStructsUserAuthMethodsAuthMethod struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
}

func QueryWithStructs(
	client graphql.Client,
) (*QueryWithStructsResponse, error) {
	var retval QueryWithStructsResponse
	err := client.MakeRequest(
		nil,
		"QueryWithStructs",
		`
query QueryWithStructs {
	user {
		authMethods {
			provider
			email
		}
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

