package usecases

import (
	"context"
	"folder_mod/internal/entities"
)

type FolderRepository interface {
	CreateFolder(ctx context.Context, folder *entities.Folder) (*entities.Folder, error)
	GetFolders(ctx context.Context) ([]*entities.Folder, error)
	FindByID(ctx context.Context, id string) (*entities.Folder, error)
	UpdateFolderData(ctx context.Context, folder *entities.Folder) error
	//DeleteFolder(ctx context.Context, folder *entities.Folder) error
	DeleteFolderByID(ctx context.Context, id string) error

	UpdateFolderParentID(ctx context.Context, baseID string, parentID string) error
	AddChildIDToParent(ctx context.Context, parentID string, childID string) error
	PositionExists(ctx context.Context, baseID string, parentID string, position float64) error
	FindFoldersByParentID(ctx context.Context, parentID string) ([]entities.Folder, error) // 新增的方法
}
