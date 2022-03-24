import React, { useState } from "react";
import PropTypes from "prop-types";

import "./index.scss";

function Router({ coords, id }) {
  const [hovered, setHovered] = useState(false);

  const mouseEnterHandler = () => {
    setHovered(true);
  };

  const mouseLeaveHandler = () => {
    setHovered(false);
  };

  return (
    <div
      className="router"
      style={{ left: `${coords.left}px`, top: `${coords.top}px` }}
      id={id}
      onMouseEnter={mouseEnterHandler}
      onMouseLeave={mouseLeaveHandler}
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
  );
}

Router.propTypes = {};

export default Router;
