import React from "react";
import { Card } from "react-bootstrap";
import { ContinuousFeedback } from "../../../interafces/continuousfeedback";


type ContinuousFeedbackCardProps = {
    feedback: ContinuousFeedback;
    toggleModalVisibility: (focusedFeedback: ContinuousFeedback) => void;
}

const renderActiveState = (isActive: boolean) => {
    if(isActive){
        return "Active";
    }
    else{
        return "Inactive";
    }
}


const ContinuousFeedbackCard: React.FC<ContinuousFeedbackCardProps> = ({feedback, toggleModalVisibility}) => {
    return (
        <div>
            <Card onClick={()=>toggleModalVisibility(feedback)} className="feedback" key={feedback.id} id={feedback.id}>
            <Card.Title>{feedback.name}</Card.Title>
            {renderActiveState(feedback.isCurrentlyActive)}
            <Card.Text>
                <span>Response Rate: {feedback.responseRate * 100}%</span>
            </Card.Text>
            </Card>
        </div>
    );
};

export default ContinuousFeedbackCard;
