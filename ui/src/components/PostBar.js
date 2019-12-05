import React from "react";

const PostBar = (props) => {
    const handleChange = event => {
        props.setPost(event.target.value);
    }
    const handleClick = () => {
        if (props.post === "") {
            return;
        }
        props.postTask(getDate(), props.post);
        props.setPost("");
    }
    const handleKeyDown = event => {
        if (event.keyCode === 13) {
            if (props.post === "") {
                return;
            }
            props.postTask(getDate(), props.post);
            props.setPost("");
        }
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
            <input value={props.post} onChange={handleChange} onKeyDown={handleKeyDown}/>
            <button onClick={handleClick}>追加</button>
        </div>
    )
};

export default PostBar;
