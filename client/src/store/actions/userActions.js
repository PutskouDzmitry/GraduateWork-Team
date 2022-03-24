import userActionTypes from "../constants/userActionTypes";

export const userLogIn = () => ({
  type: userActionTypes.USER_LOG_IN,
});

export const userLogOut = () => ({
  type: userActionTypes.USER_LOG_OUT,
});
