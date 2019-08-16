package collections

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Collection represents groups of words
type Collection struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
	Date time.Time          `bson:"date" json:"date,omitempty"`
}

// type UserModel struct {
// 	ID           uint    `gorm:"primary_key"`
// 	Username     string  `gorm:"column:username"`
// 	Email        string  `gorm:"column:email;unique_index"`
// 	Bio          string  `gorm:"column:bio;size:1024"`
// 	Image        *string `gorm:"column:image"`
// 	PasswordHash string  `gorm:"column:password;not null"`
// }

// type FollowModel struct {
// 	gorm.Model
// 	Following    UserModel
// 	FollowingID  uint
// 	FollowedBy   UserModel
// 	FollowedByID uint
// }
