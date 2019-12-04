import React from "react";

const PostBar = (props) => {
    const handleChange = event => {
        props.setPost(event.target.value);
    }
    const handleClick = () => {
        if (props.post === "") {
            return;
        }
        props.postTask(props.post, getDate());
        props.setPost("");
    }
    const getDate = () => {
        const today = new Date();
        const year = today.getFullYear();
        const month = today.getMonth() + 1;
        const day = today.getDate();
        return year + "-" + month + "-" + day
    }
    return (
        <div>
            <input value={props.post} onChange={handleChange}/>
            <button onClick={handleClick}>追加</button>
        </div>
    )
};

export default PostBar;
