import React from "react";
import PropTypes from "prop-types";
import RedirectButton from "../RedirectButton";

import "./index.scss";

function Footer() {
  return (
    <div className="footer">
      <>
        <RedirectButton path={"/"} label="Router Map" />
        <RedirectButton path={"/acs-parser"} label="ACS Parser" />
        <RedirectButton path={"/mobile-parser"} label="Mobile Parser" />
        <RedirectButton path={"/acrylic-parser"} label="Acrylic Parser" />
      </>
    </div>
  );
}

Footer.propTypes = {};

export default Footer;
