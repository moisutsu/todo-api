import React from "react";
import TaskCard from "./TaskCard";

const TaskList = (props) => {
    return (
        <React.Fragment>
            {props.tasks.map((task) => <TaskCard {...task}/>)}
        </React.Fragment>
    )
};

export default TaskList;
