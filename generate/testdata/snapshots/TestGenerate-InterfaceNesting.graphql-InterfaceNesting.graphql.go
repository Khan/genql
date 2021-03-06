package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceNestingResponse is returned by InterfaceNesting on success.
type InterfaceNestingResponse struct {
	Root InterfaceNestingRootTopic `json:"root"`
}

// InterfaceNestingRootTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                `json:"id"`
	Children []InterfaceNestingRootTopicChildrenContent `json:"-"`
}

func (v *InterfaceNestingRootTopic) UnmarshalJSON(b []byte) error {
	var firstPass struct {
		*InterfaceNestingRootTopic
		Children json.RawMessage `json:"children"`
	}
	firstPass.InterfaceNestingRootTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err = json.Unmarshal(firstPass.Children, &tn)
	if err != nil {
		return err
	}
	switch tn.TypeName {

	case "Article":

		v.Children = InterfaceNestingRootTopicChildrenArticle{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Video":

		v.Children = InterfaceNestingRootTopicChildrenVideo{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Topic":

		v.Children = InterfaceNestingRootTopicChildrenTopic{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	}
	if err != nil {
		return err
	}

	return nil
}

// InterfaceNestingRootTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNestingRootTopicChildrenArticle struct {
	// ID is the identifier of the content.
	Id     testutil.ID                                         `json:"id"`
	Parent InterfaceNestingRootTopicChildrenArticleParentTopic `json:"parent"`
}

func (v InterfaceNestingRootTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenArticleParentTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenArticleParentTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                                          `json:"id"`
	Children []InterfaceNestingRootTopicChildrenArticleParentTopicChildrenContent `json:"-"`
}

func (v *InterfaceNestingRootTopicChildrenArticleParentTopic) UnmarshalJSON(b []byte) error {
	var firstPass struct {
		*InterfaceNestingRootTopicChildrenArticleParentTopic
		Children json.RawMessage `json:"children"`
	}
	firstPass.InterfaceNestingRootTopicChildrenArticleParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err = json.Unmarshal(firstPass.Children, &tn)
	if err != nil {
		return err
	}
	switch tn.TypeName {

	case "Article":

		v.Children = InterfaceNestingRootTopicChildrenArticleParentTopicChildrenArticle{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Video":

		v.Children = InterfaceNestingRootTopicChildrenArticleParentTopicChildrenVideo{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Topic":

		v.Children = InterfaceNestingRootTopicChildrenArticleParentTopicChildrenTopic{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	}
	if err != nil {
		return err
	}

	return nil
}

// InterfaceNestingRootTopicChildrenArticleParentTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNestingRootTopicChildrenArticleParentTopicChildrenArticle struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenArticleParentTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenArticleParentTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenArticleParentTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNestingRootTopicChildrenArticleParentTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenArticleParentTopicChildrenContent()
}

// InterfaceNestingRootTopicChildrenArticleParentTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenArticleParentTopicChildrenTopic struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenArticleParentTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenArticleParentTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenArticleParentTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNestingRootTopicChildrenArticleParentTopicChildrenVideo struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenArticleParentTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenArticleParentTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNestingRootTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent()
}

// InterfaceNestingRootTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenTopic struct {
	// ID is the identifier of the content.
	Id     testutil.ID                                       `json:"id"`
	Parent InterfaceNestingRootTopicChildrenTopicParentTopic `json:"parent"`
}

func (v InterfaceNestingRootTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenTopicParentTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenTopicParentTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                                        `json:"id"`
	Children []InterfaceNestingRootTopicChildrenTopicParentTopicChildrenContent `json:"-"`
}

func (v *InterfaceNestingRootTopicChildrenTopicParentTopic) UnmarshalJSON(b []byte) error {
	var firstPass struct {
		*InterfaceNestingRootTopicChildrenTopicParentTopic
		Children json.RawMessage `json:"children"`
	}
	firstPass.InterfaceNestingRootTopicChildrenTopicParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err = json.Unmarshal(firstPass.Children, &tn)
	if err != nil {
		return err
	}
	switch tn.TypeName {

	case "Article":

		v.Children = InterfaceNestingRootTopicChildrenTopicParentTopicChildrenArticle{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Video":

		v.Children = InterfaceNestingRootTopicChildrenTopicParentTopicChildrenVideo{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Topic":

		v.Children = InterfaceNestingRootTopicChildrenTopicParentTopicChildrenTopic{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	}
	if err != nil {
		return err
	}

	return nil
}

// InterfaceNestingRootTopicChildrenTopicParentTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNestingRootTopicChildrenTopicParentTopicChildrenArticle struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenTopicParentTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenTopicParentTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenTopicParentTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNestingRootTopicChildrenTopicParentTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenTopicParentTopicChildrenContent()
}

// InterfaceNestingRootTopicChildrenTopicParentTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenTopicParentTopicChildrenTopic struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenTopicParentTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenTopicParentTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenTopicParentTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNestingRootTopicChildrenTopicParentTopicChildrenVideo struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenTopicParentTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenTopicParentTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNestingRootTopicChildrenVideo struct {
	// ID is the identifier of the content.
	Id     testutil.ID                                       `json:"id"`
	Parent InterfaceNestingRootTopicChildrenVideoParentTopic `json:"parent"`
}

func (v InterfaceNestingRootTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenVideoParentTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenVideoParentTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                                        `json:"id"`
	Children []InterfaceNestingRootTopicChildrenVideoParentTopicChildrenContent `json:"-"`
}

func (v *InterfaceNestingRootTopicChildrenVideoParentTopic) UnmarshalJSON(b []byte) error {
	var firstPass struct {
		*InterfaceNestingRootTopicChildrenVideoParentTopic
		Children json.RawMessage `json:"children"`
	}
	firstPass.InterfaceNestingRootTopicChildrenVideoParentTopic = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err = json.Unmarshal(firstPass.Children, &tn)
	if err != nil {
		return err
	}
	switch tn.TypeName {

	case "Article":

		v.Children = InterfaceNestingRootTopicChildrenVideoParentTopicChildrenArticle{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Video":

		v.Children = InterfaceNestingRootTopicChildrenVideoParentTopicChildrenVideo{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Topic":

		v.Children = InterfaceNestingRootTopicChildrenVideoParentTopicChildrenTopic{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	}
	if err != nil {
		return err
	}

	return nil
}

// InterfaceNestingRootTopicChildrenVideoParentTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNestingRootTopicChildrenVideoParentTopicChildrenArticle struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenVideoParentTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenVideoParentTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenVideoParentTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNestingRootTopicChildrenVideoParentTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenVideoParentTopicChildrenContent()
}

// InterfaceNestingRootTopicChildrenVideoParentTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNestingRootTopicChildrenVideoParentTopicChildrenTopic struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenVideoParentTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenVideoParentTopicChildrenContent() {
}

// InterfaceNestingRootTopicChildrenVideoParentTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNestingRootTopicChildrenVideoParentTopicChildrenVideo struct {
	// ID is the identifier of the content.
	Id testutil.ID `json:"id"`
}

func (v InterfaceNestingRootTopicChildrenVideoParentTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNestingRootTopicChildrenVideoParentTopicChildrenContent() {
}

func InterfaceNesting(
	client graphql.Client,
) (*InterfaceNestingResponse, error) {
	var retval InterfaceNestingResponse
	err := client.MakeRequest(
		nil,
		"InterfaceNesting",
		`
query InterfaceNesting {
	root {
		id
		children {
			id
			parent {
				id
				children {
					id
				}
			}
		}
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

