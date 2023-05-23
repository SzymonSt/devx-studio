export interface ScoreData {
  verticalId: string;
  overallScore: Score;
  surveyScores: SurveyScore[];
}

interface Score {
    mean: number;
    percentile95: number;
    percentile99: number;
}

interface SurveyScore {
    surveyName: string;
    surveyId: string;
    cfId: string;
    cfName: string;
    questionScores: QuestionScore[];
}

interface QuestionScore {
    questionId: string;
    questionContent: string;
    score: Score;
}