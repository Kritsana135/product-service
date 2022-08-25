package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"product-service/graph/generated"
	"product-service/graph/helper"
	"product-service/graph/model"
	"strings"
)

// GenerateVariants is the resolver for the generateVariants field.
func (r *mutationResolver) GenerateVariants(ctx context.Context, input model.GenerateVariantsInput) ([]*model.GenerateVariantsPayload, error) {
	// generate variants
	variants := input.Variants
	for {
		if len(variants) <= 1 {
			break
		}
		temptCombine := make([]string, 0)
		concatTitle := fmt.Sprintf("%v;%v", variants[0].Title, variants[1].Title)
		for _, option_1 := range variants[0].Options {
			for _, option_2 := range variants[1].Options {
				concatOptions := fmt.Sprintf("%v;%v", option_1, option_2)
				temptCombine = append(temptCombine, concatOptions)
			}
		}
		variants = variants[2:] // shift off first element
		variants = append([]*model.Variant{{Title: concatTitle, Options: temptCombine}}, variants...)
	}
	// result payload
	result := make([]*model.GenerateVariantsPayload, 0)
	for _, variant := range variants {
		for _, option := range variant.Options {
			splittedOpt := helper.ConvertAsToPas(strings.Split(option, ";"))
			splittedTitle := helper.ConvertAsToPas(strings.Split(variant.Title, ";"))
			result = append(result, &model.GenerateVariantsPayload{
				Sku:            "",
				VariantTitle:   splittedTitle,
				VariantOptions: splittedOpt,
				FactoryPrice:   0,
				ShopPrice:      0,
				CustomerPrice:  0,
				Quantity:       0,
				FileKey:        "",
			})
		}
	}
	return result, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
