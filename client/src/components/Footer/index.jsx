import React from "react";
import PropTypes from "prop-types";
import { useHistory } from "react-router-dom";

import "./index.scss";

function Footer(props) {
  const history = useHistory();
  return (
    <div className="footer">
      <>
        <button
          className="button button_alt button_footer"
          onClick={() => history.push("/acs-parser")}
        >
          ACS Parser
        </button>
      </>
    </div>
  );
}

Footer.propTypes = {};

export default Footer;
