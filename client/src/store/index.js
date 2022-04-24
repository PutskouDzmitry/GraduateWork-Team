import { combineReducers, configureStore } from "@reduxjs/toolkit";

import userActionsReducer from "./reducers/userActionsReducer";
import modalActionsReducer from "./reducers/modalActionsReducer";

const rootReducer = combineReducers({
  user: userActionsReducer,
  modals: modalActionsReducer,
});

const store = configureStore({
  reducer: rootReducer,
});

export default store;
