import React from "react";
import Form from "react-bootstrap/esm/Form";

type QuestionProps = {
    questionContent: string;
}

const Question: React.FC<QuestionProps> = ({questionContent}) => {
    return (
        <div className="question">
            <Form.Label><p>{questionContent}</p></Form.Label>
            <Form.Range />
        </div>
    );
};

export default Question;