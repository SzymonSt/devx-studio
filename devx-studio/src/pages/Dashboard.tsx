import React from "react";
import "./Dashboard.css";
import LineChart from "../components/chart/Chart";
import PageHeader from "../components/page-header/PageHeader";

function Dashboard() {
    const vericals = ["SDLC Process", "Onboarding", "Knowledge Sharing", "Communication", "Tooling proficiency", "Growth/Upskilling", "Wellbeing"];
    return(
        <div className="dashboard-content">
            <PageHeader title="Dashboard" />
            <div className="verticals">
                {vericals.map((vertical) => (
                    <div className="vertical">
                        <LineChart title={vertical} />
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Dashboard;