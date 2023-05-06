import React, {useState} from "react";
import { v4 as uuidv4 } from 'uuid';
import Question from "../../survey-components/Question";
import { Dropdown, Form } from "react-bootstrap";


type ScheduledSurveyProps = {
    name: string;
    index: number;
    onRemove: (i: number) => void;
}

const ScheduledSurvey: React.FC<ScheduledSurveyProps> = ({name, index, onRemove}) => {
    const [questions, setQuestions] = useState<JSX.Element[]>([]);
    const [selectedScheduleOption, setSelectedScheduleOption] = useState<string>("");

    const handleSurveyRemoval = () => {
        console.log("Removing survey at index: " + index);
        onRemove(index);
    }
    const addQuestion = () => {
        var questionContent = "New Question " + (questions.length+1);
        setQuestions([...questions, <Question questionContent={questionContent} id={uuidv4()} onRemove={removeQuestion} />]);
    }
    const removeQuestion = (index: string) => {
        setQuestions(questions => questions.filter((question, i) => question.props.id !== index));
    }
    const renderQuestions = () => {
        return questions.map((question, index) => {
            return (
                <div className="source" key={index}>
                    <div className="source-body">
                        {question}
                    </div>
                </div>
            );
        });
    }
    return (
        <div>
            <h3>{name}</h3>
            <span className="source-remove" onClick={handleSurveyRemoval}>X</span>
            <span className="question-add" onClick={() => addQuestion()}>+</span>
            <div className="form-group">
                {renderQuestions()}
            </div>
            <div>
                <h4>Schedule</h4>
                <Form.Select value={selectedScheduleOption} onChange={(event) => setSelectedScheduleOption(event.target.value)}>
                    <option value="">--Select an option--</option>
                    <option value="0 0 */7 * *">Every week</option>
                    <option value="0 0 */14 * *">Every 2 weeks</option>
                    <option value="0 0 1 * *">Every month</option>
                    <option value="0 0 1 */3 *">Every 3 months </option>
                </Form.Select>
            </div>
        </div>
    );
};

export default ScheduledSurvey;