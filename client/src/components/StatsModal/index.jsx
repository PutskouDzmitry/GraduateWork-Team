import React from "react";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";

import "./index.scss";
import { statsModalClose } from "../../store/actions/modalActions";

function StatsModal({ tableData }) {
  const dispatch = useDispatch();
  const statsModalOpened = useSelector(
    (state) => state.modals.statsModalOpened
  );

  return (
    <div className={statsModalOpened ? "stats-modal" : "stats-modal_hidden"}>
      <div className="stats-modal__block">
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>MAC</th>
              <th>Strength</th>
              <th>Quality</th>
              <th>Frequency</th>
              <th>Max speed</th>
              {/* {Object.keys(tableData[0]).map((key) => {
                return <th>{key}</th>;
              })} */}
            </tr>
          </thead>
          <tbody>
            {tableData.map((el) => {
              return (
                <tr>
                  {Object.values(el).map((val) => {
                    return <td>{val}</td>;
                  })}
                </tr>
              );
            })}
          </tbody>
        </table>
        <button
          className="button button_alt button_wide"
          onClick={() => dispatch(statsModalClose())}
        >
          Close
        </button>
      </div>
    </div>
  );
}

StatsModal.propTypes = {};

export default StatsModal;
