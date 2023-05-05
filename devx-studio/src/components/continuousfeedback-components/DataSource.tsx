import React from "react";
import ScheduledSurvey from "./scheduled-survey/ScheduledSurvey";
import EventBasedSurvey from "./eventbased-survey/EventBasedSurvey";
import Integration from "./integartion/Integration";

type DataSourceProps = {
    type: string;
    name: string;
    index: number;
    onRemove: (i: number) => void;
}

const DataSource: React.FC<DataSourceProps> = ({type, name, index, onRemove}) => {
    if (type === "scheduledsurvey") {
       return (
        <ScheduledSurvey name={name} index={index} onRemove={onRemove}/>
       );
    }
    else if (type === "eventbasedsurvey") {
        return (
            <EventBasedSurvey name={name} />
           );
    }
    else if (type === "integration") {
        return (
            <Integration name={name} />
           );
    }
    return (<div></div>);
};

export default DataSource;