import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { ContinuousFeedback, ScheduledSurvey } from "../interafces/continuousfeedback";
import { Answer, AnswerQuestion } from "../interafces/answer";
import { mockResponseGetContinuousFeedbacks } from "../mocks/continuousfeedback";
import { Button, Form } from "react-bootstrap";
import "./SurveyPage.css";

function SurveyPage() {
  const apiUri = process.env.REACT_APP_API_URI;
  const [continuousFeedback, setContinuousFeedback] = useState<ContinuousFeedback>({} as ContinuousFeedback);
  const [survey, setSurvey] = useState<ScheduledSurvey>({} as ScheduledSurvey);
  const params = useParams();

  const submitAnswer = async() =>{
    var answer = {
      verticalId: continuousFeedback.verticalId,
      surveyId: survey.id,
      continuousFeedbackName: continuousFeedback.name,
      continuousFeedbackParentId: continuousFeedback.id,
      surveyName: survey.name,
      timestamp:  new Date().toISOString(),
      questions: survey.questions.map((question) => {
        return {
          questionId: question.id,
          question: question.question,
          score: (question.score / 100),   
        } as AnswerQuestion
      })
    } as Answer;
    console.log(answer);
    await fetch(apiUri + "/continuousfeedback/answer" ,{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(answer)
    })
    .then(res => res.json())
    .then(
        (result) => {
            console.log(result);
        }    
    )
    .catch((error) => {
        console.log(error);
    })
  }

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

  const updateAnswer = (event: React.ChangeEvent<HTMLInputElement>) => {
    survey.questions.filter((question) => question.id === event.target.id)[0].score = parseFloat(event.target.value);
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
            <Form.Group key={question.id} className="question-box" controlId={question.id}>
            <Form.Label class="main-question-text">{question.question}</Form.Label>
            <Form.Label class="sub-question-text">{question.description}</Form.Label>
            <Form.Range onChange={updateAnswer} className="answer-range" />

            </Form.Group>
        ))}
        <Button onClick={submitAnswer} className="submit-survey" variant="primary" >
            Submit
        </Button>
    </Form>
  </div>
  );
}

export default SurveyPage;