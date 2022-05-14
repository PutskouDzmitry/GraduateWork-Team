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
