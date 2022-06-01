import { useState, useRef, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { addStep } from "../../store/actions/stepActions";
import {
  objectModeOn,
  objectModeOff,
  setCurrentObject,
} from "../../store/actions/objectActions";
import { statsModalOpen } from "../../store/actions/modalActions";
import { objectsInfo, dataURLtoBlob, parseACSfiles } from "../../constants";

import Step from "../Step";
import FileInput from "../FileInput";
import StatsModal from "../StatsModal";

import "./index.scss";

function ACSparser() {
  const dispatch = useDispatch();
  const canvasOld = useRef(null);
  const canvasNew = useRef(null);
  const canvasForObjects = useRef(null);
  const fileInput = useRef(null);
  const acsFileInput = useRef(null);
  const [fileName, setFileName] = useState("No map chosen");
  const [acsFileName, setAcsFileName] = useState("No ACS files chosen");
  const [acsParsed, setAcsParsed] = useState([]);
  const [isChanged, setIsChanged] = useState(false);
  const [isAcsChanged, setIsAcsChanged] = useState(false);
  const [acsFilesNumber, setAcsFilesNumber] = useState(0);
  const [isUploaded, setIsUploaded] = useState(false);
  const steps = useSelector((state) => state.steps.stepsList);
  const isObjectModeOn = useSelector(
    (state) => state.objectsInfo.isObjectModeOn
  );
  const currentObject = useSelector((state) => state.objectsInfo.currentObject);
  const [statTableData, setStatTableData] = useState([]);

  const handleChange = async () => {
    canvasOld.current.getContext("2d").clearRect(0, 0, 600, 400);
    let ctx = canvasOld.current.getContext("2d");
    let url = URL.createObjectURL(fileInput.current.files[0]);
    let img = new Image();
    img.onload = function () {
      ctx.drawImage(img, 0, 0);
    };
    img.src = url;
    setIsChanged(true);
    setFileName(fileInput.current.files[0].name);
  };

  const acsHandleChange = async () => {
    const parsedFiles = parseACSfiles(acsFileInput.current.files);
    setAcsParsed(parsedFiles);
    setIsAcsChanged(true);
    setAcsFileName(`${acsFileInput.current.files.length} files chosen`);
    setAcsFilesNumber(acsFileInput.current.files.length);
  };

  const handleUpload = async () => {
    let formData = new FormData();

    const file = dataURLtoBlob(canvasOld.current.toDataURL());
    formData.append("myFile", file);
    formData.append(
      "data",
      JSON.stringify({
        steps,
        acsParsed,
      })
    );

    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
      canvasNew.current.getContext("2d").clearRect(0, 0, 600, 400);
      let ctx = canvasNew.current.getContext("2d");
      let url = `data:image/png;base64,${xhr.response}`;
      let img = new Image();
      img.onload = function () {
        ctx.drawImage(img, 0, 0);
      };
      img.src = url;
      setIsUploaded(true);
    };

    xhr.open("POST", "http://localhost:8080/api/map/fluxMigrator", true);
    xhr.send(formData);
  };

  const handleObjectChange = (newObject) => {
    const obj = { name: newObject.name, color: newObject.color };
    dispatch(setCurrentObject(obj));
  };

  useEffect(() => {
    const currentCanvas = canvasOld.current;
    const objectsCanvas = canvasForObjects.current;
    const newCanvas = canvasNew.current;
    const ctxCurrent = currentCanvas.getContext("2d");
    const ctxObjects = objectsCanvas.getContext("2d");

    const objectClickListener = (e) => {
      if (isObjectModeOn) {
        let x1 = 0,
          y1 = 0,
          x2 = 0,
          y2 = 0;
        let isCurrentlyDrawing = false;

        objectsCanvas.addEventListener("mousedown", function (e) {
          isCurrentlyDrawing = true;
          x1 = e.offsetX;
          y1 = e.offsetY;
          x2 = e.offsetX;
          y2 = e.offsetY;
          reDrawObject();
        });

        objectsCanvas.addEventListener("mousemove", function (e) {
          x2 = e.offsetX;
          y2 = e.offsetY;
          reDrawObject();
        });

        objectsCanvas.addEventListener("mouseup", function (e) {
          finalDrawObject();
          isCurrentlyDrawing = false;
        });

        function reDrawObject() {
          if (isCurrentlyDrawing === true) {
            ctxObjects.strokeStyle = currentObject.color;
            ctxObjects.fillStyle = currentObject.color;
            ctxObjects.clearRect(0, 0, 600, 400);
            ctxObjects.beginPath();
            ctxObjects.lineWidth = "2";
            ctxObjects.rect(x1, y1, x2 - x1, y2 - y1);
            ctxObjects.fill();
          }
        }

        function finalDrawObject() {
          ctxCurrent.strokeStyle = currentObject.color;
          ctxCurrent.fillStyle = currentObject.color;
          ctxCurrent.beginPath();
          ctxCurrent.lineWidth = "2";
          ctxCurrent.rect(x1, y1, x2 - x1, y2 - y1);
          ctxCurrent.fill();
          ctxObjects.clearRect(0, 0, 600, 400);
        }
      }
    };

    const clickListener = (e) => {
      if (acsFilesNumber != steps.length) {
        let left = e.offsetX;
        let top = e.offsetY;
        let coords = { left, top };
        let id = Date.now();
        dispatch(addStep(id, coords));
      }
    };

    const statClickListener = (e) => {
      let left = e.offsetX;
      let top = e.offsetY;
      let coords = { left, top };
      var xhr = new XMLHttpRequest();
      let formData = new FormData();
      formData.append("data", JSON.stringify(acsParsed));
      formData.append("point", JSON.stringify({ coords }));
      xhr.onload = () => {
        let data = [];
        JSON.parse(xhr.response).forEach((el) => {
          let tableLine = {};
          Object.entries(el).forEach(([key, val]) => {
            tableLine[key] = val;
          });
          data.push(tableLine);
        });
        setStatTableData(data);
        dispatch(statsModalOpen());
      };
      xhr.open("POST", "http://localhost:8080/api/map/getInfo", true);
      xhr.send(formData);
    };

    currentCanvas.addEventListener("click", clickListener);
    objectsCanvas.addEventListener("click", objectClickListener);
    newCanvas.addEventListener("click", statClickListener);

    return () => {
      currentCanvas.removeEventListener("click", clickListener);
      objectsCanvas.removeEventListener("click", objectClickListener);
      newCanvas.removeEventListener("click", statClickListener);
    };
  });

  const toggleObjectMode = () => {
    isObjectModeOn ? dispatch(objectModeOff()) : dispatch(objectModeOn());
    canvasForObjects.current.focus();
  };

  return (
    <div className="main-block">
      <div
        className={isChanged ? "main-block__side" : "main-block__side_hidden"}
      >
        {isObjectModeOn ? (
          <>
            <div className="main-block__group">
              {objectsInfo.map((el) => {
                return (
                  <button
                    className={
                      currentObject.name == el.name
                        ? "button button_small button_active"
                        : "button button_small"
                    }
                    title={el.title}
                    name={el.name}
                    key={el.name}
                    onClick={() => handleObjectChange(el)}
                  >
                    {el.icon}
                  </button>
                );
              })}
            </div>
            <button className="button button_wide" onClick={toggleObjectMode}>
              Stop drawing objects
            </button>
          </>
        ) : (
          <>
            <button className="button button_wide" onClick={toggleObjectMode}>
              Start drawing objects
            </button>
            <FileInput
              id="acsInput"
              ref={acsFileInput}
              onChange={acsHandleChange}
              fileName={acsFileName}
              multiple={true}
            />
          </>
        )}
      </div>
      <div className="main-block__center">
        <p className="help-text">
          {isChanged
            ? isObjectModeOn
              ? "You are currently in object drawing mode"
              : isAcsChanged
              ? `${acsFilesNumber} ACS files chosen. ${
                  acsFilesNumber - steps.length
                } steps left`
              : "Add ACS files to continue"
            : "Add a building plan to start working"}
        </p>
        <div className={isChanged ? "canvas-wrapper" : "canvas-wrapper_hidden"}>
          <canvas
            width="600px"
            height="400px"
            className={isChanged ? "canvas" : "canvas_hidden"}
            ref={canvasOld}
          ></canvas>
          <canvas
            width="600px"
            height="400px"
            className={
              isChanged
                ? isObjectModeOn
                  ? "canvas__objects"
                  : "canvas_hidden"
                : "canvas_hidden"
            }
            ref={canvasForObjects}
          ></canvas>
          <canvas
            width="600px"
            height="400px"
            className={isUploaded ? "canvas" : "canvas_hidden"}
            ref={canvasNew}
          ></canvas>
          {steps.map((step) => {
            return <Step coords={step.coords} id={step.id} key={step.id} />;
          })}
        </div>
        <FileInput
          id="mapInput"
          ref={fileInput}
          onChange={handleChange}
          fileName={fileName}
          multiple={false}
        />
        {isChanged ? (
          <button className="button button_special" onClick={handleUpload}>
            Submit
          </button>
        ) : (
          ""
        )}
      </div>
      <StatsModal tableData={statTableData} />
    </div>
  );
}

export default ACSparser;
