import React from "react";

const TaskCard = (props) => (
    <div>
        <p>{props.date}</p>
        <p>{props.body}</p>
        <p>{props.is_finished ? "true": "false"}</p>
    </div>
);

export default TaskCard;
