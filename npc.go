package generator

// Npc defines a non-player character's traits
type Npc struct {
	ID             string   `json:"id" binding:"required"`
	UserID         string   `json:"userId" binding:"required"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Gender         string   `json:"gender"`
	Pronouns       string   `json:"pronouns"`
	Height         string   `json:"height"`
	Hook           string   `json:"hook"`
	Job            string   `json:"job"`
	Languages      []string `json:"languages" gorm:"type:varchar(64)[]"`
	LifeStage      string   `json:"lifeStage"`
	NegativeTraits []string `json:"negativeTraits" gorm:"type:varchar(64)[]"`
	NeutralTraits  []string `json:"neutralTraits" gorm:"type:varchar(64)[]"`
	PositiveTraits []string `json:"positiveTraits" gorm:"type:varchar(64)[]"`
	Quirk          string   `json:"quirk"`
	Race           string   `json:"race"`
	Wealth         string   `json:"wealth"`
}
