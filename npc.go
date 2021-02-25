package generator

import "github.com/lib/pq"

// Npc defines a non-player character's traits
type Npc struct {
	ID             string         `json:"id" gorm:"not null"`
	UserID         string         `json:"userId" binding:"required"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Gender         string         `json:"gender"`
	Pronouns       string         `json:"pronouns"`
	Height         string         `json:"height"`
	Hook           string         `json:"hook"`
	Job            string         `json:"job"`
	Languages      pq.StringArray `json:"languages" gorm:"type:varchar(64)[]"`
	LifeStage      string         `json:"lifeStage"`
	NegativeTraits pq.StringArray `json:"negativeTraits" gorm:"type:varchar(64)[]"`
	NeutralTraits  pq.StringArray `json:"neutralTraits" gorm:"type:varchar(64)[]"`
	PositiveTraits pq.StringArray `json:"positiveTraits" gorm:"type:varchar(64)[]"`
	Quirk          string         `json:"quirk"`
	Race           string         `json:"race"`
	Wealth         string         `json:"wealth"`
}
