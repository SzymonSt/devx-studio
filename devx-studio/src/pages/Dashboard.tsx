import React from "react";
import "./Dashboard.css";
import LineChart from "../components/chart/Chart";
import PageHeader from "../components/page-header/PageHeader";

function Dashboard() {
    const vericals = ["infrastructure"];
    return(
        <div className="dashboard-content">
            <PageHeader title="Dashboard" />
            <div className="verticals">
                {vericals.map((vertical) => (
                    <div className="vertical">
                        <LineChart verticalId={vertical} />
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Dashboard;