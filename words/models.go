package words

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Word represents words in database
type Word struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CollectionID primitive.ObjectID `bson:"collection_id,omitempty" json:"collection_id"`
	Word         string             `bson:"word" json:"word"`
	Date         time.Time          `bson:"date" json:"date,omitempty"`
	Translation  string             `bson:"translation" json:"translation"`
	Description  string             `bson:"description" json:"description,omitempty"`
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
