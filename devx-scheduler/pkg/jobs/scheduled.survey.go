package jobs

import (
	"fmt"
)

func ProcessScheduledSurvey(cfId string, surveyId string) {
	fmt.Println("===Processing scheduled survey===")
	// client := http.Client{}
	// var continuousFeedback map[string]interface{}
	// cfIDParsed, err := primitive.ObjectIDFromHex(cfId)
	// if err != nil {
	// 	entry := bson.M{
	// 		"log":       "ERROR",
	// 		"timestamp": primitive.DateTime(primitive.NewDateTimeFromTime(time.Now())),
	// 		"error":     err.Error(),
	// 	}
	// 	_, err = mongoClient.Database("devx-scheduler").Collection("cronlog").InsertOne(context.Background(), entry)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	// filter := bson.M{"_id": cfIDParsed}
	// err = mongoClient.Database("devx").Collection("continuousfeedback").FindOne(context.Background(), filter).Decode(&continuousFeedback)
	// if err != nil {
	// 	entry := bson.M{
	// 		"log":       "ERROR",
	// 		"timestamp": primitive.DateTime(primitive.NewDateTimeFromTime(time.Now())),
	// 		"error":     err.Error(),
	// 	}
	// 	_, err = mongoClient.Database("devx-scheduler").Collection("cronlog").InsertOne(context.Background(), entry)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	// var survey map[string]interface{}
	// for _, s := range continuousFeedback["scheduledSurveys"].(primitive.A) {
	// 	surveyTmp := s.(map[string]interface{})
	// 	if surveyTmp["_id"] == surveyId {
	// 		survey = surveyTmp
	// 		break
	// 	}
	// }
	// if survey != nil {
	// 	surveyUrl := fmt.Sprintf("http://localhost:3000/survey/%s/%s", cfId, surveyId)
	// 	method := "POST"
	// 	url := slackHook
	// 	body := map[string]interface{}{
	// 		"channel": "@szymonst2808",
	// 		"pretext": "Hello! We do our best to improve your develper experience. Please take a few minutes for a feedback:",
	// 		"text":    fmt.Sprintf("<%s|Click here to start the survey>", surveyUrl),
	// 	}
	// 	byts, _ := json.Marshal(body)
	// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(byts))
	// 	if err != nil {
	// 		entry := bson.M{
	// 			"log":       "ERROR",
	// 			"timestamp": primitive.DateTime(primitive.NewDateTimeFromTime(time.Now())),
	// 			"error":     err.Error(),
	// 		}
	// 		_, err = mongoClient.Database("devx-scheduler").Collection("cronlog").InsertOne(context.Background(), entry)
	// 		if err != nil {
	// 			return
	// 		}
	// 	}
	// 	req.Header.Set("Content-Type", "application/json")
	// 	res, err := client.Do(req)
	// 	if err != nil {
	// 		entry := bson.M{
	// 			"log":       "ERROR",
	// 			"timestamp": primitive.DateTime(primitive.NewDateTimeFromTime(time.Now())),
	// 			"error":     err.Error(),
	// 		}
	// 		_, err = mongoClient.Database("devx-scheduler").Collection("cronlog").InsertOne(context.Background(), entry)
	// 		if err != nil {
	// 			return
	// 		}
	// 	}
	// 	if res.StatusCode == 200 {
	// 		entry := bson.M{
	// 			"log":       fmt.Sprintf("OK! Survey %s processed", survey["name"]),
	// 			"timestamp": primitive.DateTime(primitive.NewDateTimeFromTime(time.Now())),
	// 		}
	// 		_, err = mongoClient.Database("devx-scheduler").Collection("cronlog").InsertOne(context.Background(), entry)
	// 		if err != nil {
	// 			return
	// 		}
	// 	}
	// 	defer res.Body.Close()

	// }
}
