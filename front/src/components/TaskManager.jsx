import React, { useEffect, useState } from 'react';
import TaskList from './TaskList';
import { Link } from 'react-router-dom';

const TaskManager = ({tasks, fetchTasks}) => {

    return (
        <div>
            <TaskList tasks={tasks} fetchTasks={fetchTasks}></TaskList>
            <Link to="/add">Add New Task</Link>
        </div>
    );
}
export default TaskManager;