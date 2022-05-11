import { useState, useRef, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { addRouter, removeAllRouters } from "../../store/actions/routerActions";
import {
  objectModeOn,
  objectModeOff,
  setCurrentObject,
} from "../../store/actions/objectActions";
import Router from "../Router";
import { objectsInfo, dataURLtoBlob } from "../../constants";

import FileInput from "../FileInput";

import "./index.scss";

function Main() {
  const dispatch = useDispatch();
  const canvasOld = useRef(null);
  const canvasNew = useRef(null);
  const canvasForObjects = useRef(null);
  const fileInput = useRef(null);
  const [fileName, setFileName] = useState("No file chosen");
  const [isChanged, setIsChanged] = useState(false);
  const [isUploaded, setIsUploaded] = useState(false);
  const routers = useSelector((state) => state.routers.routersList);
  const isObjectModeOn = useSelector(
    (state) => state.objectsInfo.isObjectModeOn
  );
  const currentObject = useSelector((state) => state.objectsInfo.currentObject);

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
    dispatch(removeAllRouters());
  };

  const handleUpload = async () => {
    let formData = new FormData();
    // let file = fileInput.current.files[0];
    // formData.append("myFile", file);
    // formData.append("data", JSON.stringify(routers));

    const file = dataURLtoBlob(canvasOld.current.toDataURL());
    formData.append("myFile", file);
    formData.append("data", JSON.stringify(routers));

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

    xhr.open("POST", "http://localhost:8080/api/map/calculation", true);
    xhr.send(formData);
  };

  const handleObjectChange = (newObject) => {
    const obj = { name: newObject.name, color: newObject.color };
    dispatch(setCurrentObject(obj));
  };

  useEffect(() => {
    const currentCanvas = canvasOld.current;
    const objectsCanvas = canvasForObjects.current;
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
      let left = e.offsetX;
      let top = e.offsetY;
      let id = Date.now();
      let coords = { left, top };
      let settings = {
        transmitterPower: 0,
        gainOfTransmittingAntenna: 0,
        gainOfReceivingAntenna: 0,
        speed: 0,
        signalLossTransmitting: 0,
        signalLossReceiving: 0,
        numberOfChannels: 0,
      };
      dispatch(addRouter(id, coords, settings));
    };
    currentCanvas.addEventListener("click", clickListener);
    objectsCanvas.addEventListener("click", objectClickListener);

    return () => {
      currentCanvas.removeEventListener("click", clickListener);
      objectsCanvas.addEventListener("click", objectClickListener);
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
          <button className="button button_wide" onClick={toggleObjectMode}>
            Start drawing objects
          </button>
        )}
      </div>
      <div className="main-block__center">
        <p className="help-text">
          {isChanged
            ? isObjectModeOn
              ? "You are currently in object drawing mode"
              : "Click anywhere on the picture to add a router"
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
          {routers.map((router) => {
            return (
              <Router coords={router.coords} id={router.id} key={router.id} />
            );
          })}
        </div>
        <FileInput
          ref={fileInput}
          onChange={handleChange}
          fileName={fileName}
        />
        {isChanged ? (
          <button className="button button_special" onClick={handleUpload}>
            Submit
          </button>
        ) : (
          ""
        )}
      </div>
    </div>
  );
}

export default Main;
