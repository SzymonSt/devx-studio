import React, {useState} from "react";
import Question from "../../survey-components/Question";


type ScheduledSurveyProps = {
    name: string;
    index: number;
    onRemove: (i: number) => void;
}

const ScheduledSurvey: React.FC<ScheduledSurveyProps> = ({name, index, onRemove}) => {
    const [questions, setQuestions] = useState<JSX.Element[]>([]);
    const handleSurveyRemoval = () => {
        onRemove(index);
    }
    const addQuestion = () => {
        var questionContent = "New Question " + (questions.length+1);
        setQuestions([...questions, <Question questionContent={questionContent} />]);
    }
    const removeQuestion = (index: number) => {
        const newQuestions = [...questions];
        newQuestions.splice(index, 1);
        setQuestions(newQuestions);
    }
    const renderQuestions = () => {
        return questions.map((question, index) => {
            return (
                <div className="source" key={index}>
                    <div className="source-body">
                        <span className="source-remove" onClick={() => removeQuestion(index)}>X</span>
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
        </div>
    );
};

export default ScheduledSurvey;