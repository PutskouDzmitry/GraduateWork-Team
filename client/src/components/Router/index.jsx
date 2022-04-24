import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { routerModalOpen } from "../../store/actions/modalActions";
import PropTypes from "prop-types";

import "./index.scss";

function Router({ coords, id }) {
  const dispatch = useDispatch();
  const [hovered, setHovered] = useState(false);
  const routerModalOpened = useSelector(
    (state) => state.modals.routerModalOpened
  );

  const mouseEnterHandler = () => {
    setHovered(true);
  };

  const mouseLeaveHandler = () => {
    setHovered(false);
  };

  const clickHandler = () => {
    dispatch(routerModalOpen());
  };

  return (
    <>
      <div
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
