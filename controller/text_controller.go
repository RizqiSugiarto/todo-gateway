package controller

import (
	"context"

	TextPB "github.com/digisata/todo-gateway/stubs/text"
)

func (c AuthController) CreateText(ctx context.Context, req *TextPB.CreateTextRequest) (*TextPB.TextBaseResponse, error) {
	res, err := c.TextUseCase.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetText(ctx context.Context, req *TextPB.GetTextByIDRequest) (*TextPB.GetTextByIDResponse, error) {
	res, err := c.TextUseCase.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) GetAllText(ctx context.Context, req *TextPB.GetAllTextByActivityIDRequest) (*TextPB.GetAllTextByActivityIDResponse, error) {
	res, err := c.TextUseCase.GetAllByUserID(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) UpdateText(ctx context.Context, req *TextPB.UpdateTextByIDRequest) (*TextPB.TextBaseResponse, error) {
	res, err := c.TextUseCase.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AuthController) DeleteText(ctx context.Context, req *TextPB.DeleteTextByIDRequest) (*TextPB.TextBaseResponse, error) {
	res, err := c.TextUseCase.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
