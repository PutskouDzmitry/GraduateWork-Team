import { useState, useRef, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { addRouter, removeAllRouters } from "../../store/actions/routerActions";
import { updateSavedMaps } from "../../store/actions/userActions";
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
  const [fileName, setFileName] = useState("No map chosen");
  const [isChanged, setIsChanged] = useState(false);
  const [isUploaded, setIsUploaded] = useState(false);
  const routers = useSelector((state) => state.routers.routersList);
  const isObjectModeOn = useSelector(
    (state) => state.objectsInfo.isObjectModeOn
  );
  const savedMaps = useSelector((state) => state.user.savedMaps);
  const isUserLoggedIn = useSelector((state) => state.user.isUserLoggedIn);
  const currentObject = useSelector((state) => state.objectsInfo.currentObject);

  // выбор файла
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


  //отрисовка нового изображения
  const handleUpload = async () => {
    let formData = new FormData();

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

  const saveMap = () => {
    let formData = new FormData();
    const file = dataURLtoBlob(canvasOld.current.toDataURL());
    const fileOutput = dataURLtoBlob(canvasNew.current.toDataURL());

    formData.append("myFile", file);
    formData.append("myFileOutput", fileOutput);
    formData.append("data", JSON.stringify(routers));

    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
      // update last save
      var xhrUpdate = new XMLHttpRequest();
      xhrUpdate.onload = function () {
        const parsedMaps = JSON.parse(xhrUpdate.response);
        dispatch(updateSavedMaps(parsedMaps));
      };

      xhrUpdate.open("POST", "http://localhost:8080/api/map/load", true);
      xhrUpdate.send();
    };

    xhr.open("POST", "http://localhost:8080/api/map/save", true);
    xhr.send(formData);
  };

  const loadMap = () => {
    dispatch(removeAllRouters());
    setFileName("Last Saved Map");
    canvasOld.current.getContext("2d").clearRect(0, 0, 600, 400);
    let ctx = canvasOld.current.getContext("2d");
    let url = `data:image/png;base64,${savedMaps.Data[0].PathInput}`;
    let img = new Image();
    img.onload = function () {
      ctx.drawImage(img, 0, 0);
    };
    img.src = url;
    setIsChanged(true);

    savedMaps.Data[0].Data.forEach((el, index) => {
      let left = el.coordinates_of_router.x;
      let top = el.coordinates_of_router.y;
      let id = `${index}-${Date.now()}`;
      let coords = { left, top };
      let settings = {
        transmitterPower: el.transmitter_power,
        gainOfTransmittingAntenna: el.gain_of_transmitting_antenna,
        gainOfReceivingAntenna: el.gain_of_receiving_antenna,
        speed: el.speed,
        signalLossTransmitting: el.signal_loss_receiving,
        signalLossReceiving: el.signal_loss_transmitting,
        numberOfChannels: el.number_of_channels,
        scale: el.scale,
        type: el.type_of_signal,
      };
      dispatch(addRouter(id, coords, settings));
    });
  };

  // useEffect for updating saves
  useEffect(() => {
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
      const parsedMaps = JSON.parse(xhr.response);
      dispatch(updateSavedMaps(parsedMaps));
    };

    xhr.open("POST", "http://localhost:8080/api/map/load", true);
    xhr.send();
  }, []);

  useEffect(() => {
    const currentCanvas = canvasOld.current;
    const objectsCanvas = canvasForObjects.current;
    const ctxCurrent = currentCanvas.getContext("2d");
    const ctxObjects = objectsCanvas.getContext("2d");

    // рисование стенок
    const objectClickListener = (e) => {
      if (isObjectModeOn) {
        let x1 = 0,
          y1 = 0,
          x2 = 0,
          y2 = 0;
        let isCurrentlyDrawing = false;

        //инициализиуер координаты на мыши
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
        scale: 0,
        typeOfSignal: "2.4",
      };
      dispatch(addRouter(id, coords, settings));
    };
    currentCanvas.addEventListener("click", clickListener);
    objectsCanvas.addEventListener("click", objectClickListener);

    //подтягивает из хранилищя роутероы, если роутеры изменлились, то и тут изменились
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
          <>
            <button className="button button_wide" onClick={toggleObjectMode}>
              Start drawing objects
            </button>
            {isUserLoggedIn ? (
              <>
                <button className="button button_wide" onClick={saveMap}>
                  Save Current Map
                </button>
                <button className="button button_wide" onClick={loadMap}>
                  Load Saved Map
                </button>
              </>
            ) : (
              ""
            )}
          </>
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
          id="mapInput"
          ref={fileInput}
          onChange={handleChange}
          fileName={fileName}
          multiple={false}
        />
        {isChanged ? (
          <>
            <button className="button button_special" onClick={handleUpload}>
              Submit
            </button>
            {isUploaded ? <div className="rainbow-image"></div> : ""}
          </>
        ) : (
          ""
        )}
      </div>
    </div>
  );
}

export default Main;
