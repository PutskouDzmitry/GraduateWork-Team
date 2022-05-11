import React from "react";
import PropTypes from "prop-types";
import "./index.scss";

const FileInput = React.forwardRef(({ onChange, fileName }, ref) => {
  return (
    <div className="file-input">
      <input
        ref={ref}
        type="file"
        name="fileInput"
        id="fileInput"
        accept="*"
        onChange={onChange}
        hidden
      />
      <label className="file-input__button" htmlFor="fileInput">
        Choose File
      </label>
      <div>{fileName}</div>
    </div>
  );
});

FileInput.propTypes = {};

export default FileInput;
