import { combineReducers, configureStore } from "@reduxjs/toolkit";

import userActionsReducer from "./reducers/userActionsReducer";

const rootReducer = combineReducers({ user: userActionsReducer });

const store = configureStore({
  reducer: rootReducer,
});

export default store;
