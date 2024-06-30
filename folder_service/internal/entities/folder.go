package entities

import (
	"time"
)

// omitempty 標籤用於JSON 或 BSON 編碼器，
// struct的值為零值 對於指針類型為 nil，於數字類型為 0，
// 對於字符串類型為空字符串，該字段應該被省略。
// MongoDB 中，無法直接設定預設值
// 方法：程式自己檢查處理/Triggers/validation rules/aggregation pipeline
type Folder struct {
	ID        string     `bson:"_id,omitempty"`
	BaseID    string     `bson:"base_id"`
	ParentID  string     `bson:"parent_id,omitempty"`
	Position  float64    `bson:"position"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	Data      FolderData `bson:"data"` // 前端custome任意儲存的資料(color/name et cetera)
	Type      string     `bson:"type"`
}

type FolderData struct {
	Name  string `bson:"name"`
	Color string `bson:"color"`
}

// Initialize sets the default values for a Folder
func (f *Folder) Initialize() {
	now := time.Now()
	f.CreatedAt = now
	f.UpdatedAt = now
	f.Type = "folder"
}

type Board struct {
	ID        string    `bson:"_id,omitempty"`
	BaseID    string    `bson:"base_id"`
	Position  float64   `bson:"position"`
	ParentID  string    `bson:"parent_id,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Tabs      []Tab     `bson:"tabs"`
	Data      BoardData `bson:"data"`
	Type      string    `bson:"type"` // 用于存储类型信息
}

type BoardData struct {
	Name  string `bson:"name"`
	Color string `bson:"color"`
}

type Tab struct {
	Permission Permission `bson:"permission"`
}

type Permission struct {
	Read     bool   `bson:"read"`
	UserList []User `bson:"userlist"`
}

type User struct {
	UserID string `bson:"userid"`
}