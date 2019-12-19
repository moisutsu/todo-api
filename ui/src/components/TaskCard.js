import React from "react";

import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Checkbox from "@material-ui/core/Checkbox";

const useStyles = makeStyles({
    card: {
      minWidth: 275,
      backgroundColor: "silver",
      margin: 5, // 周囲との隙間
    },
    bullet: {
      display: 'inline-block',
      margin: '0 2px',
      transform: 'scale(0.8)',
    },
    title: {
      fontSize: 14,
    },
    pos: {
      marginBottom: 12,
    },
  });

const TaskCard = (props) => {
    const classes = useStyles();
    const handleDeleteClick = () => {
        props.deleteTask(props.index);
    }
    const handleCompleteClick = () => {
        props.putTask(props.index, props.date, props.body, !props.is_finished);
    }
    
    return (
    <Card className={classes.card}>
        <CardContent>
            <Typography className={classes.title} color="textSecondary" gutterBottom>
                {props.date}
            </Typography>
            <Typography variant="h5" component="h2">
                <Checkbox
                    checked={props.is_finished}
                    onClick={handleCompleteClick}
                    value="secondary"
                    color="primary"
                />
                {props.body}
            </Typography>
        </CardContent>
        <CardActions>
        <Button onClick={handleDeleteClick}>削除</Button>
        </CardActions>
    </Card>
)};

export default TaskCard;
