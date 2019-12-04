import React, {useState, useEffect} from "react"

import axios from "axios";

import Header from "./Header"
import TaskList from "./TaskList"
import PostBar from "./PostBar"

const App = () => {
    const [tasks, setTasks] = useState([])
    const [post, setPost] = useState("")

    const fetchTasks = async () => {
        const response = await axios.get("http://localhost:8080/todo-api/tasks")
        setTasks(response.data)
    }

    const postTask = async (body, date) => {
        await axios.post("http://localhost:8080/todo-api/tasks", {
            date: date,
            body: body,
            is_finished: 0
        })
        fetchTasks()
    }

    useEffect(() => {
        fetchTasks()
    }, [])

    return (
        <React.Fragment>
            <Header />
            <PostBar post={post} setPost={setPost} postTask={postTask}/>
            <TaskList tasks={tasks}/>
        </React.Fragment>
    )
}

export default App;
