import React from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { useHistory } from "react-router-dom";
import { userLogIn, userLogOut } from "../../store/actions/userActions";

function LoginLogoutButton({ className }) {
  const dispatch = useDispatch();
  const isUserLoggedIn = useSelector((state) => state.user.isUserLoggedIn);
  const history = useHistory();

  if (isUserLoggedIn) {
    return (
      <button
        className={className}
        onClick={() => {
          dispatch(userLogOut());
          history.push("/");
        }}
      >
        Log out
      </button>
    );
  } else {
    return (
      <button
        className={className}
        onClick={() => {
          history.push("/login");
          // dispatch(userLogIn());
        }}
      >
        Log in
      </button>
    );
  }
}

LoginLogoutButton.propTypes = {};

export default LoginLogoutButton;
