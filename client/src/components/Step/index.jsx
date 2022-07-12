import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { removeStep } from "../../store/actions/stepActions";
import PropTypes from "prop-types";

import "./index.scss";

function Router({ coords, id }) {
  const dispatch = useDispatch();
  const steps = useSelector((state) => state.steps.stepsList);

  const clickHandler = (e) => {
    let currentStepId = e.nativeEvent.path.find((el) => {
      return el.getAttribute("name") == "step";
    }).id;
    dispatch(removeStep(currentStepId));
  };

  return (
    <>
      <div
        name="step"
        className="step"
        style={{ left: `${coords.left}px`, top: `${coords.top}px` }}
        id={id}
        onClick={clickHandler}
      ></div>
    </>
  );
}

Router.propTypes = {};

export default Router;
