import React from "react";
import PropTypes from "prop-types";
import { useHistory, useLocation } from "react-router-dom";

import "./index.scss";

function RedirectButton({ path, label }) {
  const history = useHistory();
  const location = useLocation();

  return (
    <button
      className={
        location.pathname == path
          ? "button button_footer"
          : "button button_alt button_footer"
      }
      onClick={() => history.push(path)}
    >
      {label}
    </button>
  );
}

RedirectButton.propTypes = {};

export default RedirectButton;
