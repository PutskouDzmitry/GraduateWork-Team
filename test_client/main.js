function doUpload(fileElements) {
    let formData = new FormData();
    let file = fileElements.files; //return a JSON array aka FileList which contains of information like filename, size, type and etc
    formData.append("myFile", file[0]);

    var xhr = new XMLHttpRequest();

    xhr.onload = function () {
        console.log("Message from backend", xhr.response);
    };

    xhr.open("POST", "http://localhost:8080/api/map/calculation", true);
    xhr.send(formData);
}