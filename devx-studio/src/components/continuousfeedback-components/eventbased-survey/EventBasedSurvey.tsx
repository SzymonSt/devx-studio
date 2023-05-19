import React from "react";

type EventBasedSurveyProps = {
    name: string;
}

const EventBasedSurvey: React.FC<EventBasedSurveyProps> = ({name}) => {
    return (
        <div>{name}</div>
    );
};

export default EventBasedSurvey;