import React from "react";

const TaskCard = (props) => {
    const handleClick = () => {
        props.deleteTask(props.index)
    }

    return (
    <div>
        <p>{props.date}</p>
        <p>{props.body}</p>
        <p>{props.is_finished ? "true": "false"}</p>
        <button onClick={handleClick}>削除</button>
    </div>
)};

export default TaskCard;
