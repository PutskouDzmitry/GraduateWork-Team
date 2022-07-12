import React from "react";
import PropTypes from "prop-types";
import "./index.scss";

const FileInput = React.forwardRef(
  ({ id, onChange, fileName, multiple }, ref) => {
    const dateNow = Date.now();
    return (
      <div className="file-input">
        <input
          ref={ref}
          type="file"
          name={`fileInput-${id}`}
          id={`${id}-${dateNow}`}
          accept="*"
          onChange={onChange}
          multiple={multiple}
          hidden
        />
        <label className="file-input__button" htmlFor={`${id}-${dateNow}`}>
          Choose File(s)
        </label>
        <div>{fileName}</div>
      </div>
    );
  }
);

FileInput.propTypes = {};

export default FileInput;
