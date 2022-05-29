import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import "./index.scss";
import { useSelector } from "react-redux";
import LoginLogoutButton from "../LoginLogoutButton";

function Sidebar(props) {
  const [isOpened, setIsOpened] = useState(false);
  const isUserLoggedIn = useSelector((state) => state.user.isUserLoggedIn);

  return (
    <div className={`sidebar ${isOpened ? "sidebar_opened" : ""}`}>
      <div className="sidebar__body">
        {isUserLoggedIn ? (
          <>You are logged in as: dimaputkovvv@gmail.com</>
        ) : (
          <>
            <p className="sidebar__body__text">
              You are not authenicated. Please,
              <LoginLogoutButton className="button" /> to enable saving and
              loading
            </p>
          </>
        )}
      </div>
      <button
        className={`sidebar__button ${
          isOpened ? "sidebar__button_opened" : ""
        }`}
        onClick={() => {
          setIsOpened((curr) => !curr);
        }}
      >
        {isOpened ? (
          <i className="fa-solid fa-angle-left"></i>
        ) : (
          <i className="fa-solid fa-angle-right"></i>
        )}
      </button>
    </div>
  );
}

Sidebar.propTypes = {};

export default Sidebar;
