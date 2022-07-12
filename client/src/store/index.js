import { combineReducers, configureStore } from "@reduxjs/toolkit";

import userActionsReducer from "./reducers/userActionsReducer";
import modalActionsReducer from "./reducers/modalActionsReducer";
import routerActionsReducer from "./reducers/routerActionsReducer";
import stepActionsReducer from "./reducers/stepActionsReducer";
import objectActionsReducer from "./reducers/objectActionsReducer";

const rootReducer = combineReducers({
  user: userActionsReducer,
  modals: modalActionsReducer,
  routers: routerActionsReducer,
  steps: stepActionsReducer,
  objectsInfo: objectActionsReducer,
});

const store = configureStore({
  reducer: rootReducer,
});

export default store;
