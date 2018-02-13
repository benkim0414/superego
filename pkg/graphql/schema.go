package graphql

import (
	"errors"
	"fmt"

	"github.com/benkim0414/superego/pkg/profile"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

var (
	nodeDefinitions *relay.NodeDefinitions
	nameType        *graphql.Object
	profileType     *graphql.Object
)

func NewSchema(resolver Resolver) (graphql.Schema, error) {

	// interface Node {
	//   id: ID!
	// }
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)

			switch resolvedID.Type {
			case "Profile":
				return resolver.GetProfile(ctx, resolvedID.ID)
			default:
				return nil, errors.New("Unknown node type")
			}
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			default:
				return profileType
			}
		},
	})

	// type Name {
	//   formatted: String!
	//   familyName: String!
	//   givenName: String!
	// }
	nameType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Name",
		Fields: graphql.Fields{
			"formatted": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if name, ok := p.Source.(profile.Name); ok {
						formatted := fmt.Sprintf("%s %s", name.GivenName, name.FamilyName)
						return formatted, nil
					}
					return nil, nil
				},
			},
			"familyName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"givenName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})

	// type Profile : Node {
	//   id: ID!
	//   displayName: String
	//   name: Name
	//   email: String!
	//   imageUrl: String
	//   aboutMe: String
	// }
	profileType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Profile",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Profile", nil),
			"displayName": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: nameType,
			},
			"email": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"imageUrl": &graphql.Field{
				Type: graphql.String,
			},
			"aboutMe": &graphql.Field{
				Type: graphql.String,
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	// type Query {
	//   node(id: String!): Node
	// }
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": nodeDefinitions.NodeField,
		},
	})

	// input CreateProfileInput {
	//   clientMutationID: String!
	//   displayName: String
	//   name: Name
	//   email: String!
	//   imageUrl: String
	//   aboutMe: String
	// }
	//
	// input CreateProfilePayload {
	//   clientMutationID: String!
	//   profile: Profile
	// }
	profileMutation := relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "CreateProfile",
		InputFields: graphql.InputObjectConfigFieldMap{
			"displayName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"familyName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"givenName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"imageUrl": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"aboutMe": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
		OutputFields: graphql.Fields{
			"profile": &graphql.Field{
				Type: profileType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if payload, ok := p.Source.(map[string]interface{}); ok {
						ctx := p.Context
						profileID := payload["profileId"].(string)
						return resolver.GetProfile(ctx, profileID)
					}
					return nil, nil
				},
			},
		},
		MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			email := inputMap["email"].(string)
			p := &profile.Profile{
				Email: email,
			}

			if displayName, ok := inputMap["displayName"].(string); ok {
				p.DisplayName = displayName
			}

			name := profile.Name{}
			if familyName, ok := inputMap["familyName"].(string); ok {
				name.FamilyName = familyName
			}
			if givenName, ok := inputMap["givenName"].(string); ok {
				name.GivenName = givenName
			}
			p.Name = name

			if imageURL, ok := inputMap["imageUrl"].(string); ok {
				p.ImageURL = imageURL
			}

			if aboutMe, ok := inputMap["aboutMe"].(string); ok {
				p.AboutMe = aboutMe
			}

			profile, err := resolver.PostProfile(ctx, p)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"profileId": profile.ID,
			}, nil
		},
	})

	// type Mutation {
	//   createProfile(input CreateProfileInput!): CreateProfilePayload
	// }
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createProfile": profileMutation,
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
