package DB

import (
	"chat/dal/initDB"
	"chat/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StorageMessage(message *model.Message) error {
	_, err := initDB.MONGODB_DB.Database("chat").Collection("message").InsertOne(context.TODO(), &message)
	return err
}

func GetRemarkById(mid string, fid string) *mongo.SingleResult {
	findOptions := options.Find()
	findOptions.SetProjection(bson.D{{"mid", 1}, {"fid", 1}, {"remark", 1}, {"_id", 0}})
	result := initDB.MONGODB_DB.Database("chat").Collection("session").FindOne(context.TODO(), bson.D{
		{"mid", mid},
		{"fid", fid},
	})
	return result
}
func GetMessages(mid int64, fid int64) (*mongo.Cursor, error) {
	filter := bson.D{
		{
			"$or", bson.A{
				bson.D{{"fromID", mid}, {"toID", fid}},
				bson.D{{"fromID", fid}, {"toID", mid}},
			},
		},
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"createAt", -1}})
	cursor, err := initDB.MONGODB_DB.Database("chat").Collection("message").Find(
		context.TODO(),
		filter,
		findOptions,
	)
	return cursor, err
}
func GetFriends(mid string) (friendList *mongo.Cursor, err error) {

	filter := bson.D{
		{"mid", mid},
	}
	opts := options.Find().SetProjection(bson.D{{"fid", 1}, {"remark", 1}, {"friendName", 1}, {"friendAvatar", 1}, {"_id", 0}})
	cursor, err := initDB.MONGODB_DB.Database("chat").Collection("session").Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	return cursor, err
}
