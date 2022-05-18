import Tesseract from "tesseract.js";

export const objectsInfo = [
  {
    title: "Window (No Toning)",
    name: "windowNoToning",
    icon: <i className="fa-solid fa-table-cells-large"></i>,
    color: "#ccf2ff",
  },
  {
    title: "Window (With Toning)",
    name: "windowToning",
    icon: <i className="fa-brands fa-windows"></i>,
    color: "#00ace6",
  },
  {
    title: "Wooden Wall",
    name: "woodenWall",
    icon: <i className="fa-solid fa-tree"></i>,
    color: "#8a4928",
  },
  {
    title: "Interior Wall",
    name: "interiorWall",
    icon: <i className="fa-solid fa-table-cells"></i>,
    color: "#f8bf00",
  },
  {
    title: "Main Wall",
    name: "mainWall",
    icon: <i className="fa-solid fa-square"></i>,
    color: "#009933",
  },
  {
    title: "Aquarium",
    name: "aquarium",
    icon: <i className="fa-solid fa-fish-fins"></i>,
    color: "#00e6e6",
  },
];

export function dataURLtoBlob(dataurl) {
  var arr = dataurl.split(","),
    mime = arr[0].match(/:(.*?);/)[1],
    bstr = atob(arr[1]),
    n = bstr.length,
    u8arr = new Uint8Array(n);
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  return new Blob([u8arr], { type: mime });
}

export function parseACSfiles(fileList) {
  let result = [];

  for (let i = 0; i < fileList.length; i++) {
    let reader = new FileReader();
    let parser = new DOMParser();
    let file = fileList[i];

    reader.onload = (() => {
      return (fileContents) => {
        let resultItem = {
          id: i,
          signals: [],
        };
        const xmlContents = parser.parseFromString(
          fileContents.target.result,
          "text/xml"
        );
        const signals = xmlContents.getElementsByTagName("Client");
        const signalsArray = Array.from(signals);

        for (let j = 0; j < signalsArray.length; j++) {
          const signal = signalsArray[j];
          const nodes = Array.from(signal.childNodes);
          const AT_ID = nodes.find((el) => {
            return el.tagName === "AT_ID";
          }).textContent;
          const MAC = nodes.find((el) => {
            return el.tagName === "MAC";
          }).textContent;
          const LastSignalStrength = nodes.find((el) => {
            return el.tagName === "LastSignalStrength";
          }).textContent;
          const obj = { AT_ID, MAC, LastSignalStrength };
          resultItem.signals.push({
            id: j,
            obj,
          });
        }

        result.push(resultItem);
      };
    })(file);

    reader.readAsText(file);
  }

  return result;
}

export async function recognize(fileList, lang) {
  let parsedArray = [];

  for (let i = 0; i < fileList.length; i++) {
    const file = fileList[i];

    let parsedText = await Tesseract.recognize(file, lang).then(
      ({ data: { text } }) => {
        return text;
      }
    );
    parsedArray.push(parsedText);
  }
  return parsedArray;
}
