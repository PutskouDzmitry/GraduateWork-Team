import objectActionTypes from "../constants/objectActionTypes";

export const objectModeOn = () => ({
  type: objectActionTypes.OBJECT_MODE_ON,
});

export const objectModeOff = () => ({
  type: objectActionTypes.OBJECT_MODE_OFF,
});

export const setCurrentObject = (currentObject) => ({
  type: objectActionTypes.SET_CURRENT_OBJECT,
  currentObject,
});

export const removeCurrentObject = () => ({
  type: objectActionTypes.REMOVE_CURRENT_OBJECT,
});
