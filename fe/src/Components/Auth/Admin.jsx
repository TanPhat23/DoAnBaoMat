import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';

const Admin = () => {
    const navigate = useNavigate();
    const [todos, setTodos] = useState([]); // Use state to hold todos

    const getTodos = async () => {
        try {
        const response = await axios.get("http://localhost:8080/todos", {
            withCredentials: true,
        });

        console.log(response.data);
        setTodos(response.data);
        } catch (error) {
        console.log("Error fetching todos:", error);
        }
    };
  
    useEffect(() => {
        getTodos();
    }, []); // Run once on mount
  return (
    <div className="todo-list">
        <h2>Your Todos:</h2>
        {todos.length > 0 ? (
          todos.map(todo => (
            <div key={todo._id} className="todo-item"> {/* Use unique ID for key */}
              <p>Text: {todo.Text}</p>
              <p>Done: {todo.Done ? "Yes" : "No"}</p> {/* Display whether done */}
            </div>
          ))
        ) : (
          <p>No todos available.</p>
        )}
    </div>
  )
}

export default Admin
