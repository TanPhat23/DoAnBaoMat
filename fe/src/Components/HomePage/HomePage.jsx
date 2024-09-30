import React, { useEffect, useState } from 'react';
import Narbar from "../Navbar/Navbar";
import bannerBackGound from "./banner.jpg";
import { FiArrowRight } from "react-icons/fi";
import { useNavigate } from "react-router-dom";
import axios from 'axios';

const HomePage = () => {
  const navigate = useNavigate();
  const [todos, setTodos] = useState([]); // Use state to hold todos

  const getTodos = async () => {
    try {
      const response = await axios.get("http://localhost:8080/todos", {
        withCredentials: true,
      });

      console.log(response.data);
      setTodos(response.data); // Set the fetched todos in state
    } catch (error) {
      console.log("Error fetching todos:", error);
    }
  };

  useEffect(() => {
    getTodos();
  }, []); // Run once on mount

  const getCurrentUser = async() =>{
    try{
      const userResponse = await axios.get("http://localhost:8080/currentuser", {
        withCredentials: true,
      }); 
      console.log("Current user data:", userResponse.data);
    }catch (error) {
      console.log("Error fetching currentuser:", error);
    }
  };

  useEffect(() => {
    getCurrentUser();
  }, []); // Run once on mount

  

  return (
    <div className="Home-container">
      <Narbar />
      <div className="home-banner-container">
        <div className="home-bannerImage-container">
          {/* <img src={bannerBackGound} alt=""/> */}
        </div>
        <div className="home-text-section">
          <h1 className="primary-heading">YOUR SAFE IS MY WORK</h1>
          <button
            className="secondary-button"
            onClick={() => navigate("/loginPage")}
          >
            LET TRY NOW
            <FiArrowRight />
          </button>
        </div>
      </div>  

      {/* Displaying the todos */}
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
    </div>
  );
}

export default HomePage;
