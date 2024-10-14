import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';
import todoImage from "./todo.jpg"
const Admin = () => {
    const navigate = useNavigate();
    const [todos, setTodos] = useState([]);

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
    }, []); 
    return (
      <div className="max-w-2xl mx-auto p-4 mt-32 ">
        <div className="fixed top-0 left-0 w-full h-full z-[-2]">
          <img src={todoImage} alt="Background" className="object-container w-full h-full" />
        </div>
        <h2 className="text-2xl font-bold text-center mb-4">Your Todos:</h2>
          {todos.length > 0 ? (
            <table className="w-full border border-blue-500">
              <thead>
                <tr className="bg-blue-300">
                  <th className="border border-blue-700 p-2">Text</th> 
                  <th className="border border-blue-700 p-2">Done</th>
                </tr>
              </thead>
                <tbody>
                  {todos.map((todo) => (
                    <tr key={todo._id} > 
                      <td className="border border-blue-700 p-2">{todo.Text}</td>
                      <td className="border border-blue-700 p-2">
                      <input type="checkbox" checked = {todo.Done}></input>
                      </td>
                    </tr>
                  ))}
                </tbody>
            </table>
          ) : (
            <p  className="text-center">No todos available.</p>
          )}
      </div>
  );
}

export default Admin