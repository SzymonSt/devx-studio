import React, { useState } from "react";
import "./Question.css"
import {Form} from "react-bootstrap";
import {RiDeleteBin5Line} from 'react-icons/ri';

type QuestionProps = {
    questionContent: string;
    id: string;
    onRemove: (id: string) => void;
}

const Question: React.FC<QuestionProps> = ({questionContent, id, onRemove}) => {
    const [questionContentState, setQuestionContent] = useState("");
    const [isCountetdToOverallScore, setIsCountetdToOverallScore] = useState(true);
    const handleRemove = () => {
        onRemove(id);
    }
    return (
        <div className="question">
            <div className="question-header">
                <button className="btn bin-icon" onClick={handleRemove}><RiDeleteBin5Line size={25} /></button>
                <Form.Control
                    placeholder={questionContent}
                    value={questionContentState}
                    onChange={(event) => setQuestionContent(event.target.value)}>
                </Form.Control>
            </div>
            <Form.Check type="switch" label="Count into overall vertical score" checked={isCountetdToOverallScore} onChange={(event) => setIsCountetdToOverallScore(event.target.checked)}/>
            <Form.Range />
        </div>
    );
};

export default Question;