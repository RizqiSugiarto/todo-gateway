package controller

import (
	"context"

	invitationCategoryPB "github.com/digisata/auth-service/stubs/invitation-category"
)

func (c AuthController) CreateInvitationCategory(ctx context.Context, req *invitationCategoryPB.CreateInvitationCategoryRequest) (*invitationCategoryPB.InvitationCategoryBaseResponse, error) {
	res, err := c.InvitationCategoryUseCase.CreateInvitationCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetInvitationCategory(ctx context.Context, req *invitationCategoryPB.GetInvitationCategoryByIDRequest) (*invitationCategoryPB.InvitationCategoryBaseResponse, error) {
	res, err := c.InvitationCategoryUseCase.GetInvitationCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetAllInvitationCategory(ctx context.Context, req *invitationCategoryPB.GetAllInvitationCategoryRequest) (*invitationCategoryPB.GetAllInvitationCategoryResponse, error) {
	res, err := c.InvitationCategoryUseCase.GetAllInvitationCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) UpdateInvitationCategory(ctx context.Context, req *invitationCategoryPB.UpdateInvitationCategoryByIDRequest) (*invitationCategoryPB.InvitationCategoryBaseResponse, error) {
	res, err := c.InvitationCategoryUseCase.UpdateInvitationCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) DeleteInvitationCategory(ctx context.Context, req *invitationCategoryPB.DeleteInvitationCategoryByIDRequest) (*invitationCategoryPB.InvitationCategoryBaseResponse, error) {
	res, err := c.InvitationCategoryUseCase.DeleteInvitationCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
