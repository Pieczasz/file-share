package models

import "time"

type File struct {
    ID        string    `gorm:"primaryKey;type:uuid" json:"id"`
    Name      string    `json:"name" gorm:"not null"`
    Extension string    `json:"extension" gorm:"not null"`
    Path      string    `json:"path" gorm:"not null;unique"`
    Size      int64     `json:"size" gorm:"not null"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
    ExpiresAt time.Time `json:"expires_at" gorm:"index"`
}