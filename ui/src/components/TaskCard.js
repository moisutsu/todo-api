import React from "react";

const TaskCard = (props) => (
    <div>
        Cardの始まり
        <p>{props.date}</p>
        <p>{props.body}</p>
        <p>{props.is_finished ? "true": "false"}</p>
        Cardの終わり
    </div>
);

export default TaskCard;
