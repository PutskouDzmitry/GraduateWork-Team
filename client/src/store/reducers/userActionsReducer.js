import userActionTypes from "../constants/userActionTypes";

let initialState = {
  isUserLoggedIn: false,
  savedMaps: [],
};

const userActionsReducer = (state = initialState, action) => {
  let newState = {};

  switch (action.type) {
    case userActionTypes.USER_LOG_IN:
      newState = {
        ...state,
        isUserLoggedIn: true,
      };
      return newState;

    case userActionTypes.USER_LOG_OUT:
      newState = {
        ...state,
        isUserLoggedIn: false,
      };
      return newState;

    case userActionTypes.UPDATE_SAVED_MAPS:
      newState = {
        ...state,
        savedMaps: action.maps,
      };
      return newState;

    default:
      return state;
  }
};

export default userActionsReducer;
