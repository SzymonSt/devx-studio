import React, {useState} from "react";
import "./ContinuousFeedbackCreator.css";
import ReactDOM from 'react-dom';
import Dropdown from 'react-bootstrap/Dropdown';
import DataSource from "../continuousfeedback-components/DataSource";

type ContinuousFeedbackCreatorProps = {
    isVisible: boolean;
}

const ContinuousFeedbackCreator: React.FC<ContinuousFeedbackCreatorProps> = ({isVisible}) => {
    const [sources, setSources] = useState<JSX.Element[]>([]);

    const addSource = (type: string) => {
        var name  = "New_" + type + (sources.length+1);
        setSources([...sources, <DataSource type={type} name={name} index={sources.length} onRemove={removeSource} />]);
    }

    const removeSource = (index: number) => {
        console.log("Removing source at index: " + index);
        const newSources = [...sources];
        var deleted = newSources.splice(index, 1);
        console.log("Deleted: " + deleted.length);
        setSources(newSources);
    }

    const renderSources = () => {
        return sources.map((source, index) => {
            return (
                <div className="source" key={index}>
                    <div className="source-body">
                        {source}
                    </div>
                </div>
            );
        });
    }

    if (!isVisible) {
        return null;
    }
    return ReactDOM.createPortal(
        <div className="cf-creator">
            <div className="cf-cretor-header">
            <Dropdown>
                <Dropdown.Toggle variant="primary" id="dropdown-basic">
                    Datasource
                </Dropdown.Toggle>

                <Dropdown.Menu>
                    <Dropdown.Item onClick={() => addSource("scheduledsurvey")}>Scheduled Survey</Dropdown.Item>
                    <Dropdown.Item disabled>Event Based Survey </Dropdown.Item>
                    <Dropdown.Item disabled>Integration</Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>
            <button className="btn btn-success save">Save</button>
            </div>
            <div className="sources">
                {renderSources()}
            </div>
        </div>, document.getElementById('root')! as HTMLElement
    );
};
export default ContinuousFeedbackCreator;