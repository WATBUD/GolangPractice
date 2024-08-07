package repository

import (
	"context"
	"errors"
	"fmt"
	"folder_mod/internal/entities"
	"folder_mod/internal/usecases"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFolderRepository struct {
	collection *mongo.Collection
}

func NewMongoFolderRepository(client *mongo.Client) usecases.FolderRepository {
	collection := client.Database("mai_dev").Collection("folders")
	return &MongoFolderRepository{collection: collection}
}

func (r *MongoFolderRepository) CreateFolder(ctx context.Context, folder *entities.Folder) (*entities.Folder, error) {
	// Insert the new folder
	insertResult, err := r.collection.InsertOne(ctx, folder)
	if err != nil {
		return nil, fmt.Errorf("failed to insert folder: %w", err)
	}
	// Retrieve the inserted document
	var insertedFolder entities.Folder
	err = r.collection.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&insertedFolder)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve inserted folder: %w", err)
	}
	// newFolderID := insertResult.InsertedID.(primitive.ObjectID).Hex()
	// folder.ID = newFolderID
	return &insertedFolder, nil
}

func (r *MongoFolderRepository) GetFolders(ctx context.Context) ([]*entities.Folder, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var folders []*entities.Folder
	if err := cursor.All(ctx, &folders); err != nil {
		return nil, err
	}
	return folders, nil
}

func (r *MongoFolderRepository) FindByID(ctx context.Context, id string) (*entities.Folder, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var folder entities.Folder
	if err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&folder); err != nil {
		return nil, err
	}
	return &folder, nil
}

func (r *MongoFolderRepository) UpdateFolderData(ctx context.Context, folder *entities.Folder) error {
	objectID, err := primitive.ObjectIDFromHex(folder.ID)
	if err != nil {
		return fmt.Errorf("invalid parent_id: %w", err)
	}
	filter := bson.M{"_id": objectID, "base_id": folder.BaseID}
	update := bson.M{
		"$set": bson.M{
			"parent_id":  folder.ParentID,
			"position":   folder.Position,
			"data":       folder.Data,
			"updated_at": folder.UpdatedAt,
		},
	}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update folder: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no matching document found for ID: %s and BaseID: %s", folder.ID, folder.BaseID)
	}
	return nil
}

func (r *MongoFolderRepository) DeleteFolderByID(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID: %w", err)
	}

	// Find the folder to be deleted based on the _id
	var folderToDelete entities.Folder
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&folderToDelete)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("folder not found")
		}
		return fmt.Errorf("failed to find folder: %w", err)
	}

	// Check if the folder type is "folder"
	if folderToDelete.Type != "folder" {
		return fmt.Errorf("item to delete is not a folder")
	}

	// Recursively delete folder and its children
	return r.deleteFolderAndChildren(ctx, objectID)
}

func (r *MongoFolderRepository) DeleteFolder(ctx context.Context, folder *entities.Folder) error {
	objectID, err := primitive.ObjectIDFromHex(folder.ID)
	if err != nil {
		return fmt.Errorf("invalid parent_id: %w", err)
	}
	// Find the folder to be deleted based on the _id
	var folderToDelete entities.Folder
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&folderToDelete)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("folder not found")
		}
		return fmt.Errorf("failed to find folder: %w", err)
	}

	// Recursively delete folder and its children
	return r.deleteFolderAndChildren(ctx, objectID)
}

func (r *MongoFolderRepository) deleteFolderAndChildren(ctx context.Context, folderID primitive.ObjectID) error {
	// Find all children folders
	children, err := r.FindFoldersByParentID(ctx, folderID.Hex())
	if err != nil {
		return err
	}

	// Recursively delete all children folders
	for _, child := range children {
		childID, err := primitive.ObjectIDFromHex(child.ID)
		if err != nil {
			return fmt.Errorf("invalid child ID: %w", err)
		}
		if err := r.deleteFolderAndChildren(ctx, childID); err != nil {
			return err
		}
	}

	// Delete the folder itself
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": folderID})
	return err
}

func (r *MongoFolderRepository) FindFoldersByParentID(ctx context.Context, parentID string) ([]entities.Folder, error) {
	filter := bson.M{"parent_id": parentID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var folders []entities.Folder
	for cursor.Next(ctx) {
		var folder entities.Folder
		if err := cursor.Decode(&folder); err != nil {
			return nil, err
		}
		folders = append(folders, folder)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return folders, nil
}

func (r *MongoFolderRepository) UpdateFolderParentID(ctx context.Context, objectID string, parentID string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"base_id": objectID},
		bson.M{"$set": bson.M{"parent_id": parentID}},
	)
	return err
}

func (r *MongoFolderRepository) AddChildIDToParent(ctx context.Context, parentID string, childID string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"base_id": parentID},
		bson.M{"$addToSet": bson.M{"child_ids": childID}},
	)
	return err
}

func (r *MongoFolderRepository) PositionExists(ctx context.Context, baseID string, parentID string, position float64) error {
	filter := bson.M{
		"base_id":   baseID,
		"position":  position,
		"parent_id": parentID,
	}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("position already exists for the same baseID and parentID")
	}
	return nil
}
