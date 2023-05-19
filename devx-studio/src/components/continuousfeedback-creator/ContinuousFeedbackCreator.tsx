import React, {useState, useRef, useEffect} from "react";
import "./ContinuousFeedbackCreator.css";
import ReactDOM from 'react-dom';
import Dropdown from 'react-bootstrap/Dropdown';
import DataSource from "../continuousfeedback-components/DataSource";
import { Form, FormText } from "react-bootstrap";
import { ContinuousFeedback, ScheduledSurvey } from "../../interafces/continuousfeedback";

type ContinuousFeedbackCreatorProps = {
    isVisible: boolean;
    feedback?: ContinuousFeedback;
    visibilityHandler: (isVisible: boolean) => void;
}

const ContinuousFeedbackCreator: React.FC<ContinuousFeedbackCreatorProps> = ({isVisible, feedback, visibilityHandler}) => {
    const apiUri = process.env.REACT_APP_API_URI;
    const childRefs = useRef<Array<any>>([]);
    const [sources, setSources] = useState<JSX.Element[]>([]);
    const [scheduledSurveys, setScheduledSurveys] = useState<ScheduledSurvey[]>([]);
    const [selectedVerticalOption, setSelectedVerticalOption] = useState<string>("");
    const [continuousfeedback, setContinuousFeedback] = useState<ContinuousFeedback>({} as ContinuousFeedback);

    const handleVerticalChange = (event: React.ChangeEvent<HTMLSelectElement>) => {   
        childRefs.current.forEach((childRef) => {
            childRef.current.handleVerticalChange(event);
        });
        setSelectedVerticalOption(event.target.value);
    }

    const addSource = (type: string) => {
        var name  = "New_" + type + (sources.length);
        setScheduledSurveys([...scheduledSurveys, 
            {
                id: "",
                name: name,
                lastOpened: "",
                openPeriod: "",
                interval: "",
                responseRate: 0.22,
                audience: [],
                questions: [],
            }
        ]);
    }

    const removeSource = (index: string) => {
    }

    const renderScheduledSurveys = () => {
        return scheduledSurveys.map((survey, index) => {
            return (
                <DataSource type={"scheduledsurvey"} name={survey.name} 
                index={survey.id}
                onRemove={removeSource} />
            );
        });
    }

    const saveCF = async() => {
        console.log(continuousfeedback);
        continuousfeedback.scheduledSurveys = scheduledSurveys;
        continuousfeedback.verticalId = selectedVerticalOption;
        continuousfeedback.id ? updateCF() : createCF();
        closeCFComponent();
    }

    const updateCF = async() => {
        console.log("update")
        const resp = await fetch(apiUri + "/continuousfeedback/" + continuousfeedback.id ,{
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(continuousfeedback)
        });
        resp.ok ? console.log(resp.status + "|" + resp.statusText) : console.log("ERROR: " + resp.status + "|" + resp.statusText);
    }

    const createCF = async() => {
        console.log("create")
        const resp = await fetch(apiUri + "/continuousfeedback/",{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(continuousfeedback)
        });
        resp.ok ? console.log(resp.status + "|" + resp.statusText) : console.log("ERROR: " + resp.status + "|" + resp.statusText);
    }

    const closeCFComponent = () => {
        visibilityHandler(false);
    }

    useEffect(() => {
        if(feedback?.id != undefined){
            setSelectedVerticalOption(feedback.verticalId);
            setContinuousFeedback(feedback);
            setScheduledSurveys(feedback.scheduledSurveys);
        }else {
            setSelectedVerticalOption("");
            setContinuousFeedback({name: "New Continuous Feedback"} as ContinuousFeedback); 
            setScheduledSurveys([]);
        }

      }, [isVisible]);

    if (!isVisible) {
        return null;
    }
    return ReactDOM.createPortal(
        <div className="cf-creator">
            <div className="cf-cretor-header">
            <h2>{continuousfeedback.name}</h2>
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
            <button onClick={()=> saveCF()} className="btn btn-success save">Save</button>
            </div>
            <div className="sources">
                {renderScheduledSurveys()}
            </div>
        </div>, document.getElementById('root')! as HTMLElement
    );
};
export default ContinuousFeedbackCreator;