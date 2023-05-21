import {rest} from 'msw'

const apiUri = process.env.REACT_APP_API_URI;

export const mockResponseGetContinuousFeedbacks = [
    {
    id: "4567899",
    name: "Continuous Feedback 1 - Infrastructure",
    verticalId: "infrastucture",
    isCurrentlyActive: true,
    responseRate: 0.66,
    scheduledSurveys: [
        {
            id: "cgf4567",
            name: "IaC",
            lastOpened: "2021-01-01T00:00:00.000Z",
            openPeriod: "24h",
            interval: "",
            responseRate: 0.77,
            audience: [
                {
                    id: "3456789",
                }
            ],
            questions: [
                {
                    id: "1",
                    question: "How would you rate the overall ease of using the IaC modules?",
                    isCalculatedInOverallScore: true
                },
                {
                    id: "2",
                    question: "How would you rate the quality and coverage of the IaC modules wiki/documentation?",
                    isCalculatedInOverallScore: true
                },
                {
                    id: "3",
                    question: "How would you rate covergare of cloud provider resources by the IaC modules ?",
                    description: "Are there any missing resources? Are there any modules that are not up to date?",
                    isCalculatedInOverallScore: true
                }
            ]
        },
        {
            id: "hjg4527",
            name: "Infrastructure observability",
            lastOpened: "2021-01-01T00:00:00.000Z",
            openPeriod: "48h",
            interval: "",
            responseRate: 0.56,
            audience: [
                {
                    id: "3a563h4",
                }
            ],
            questions: [
                {
                    id: "1",
                    question: "How would you rate quality and accuracy of infrastructure alerts?",
                    isCalculatedInOverallScore: true
                },{
                    id: "2",
                    question: "How would you rate UX of infrastructure metrics dashboards?",
                    description: "Are all dashabords easy to find and use? Are there any dahsbaords missing?",
                    isCalculatedInOverallScore: false
                },
                {
                    id: "3",
                    question: "How would you rate possibility to create custom infrastructure dashboards?",
                    description: "Is it easy to create custom dashboards? Are there any missing details in the wiki? Do you know that you can create custom dashboard as code?",
                    isCalculatedInOverallScore: true
                }
            ]
        }
    ],
    eventSurveys: [],
    integrationSurveys: []
},
{
    id: "45678ij",
    name: "Continuous Feedback 2 - CI/CD",
    verticalId: "infrastucture",
    isCurrentlyActive: false,
    responseRate: 0.78,
    scheduledSurveys: [
        {
            id: "cgf4567",
            name: "IaC",
            lastOpened: "2021-01-01T00:00:00.000Z",
            openPeriod: "24h",
            interval: "",
            responseRate: 0.77,
            audience: [
                {
                    id: "3456789",
                }
            ],
            questions: [
                {
                    id: "1",
                    question: "How would you rate the overall ease of using the IaC modules?",
                    isCalculatedInOverallScore: true
                },
                {
                    id: "2",
                    question: "How would you rate the quality and coverage of the IaC modules wiki/documentation?",
                    isCalculatedInOverallScore: true
                },
                {
                    id: "3",
                    question: "How would you rate covergare of cloud provider resources by the IaC modules ?",
                    description: "Are there any missing resources? Are there any modules that are not up to date?",
                    isCalculatedInOverallScore: true
                }
            ]
        },
        {
            id: "hjg4527",
            name: "Infrastructure observability",
            lastOpened: "2021-01-01T00:00:00.000Z",
            openPeriod: "48h",
            interval: "",
            responseRate: 0.56,
            audience: [
                {
                    id: "3a563h4",
                }
            ],
            questions: [
                {
                    id: "1",
                    question: "How would you rate quality and accuracy of infrastructure alerts?",
                    isCalculatedInOverallScore: true
                },{
                    id: "2",
                    question: "How would you rate UX of infrastructure metrics dashboards?",
                    description: "Are all dashabords easy to find and use? Are there any dahsbaords missing?",
                    isCalculatedInOverallScore: false
                },
                {
                    id: "3",
                    question: "How would you rate possibility to create custom infrastructure dashboards?",
                    description: "Is it easy to create custom dashboards? Are there any missing details in the wiki? Do you know that you can create custom dashboard as code?",
                    isCalculatedInOverallScore: true
                }
            ]
        }
    ],
    eventSurveys: [],
    integrationSurveys: []
},
{
    id: "4hk78ij",
    name: "Continuous Feedback 1 - Knowledge sharing",
    verticalId: "infrastucture",
    isCurrentlyActive: true,
    responseRate: 0.33,
    scheduledSurveys: [
        {
            id: "cgf4567",
            name: "IaC",
            lastOpened: "2021-01-01T00:00:00.000Z",
            openPeriod: "24h",
            interval: "",
            responseRate: 0.77,
            audience: [
                {
                    id: "3456789",
                }
            ],
            questions: [
                {
                    id: "1",
                    question: "How would you rate the overall ease of using the IaC modules?",
                    isCalculatedInOverallScore: true
                },
                {
                    id: "2",
                    question: "How would you rate the quality and coverage of the IaC modules wiki/documentation?",
                    isCalculatedInOverallScore: true
                },
                {
                    id: "3",
                    question: "How would you rate covergare of cloud provider resources by the IaC modules ?",
                    description: "Are there any missing resources? Are there any modules that are not up to date?",
                    isCalculatedInOverallScore: true
                }
            ]
        },
        {
            id: "hjg4527",
            name: "Infrastructure observability",
            lastOpened: "2021-01-01T00:00:00.000Z",
            openPeriod: "48h",
            interval: "",
            responseRate: 0.56,
            audience: [
                {
                    id: "3a563h4",
                }
            ],
            questions: [
                {
                    id: "1",
                    question: "How would you rate quality and accuracy of infrastructure alerts?",
                    isCalculatedInOverallScore: true
                },{
                    id: "2",
                    question: "How would you rate UX of infrastructure metrics dashboards?",
                    description: "Are all dashabords easy to find and use? Are there any dahsbaords missing?",
                    isCalculatedInOverallScore: false
                },
                {
                    id: "3",
                    question: "How would you rate possibility to create custom infrastructure dashboards?",
                    description: "Is it easy to create custom dashboards? Are there any missing details in the wiki? Do you know that you can create custom dashboard as code?",
                    isCalculatedInOverallScore: true
                }
            ]
        }
    ],
    eventSurveys: [],
    integrationSurveys: []
}
]

export const cfhandlers = [
    rest.get(apiUri + '/continuousfeedback', (req, res, ctx) => {
        return res(ctx.json(mockResponseGetContinuousFeedbacks))
    }),
    rest.get(apiUri + '/continuousfeedback/:id', (req, res, ctx) => {
        const { id } = req.params
        const cf = mockResponseGetContinuousFeedbacks.find(cf => cf.id === id)
        return res(ctx.json(cf))
    })
]