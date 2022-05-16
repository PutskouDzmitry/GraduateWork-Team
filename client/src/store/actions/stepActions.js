import stepActionTypes from "../constants/stepActionTypes";

export const addStep = (id, coords) => ({
  type: stepActionTypes.ADD_STEP,
  id,
  coords,
});

export const removeStep = (id) => ({
  type: stepActionTypes.REMOVE_STEP,
  id,
});

export const removeAllSteps = () => ({
  type: stepActionTypes.REMOVE_ALL_STEPS,
});
