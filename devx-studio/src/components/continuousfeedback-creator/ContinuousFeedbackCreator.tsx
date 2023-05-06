import React, {useState, useRef} from "react";
import "./ContinuousFeedbackCreator.css";
import ReactDOM from 'react-dom';
import Dropdown from 'react-bootstrap/Dropdown';
import DataSource from "../continuousfeedback-components/DataSource";
import { Form } from "react-bootstrap";

type ContinuousFeedbackCreatorProps = {
    isVisible: boolean;
}

const ContinuousFeedbackCreator: React.FC<ContinuousFeedbackCreatorProps> = ({isVisible}) => {
    const childRefs = useRef<Array<any>>([]);
    const [sources, setSources] = useState<JSX.Element[]>([]);
    const [selectedVerticalOption, setSelectedVerticalOption] = useState<string>("");

    const handleVerticalChange = (event: React.ChangeEvent<HTMLSelectElement>) => {   
        childRefs.current.forEach((childRef) => {
            childRef.current.handleVerticalChange(event);
        });
        setSelectedVerticalOption(event.target.value);
    }

    const addSource = (type: string) => {
        var name  = "New_" + type + (sources.length);
        setSources([...sources, 
            <DataSource type={type} name={name} 
                    index={sources.length}
                    onRemove={removeSource} />]);
    }

    const removeSource = (index: number) => {
        const newSources = [...sources];
        newSources.splice(index, 1);
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
            <Form.Select className="vertical-select" value={selectedVerticalOption} onChange={handleVerticalChange}>
                    <option value="">--Select a vertical--</option>
                    <option value="sdlcprocess">SDLC Process</option>
                    <option value="tooling-proficiency">Tooling proficiency</option>
                    <option value="onboarding">Onboarding</option>
                    <option value="communication">Communication</option>
                    <option value="knowledge-sharing">Knowledge Sharing</option>
                    <option value="growth">Growth/Upskill</option>
                    <option value="wellbeing">Wellbeing</option>
            </Form.Select>
            <button className="btn btn-success save">Save</button>
            </div>
            <div className="sources">
                {renderSources()}
            </div>
        </div>, document.getElementById('root')! as HTMLElement
    );
};
export default ContinuousFeedbackCreator;