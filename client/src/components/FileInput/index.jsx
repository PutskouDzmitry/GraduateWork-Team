import React from "react";
import PropTypes from "prop-types";
import "./index.scss";

const FileInput = React.forwardRef(({ onChange, fileName, multiple }, ref) => {
  const id = Date.now();
  return (
    <div className="file-input">
      <input
        ref={ref}
        type="file"
        name={`fileInput-${id}`}
        id={`fileInput-${id}`}
        accept="*"
        onChange={onChange}
        multiple={multiple}
        hidden
      />
      <label className="file-input__button" htmlFor={`fileInput-${id}`}>
        Choose File(s)
      </label>
      <div>{fileName}</div>
    </div>
  );
});

FileInput.propTypes = {};

export default FileInput;
