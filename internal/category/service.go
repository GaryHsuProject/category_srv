package category

import (
	"context"
	"shop/driver"
	"shop/models"
	. "shop/proto/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryService struct {
	categoryRepo Repository
	UnimplementedCategoryServer
}

func NewCategoryService(categoryRepo Repository) CategoryServer {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) GetAllCategory(context.Context, *emptypb.Empty) (*Categories, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		driver.Logger.Error(err)
		return nil, err
	}
	pbCateogories := &Categories{}
	for _, category := range categories {
		pbCategory := &CategoryDetail{
			Id:       int32(category.ID),
			Name:     category.Name,
			Sort:     int32(category.Sort),
			IsParent: category.IsParent,
			ParentId: int32(category.ParentId),
		}
		pbCateogories.Categories = append(pbCateogories.Categories, pbCategory)
	}
	return pbCateogories, nil

}

func (s *CategoryService) Create(ctx context.Context, req *CreateCategory) (*emptypb.Empty, error) {
	category := &models.Category{
		Name:     req.Name,
		Sort:     int(req.Sort),
		ParentId: int(req.ParentId),
		IsParent: req.IsParent,
	}
	err := s.categoryRepo.Insert(category)
	if err != nil {
		driver.Logger.Error(err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
