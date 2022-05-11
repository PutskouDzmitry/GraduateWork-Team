import routerActionTypes from "../constants/routerActionTypes";

export const addRouter = (id, coords, settings) => ({
  type: routerActionTypes.ADD_ROUTER,
  id,
  coords,
  settings,
});

export const updateRouter = (id, settings) => ({
  type: routerActionTypes.UPDATE_ROUTER,
  id,
  settings,
});

export const removeRouter = (id) => ({
  type: routerActionTypes.REMOVE_ROUTER,
  id,
});

export const removeAllRouters = () => ({
  type: routerActionTypes.REMOVE_ALL_ROUTERS,
});

export const setCurrentRouterId = (id) => ({
  type: routerActionTypes.SET_CURRENT_ROUTER_ID,
  id,
});

export const setCurrentRouterSettings = (settings) => ({
  type: routerActionTypes.SET_CURRENT_ROUTER_SETTINGS,
  settings,
});
