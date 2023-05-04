import React, { useState } from "react";
import "./ContinuousFeedback.css";
import PageHeader from "../components/page-header/PageHeader";
import ContinuousFeedbackCreator from "../components/continuousfeedback-creator/ContinuousFeedbackCreator";

function ContinuousFeedback() {
    const [isModalVisible, setModalVisibility] = useState(false);
  
    const toggleModal = () => {
        setModalVisibility(!isModalVisible);
    }
    return (
        <div>
            <PageHeader title="Continuous Feedback" />
            <div>
                <button onClick={toggleModal} className="btn">Create New Stream</button>
                <ContinuousFeedbackCreator isVisible={isModalVisible} />
            </div>
            <div className="streams">

            </div>
        </div>
    );
    }

export default ContinuousFeedback;