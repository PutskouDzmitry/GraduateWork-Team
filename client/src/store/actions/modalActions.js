import modalActionTypes from "../constants/modalActionTypes";

export const routerModalOpen = () => ({
  type: modalActionTypes.ROUTER_MODAL_OPEN,
});

export const routerModalClose = () => ({
  type: modalActionTypes.ROUTER_MODAL_CLOSE,
});

export const loaderModalOpen = () => ({
  type: modalActionTypes.LOADER_MODAL_OPEN,
});

export const loaderModalClose = () => ({
  type: modalActionTypes.LOADER_MODAL_CLOSE,
});
