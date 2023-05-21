import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { ContinuousFeedback, ScheduledSurvey } from "../interafces/continuousfeedback";
import { mockResponseGetContinuousFeedbacks } from "../mocks/continuousfeedback";
import { Button, Form } from "react-bootstrap";
import "./SurveyPage.css";

function SurveyPage() {
  const apiUri = process.env.REACT_APP_API_URI;
  const [continuousFeedback, setContinuousFeedback] = useState<ContinuousFeedback>({} as ContinuousFeedback);
  const [survey, setSurvey] = useState<ScheduledSurvey>({} as ScheduledSurvey);
  const params = useParams();

  const fetchSurvey = async() => {
    await fetch(apiUri + "/continuousfeedback/" + params.cfId ,{method: 'GET'})
    .then(res => res.json())
    .then(
        (result) => {
            setContinuousFeedback(result);
            setSurvey(result.scheduledSurveys.filter((survey: any) => survey.id === params.surveyId)[0]);
        }    
    )
    .catch((error) => {
        console.log(error);
    })
  }

    useEffect(() => {
        // var cf = mockResponseGetContinuousFeedbacks.filter((cf) => cf.id === params.cfId)[0];
        // if (!cf) {
        //     return;
        // }
        // console.log(cf);
        // setContinuousFeedback(cf);
        // setSurvey(cf.scheduledSurveys.filter((survey) => survey.id === params.surveyId)[0]);
        fetchSurvey();
    }, []);

  if (survey?.id === undefined) {
    return <div>404</div>
  }

  return( <div>
    <h1>Continuous Feedback: {params.cfId}</h1>
    <h1>Survey: {params.surveyId}</h1>
    <h1>Name: {survey.name}</h1>
    <Form className="questions">
        {survey.questions.map((question) => (
            <Form.Group className="question-box" controlId={question.id}>
            <Form.Label class="main-question-text">{question.question}</Form.Label>
            <Form.Label class="sub-question-text">{question.description}</Form.Label>
            <Form.Range className="answer-range" />

            </Form.Group>
        ))}
        <Button className="submit-survey" variant="primary" type="submit">
            Submit
        </Button>
    </Form>
  </div>
  );
}

export default SurveyPage;