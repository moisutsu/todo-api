import React from "react";

const TaskCard = (props) => {
    const handleDeleteClick = () => {
        props.deleteTask(props.index)
    }
    const handleCompleteClick = () => {
        props.putTask(props.index, props.date, props.body, !props.is_finished)
    }

    const CompleteButton = props.is_finished ? <button onClick={handleCompleteClick}>未完了</button> : <button onClick={handleCompleteClick}>完了</button>

    return (
    <div>
        <p>{props.date}</p>
        <p>{props.body}</p>
        <p>{props.is_finished ? "true": "false"}</p>
        {CompleteButton}
        <button onClick={handleDeleteClick}>削除</button>
    </div>
)};

export default TaskCard;
