import React, { useEffect, useState } from "react";
import "./ContinuousFeedbackPage.css";
import PageHeader from "../components/page-header/PageHeader";
import {ContinuousFeedback} from "../interafces/continuousfeedback";
import ContinuousFeedbackCreator from "../components/continuousfeedback-creator/ContinuousFeedbackCreator";

function ContinuousFeedbackPage() {
    const apiUri = process.env.REACT_APP_API_URI;
    const [isModalVisible, setModalVisibility] = useState(false);
    const [continuousFeedbacks, setContinuousFeedbacks] = useState<ContinuousFeedback[]>([]);
  
    const toggleModal = () => {
        setModalVisibility(!isModalVisible);
    }
    const fetchContinuousFeedbacks = () => {
        fetch(apiUri + "/continuousfeedback",{method: 'GET'})
        .then(res => res.json())
        .then(
            (result) => {
                setContinuousFeedbacks(result);
            }
        )
        .catch((error) => {
            console.log(error);
        })
    }

    useEffect(() => {
        fetchContinuousFeedbacks();
      }, []);
    return (
        <div>
            <PageHeader title="Continuous Feedback" />
            <div>
                <button onClick={toggleModal} className="btn">Create New Stream</button>
                <ContinuousFeedbackCreator isVisible={isModalVisible} />
            </div>
            <div className="feedbacks">
                {continuousFeedbacks.map((feedback) => (
                    <div className="feedback" key={feedback.id}>
                        <div className="feedback-title">{feedback.name}</div>
                        <div className="feedback-isActive">Active: {feedback.isCurrentlyActive}</div>
                    </div>))}

            </div>
        </div>
    );
    }

export default ContinuousFeedbackPage;