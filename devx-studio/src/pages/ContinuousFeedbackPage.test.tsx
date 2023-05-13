import React, { useEffect, useState } from "react";
import { act, render, screen} from "@testing-library/react";
import ContinuousFeedbackPage from "./ContinuousFeedbackPage";
import {server } from "../mocks/server";

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

test("renders ContinuousFeedbackPage with continuous feedbacks fetched from api", async () => {
    render(<ContinuousFeedbackPage />);
    const cfElement1Title = await screen.findByText("Continuous Feedback 1 - Infrastructure");
    const cfElement1IsActive = await screen.findByText("Active: true");
    const cfElement1ResponseRate = await screen.findByText("Response Rate: 66 %");

    expect(cfElement1Title).toBeInTheDocument();
    expect(cfElement1IsActive).toBeInTheDocument();
    expect(cfElement1ResponseRate).toBeInTheDocument();
});