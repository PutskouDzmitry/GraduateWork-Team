import userActionTypes from "../constants/userActionTypes";

export const userLogIn = () => ({
  type: userActionTypes.USER_LOG_IN,
});

export const userLogOut = () => ({
  type: userActionTypes.USER_LOG_OUT,
});

export const updateSavedMaps = (maps) => ({
  type: userActionTypes.UPDATE_SAVED_MAPS,
  maps,
});
