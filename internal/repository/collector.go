package repository

import (
	"context"

	"content_collector/internal/apperrors"
	"content_collector/internal/config"
	"content_collector/internal/domain/model"
	"content_collector/internal/infrastructure/datastore"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectorCollectionName = "collectors"
	defaultInDaysDeleteAt   = 1
)

type CollectorMongoDBRepository struct {
	CollectorCollection *mongo.Collection
	context             context.Context
}

func NewCollectorMongoDBRepository(cfg *config.Config, client *datastore.MongoClient, context context.Context) ICollectorRepository {
	return &CollectorMongoDBRepository{
		CollectorCollection: client.GetClient().Database(cfg.Mongo.MongoDbName).Collection(collectorCollectionName),
		context:             context,
	}
}

func (r *CollectorMongoDBRepository) GetById(id string) (*model.CollectorRepository, error) {
	filter := bson.M{"_id": id}
	var collector model.CollectorRepository
	err := r.CollectorCollection.FindOne(r.context, filter).Decode(&collector)
	if err != nil {
		return nil, apperrors.MongoCollectorRepositoryGetByIdError.AppendMessage(err)
	}
	return &collector, nil
}

func (r *CollectorMongoDBRepository) GetByUrl(url string) (*model.CollectorRepository, error) {
	filter := bson.M{"url": url}
	var collector model.CollectorRepository
	err := r.CollectorCollection.FindOne(r.context, filter).Decode(&collector)
	if err != nil {
		return nil, apperrors.MongoCollectorRepositoryGetByIdError.AppendMessage(err)
	}
	return &collector, nil
}

func (r *CollectorMongoDBRepository) GetForDelete() ([]*model.CollectorRepository, error) {
	filter := bson.M{"delete_at": bson.M{"$lte": model.NewTime()}}
	var collectors []*model.CollectorRepository
	cursor, err := r.CollectorCollection.Find(r.context, filter)
	if err != nil {
		return nil, apperrors.MongoCollectorRepositoryGetByIdError.AppendMessage(err)
	}
	err = cursor.All(r.context, &collectors)
	if err != nil {
		return nil, apperrors.MongoCollectorRepositoryGetByIdError.AppendMessage(err)
	}
	return collectors, nil
}

func (r *CollectorMongoDBRepository) Create(collector *model.CollectorRepository) error {
	if collector.ID == "" {
		collector.ID = model.NewUUID()
	}

	createdTime := model.NewTime()
	collector.CreatedAt = &createdTime
	if collector.DeleteAt == nil {
		deleteTime := createdTime.AddDate(0, 0, defaultInDaysDeleteAt)
		collector.DeleteAt = &deleteTime
	}

	_, err := r.CollectorCollection.InsertOne(r.context, collector)
	if err != nil {
		return apperrors.MongoCollectorRepositoryCreateError.AppendMessage(err)
	}

	return nil
}

func (r *CollectorMongoDBRepository) Update(collector *model.CollectorRepository) error {
	filter := bson.M{"_id": collector.ID}
	collectorByte, err := bson.Marshal(collector)
	if err != nil {
		return apperrors.MongoCollectorRepositoryUpdateMarshalError.AppendMessage(err)
	}

	var update bson.M
	err = bson.Unmarshal(collectorByte, &update)
	if err != nil {
		return apperrors.MongoCollectorUpdateUnmarshalError.AppendMessage(err)
	}

	_, err = r.CollectorCollection.UpdateOne(r.context, filter, update)
	if err != nil {
		return apperrors.MongoCollectorRepositoryUpdateError.AppendMessage(err)
	}
	return nil
}

func (r *CollectorMongoDBRepository) Delete(id string) error {
	filter := bson.M{"_id": id}
	_, err := r.CollectorCollection.DeleteOne(r.context, filter)
	if err != nil {
		return apperrors.MongoCollectorRepositoryDeleteError.AppendMessage(err)
	}
	return nil
}
