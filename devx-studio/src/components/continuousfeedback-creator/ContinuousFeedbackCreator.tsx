import React from "react";
import "./ContinuousFeedbackCreator.css";
import ReactDOM from 'react-dom';
import Dropdown from 'react-bootstrap/Dropdown';

type ContinuousFeedbackCreatorProps = {
    isVisible: boolean;
}

const ContinuousFeedbackCreator: React.FC<ContinuousFeedbackCreatorProps> = ({isVisible}) => {
    if (!isVisible) {
        return null;
    }
    return ReactDOM.createPortal(
        <div className="cf-creator">
            <Dropdown>
                <Dropdown.Toggle variant="success" id="dropdown-basic">
                    Dropdown Button
                </Dropdown.Toggle>

                <Dropdown.Menu>
                    <Dropdown.Item href="#/action-1">Action</Dropdown.Item>
                    <Dropdown.Item href="#/action-2">Another action</Dropdown.Item>
                    <Dropdown.Item href="#/action-3">Something else</Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>
            <div className="sources">
            </div>
        </div>, document.getElementById('root')! as HTMLElement
    );
};

// function ContinuousFeedbackCreator() {

// }

export default ContinuousFeedbackCreator;