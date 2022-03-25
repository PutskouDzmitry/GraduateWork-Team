import { useState, useRef, useEffect } from "react";
import { FileUploader } from "react-drag-drop-files";
import Router from "../Router";

import "./index.scss";

const fileTypes = ["JPG", "JPEG", "PNG", "GIF"];

function Main() {
  const canvas = useRef(null);
  const [isUploaded, setIsUploaded] = useState(false);
  const [routers, setRouters] = useState([]);
  const [formData, setFormData] = useState(new FormData());

  const handleChange = async (file) => {
    canvas.current.getContext("2d").clearRect(0, 0, 600, 400);
    let ctx = canvas.current.getContext("2d");
    let url = URL.createObjectURL(file[0]);
    let img = new Image();
    img.onload = function () {
      ctx.drawImage(img, 0, 0);
    };
    img.src = url;

    let imageBlob = await new Promise((resolve) =>
      canvas.current.toBlob(resolve, "image/png")
    );
    let formData = new FormData();
    formData.append("file", imageBlob, file[0].name);
    console.log(formData);
    setIsUploaded(true);
    setFormData(formData);
  };

  useEffect(() => {
    const currentCanvas = canvas.current;

    const clickListener = (e) => {
      let left = e.offsetX;
      let top = e.offsetY;
      console.log(e);
      let newRouters = routers.slice();
      newRouters.push({ coords: { left, top }, id: Date.now() });
      setRouters(newRouters);
    };
    currentCanvas.addEventListener("click", clickListener);

    return () => {
      currentCanvas.removeEventListener("click", clickListener);
    };
  });

  return (
    <div className="main-block">
      {isUploaded ? (
        <p className={isUploaded ? "help-text" : "help-text_hidden"}>
          Click anywhere on the picture to add a router
        </p>
      ) : (
        <p className={isUploaded ? "help-text_hidden" : "help-text"}>
          Add a building plan to start working
        </p>
      )}
      <div className={isUploaded ? "canvas-wrapper" : "canvas-wrapper_hidden"}>
        <canvas
          width="600px"
          height="400px"
          className={isUploaded ? "canvas" : "canvas canvas_hidden"}
          ref={canvas}
        ></canvas>
        {routers.map((router) => {
          return (
            <Router coords={router.coords} id={router.id} key={router.id} />
          );
        })}
      </div>
      <FileUploader
        multiple={true}
        handleChange={handleChange}
        name="file"
        types={fileTypes}
      />
      <button
        onClick={async () => {
          const picture = await fetch("http://localhost:8080/api/wifi/kek", {
            method: "POST",
            body: formData,
          }).then((response) => response.json());
          console.log(picture);
        }}
      >
        Kek
      </button>
    </div>
  );
}

export default Main;
