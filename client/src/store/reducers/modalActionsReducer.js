import modalActionTypes from "../constants/modalActionTypes";

let initialState = {
  routerModalOpened: false,
  loaderModalOpened: false,
};

const modalActionsReducer = (state = initialState, action) => {
  let newState = {};

  switch (action.type) {
    case modalActionTypes.ROUTER_MODAL_OPEN:
      newState = {
        ...state,
        routerModalOpened: true,
      };
      return newState;

    case modalActionTypes.ROUTER_MODAL_CLOSE:
      newState = {
        ...state,
        routerModalOpened: false,
      };
      return newState;

    case modalActionTypes.LOADER_MODAL_OPEN:
      newState = {
        ...state,
        loaderModalOpened: true,
      };
      return newState;

    case modalActionTypes.LOADER_MODAL_CLOSE:
      newState = {
        ...state,
        loaderModalOpened: false,
      };
      return newState;

    default:
      return state;
  }
};

export default modalActionsReducer;
