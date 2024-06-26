import React, { useEffect, useState } from "react";
import "./ContinuousFeedbackPage.css";
import PageHeader from "../components/page-header/PageHeader";
import {ContinuousFeedback} from "../interafces/continuousfeedback";
import ContinuousFeedbackCreator from "../components/continuousfeedback-creator/ContinuousFeedbackCreator";
import { mockResponseGetContinuousFeedbacks } from "../mocks/continuousfeedback";
import ContinuousFeedbackCard from "../components/continuousfeedback-components/continuousfeedback-card/ContinuousFeedbackCard";

function ContinuousFeedbackPage() {
    const apiUri = process.env.REACT_APP_API_URI;
    const [isModalVisible, setModalVisibility] = useState(false);
    const [focusedFeedback, setFocusedFeedback] = useState<ContinuousFeedback>({} as ContinuousFeedback); //TODO: Change to null when not focused on a feedback
    const [continuousFeedbacks, setContinuousFeedbacks] = useState<ContinuousFeedback[]>([]);
  
    const toggleModal = (focusedFeedback: ContinuousFeedback) => {
        setModalVisibility(!isModalVisible);
        setFocusedFeedback(focusedFeedback);
    }
    const fetchContinuousFeedbacks = async() => {
        await fetch(apiUri + "/continuousfeedback/",{method: 'GET'})
        .then(res => res.json())
        .then(
            (result) => {
                setContinuousFeedbacks(result);
            }
            
        )
        .catch((error) => {
            console.log(error);
        })
        console.log(continuousFeedbacks);
    }

    useEffect(() => {
        //setContinuousFeedbacks(mockResponse);
        fetchContinuousFeedbacks();
      }, []);
    return (
        <div>
            <PageHeader title="Continuous Feedback" />
            <div>
                <button onClick={()=>toggleModal({} as ContinuousFeedback)} className="btn">Create New Stream</button>
                <ContinuousFeedbackCreator isVisible={isModalVisible} feedback={focusedFeedback} visibilityHandler={setModalVisibility}/>
            </div>
            <div className="feedbacks">
                {continuousFeedbacks.map((feedback) => (
                    <ContinuousFeedbackCard key={feedback.id} toggleModalVisibility={toggleModal} feedback={feedback} />
                    ))}
            </div>
        </div>
    );
    }

export default ContinuousFeedbackPage;