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
        if (response.data === null || response.data === undefined) {
            setTasks([])
        } else {
            setTasks(response.data)
        }
    }

    const putTask = async (index, date, body, is_finished) => {
        await axios.put("http://localhost:8080/todo-api/tasks/" + index, {
            date: date,
            body: body,
            is_finished: is_finished
        })
        fetchTasks()
    }

    const postTask = async (date, body) => {
        await axios.post("http://localhost:8080/todo-api/tasks", {
            headers: {
                "Content-Type": "application/json"
            },
            date: date,
            body: body,
            is_finished: 0
        })
        fetchTasks()
    }

    const deleteTask = async (index) => {
        await axios.delete("http://localhost:8080/todo-api/tasks/" + index)
        fetchTasks()
    }

    useEffect(() => {
        fetchTasks()
    }, [])

    return (
        <React.Fragment>
            <Header />
            <PostBar post={post} setPost={setPost} postTask={postTask}/>
            <TaskList tasks={tasks} deleteTask={deleteTask} putTask={putTask}/>
        </React.Fragment>
    )
}

export default App;
