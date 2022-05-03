import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { routerModalOpen } from "../../store/actions/modalActions";
import {
  setCurrentRouterId,
  setCurrentRouterSettings,
} from "../../store/actions/routerActions";
import PropTypes from "prop-types";

import "./index.scss";

function Router({ coords, id }) {
  const dispatch = useDispatch();
  const [hovered, setHovered] = useState(false);
  const routers = useSelector((state) => state.routers.routersList);

  const mouseEnterHandler = () => {
    setHovered(true);
  };

  const mouseLeaveHandler = () => {
    setHovered(false);
  };

  const clickHandler = (e) => {
    let currentRouterId = e.nativeEvent.path.find((el) => {
      return el.getAttribute("name") == "router";
    }).id;
    let currentRouterSettings = routers.find((router) => {
      return router.id == currentRouterId;
    }).settings;

    dispatch(setCurrentRouterId(currentRouterId));
    dispatch(setCurrentRouterSettings(currentRouterSettings));
    dispatch(routerModalOpen());
  };

  return (
    <>
      <div
        name="router"
        className="router"
        style={{ left: `${coords.left}px`, top: `${coords.top}px` }}
        id={id}
        onMouseEnter={mouseEnterHandler}
        onMouseLeave={mouseLeaveHandler}
        onClick={clickHandler}
      >
        <div className="router__main">
          {hovered ? (
            <button className="router__button">
              <i className="fa-solid fa-gears"></i>
            </button>
          ) : (
            <p className="router__text">.</p>
          )}
        </div>
      </div>
    </>
  );
}

Router.propTypes = {};

export default Router;
