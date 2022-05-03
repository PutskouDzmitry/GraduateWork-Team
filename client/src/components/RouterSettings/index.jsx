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

  const [settings, setSettings] = useState({
    transmitterPower: 0,
    gainOfTransmittingAntenna: 0,
    gainOfReceivingAntenna: 0,
    speed: 0,
    signalLossTransmitting: 0,
    signalLossReceiving: 0,
    numberOfChannels: 0,
  });

  useEffect(() => {
    const {
      transmitterPower,
      gainOfTransmittingAntenna,
      gainOfReceivingAntenna,
      speed,
      signalLossTransmitting,
      signalLossReceiving,
      numberOfChannels,
    } = currentRouterSettings;

    setSettings({
      transmitterPower,
      gainOfTransmittingAntenna,
      gainOfReceivingAntenna,
      speed,
      signalLossTransmitting,
      signalLossReceiving,
      numberOfChannels,
    });
  }, [routerModalOpened]);

  const handleSubmit = (e) => {
    e.preventDefault();
    let {
      transmitterPower,
      gainOfTransmittingAntenna,
      gainOfReceivingAntenna,
      speed,
      signalLossTransmitting,
      signalLossReceiving,
      numberOfChannels,
    } = settingsForm.current;

    const settings = {
      transmitterPower: transmitterPower.value,
      gainOfTransmittingAntenna: gainOfTransmittingAntenna.value,
      gainOfReceivingAntenna: gainOfReceivingAntenna.value,
      speed: speed.value,
      signalLossTransmitting: signalLossTransmitting.value,
      signalLossReceiving: signalLossReceiving.value,
      numberOfChannels: numberOfChannels.value,
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
        {Object.keys(settings).map((key) => {
          return (
            <input
              key={key}
              type="number"
              name={`${key}`}
              placeholder={`${key}`}
              value={settings[key]}
              onChange={(e) =>
                setSettings({ ...settings, [key]: e.target.value })
              }
            />
          );
        })}
        <button onClick={handleSubmit}>submit</button>
        <button onClick={handleClose}>close</button>
        <button onClick={handleRemove}>remove router</button>
      </form>
    </div>
  );
}

RouterSettings.propTypes = {};

export default RouterSettings;
