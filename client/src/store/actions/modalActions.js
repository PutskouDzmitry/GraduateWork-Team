import modalActionTypes from "../constants/modalActionTypes";

export const routerModalOpen = () => ({
  type: modalActionTypes.ROUTER_MODAL_OPEN,
});

export const routerModalClose = () => ({
  type: modalActionTypes.ROUTER_MODAL_CLOSE,
});
