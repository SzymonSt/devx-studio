import React from "react";
import './PageHeader.css';

type PageHeaderProps = {
    title: string;
}

const PageHeader: React.FC<PageHeaderProps> = ({title}) => {
    return(
        <div className="page-header">
            <h1 className="header-text">{title}</h1>
        </div>
    );
}

export default PageHeader;