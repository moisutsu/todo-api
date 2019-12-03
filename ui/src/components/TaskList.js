import React from "react";
import TaskCard from "./TaskCard";

const TaskList = (props) => {
    return (
        <React.Fragment>
            {props.tasks.map((task, index) => <TaskCard {...task} key={index}/>)}
        </React.Fragment>
    )
};

export default TaskList;
