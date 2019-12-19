import React from "react";
import Button from "@material-ui/core/Button"
import Input from "@material-ui/core/Input"

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
            <Input value={props.post} onChange={handleChange} onKeyDown={handleKeyDown}/>
            <Button onClick={handleClick}>追加</Button>
        </div>
    )
};

export default PostBar;
