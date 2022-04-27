import { combineReducers, configureStore } from "@reduxjs/toolkit";

import userActionsReducer from "./reducers/userActionsReducer";
import modalActionsReducer from "./reducers/modalActionsReducer";
import routerActionsReducer from "./reducers/routerActionsReducer";

const rootReducer = combineReducers({
  user: userActionsReducer,
  modals: modalActionsReducer,
  routers: routerActionsReducer,
});

const store = configureStore({
  reducer: rootReducer,
});

export default store;
