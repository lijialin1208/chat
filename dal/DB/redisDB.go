package DB

import (
	"chat/dal/initDB"
	"context"
	"fmt"
	"github.com/hertz-contrib/websocket"
	"github.com/redis/go-redis/v9"
)

func StorageMacAndSession(macAddress string, session *websocket.Conn) error {
	i := &session
	Session := fmt.Sprintf("%p", i)
	intCmd := initDB.REDIS_DB.Set(context.Background(), macAddress, Session, 0)
	return intCmd.Err()
}
func GetMacAndSession(mac string) *redis.StringCmd {
	cmd := initDB.REDIS_DB.Get(context.Background(), mac)
	return cmd
}

func StorageUserIDAndMac(userID string, mac string) error {
	intCmd := initDB.REDIS_DB.SAdd(context.Background(), userID, mac)
	return intCmd.Err()
}
func DeleteUserIDAndMac(userID string, mac string) error {
	intCmd := initDB.REDIS_DB.SRem(context.Background(), userID, mac)
	return intCmd.Err()
}
func GetUsersMac(userID string) *redis.StringSliceCmd {
	members := initDB.REDIS_DB.SMembers(context.Background(), userID)
	return members
}
