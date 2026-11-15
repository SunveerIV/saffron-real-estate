package service

import (
	"context"
	"errors"

	"saffronreal.estate/backend/internal/model"
	"saffronreal.estate/backend/internal/repository"
)

type HouseBuyRequestService struct {
	houseBuyRequests *repository.HouseBuyRequestRepository
	users            *UserService
}

var (
	ErrInvalidFirstName         = errors.New("first_name length must be nonzero")
	ErrInvalidLastName          = errors.New("last_name length must be nonzero")
	ErrInvalidAddress           = errors.New("address length must be nonzero")
	ErrInvalidPriceOfferedCents = errors.New("price_offered_cents must be nonnegative")
	ErrInvalidComments          = errors.New("comments must not be more than 200 chars")
)

type CreateHouseBuyRequestParams struct {
	Email             string
	FirstName         string
	LastName          string
	Address           string
	PriceOfferedCents int64
	Comments          string
}

func NewHouseBuyRequestService(
	houseBuyRequests *repository.HouseBuyRequestRepository,
	users *UserService,
) *HouseBuyRequestService {
	return &HouseBuyRequestService{
		houseBuyRequests: houseBuyRequests,
		users:            users,
	}
}

func (s *HouseBuyRequestService) Create(
	ctx context.Context,
	params CreateHouseBuyRequestParams,
) (model.HouseBuyRequest, error) {
	err := validateHouseBuyRequest(params)
	if err != nil {
		return model.HouseBuyRequest{}, err
	}

	user, err := s.users.Get(ctx, params.Email)
	if errors.Is(err, repository.ErrUserNotFound) {
		user, err = s.users.Create(ctx, params.Email)
	}
	if err != nil {
		return model.HouseBuyRequest{}, err
	}

	return s.houseBuyRequests.NewHouseBuyRequest(
		ctx,
		repository.CreateHouseBuyRequestParams{
			UserId:            user.Id,
			FirstName:         params.FirstName,
			LastName:          params.LastName,
			Address:           params.Address,
			PriceOfferedCents: params.PriceOfferedCents,
			Comments:          params.Comments,
		},
	)
}

func (s *HouseBuyRequestService) GetHouseBuyRequests(
	ctx context.Context,
	email string,
) ([]model.HouseBuyRequest, error) {
	user, err := s.users.Get(ctx, email)
	if err != nil {
		return []model.HouseBuyRequest{}, err
	}

	return s.houseBuyRequests.GetHouseBuyRequests(ctx, user.Id)
}

func validateHouseBuyRequest(
	params CreateHouseBuyRequestParams,
) error {
	if params.Address == "" {
		return ErrInvalidAddress
	}

	if params.FirstName == "" {
		return ErrInvalidFirstName
	}

	if params.LastName == "" {
		return ErrInvalidLastName
	}

	if len(params.Comments) > 200 {
		return ErrInvalidComments
	}

	if params.PriceOfferedCents < 0 {
		return ErrInvalidPriceOfferedCents
	}

	return nil
}
