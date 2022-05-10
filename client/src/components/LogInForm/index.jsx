import React from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { userLogIn } from "../../store/actions/userActions";
import { useHistory } from "react-router-dom";

function LogInForm(props) {
  const dispatch = useDispatch();
  const history = useHistory();

  const handleLogIn = async () => {
      var xhr = new XMLHttpRequest();
      let formData = new FormData();
      const username = document.getElementById('login').value;
      const password = document.getElementById('password').value;
      console.log(username, password)
      formData.append("login", username);
      formData.append("password", password);
      xhr.open("POST", "http://localhost:8080/auth/login", true);
      xhr.send(formData);
  }

  return (
    <div>
        Log In
        <input type="text" id="login" placeholder="Login" />
        <input type="password" id="password" placeholder="Password" />
        <button onClick={handleLogIn}>Submit</button>
    </div>
  );
}

LogInForm.propTypes = {};

export default LogInForm;
