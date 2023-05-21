
export interface Answer{
    verticalId: string;
    continuousFeedbackParentId: string;
    continuousFeedbackName: string;
    surveyId: string;
    surveyName: string;
    timestamp: string;
    questions: AnswerQuestion[];
}

export interface AnswerQuestion{
    questionId: string;
    question: string;
    description?: string;
    score: number;
    isCalculatedInOverallScore: boolean;
}