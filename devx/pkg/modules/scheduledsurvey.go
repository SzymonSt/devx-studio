package modules

type ScheduledSurvey struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	//...
}
