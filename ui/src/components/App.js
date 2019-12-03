import React, {useState, useEffect} from "react"

import axios from "axios";

import Header from "./Header"
import TaskList from "./TaskList"

const App = () => {
    const [tasks, setTasks] = useState([])

    const fetchTasks = async () => {
        const response = await axios.get("http://localhost:8080/todo-api/tasks")
        setTasks(response.data)
    }

    useEffect(() => {
        fetchTasks()
    }, [])

    return (
        <React.Fragment>
            <Header />
            <TaskList tasks={tasks}/>
        </React.Fragment>
    )
}

export default App;
