import React, { useState, useRef, useEffect } from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { routerModalClose } from "../../store/actions/modalActions";
import { updateRouter, removeRouter } from "../../store/actions/routerActions";

import "./index.scss";

function RouterSettings({}) {
  const dispatch = useDispatch();
  const settingsForm = useRef(null);
  const submitButton = useRef(null);
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
    e?.preventDefault();
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

  const handleKeyPress = (e) => {
    if (e.charCode === 13) {
      handleSubmit();
    }
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

  const toNormalString = (camelCaseString) => {
    const normalString = camelCaseString
      .replace(/([A-Z])/g, " $1")
      .replace(/^./, function (str) {
        return str.toUpperCase();
      });
    return normalString;
  };

  return (
    <div className={routerModalOpened ? "settings" : "settings_hidden"}>
      <form className="settings__form" ref={settingsForm}>
        {Object.keys(settings).map((key) => {
          return (
            <div key={key} className="settings__form__block">
              <label
                htmlFor={`${key}`}
                className="settings__form__block__label"
              >
                {toNormalString(key)}
              </label>
              <input
                type="number"
                className="settings__form__block__input"
                name={`${key}`}
                placeholder={`${key}`}
                value={settings[key]}
                onChange={(e) =>
                  setSettings({ ...settings, [key]: e.target.value })
                }
                onKeyPress={handleKeyPress}
              />
            </div>
          );
        })}
        <div className="settings__form__buttons">
          <button className="button button_gray" onClick={handleClose}>
            <i class="fa-solid fa-xmark"></i>
          </button>
          <button className="button button_red" onClick={handleRemove}>
            <i class="fa-solid fa-trash-can"></i>
          </button>
          <button
            ref={submitButton}
            className="button"
            onClick={handleSubmit}
            type="submit"
          >
            <i class="fa-solid fa-check"></i>
          </button>
        </div>
      </form>
    </div>
  );
}

RouterSettings.propTypes = {};

export default RouterSettings;
