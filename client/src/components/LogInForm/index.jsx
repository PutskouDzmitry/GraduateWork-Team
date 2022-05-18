import React, { useRef } from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { userLogIn } from "../../store/actions/userActions";
import { useHistory } from "react-router-dom";

import "./index.scss";

function LogInForm({ type }) {
  const dispatch = useDispatch();
  const history = useHistory();
  const loginUsername = useRef(null);
  const loginPassword = useRef(null);
  const signinUsername = useRef(null);
  const signinPassword = useRef(null);

  const handleLogIn = async () => {
    var xhr = new XMLHttpRequest();
    let formData = new FormData();
    const username = loginUsername.current.value;
    const password = loginPassword.current.value;
    console.log(username, password);
    formData.append("login", username);
    formData.append("password", password);
    xhr.onload = () => {
      dispatch(userLogIn());
      history.push("/");
    };
    xhr.open("POST", "http://localhost:8080/auth/login", true);
    xhr.send(formData);
  };

  const handleSignIn = async () => {
    var xhr = new XMLHttpRequest();
    let formData = new FormData();
    const username = signinUsername.current.value;
    const password = signinPassword.current.value;
    console.log(username, password);
    formData.append("login", username);
    formData.append("password", password);
    xhr.onload = () => {
      dispatch(userLogIn());
      history.push("/");
    };
    xhr.open("POST", "http://localhost:8080/auth/sign-up", true);
    xhr.send(formData);
  };

  return (
    <div className="main-block_login">
      {type === "login" ? (
        <>
          <div className="login-block">
            <p className="login-block__header">Log in</p>
            <hr className="login-block__line" />
            <div className="login-block__part">
              <i className="fa-solid fa-user icon"></i>
              <input
                ref={loginUsername}
                className="login-block__input"
                type="text"
                id="login"
                placeholder="Username"
              />
            </div>
            <div className="login-block__part">
              <i className="fa-solid fa-key icon"></i>
              <input
                ref={loginPassword}
                className="login-block__input"
                type="password"
                id="password"
                placeholder="Password"
              />
            </div>
            <button
              className="button button_special login-block__button"
              onClick={handleLogIn}
            >
              Submit
            </button>
          </div>
          <div className="block-divider">OR</div>
          <div className="main-block_login__part">
            <div className="login-block secondary">
              <p className="login-block__header">First time?</p>
              <hr className="login-block__line" />
              <button
                className="button login-block__button single-button"
                onClick={() => history.push("/register")}
              >
                Sign in
              </button>
            </div>
            <div className="block-divider small"></div>
            <div className="login-block secondary">
              <a
                className="button login-block__button single-button link"
                href="http://localhost:8080/auth/google"
              >
                Log in with <i className="fa-brands fa-google"></i>
              </a>
            </div>
          </div>
        </>
      ) : (
        <div className="login-block">
          <p className="login-block__header">Sign in</p>
          <hr className="login-block__line" />
          <div className="login-block__part">
            <i className="fa-solid fa-user icon"></i>
            <input
              ref={signinUsername}
              className="login-block__input"
              type="text"
              id="login"
              placeholder="Username"
            />
          </div>
          <div className="login-block__part">
            <i className="fa-solid fa-key icon"></i>
            <input
              ref={signinPassword}
              className="login-block__input"
              type="password"
              id="password"
              placeholder="Password"
            />
          </div>
          <button
            className="button button_special login-block__button"
            onClick={handleSignIn}
          >
            Submit
          </button>
        </div>
      )}
    </div>
  );
}

LogInForm.propTypes = {};

export default LogInForm;
