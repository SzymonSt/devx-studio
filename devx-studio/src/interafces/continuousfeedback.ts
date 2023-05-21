export interface ContinuousFeedback{
    id: string;
    name: string;
    verticalId: string;
    isCurrentlyActive: boolean;
    responseRate: number;
    scheduledSurveys: ScheduledSurvey[];
    eventSurveys: EventSurvey[];
    integrationSurveys: IntegrationSurvey[];
}

export interface ScheduledSurvey{
    id: string;
    name: string;
    lastOpened: string;
    openPeriod: string;
    interval: string;
    responseRate: number;
    audience: AudienceItem[];
    questions: Question[];
}
export interface EventSurvey{}
export interface IntegrationSurvey{}

export interface AudienceItem{
    id: string;
}
export interface Question{
    id: string;
    question: string;
    description?: string;
    score: number;
    isCalculatedInOverallScore: boolean;
}
