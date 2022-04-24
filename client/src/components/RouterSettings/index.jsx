import React, { useState } from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { routerModalClose } from "../../store/actions/modalActions";

import "./index.scss";

function RouterSettings({}) {
  const dispatch = useDispatch();
  const routerModalOpened = useSelector(
    (state) => state.modals.routerModalOpened
  );

  const handleSubmit = (e) => {
    e.preventDefault();
    // dispatch(routerModalClose());
  };

  const handleRemove = () => {
    // dispatch(routerModalClose());
  };

  const handleClose = () => {
    dispatch(routerModalClose());
  };

  return (
    <div className={routerModalOpened ? "settings" : "settings_hidden"}>
      <form className="settings__form">
        <input type="text" />
        <input type="number" />
        <button onClick={handleSubmit}>submit</button>
        <button onClick={handleClose}>close</button>
        <button onClick={handleRemove}>remove router</button>
      </form>
    </div>
  );
}

RouterSettings.propTypes = {};

export default RouterSettings;
