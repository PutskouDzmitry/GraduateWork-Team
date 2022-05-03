import { useState, useRef, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { addRouter } from "../../store/actions/routerActions";
import Router from "../Router";

import "./index.scss";

function Main() {
  const dispatch = useDispatch();
  const canvasOld = useRef(null);
  const canvasNew = useRef(null);
  const fileInput = useRef(null);
  const [isChanged, setIsChanged] = useState(false);
  const [isUploaded, setIsUploaded] = useState(false);
  const routers = useSelector((state) => state.routers.routersList);

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
  };

  const handleUpload = async () => {
    let formData = new FormData();
    let file = fileInput.current.files[0];
    formData.append("myFile", file);

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

  useEffect(() => {
    const currentCanvas = canvasOld.current;

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

    return () => {
      currentCanvas.removeEventListener("click", clickListener);
    };
  });

  return (
    <div className="main-block">
      {isChanged ? (
        <p className={isChanged ? "help-text" : "help-text_hidden"}>
          Click anywhere on the picture to add a router
        </p>
      ) : (
        <p className={isChanged ? "help-text_hidden" : "help-text"}>
          Add a building plan to start working
        </p>
      )}
      <div className={isChanged ? "canvas-wrapper" : "canvas-wrapper_hidden"}>
        <canvas
          width="600px"
          height="400px"
          className={isChanged ? "canvas" : "canvas canvas_hidden"}
          ref={canvasOld}
        ></canvas>
        <canvas
          width="600px"
          height="400px"
          className={isUploaded ? "canvas" : "canvas canvas_hidden"}
          ref={canvasNew}
        ></canvas>
        {routers.map((router) => {
          return (
            <Router coords={router.coords} id={router.id} key={router.id} />
          );
        })}
      </div>
      <input
        ref={fileInput}
        type="file"
        name="file"
        accept="*"
        onChange={handleChange}
      />
      <button onClick={handleUpload}>submit</button>
    </div>
  );
}

export default Main;
