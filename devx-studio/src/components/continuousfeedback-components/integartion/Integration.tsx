import React from "react";

type IntegrationProps = {
    name: string;
}

const Integration: React.FC<IntegrationProps> = ({name}) => {
    return (
        <div>{name}</div>
    );
};

export default Integration;