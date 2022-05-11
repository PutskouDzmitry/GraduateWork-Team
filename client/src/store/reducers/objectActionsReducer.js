import objectActionTypes from "../constants/objectActionTypes";
import { objectsInfo } from "../../constants";

let initialState = {
  isObjectModeOn: false,
  currentObject: {
    name: objectsInfo[0].name,
    color: objectsInfo[0].color,
  },
};

const objectActionsReducer = (state = initialState, action) => {
  let newState = {};

  switch (action.type) {
    case objectActionTypes.OBJECT_MODE_ON:
      newState = {
        ...state,
        isObjectModeOn: true,
      };
      return newState;

    case objectActionTypes.OBJECT_MODE_OFF:
      newState = {
        ...state,
        isObjectModeOn: false,
      };
      return newState;

    case objectActionTypes.SET_CURRENT_OBJECT:
      newState = {
        ...state,
        currentObject: action.currentObject,
      };
      return newState;

    case objectActionTypes.REMOVE_CURRENT_OBJECT:
      newState = {
        ...state,
        currentObject: {},
      };
      return newState;

    default:
      return state;
  }
};

export default objectActionsReducer;
