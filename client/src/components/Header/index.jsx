import React from "react";
import PropTypes from "prop-types";
import LoginLogoutButton from "../LoginLogoutButton";

import "./index.scss";

function Header({ withLoginButton }) {
  return (
    <div className="header">
      {withLoginButton ? (
        <>
          <button
            style={{ visibility: "hidden" }}
            className="button button_alt"
            disabled
          ></button>
          <p className="header__text">Wi-Fi Radar</p>
          <LoginLogoutButton className="button button_alt" />
        </>
      ) : (
        <>
          <button
            style={{ visibility: "hidden" }}
            className="button button_alt"
            disabled
          >
            Log In
          </button>
          <p className="header__text">Wi-Fi Radar</p>
          <button
            style={{ visibility: "hidden" }}
            className="button button_alt"
            disabled
          ></button>
        </>
      )}
    </div>
  );
}

Header.propTypes = {};

export default Header;
