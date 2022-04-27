import React, { useState, useRef, useEffect } from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { routerModalClose } from "../../store/actions/modalActions";
import { updateRouter, removeRouter } from "../../store/actions/routerActions";

import "./index.scss";

function RouterSettings({}) {
  const dispatch = useDispatch();
  const settingsForm = useRef(null);
  const routerModalOpened = useSelector(
    (state) => state.modals.routerModalOpened
  );
  const currentRouterId = useSelector(
    (state) => state.routers.currentRouter.id
  );
  const currentRouterSettings = useSelector(
    (state) => state.routers.currentRouter.settings
  );
  const [val1, setVal1] = useState(currentRouterSettings.val1);
  const [val2, setVal2] = useState(currentRouterSettings.val2);

  useEffect(() => {
    setVal1(currentRouterSettings.val1);
    setVal2(currentRouterSettings.val2);
  }, [routerModalOpened]);

  const handleSubmit = (e) => {
    e.preventDefault();
    let { field1, field2 } = settingsForm.current;
    const settings = {
      val1: field1.value,
      val2: field2.value,
    };
    dispatch(updateRouter(currentRouterId, settings));
    dispatch(routerModalClose());
  };

  const handleRemove = (e) => {
    e.preventDefault();
    dispatch(removeRouter(currentRouterId));
    dispatch(routerModalClose());
  };

  const handleClose = (e) => {
    e.preventDefault();
    dispatch(routerModalClose());
  };

  return (
    <div className={routerModalOpened ? "settings" : "settings_hidden"}>
      <form className="settings__form" ref={settingsForm}>
        <input
          type="text"
          name="field1"
          placeholder="val1"
          value={val1}
          onChange={(e) => setVal1(e.target.value)}
        />
        <input
          type="number"
          name="field2"
          placeholder="val2"
          value={val2}
          onChange={(e) => setVal2(e.target.value)}
        />
        <button onClick={handleSubmit}>submit</button>
        <button onClick={handleClose}>close</button>
        <button onClick={handleRemove}>remove router</button>
      </form>
    </div>
  );
}

RouterSettings.propTypes = {};

export default RouterSettings;
