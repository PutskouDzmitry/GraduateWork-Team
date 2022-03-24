import React from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { userLogIn } from "../../store/actions/userActions";
import { useHistory } from "react-router-dom";

function LogInForm(props) {
  const dispatch = useDispatch();
  const history = useHistory();

  function handleLogIn() {
    dispatch(userLogIn());
    history.push("/");
  }

  return (
    <div>
      Log In
      <input type="text" />
      <input type="password" />
      <button onClick={handleLogIn}>Submit</button>
    </div>
  );
}

LogInForm.propTypes = {};

export default LogInForm;
