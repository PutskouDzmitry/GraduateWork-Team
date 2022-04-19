import React from "react";
import PropTypes from "prop-types";
import LoginLogoutButton from "../LoginLogoutButton";

import "./index.scss";

function Header() {
  return (
    <div className="header">
      <button
        style={{ visibility: "hidden" }}
        className="button button_alt"
        disabled
      ></button>
      <p className="header__text">Wi-Fi Radar</p>
      <LoginLogoutButton className="button button_alt" />
    </div>
  );
}

Header.propTypes = {};

export default Header;
