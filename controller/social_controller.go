package controller

import (
	"context"

	socialPB "github.com/digisata/auth-service/stubs/social"
)

func (c AuthController) CreateSocial(ctx context.Context, req *socialPB.CreateSocialRequest) (*socialPB.SocialBaseResponse, error) {
	res, err := c.SocialUseCase.CreateSocial(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetSocial(ctx context.Context, req *socialPB.GetSocialByIDRequest) (*socialPB.SocialBaseResponse, error) {
	res, err := c.SocialUseCase.GetSocial(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetAllSocial(ctx context.Context, req *socialPB.GetAllSocialRequest) (*socialPB.GetAllSocialResponse, error) {
	res, err := c.SocialUseCase.GetAllSocial(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) UpdateSocial(ctx context.Context, req *socialPB.UpdateSocialByIDRequest) (*socialPB.SocialBaseResponse, error) {
	res, err := c.SocialUseCase.UpdateSocial(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) DeleteSocial(ctx context.Context, req *socialPB.DeleteSocialByIDRequest) (*socialPB.SocialBaseResponse, error) {
	res, err := c.SocialUseCase.DeleteSocial(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
