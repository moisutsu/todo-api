import React, {useState} from "react"

import Header from "./Header"
import TaskList from "./TaskList"

const App = () => {
    const [tasks, setTasks] = useState([])
    return (
        <React.Fragment>
            <Header />
            <TaskList tasks={tasks}/>
            <button onClick={() => setTasks(tasks.concat({
        date: "2019年12月1日",
        body: "Goの勉強",
        is_finished: false,
    }))}>
                Add Task
            </button>
        </React.Fragment>
    )
}

export default App;
