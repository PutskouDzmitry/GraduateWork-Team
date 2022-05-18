import React from "react";
import PropTypes from "prop-types";
import RedirectButton from "../RedirectButton";
import { useSelector } from "react-redux";

import "./index.scss";

function LoaderModal() {
  const loaderModalOpened = useSelector(
    (state) => state.modals.loaderModalOpened
  );

  return (
    <div className={loaderModalOpened ? "loader-modal" : "loader-modal_hidden"}>
      <div className="lds-roller">
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
      </div>
    </div>
  );
}

LoaderModal.propTypes = {};

export default LoaderModal;
